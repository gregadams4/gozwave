package gozwave

import (
	"io"
	"strconv"
	"sync"
	"time"

	"github.com/gregadams4/gozwave/commands"
	"github.com/gregadams4/gozwave/functions"
	"github.com/gregadams4/gozwave/interfaces"
	"github.com/gregadams4/gozwave/serialapi"

	"github.com/Sirupsen/logrus"
	"github.com/pborman/uuid"
)

type Connection struct {
	readWriteCloser io.ReadWriteCloser
	portOpener      PortOpener
	Connected       bool

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

func (conn *Connection) RegisterNode(address byte) chan interface{} {
	c := make(chan interface{})

	conn.Lock()
	conn.updateChans[strconv.Itoa(int(address))] = c
	conn.Unlock()

	return c
}

func (conn *Connection) Write(msg interfaces.Encodable) error {
	pkg := newSendPackage(msg.Encode())
	conn.send <- pkg
	return nil
}

func (conn *Connection) WriteEncoded(msg []byte) error {
	pkg := newSendPackage(msg)
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

func (conn *Connection) WriteEncodedWithTimeout(msg []byte, t time.Duration) (<-chan *serialapi.Message, error) {
	pkg := newSendPackage(msg)
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

			logrus.Errorf("WriteAndWaitForReport: Received wrong type: %v", msg)
		}
	}()

	return returnChan, nil
}

func (conn *Connection) WriteEncodedAndWaitForReport(msg []byte, t time.Duration, er byte) (<-chan commands.Report, error) {
	pkg := newSendPackage(msg)
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

			logrus.Errorf("WriteAndWaitForReport: Received wrong type: %v", msg)
		}
	}()

	return returnChan, nil
}

func (conn *Connection) Connect(connectChan chan error) (err error) {
	defer func() {
		logrus.Error("Disonnected")
		conn.Connected = false
	}()

	conn.readWriteCloser, err = conn.portOpener.Open()

	if err != nil {
		select {
		case connectChan <- err:
		default:
		}
		return
	}

	go func() {
		<-time.After(time.Millisecond * 200)
		logrus.Debug("Connected")
		select {
		case connectChan <- nil:
		default:
		}
		conn.Connected = true
	}()

	go conn.Writer()
	return conn.Reader()
}

func (conn *Connection) Writer() {
	logrus.Infof("Starting send worker")
	for {
		<-time.After(time.Millisecond * 50)
		select {
		case pkg := <-conn.send:
			conn.Lock()
			conn.lastCommand = ""

			conn.inFlight[pkg.uuid] = pkg
			conn.lastCommand = pkg.uuid

			if pkg.timeout != 0 { // Only add the message to the inflight list if someone is waiting for an response
				go conn.timeoutWorker(pkg)
			}

			logrus.Debugf("Write: %x", pkg.message)
			conn.readWriteCloser.Write(pkg.message)
			conn.Unlock()

			select {
			case result := <-conn.lastResult:
				pkg.result = result
			case <-time.After(time.Second):
				// SEND TIMEOUT
			}
			conn.lastCommand = ""
		}
	}
}

func (conn *Connection) Reader() error {
	incomming := make([]byte, 0)
	age := time.Now()

	for {
		buf := make([]byte, 128)

		n, err := conn.readWriteCloser.Read(buf)
		if err != nil {
			logrus.Error("Serial read failed: ", err)
			conn.readWriteCloser.Close()
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
			l, msg := serialapi.Decode(incomming)
			// logrus.Info(msg)
			if l == 1 {
				for index, c := range conn.inFlight {
					conn.RLock()
					if c.uuid == conn.lastCommand {
						select {
						case conn.lastResult <- incomming[0]:
						case <-time.After(time.Millisecond * 50):
						}

						if msg.IsNAK() {
							logrus.Warnf("Command failed: %s - %#v", c.uuid, c)
							delete(conn.inFlight, index)
							close(c.returnChan)
							c.returnChan = nil
						}

						if msg.IsCAN() {
							go func() {
								<-time.After(time.Millisecond * 100)
								conn.send <- c
							}()
						}
					}
					conn.RUnlock()
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
				conn.Lock()
				for index, c := range conn.inFlight {
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
					delete(conn.inFlight, index)
				}
				conn.Unlock()
			}

			if cmd, ok := msg.Data.(*functions.FuncApplicationCommandHandler); ok {
				// Deliver the update to the parent
				go conn.reportCallback(msg.NodeId, cmd.Data)
			}

			//logrus.Infof("Recived: %s", strings.TrimSpace(hex.Dump(incomming)))
			logrus.Debugf("Recived: %x", incomming)
			incomming = incomming[l:]
			if l > 1 {
				conn.readWriteCloser.Write([]byte{0x06}) // Send ACK to stick
			}

		}
	}

}

func (conn *Connection) DeliverUpdate(id byte, msg interface{}) {
	logrus.Infof("DeliverUpdate: %T", msg)
	c, ok := conn.updateChans[strconv.Itoa(int(id))]
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
