package nodes

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/stampzilla/gozwave/commands"
	database "github.com/stampzilla/gozwave/database"
	"github.com/stampzilla/gozwave/events"
	"github.com/stampzilla/gozwave/functions"
	"github.com/stampzilla/gozwave/interfaces"

	"github.com/Sirupsen/logrus"
)

func New(address int) *Node {
	n := &Node{
		Id:            address,
		StateBool:     make(map[string]bool),
		StateFloat:    make(map[string]float64),
		pushEventFunc: func(interface{}) {},
	}

	//go n.Worker()

	return n
}

type Node struct {
	Id      int  `json:"id"`
	IsAwake bool `json:"is_awake"`

	ProtocolInfo        *functions.FuncGetNodeProtocolInfo
	ManufacurerSpecific *commands.ManufacturerSpecificReport

	Device *database.Device

	Endpoints []*Endpoint

	StateBool  map[string]bool
	StateFloat map[string]float64
	statesOk   bool

	connection    interfaces.Writer
	pushEventFunc func(interface{})

	awake      chan struct{}
	identified bool

	sync.RWMutex
}

//TODO not use empty interface
//type ProcessEventFunc func(commands.Report)

func (n *Node) Setup(connection interfaces.Writer, pushEventFunc func(interface{})) {
	n.connection = connection
	n.pushEventFunc = pushEventFunc
}

func (n *Node) isAwake() chan struct{} {
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

func (n *Node) ProcessEvent(event commands.Report) {
	switch data := event.(type) {
	case *commands.AlarmReport:
		if data.Status == 0xFF {
			n.StateBool["alarm"] = true
			if data.Type != 0 {
				n.StateFloat["alarm"] = float64(data.Type)
				n.StateFloat["alarmLastType"] = float64(data.Type)
			}
			n.pushEvent(events.NodeUpdated{
				Address: n.Id,
			})

			<-time.After(time.Second)
			n.StateBool["alarm"] = false
			n.StateFloat["alarm"] = 0
			n.pushEvent(events.NodeUpdated{
				Address: n.Id,
			})
		}
	case *commands.ConfigurationReport:
		n.pushEvent(events.NodeUpdated{
			Address: n.Id,
		})
	case *commands.WakeUpReport:
		logrus.Error("NODE RECEIVED WAKEUP")
		if n.awake != nil {
			close(n.awake)
			n.awake = nil
		}
	case *commands.SwitchBinaryReport:
		n.StateBool["on"] = data.Status
		n.pushEvent(events.NodeUpdated{
			Address: n.Id,
		})
	case *commands.SwitchMultilevelReport:
		n.StateBool["on"] = data.Level > 0
		n.StateFloat["level"] = float64(data.Level)
		n.pushEvent(events.NodeUpdated{
			Address: n.Id,
		})
	case *commands.SensorMultiLevelReport:
		//n.StateFloat[data.TypeString+" ("+data.Unit+")"] = data.Value
		key := strings.ToLower(data.TypeString + "_" + data.Unit)
		n.StateFloat[key] = data.Value
		n.pushEvent(events.NodeUpdated{
			Address: n.Id,
		})
	}

}

func (n *Node) Identify() {
	logrus.Infof("Started identification on node %d", n.Id)
	defer logrus.Infof("Ended identification on node %d", n.Id)

	for {
		if n.ProtocolInfo == nil {
			resp, err := n.RequestProtocolInfo()
			if err != nil {
				logrus.Errorf("Node ident: Failed RequestProtocolInfo: %s", err.Error())
				<-time.After(time.Second * 10)
				continue
			}

			n.Lock()
			n.ProtocolInfo = resp
			n.IsAwake = resp.Listening
			n.pushEvent(events.NodeUpdated{
				Address: n.Id,
			})
			n.Unlock()
		}

		//<-self.Connection.SendRaw([]byte{functions.GetNodeProtocolInfo, byte(index + 1)}) // Request node information
		//		nodeinfo := self.WaitForGetNodeProtocolInfo()

		// All manufacturer specific information such as vendor, vendors product ID and product type
		if n.ManufacurerSpecific == nil {
			<-n.isAwake()
			resp, err := n.RequestManufacturerSpecific()
			if err != nil {
				logrus.Errorf("Node ident: Failed ManufacurerSpecific: %s", err.Error())
				<-time.After(time.Second * 10)
				continue
			}

			n.Lock()
			n.ManufacurerSpecific = resp
			n.Device = database.New(n.ManufacurerSpecific.Manufacturer, n.ManufacurerSpecific.Type, n.ManufacurerSpecific.Id)
			n.pushEvent(events.NodeUpdated{
				Address: n.Id,
			})
			n.Unlock()
		}

		// Get node version
		//resp := <-n.connection.SendRaw([]byte{
		//functions.SendData, // Function
		//byte(n.Id),         // Node id
		//0x02,               // Length
		//commands.Version,   // Command
		//0x11,               // VersionCmd_Get
		//0x00,
		////0x05, // TransmitOptions?
		////0x23, // Callback?
		//}, time.Second*10) // Request node information
		//logrus.Println("%#v", resp)

		// associoation with controller
		//if n.Id == 5 {
		//logrus.Println("")
		//logrus.Println("")
		//resp := <-n.connection.SendRaw([]byte{
		//functions.SendData,   // Function
		//byte(n.Id),           // Node id
		//0x04,                 // Length
		//commands.Association, // Command
		//0x01,                 // AssociationCmd_Set
		//0x01,                 // Group
		//0x01,                 // Target node (controller)
		//0x25,                 // TransmitOptions?
		////0x23, // Callback?
		//}, time.Second*10) // Request node information
		//logrus.Println("")
		//logrus.Println("")
		//logrus.Println("%#v", resp)
		//}

		// Request node endpoints
		n.RLock()
		if n.Endpoints == nil {
			n.RUnlock()
			<-n.isAwake()
			err := n.RequestEndpoints()
			if err != nil {
				<-time.After(time.Second * 10)
				continue
			}

			n.RLock()
			n.pushEvent(events.NodeUpdated{
				Address: n.Id,
			})
			n.RUnlock()
			continue
		}
		n.RUnlock()

		n.RLock()
		if !n.statesOk {
			n.RUnlock()
			<-n.isAwake()
			err := n.RequestStates()
			if err != nil {
				<-time.After(time.Second * 10)
				continue
			}
			continue
		}
		n.RUnlock()

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

		n.RLock()
		n.pushEvent(events.NodeUpdated{
			Address: n.Id,
		})
		n.RUnlock()

		n.Lock()
		n.identified = true
		n.Unlock()
		return
	}
}

func (n *Node) pushEvent(event interface{}) {
	n.pushEventFunc(event)
}

func (n *Node) HasCommand(c commands.ZWaveCommand) bool {
	if n.Device == nil {
		return false
	}

	for _, v := range n.Device.CommandClasses {
		if v.ID == c {
			return true
		}
	}

	return false
}

func (n *Node) HasConfigParam(p int) bool {
	for _, cp := range n.Device.ConfigParams {
		if p == cp.ID {
			return true
		}
	}

	return false
}

func (n *Node) BuildConfigParamData(p int, v ...int) ([]byte, error) {
	var err error
	var data []byte
	for _, cp := range n.Device.ConfigParams {
		if cp.ID == p {
			err = validateConfigParamValues(v, cp)
			if err != nil {
				return data, err
			}
			data, err = buildData(p, v, cp)
			if err != nil {
				return data, err
			}
		}
	}

	return data, nil
}

func (n *Node) IsDeviceClass(generic, specific byte) bool {
	if n.ProtocolInfo == nil {
		return false
	}

	return n.ProtocolInfo.Generic == generic && n.ProtocolInfo.Specific == specific
}

func buildData(p int, v []int, cp database.ConfigParam) ([]byte, error) {
	var data []byte

	data = append(data, byte(p))
	data = append(data, byte(cp.Size))

	switch cp.Size {
	case 1:
		data = append(data, byte(v[0]))
	case 2:
		if len(v) == 1 {
			var first, second uint8 = uint8(v[0] >> 8), uint8(v[0] & 0xff)
			data = append(data, byte(first), byte(second))
		} else {
			data = append(data, byte(v[0]), byte(v[1]))
		}
	case 3:
		if len(v) == 1 {
			var first, second, third uint8 = uint8(v[0] >> 16), uint8(v[0] >> 8), uint8(v[0] & 0xff)
			data = append(data, byte(first), byte(second), byte(third))
		} else {
			data = append(data, byte(v[0]), byte(v[1]), byte(v[2]))
		}
	case 4:
		if len(v) == 1 {
			var first, second, third, fourth uint8 = uint8(v[0] >> 24), uint8(v[0] >> 16), uint8(v[0] >> 8), uint8(v[0] & 0xff)
			data = append(data, byte(first), byte(second), byte(third), byte(fourth))
		} else {
			data = append(data, byte(v[0]), byte(v[1]), byte(v[2]), byte(v[3]))
		}
	}
	return data, nil
}

func validateConfigParamValues(v []int, cp database.ConfigParam) error {
	var err error

	if len(v) == 0 {
		return errors.New("no config value supplied")
	} else if len(v) == 1 {
		err = validateSingleValue(v[0], cp)
		if err != nil {
			return err
		}
	} else {
		err = validateMultipleValues(v, cp)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateSingleValue(v int, cp database.ConfigParam) error {
	if v < (256^cp.Size)-1 && v >= 0 {
		for _, cv := range cp.Values {
			if v >= cv.From && v <= cv.To {
				return nil
			}
		}
	} else {
		return fmt.Errorf("%d is not a valid value for %d", v, cp.ID)
	}

	return fmt.Errorf("%d is not a valid value for %d", v, cp.ID)
}

func validateMultipleValues(values []int, cp database.ConfigParam) error {
	if len(values) != cp.Size {
		return fmt.Errorf("expected %d values only got %d", cp.Size, len(values))
	}

	for _, v := range values {
		if v < 0 || v > 255 {
			return fmt.Errorf("each value must be between 0 and 255")
		}
	}

	return nil
}
