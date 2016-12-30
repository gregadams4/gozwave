package functions

import "github.com/stampzilla/gozwave/commands"

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

func (self *FuncApplicationCommandHandler) Decode(data []byte) {
	self.Command = commands.ZWaveCommand(data[0])
	self.Class = data[1]

	// logrus.Info("decoding application command handler")
	switch self.Command {
	case commands.Alarm:
		switch self.Class {
		case 0x05: // Report
			self.Data = commands.NewAlarmReport(data[2:])
		}
	case commands.Configuration:
		switch self.Class {
		case 0x06: // Report
			self.Data = commands.NewConfigurationReport(data[2:])
		}
	case commands.ManufacturerSpecific:
		switch self.Class {
		case 0x05: // Report
			self.Data = commands.NewManufacturerSpecificReport(data[2:])
		}
	case commands.MultiInstance:
		switch self.Class {
		case 0x08: // MultiChannelCmd_EndPointReport
			self.Data = commands.NewMultiChannelCmdEndPointReport(data[2:])
		}
	case commands.SensorMultiLevel:
		switch self.Class {
		case 0x05: // Report
			self.Data = commands.NewSensorMultiLevelReport(data[2:])
		}
	case commands.SwitchBinary:
		// logrus.Info("was switch binary")
		switch self.Class {
		case 0x03: // Report
			self.Data = commands.NewSwitchBinaryReport(data[2:])
		}
	case commands.SwitchMultilevel:
		switch self.Class {
		case 0x03: // Report
			self.Data = commands.NewSwitchMultilevelReport(data[2:])
		}
	case commands.WakeUp:
		//	self.Node = data[2];
		self.Data = commands.NewWakeUpReport()
		//default:
		//self.Data = data
	}

	if report, ok := self.Data.(commands.Report); ok && report != nil {
		report.SetNode(self.Node)
	}
}
