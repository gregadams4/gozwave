package commands
type ThermostatFanModeGetV2 struct {
	node byte
}

func NewThermostatFanModeGetV2() ThermostatFanModeGetV2 {
	return ThermostatFanModeGetV2{}
}

func (c *ThermostatFanModeGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanModeGetV2) Set() error {

	return nil
}

func (c *ThermostatFanModeGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatFanModeV2),
		0x02,
		0x25,
	}
}
type ThermostatFanModeReportV2 struct {
    *report
    Level byte
    data []byte
}

func NewThermostatFanModeReportV2(data []byte) *ThermostatFanModeReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatFanModeReportV2{
				Level: data[0],
        data: data,
    }
}

type ThermostatFanModeSetV2 struct {
	node byte
	Level byte
}

func NewThermostatFanModeSetV2() ThermostatFanModeSetV2 {
	return ThermostatFanModeSetV2{}
}

func (c *ThermostatFanModeSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanModeSetV2) Set(Level byte,) error {
	c.Level = Level

	return nil
}

func (c *ThermostatFanModeSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatFanModeV2),
		0x01,
		c.Level,
		0x25,
	}
}
type ThermostatFanModeSupportedGetV2 struct {
	node byte
}

func NewThermostatFanModeSupportedGetV2() ThermostatFanModeSupportedGetV2 {
	return ThermostatFanModeSupportedGetV2{}
}

func (c *ThermostatFanModeSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanModeSupportedGetV2) Set() error {

	return nil
}

func (c *ThermostatFanModeSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatFanModeV2),
		0x04,
		0x25,
	}
}
type ThermostatFanModeSupportedReportV2 struct {
    *report
    BitMask byte
    data []byte
}

func NewThermostatFanModeSupportedReportV2(data []byte) *ThermostatFanModeSupportedReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatFanModeSupportedReportV2{
				BitMask: data[0],
        data: data,
    }
}

