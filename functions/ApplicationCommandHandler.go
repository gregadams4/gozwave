package functions

import (
	"log"

	"github.com/Sirupsen/logrus"
	"github.com/gregadams4/gozwave/commands"
)

type FuncApplicationCommandHandler struct {
	Command commands.ZWaveCommand
	Class   byte // Command Class
	Node    byte

	Data commands.Report
}

func NewApplicationCommandHandler(data []byte) *FuncApplicationCommandHandler {
	f := &FuncApplicationCommandHandler{}
	f.Decode(data)

	return f
}

func (a *FuncApplicationCommandHandler) Decode(data []byte) {
	a.Command = commands.ZWaveCommand(data[0])
	a.Class = data[1]

	log.Println("here")
	logrus.Infof("%v", a.Command)
	switch a.Command {
	case commands.Alarm:
		switch a.Class {
		case 0x05: // Report
			a.Data = commands.NewAlarmReport(data[2:])
		}
	case commands.Configuration:
		switch a.Class {
		case 0x06: // Report
			a.Data = commands.NewConfigurationReport(data[2:])
		}
	case commands.ManufacturerSpecific:
		switch a.Class {
		case 0x05: // Report
			a.Data = commands.NewManufacturerSpecificReport(data[2:])
		}
	case commands.MultiInstance:
		switch a.Class {
		case 0x08: // MultiChannelCmd_EndPointReport
			a.Data = commands.NewMultiInstanceReport(data[2:])
			// a.Data = commands.NewMultiChannelCmdEndPointReport(data[2:])
		}
	case commands.SensorMultilevel:
		switch a.Class {
		case 0x05: // Report
			a.Data = commands.NewSensorMultilevelReport(data[2:])
		}
	case commands.SwitchBinary:
		switch a.Class {
		case 0x03: // Report
			a.Data = commands.NewSwitchBinaryReport(data[2:])
		}
	case commands.SwitchMultilevel:
		switch a.Class {
		case 0x03: // Report
			a.Data = commands.NewSwitchMultilevelReport(data[2:])
		}
	case commands.SwitchColorV2:
		switch a.Class {
		case 0x04: // Report
			a.Data = commands.NewSwitchColorReportV2(data[2:])
		}
	case commands.WakeUp:
		//	a.Node = data[2];
		a.Data = commands.NewWakeUpNotification(data[2:])
		//default:
		//a.Data = data
	}

	if report, ok := a.Data.(commands.Report); ok && report != nil {
		report.SetNode(a.Node)
	}
}
