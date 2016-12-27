package gozwave

import (
	"io"
	"strconv"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/pborman/uuid"
	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/functions"
	"github.com/stampzilla/gozwave/interfaces"
	"github.com/stampzilla/gozwave/serialapi"
	"github.com/tarm/serial"
)

type Connection struct {
	Name      string
	Baud      int
	port      io.ReadWriteCloser
	Connected bool

	// Keep track of requests wating a response
	inFlight    map[string]*sendPackage
	updateChans map[string]chan interface{}
	send        chan *sendPackage
	lastCommand string // Uuid of last sent command
	lastResult  chan byte

	reportCallback func(byte, commands.Report)

	sync.RWMutex
}

type sendPackage struct {
	message []byte
	uuid    string
	result  byte // ACK, NAC, CAN

	function       byte
	commandclass   byte
	expectedReport byte
	node           byte

	timeout time.Duration

	returnChan chan *serialapi.Message
}

func NewConnection() *Connection {
	z := &Connection{
		inFlight:       make(map[string]*sendPackage),
		updateChans:    make(map[string]chan interface{}),
		send:           make(chan *sendPackage),
		lastResult:     make(chan byte),
		reportCallback: func(byte, commands.Report) {},
	}
	return z
}

func (self *Connection) RegisterNode(address byte) chan interface{} {
	c := make(chan interface{})

	self.Lock()
	self.updateChans[strconv.Itoa(int(address))] = c
	self.Unlock()

	return c
}

func (conn *Connection) Write(msg interfaces.Encodable) error {
	pkg := newSendPackage(msg.Encode())
	conn.send <- pkg
	return nil
}
func (conn *Connection) WriteWithTimeout(msg interfaces.Encodable, t time.Duration) (<-chan *serialapi.Message, error) {
	pkg := newSendPackage(msg.Encode())
	pkg.returnChan = make(chan *serialapi.Message)
	pkg.timeout = t

	conn.send <- pkg

	return pkg.returnChan, nil
}
func (conn *Connection) WriteAndWaitForReport(msg interfaces.Encodable, t time.Duration, er byte) (<-chan commands.Report, error) {
	pkg := newSendPackage(msg.Encode())
	returnChan := make(chan commands.Report)
	pkg.returnChan = make(chan *serialapi.Message)
	pkg.timeout = t
	pkg.expectedReport = er

	conn.send <- pkg

	go func() {
		defer close(returnChan)
		for msg := range pkg.returnChan {
			if f, ok := msg.Data.(*functions.FuncApplicationCommandHandler); ok {
				returnChan <- f.Data
				return
			}

			logrus.Errorf("Received wrong type: %t", msg)
		}
	}()

	return returnChan, nil
}

/*func (self *Connection) Send(data interfaces.Encodable, timeout time.Duration) chan *Message {
	return self.SendRaw(data.Encode(), timeout)
}

func (self *Connection) SendRaw(data []byte, timeout time.Duration) chan *Message {
	return self.SendRawAndWaitForResponse(data, timeout, 0)
}

func (self *Connection) SendRawAndWaitForResponse(data []byte, timeout time.Duration, command byte) chan *Message {
	returnChan := make(chan *Message)

	// Compile message
	msg := append([]byte{0x00, 0x00}, data...)
	msg = append(msg, 0x00)

	// Add length
	msg[0] = byte(len(msg) - 1)

	// Add checksum
	msg[len(msg)-1] = generateChecksum(msg)

	// Add header
	msg = append([]byte{0x01}, msg...)

	uuid := uuid.New()
	//logrus.Infof("Sending: %s", strings.TrimSpace(hex.Dump(msg)))
	pkg := &sendPackage{
		message:    msg,
		uuid:       uuid,
		returnChan: returnChan,
	}

	if len(data) > 0 {
		pkg.function = data[0]
	}
	if len(data) > 1 {
		pkg.node = data[1]
	}
	if len(data) > 4 {
		pkg.commandclass = data[3]

		if command != 0 {
			pkg.command = command
		}
	}

	self.send <- pkg

	// Timeout
	go func() {
		<-time.After(timeout)
		self.Lock()
		for index, c := range self.inFlight {
			if index == uuid {
				logrus.Warnf("TIMEOUT: %s - %#v", uuid, c)
				delete(self.inFlight, uuid)
				close(c.returnChan)
				c.returnChan = nil
			}
		}
		self.Unlock()
	}()

	return returnChan
}*/

func (self *Connection) Connect(connectChan chan error) (err error) {
	defer func() {
		logrus.Error("Disonnected\n\n")
		self.Connected = false
	}()

	self.Lock()
	c := &serial.Config{Name: self.Name, Baud: self.Baud}
	self.port, err = serial.OpenPort(c)
	self.Unlock()

	if err != nil {
		select {
		case connectChan <- err:
		default:
		}
		return
	}

	go func() {
		<-time.After(time.Millisecond * 200)
		logrus.Debug("Connected\n\n")
		select {
		case connectChan <- nil:
		default:
		}
		self.Connected = true
	}()

	go self.Writer()
	return self.Reader()
}

func (self *Connection) Writer() {
	logrus.Infof("Starting send worker")
	for {
		<-time.After(time.Millisecond * 50)
		select {
		case pkg := <-self.send:
			logrus.Infof("Sending: %x", pkg.message)

			self.Lock()
			self.lastCommand = ""

			self.inFlight[pkg.uuid] = pkg
			self.lastCommand = pkg.uuid

			if pkg.timeout != 0 { // Only add the message to the inflight list if someone is waiting for an response
				go self.timeoutWorker(pkg)
			}

			logrus.Infof("Write: %x", pkg.message)
			self.port.Write(pkg.message)
			self.Unlock()

			select {
			case result := <-self.lastResult:
				//logrus.Infof("RESULT: %s", pkg.uuid)
				logrus.Infof("Send result: %x, %s", pkg.message, result)
				pkg.result = result
			case <-time.After(time.Second):
				logrus.Infof("Send timeout: %x", pkg.message)
				// SEND TIMEOUT
			}
			self.lastCommand = ""
		}
	}
}

func (self *Connection) Reader() error {
	incomming := make([]byte, 0)
	age := time.Now()

	for {
		buf := make([]byte, 128)

		n, err := self.port.Read(buf)
		if err != nil {
			logrus.Error("Serial read failed: ", err)
			self.port.Close()
			return err
		}

		//fmt.Println(hex.Dump(buf[:n]))

		// If data in buffer is older than 40ms, clear buffer before appending data
		if time.Now().Sub(age) > (time.Millisecond * 40) {
			incomming = make([]byte, 0)
		}

		incomming = append(incomming, buf[:n]...)
		age = time.Now()

		for len(incomming) > 0 {
			logrus.Warnf("Try to decode: %x", incomming)
			l, msg := serialapi.Decode(incomming)

			if l == 1 {
				for index, c := range self.inFlight {
					self.RLock()
					if c.uuid == self.lastCommand {
						select {
						case self.lastResult <- incomming[0]:
						case <-time.After(time.Millisecond * 50):
						}

						logrus.Warnf("Command result!?: %s - %#v", c.uuid, c)
						if msg.IsNAK() {
							logrus.Warnf("Command failed: %s - %#v", c.uuid, c)
							delete(self.inFlight, index)
							close(c.returnChan)
							c.returnChan = nil
						}

						if msg.IsCAN() {
							go func() {
								<-time.After(time.Millisecond * 100)
								self.send <- c
							}()
						}
					}
					self.RUnlock()
				}
			}

			// The message is not compleatly read yet, wait for some more data
			if l > len(incomming) {
				break
			}

			if l == -1 { // Invalid checksum
				incomming = incomming[1:] // Remove first char and try again
				logrus.Infof("Removing first byte, buffer len=%d", len(incomming))

				continue
			}

			if l > 1 { // A message was received
				self.Lock()
				for index, c := range self.inFlight {
					if !c.Match(incomming) {
						continue
					}

					if c.returnChan != nil {
						select {
						case c.returnChan <- msg: // Try to deliver message
						default:
						}

						close(c.returnChan)
						c.returnChan = nil
					}
					delete(self.inFlight, index)
				}
				self.Unlock()
			}

			if cmd, ok := msg.Data.(*functions.FuncApplicationCommandHandler); ok {
				// Deliver the update to the parent
				go self.reportCallback(msg.NodeId, cmd.Data)

				//self.DeliverUpdate(msg.NodeId, cmd.Data)
				//go self.eventCallback(msg.NodeId, cmd.Data);
			}

			//switch cmd := msg.Data.(type) {
			//case *functions.FuncApplicationCommandHandler:
			//}

			//logrus.Infof("Recived: %s", strings.TrimSpace(hex.Dump(incomming)))
			logrus.Infof("Recived: %x", incomming)
			incomming = incomming[l:]
			if l > 1 {
				self.port.Write([]byte{0x06}) // Send ACK to stick
			}

		}
	}

}

func (self *Connection) DeliverUpdate(id byte, msg interface{}) {
	logrus.Infof("DeliverUpdate: %T", msg)
	c, ok := self.updateChans[strconv.Itoa(int(id))]
	if ok {
		go func() {
			select {
			case c <- msg:
			case <-time.After(time.Second):
			}
		}()
	}
}

func newSendPackage(data []byte) *sendPackage {
	pkg := &sendPackage{
		message: serialapi.CompileMessage(data),
		uuid:    uuid.New(),
	}

	if len(data) > 0 {
		pkg.function = data[0]
	}
	if len(data) > 1 {
		pkg.node = data[1]
	}
	if len(data) > 4 {
		pkg.commandclass = data[3]
	}

	return pkg
}

func (c *sendPackage) Match(incomming []byte) bool {
	//if len(incomming) > 3 && c.function != 0 && c.function != incomming[3] && !(c.function == functions.SendData && incomming[3] == 0x04) {

	// SerialAPI specific
	if !MatchByteAt(incomming, c.function, 3) && !(c.function == functions.SendData && MatchByteAt(incomming, functions.ApplicationCommandHandler, 3)) {
		//logrus.Warnf("Skipping pkg %s, function %x != %x", c.uuid, c.function, incomming[3])
		return false
	}

	if c.function == functions.GetNodeProtocolInfo {
		return true
	}

	// Z-wave specific
	if !MatchByteAt(incomming, c.node, 5) {
		//logrus.Warnf("Skipping pkg %s, node %x != %x", c.uuid, c.node, incomming[5])
		return false
	}
	if !MatchByteAt(incomming, c.commandclass, 7) {
		//logrus.Warnf("Skipping pkg %s, commandclass %x is not %x", c.uuid, c.commandclass, incomming[7])
		return false
	}
	if !MatchByteAt(incomming, c.expectedReport, 8) {
		logrus.Errorf("Skipping pkg %s, expectedReport %x is not %x", c.uuid, c.expectedReport, incomming[8])
		return false
	}

	//if len(incomming) > 5 && c.node != 0 && c.node != incomming[5] {
	//return false
	//}
	//if len(incomming) > 7 && c.commandclass != 0 && c.commandclass != incomming[7] {
	//return false
	//}
	//if len(incomming) > 8 && c.command != 0 && c.command != incomming[8] {
	//return false
	//}

	return true
}

func (conn *Connection) timeoutWorker(sp *sendPackage) {
	<-time.After(sp.timeout)

	conn.Lock()
	for index, c := range conn.inFlight {
		if index == sp.uuid {
			logrus.Warnf("TIMEOUT: %s", sp.uuid)
			delete(conn.inFlight, sp.uuid)
			close(c.returnChan)
			c.returnChan = nil
		}
	}
	conn.Unlock()
}

func MatchByteAt(message []byte, b byte, position int) bool {
	if b == 0 {
		return true
	}

	if len(message) < position {
		return false
	}

	return message[position] == b
}
