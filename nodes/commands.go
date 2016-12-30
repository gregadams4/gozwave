package nodes

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/stampzilla/gozwave/commands"
	"github.com/stampzilla/gozwave/interfaces"
)

func (n *Node) GetConfiguration(parameter int) error {
	var send interfaces.Encodable

	if n.HasCommand(commands.Configuration) {
		if n.HasConfigParam(parameter) {
			cmd := commands.NewConfigurationGet()
			cmd.SetParameter(parameter)
			cmd.SetNode(n.Id)

			send = cmd
		} else {
			return fmt.Errorf("%d does not have the configuration parameter %d", n.Id, parameter)
		}
	} else {
		return fmt.Errorf("%d has no configuration command class", n.Id)
	}

	// n.connection.Write(send)
	t, err := n.connection.WriteAndWaitForReport(send, 5*time.Second, 0x06)
	if err != nil {
		return err
	}

	report := <-t

	if report == nil {
		return errors.New("timed out")
	}
	configReport := report.(*commands.ConfigurationReport)

	log.Printf("report: %v - %v", configReport.Parameter, configReport.Value)

	return nil
}

func (n *Node) On() {
	var send interfaces.Encodable

	switch {
	case n.HasCommand(commands.SwitchBinary):
		cmd := commands.NewSwitchBinary()
		cmd.SetValue(true)
		cmd.SetNode(n.Id)

		send = cmd
	case n.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevel()
		cmd.SetValue(100)
		cmd.SetNode(n.Id)

		send = cmd
	default:
		return
	}

	n.connection.Write(send)
}

func (n *Node) Off() {
	var send interfaces.Encodable

	switch {
	case n.HasCommand(commands.SwitchBinary):
		cmd := commands.NewSwitchBinary()
		cmd.SetValue(false)
		cmd.SetNode(n.Id)

		send = cmd
	case n.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevel()
		cmd.SetValue(0)
		cmd.SetNode(n.Id)

		send = cmd
	default:
		return
	}

	n.connection.Write(send)
}

func (n *Node) Level(value float64) {
	var send interfaces.Encodable

	switch {
	case n.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevel()
		cmd.SetValue(value)
		cmd.SetNode(n.Id)

		send = cmd
	default:
		return
	}

	n.connection.Write(send)
}
