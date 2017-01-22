package nodes

import database "github.com/gregadams4/gozwave/database"

type Endpoint struct {
	Id             int
	CommandClasses []*database.CommandClass

	StateBool  map[string]bool
	StateFloat map[string]float64

	node *Node
}

// func (n *Node) Endpoint(id int) *Endpoint {
// 	if n == nil {
// 		logrus.Errorf("Failed to get endpoint from NIL node")
// 		return nil
// 	}

// 	if id > len(n.Endpoints) {
// 		return nil
// 	}

// 	return n.Endpoints[id]
// }

// func (e *Endpoint) Write(msg interfaces.Encodable) {
// 	logrus.Debugf("Send to endpoint %d ", e.Id)
// 	e.node.connection.Write(commands.NewMultiChannelEncap(msg.Encode(), e.Id))
// }

// func (e *Endpoint) On() {
// 	var send interfaces.Encodable

// 	switch {
// 	case e.node.HasCommand(commands.SwitchBinary):
// 		cmd := commands.NewSwitchBinarySet()
// 		cmd.Set(0xff)
// 		cmd.SetNode(e.node.Id)

// 		send = &cmd
// 	case e.node.HasCommand(commands.SwitchMultilevel):
// 		cmd := commands.NewSwitchMultilevelSet()
// 		cmd.Set(0xff)
// 		cmd.SetNode(e.node.Id)

// 		send = &cmd
// 	default:
// 		return
// 	}

// 	e.Write(send)
// }

// func (e *Endpoint) Off() {
// 	var send interfaces.Encodable

// 	switch {
// 	case e.node.HasCommand(commands.SwitchBinary):
// 		cmd := commands.NewSwitchBinarySet()
// 		cmd.Set(0x00)
// 		cmd.SetNode(e.node.Id)

// 		send = &cmd
// 	case e.node.HasCommand(commands.SwitchMultilevel):
// 		cmd := commands.NewSwitchMultilevelSet()
// 		cmd.Set(0x00)
// 		cmd.SetNode(e.node.Id)

// 		send = &cmd
// 	default:
// 		return
// 	}

// 	e.Write(send)
// }

// func (e *Endpoint) Level(value float64) {
// 	var send interfaces.Encodable

// 	switch {
// 	case e.node.HasCommand(commands.SwitchMultilevel):
// 		cmd := commands.NewSwitchMultilevel()
// 		cmd.SetValue(value)
// 		cmd.SetNode(e.node.Id)

// 		send = cmd
// 	default:
// 		return
// 	}

// 	e.Write(send)
// }
