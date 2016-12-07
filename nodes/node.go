package nodes

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/database"
	"github.com/stampzilla/gozwave/events"
	"github.com/stampzilla/gozwave/functions"
	"github.com/stampzilla/gozwave/serialapi"
)

func New(address byte, connection *serialapi.Connection, eventQue chan interface{}) (*node, chan struct{}) {
	n := &node{
		Id_:        address,
		connection: connection,
		eventQue:   eventQue,
	}

	c := make(chan struct{})

	go n.Worker(c)

	return n, c
}

type node struct {
	Id_     byte `json:"id"`
	IsAwake bool `json:"is_awake"`

	ProtocolInfo        *functions.FuncGetNodeProtocolInfo
	ManufacurerSpecific *commands.CmdManufacturerSpecific

	Device interface{}

	connection *serialapi.Connection
	eventQue   chan interface{}

	awake      chan struct{}
	identified bool
}

func (n *node) Id() byte {
	return n.Id_
}

func (n *node) isAwake() chan struct{} {
	if n.IsAwake {
		c := make(chan struct{})
		close(c)
		return c
	}

	if n.awake == nil {
		n.awake = make(chan struct{})
	}

	return n.awake
}

func (n *node) Worker(basicDone chan struct{}) {

	updateChan := n.connection.RegisterNode(n.Id())
	go func() {
		for {
			select {
			case update := <-updateChan:
				//logrus.Errorf("Node received update: %T", update)
				switch msg := update.(type) {
				case *serialapi.Message:
					switch body := msg.Data.(type) {
					case *functions.FuncApplicationCommandHandler:
						switch body.Data.(type) {
						case *commands.CmdWakeUp:
							logrus.Error("NODE RECEIVED WAKEUP")
							if n.awake != nil {
								close(n.awake)
								n.awake = nil
							}
						}
					}
				}
			}
		}
	}()

	go n.Identify(basicDone)

	if n.Id() != 1 {
		n.pushEvent(events.NodeDiscoverd{
			Address: n.Id(),
		})
	}

	for {
		select {}
	}
}
func (n *node) Identify(basicDone chan struct{}) {
	logrus.Warnf("Started identification on node %d", n.Id())
	defer logrus.Warnf("Ended identification on node %d", n.Id())

	defer func() {
		if basicDone != nil {
			close(basicDone)
			basicDone = nil
		}
	}()

	if n.Id() == 1 {
		return
	}

	for {
		if n.ProtocolInfo == nil {
			err := n.RequestProtocolInfo()
			if err != nil {
				<-time.After(time.Second * 10)
				continue
			}
			n.pushEvent(events.NodeUpdated{
				Address: n.Id(),
			})
		}

		if n.ProtocolInfo != nil {
			n.IsAwake = n.ProtocolInfo.Listening
		}

		if basicDone != nil {
			close(basicDone)
			basicDone = nil
		}

		<-n.isAwake()
		//<-self.Connection.SendRaw([]byte{functions.GetNodeProtocolInfo, byte(index + 1)}) // Request node information
		//		nodeinfo := self.WaitForGetNodeProtocolInfo()

		// All manufacturer specific information such as vendor, vendors product ID and product type
		if n.ManufacurerSpecific == nil {
			err := n.RequestManufacturerSpecific()
			if err != nil {
				<-time.After(time.Second * 10)
				continue
			}

			n.Device = database.New(n.ManufacurerSpecific.Manufacturer, n.ManufacurerSpecific.Type, n.ManufacurerSpecific.Id)

			n.pushEvent(events.NodeUpdated{
				Address: n.Id(),
			})

		}

		resp := <-n.connection.SendRaw([]byte{
			functions.SendData, // Function
			byte(n.Id()),       // Node id
			0x02,               // Length
			commands.Version,   // Command
			0x11,               // VersionCmd_Get
			0x00,
			//0x05, // TransmitOptions?
			//0x23, // Callback?
		}, time.Second*10) // Request node information

		logrus.Println("%#v", resp)

		//<-self.Connection.SendRaw([]byte{functions.IsFailedNode, byte(index + 1)}) // Request is failed node
		//	<-self.SendRaw([]byte{0xa0, byte(index + 1)}) // Request ?

		// All firmware and Z-Wave version information
		// All switching and reporting capabilities including current switching states and sensor values.
		// Number and maximum size of association groups
		// Configuration values if known
		// The Gateways will then do some initial setup:
		// If needed a certain predefined wakeup interval is set for battery operated devices
		// If available and requested certain standard defaults behaviours are configured in the device
		// If association groups are available the gateway will always put its own Node ID into these association groups in order to be informed about status changes and to be prepared for using associations for scene switching
		// The user can change all values. However it needs to be clear that particularly removing the gateways Node ID from the association groups may cause malfunctions of the gateway.

		n.identified = true
		return
	}
}

func (self *node) pushEvent(event interface{}) {
	select {
	case self.eventQue <- event:
	default:
	}
}
