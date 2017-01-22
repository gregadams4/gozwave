package commands
type ThermostatFanModeGetV4 struct {
	node byte
}

func NewThermostatFanModeGetV4() ThermostatFanModeGetV4 {
	return ThermostatFanModeGetV4{}
}

func (c *ThermostatFanModeGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanModeGetV4) Set() error {

	return nil
}

func (c *ThermostatFanModeGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatFanModeV4),
		0x02,
		0x25,
	}
}
type ThermostatFanModeReportV4 struct {
    *report
    Properties1 byte
    data []byte
}

func NewThermostatFanModeReportV4(data []byte) *ThermostatFanModeReportV4 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatFanModeReportV4{
				Properties1: data[0],
        data: data,
    }
}

type ThermostatFanModeSetV4 struct {
	node byte
	Properties1 byte
}

func NewThermostatFanModeSetV4() ThermostatFanModeSetV4 {
	return ThermostatFanModeSetV4{}
}

func (c *ThermostatFanModeSetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanModeSetV4) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *ThermostatFanModeSetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatFanModeV4),
		0x01,
		c.Properties1,
		0x25,
	}
}
type ThermostatFanModeSupportedGetV4 struct {
	node byte
}

func NewThermostatFanModeSupportedGetV4() ThermostatFanModeSupportedGetV4 {
	return ThermostatFanModeSupportedGetV4{}
}

func (c *ThermostatFanModeSupportedGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanModeSupportedGetV4) Set() error {

	return nil
}

func (c *ThermostatFanModeSupportedGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatFanModeV4),
		0x04,
		0x25,
	}
}
type ThermostatFanModeSupportedReportV4 struct {
    *report
    BitMask byte
    data []byte
}

func NewThermostatFanModeSupportedReportV4(data []byte) *ThermostatFanModeSupportedReportV4 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatFanModeSupportedReportV4{
				BitMask: data[0],
        data: data,
    }
}

