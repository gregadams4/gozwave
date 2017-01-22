package commands
type CommandClassSecurityPanelZoneSensorInstalledReport struct {
    *report
    Zonenumber byte
    NumberofSensors byte
    data []byte
}

func NewCommandClassSecurityPanelZoneSensorInstalledReport(data []byte) *CommandClassSecurityPanelZoneSensorInstalledReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &CommandClassSecurityPanelZoneSensorInstalledReport{
				Zonenumber: data[0],
				NumberofSensors: data[1],
        data: data,
    }
}

type SecurityPanelZoneSensorTypeGet struct {
	node byte
	ZoneNumber byte
	SensorNumber byte
}

func NewSecurityPanelZoneSensorTypeGet() SecurityPanelZoneSensorTypeGet {
	return SecurityPanelZoneSensorTypeGet{}
}

func (c *SecurityPanelZoneSensorTypeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityPanelZoneSensorTypeGet) Set(ZoneNumber byte,SensorNumber byte,) error {
	c.ZoneNumber = ZoneNumber
	c.SensorNumber = SensorNumber

	return nil
}

func (c *SecurityPanelZoneSensorTypeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SecurityPanelZoneSensor),
		0x03,
		c.ZoneNumber,
		c.SensorNumber,
		0x25,
	}
}
type SecurityPanelZoneSensorTypeReport struct {
    *report
    ZoneNumber byte
    SensorNumber byte
    ZWaveAlarmType byte
    data []byte
}

func NewSecurityPanelZoneSensorTypeReport(data []byte) *SecurityPanelZoneSensorTypeReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SecurityPanelZoneSensorTypeReport{
				ZoneNumber: data[0],
				SensorNumber: data[1],
				ZWaveAlarmType: data[2],
        data: data,
    }
}

type SecurityPanelZoneSensorInstalledGet struct {
	node byte
	Zonenumber byte
}

func NewSecurityPanelZoneSensorInstalledGet() SecurityPanelZoneSensorInstalledGet {
	return SecurityPanelZoneSensorInstalledGet{}
}

func (c *SecurityPanelZoneSensorInstalledGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityPanelZoneSensorInstalledGet) Set(Zonenumber byte,) error {
	c.Zonenumber = Zonenumber

	return nil
}

func (c *SecurityPanelZoneSensorInstalledGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SecurityPanelZoneSensor),
		0x01,
		c.Zonenumber,
		0x25,
	}
}
type SecurityPanelZoneSensorStateGet struct {
	node byte
	ZoneNumber byte
	SensorNumber byte
}

func NewSecurityPanelZoneSensorStateGet() SecurityPanelZoneSensorStateGet {
	return SecurityPanelZoneSensorStateGet{}
}

func (c *SecurityPanelZoneSensorStateGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityPanelZoneSensorStateGet) Set(ZoneNumber byte,SensorNumber byte,) error {
	c.ZoneNumber = ZoneNumber
	c.SensorNumber = SensorNumber

	return nil
}

func (c *SecurityPanelZoneSensorStateGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SecurityPanelZoneSensor),
		0x05,
		c.ZoneNumber,
		c.SensorNumber,
		0x25,
	}
}
type SecurityPanelZoneSensorStateReport struct {
    *report
    ZoneNumber byte
    SensorNumber byte
    ZWaveAlarmType byte
    ZWaveAlarmEvent byte
    EventParameters byte
    data []byte
}

func NewSecurityPanelZoneSensorStateReport(data []byte) *SecurityPanelZoneSensorStateReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &SecurityPanelZoneSensorStateReport{
				ZoneNumber: data[0],
				SensorNumber: data[1],
				ZWaveAlarmType: data[2],
				ZWaveAlarmEvent: data[3],
				EventParameters: data[4],
        data: data,
    }
}

