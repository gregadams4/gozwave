package nodes

import (
	"log"
	"time"

	"github.com/gregadams4/gozwave/commands"
	"github.com/gregadams4/gozwave/interfaces"
)

// func (n *Node) SetConfiguration(parameter int, value ...int) error {
// 	var send interfaces.Encodable

// 	if n.HasCommand(commands.Configuration) {
// 		if data, err := n.BuildConfigParamData(parameter, value...); err == nil {
// 			logrus.Infof("%d - setting %d to %d", n.Id, parameter, value)
// 			logrus.Infof("bytes - %v", data)
// 			cmd := commands.NewConfigurationSet()

// 			// cmd.SetParameter(parameter)
// 			// cmd.SetValue(value)
// 			cmd.SetData(data)
// 			cmd.SetNode(n.Id)

// 			send = cmd
// 		} else {
// 			return err
// 		}
// 	} else {
// 		return fmt.Errorf("%d has no configuration command class", n.Id)
// 	}

// 	n.connection.Write(send)

// 	return nil
// }

// func (n *Node) GetConfiguration(parameter int) ([]byte, error) {
// 	var send interfaces.Encodable

// 	if n.HasCommand(commands.Configuration) {
// 		if n.HasConfigParam(parameter) {
// 			cmd := commands.NewConfigurationGet()
// 			cmd.SetParameter(parameter)
// 			cmd.SetNode(n.Id)

// 			send = cmd
// 		} else {
// 			return []byte{}, fmt.Errorf("%d does not have the configuration parameter %d", n.Id, parameter)
// 		}
// 	} else {
// 		return []byte{}, fmt.Errorf("%d has no configuration command class", n.Id)
// 	}

// 	// n.connection.Write(send)
// 	t, err := n.connection.WriteAndWaitForReport(send, 5*time.Second, 0x06)
// 	if err != nil {
// 		return nil, err
// 	}

// 	report := <-t

// 	if report == nil {
// 		return []byte{}, errors.New("timed out")
// 	}
// 	configReport := report.(*commands.ConfigurationReport)

// 	// log.Printf("report: %v - %v", configReport.Parameter, configReport.Value)

// 	return configReport.Value, nil
// }

func (n *Node) Color() {
	var send interfaces.Encodable

	if n.HasCommand(commands.SwitchColor) {
		cmd := commands.NewSwitchColorSetV2()
		cmd.Set(0x01, 0x02, []commands.SwitchColorSetV2vg1{
			commands.SwitchColorSetV2vg1{
				ColorComponentID: 0x04,
				Value:            0x00,
			},
		})
		cmd.SetNode(n.Id)

		send = &cmd
	}
	// switch {
	// case n.HasCommand(commands.SwitchBinary):
	// 	cmd := commands.NewSwitchBinarySet()
	// 	cmd.Set(0xff)
	// 	cmd.SetNode(n.Id)
	//
	// 	send = &cmd
	// case n.HasCommand(commands.SwitchMultilevel):
	// 	cmd := commands.NewSwitchMultilevelSet()
	// 	cmd.Set(0xff)
	// 	cmd.SetNode(n.Id)
	//
	// 	send = &cmd
	// default:
	// 	return
	// }

	n.connection.Write(send)
}

func (n *Node) GetColor() {
	var send interfaces.Encodable
	if n.HasCommand(commands.SwitchColor) {
		cmd := commands.NewSwitchColorGetV2()
		cmd.Set(0x04)
		cmd.SetNode(n.Id)

		send = &cmd
	}

	t, err := n.connection.WriteAndWaitForReport(send, 5*time.Second, 0x04)
	if err != nil {
		log.Println(err)
	}

	report := <-t

	if report == nil {
		log.Println(err)
	}

	log.Printf("%v", report)
	//  := report.(*commands.ConfigurationReport)

	// log.Printf("report: %v - %v", configReport.Parameter, configReport.Value)

	// return configReport.Value, nil
}

func (n *Node) On() {
	var send interfaces.Encodable

	switch {
	case n.HasCommand(commands.SwitchBinary):
		cmd := commands.NewSwitchBinarySet()
		cmd.Set(0xff)
		cmd.SetNode(n.Id)

		send = &cmd
	case n.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevelSet()
		cmd.Set(0xff)
		cmd.SetNode(n.Id)

		send = &cmd
	default:
		return
	}

	n.connection.Write(send)
}

func (n *Node) Off() {
	var send interfaces.Encodable

	switch {
	case n.HasCommand(commands.SwitchBinary):
		cmd := commands.NewSwitchBinarySet()
		cmd.Set(0x00)
		cmd.SetNode(n.Id)

		send = &cmd
	case n.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevelSet()
		cmd.Set(0x00)
		cmd.SetNode(n.Id)

		send = &cmd
	default:
		return
	}

	n.connection.Write(send)
}

// func (n *Node) Level(value float64) {
// 	var send interfaces.Encodable

// 	switch {
// 	case n.HasCommand(commands.SwitchMultilevel):
// 		cmd := commands.NewSwitchMultilevel()
// 		cmd.SetValue(value)
// 		cmd.SetNode(n.Id)

// 		send = cmd
// 	default:
// 		return
// 	}

// 	n.connection.Write(send)
// }
