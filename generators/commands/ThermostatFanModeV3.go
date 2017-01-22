package commands
type ThermostatFanModeGetV3 struct {
	node byte
}

func NewThermostatFanModeGetV3() ThermostatFanModeGetV3 {
	return ThermostatFanModeGetV3{}
}

func (c *ThermostatFanModeGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanModeGetV3) Set() error {

	return nil
}

func (c *ThermostatFanModeGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatFanModeV3),
		0x02,
		0x25,
	}
}
type ThermostatFanModeReportV3 struct {
    *report
    Properties1 byte
    data []byte
}

func NewThermostatFanModeReportV3(data []byte) *ThermostatFanModeReportV3 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatFanModeReportV3{
				Properties1: data[0],
        data: data,
    }
}

type ThermostatFanModeSetV3 struct {
	node byte
	Properties1 byte
}

func NewThermostatFanModeSetV3() ThermostatFanModeSetV3 {
	return ThermostatFanModeSetV3{}
}

func (c *ThermostatFanModeSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanModeSetV3) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *ThermostatFanModeSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatFanModeV3),
		0x01,
		c.Properties1,
		0x25,
	}
}
type ThermostatFanModeSupportedGetV3 struct {
	node byte
}

func NewThermostatFanModeSupportedGetV3() ThermostatFanModeSupportedGetV3 {
	return ThermostatFanModeSupportedGetV3{}
}

func (c *ThermostatFanModeSupportedGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanModeSupportedGetV3) Set() error {

	return nil
}

func (c *ThermostatFanModeSupportedGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatFanModeV3),
		0x04,
		0x25,
	}
}
type ThermostatFanModeSupportedReportV3 struct {
    *report
    BitMask byte
    data []byte
}

func NewThermostatFanModeSupportedReportV3(data []byte) *ThermostatFanModeSupportedReportV3 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatFanModeSupportedReportV3{
				BitMask: data[0],
        data: data,
    }
}

