package commands
type ThermostatModeGetV2 struct {
	node byte
}

func NewThermostatModeGetV2() ThermostatModeGetV2 {
	return ThermostatModeGetV2{}
}

func (c *ThermostatModeGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatModeGetV2) Set() error {

	return nil
}

func (c *ThermostatModeGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatModeV2),
		0x02,
		0x25,
	}
}
type ThermostatModeReportV2 struct {
    *report
    Level byte
    data []byte
}

func NewThermostatModeReportV2(data []byte) *ThermostatModeReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatModeReportV2{
				Level: data[0],
        data: data,
    }
}

type ThermostatModeSetV2 struct {
	node byte
	Level byte
}

func NewThermostatModeSetV2() ThermostatModeSetV2 {
	return ThermostatModeSetV2{}
}

func (c *ThermostatModeSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatModeSetV2) Set(Level byte,) error {
	c.Level = Level

	return nil
}

func (c *ThermostatModeSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatModeV2),
		0x01,
		c.Level,
		0x25,
	}
}
type ThermostatModeSupportedGetV2 struct {
	node byte
}

func NewThermostatModeSupportedGetV2() ThermostatModeSupportedGetV2 {
	return ThermostatModeSupportedGetV2{}
}

func (c *ThermostatModeSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatModeSupportedGetV2) Set() error {

	return nil
}

func (c *ThermostatModeSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatModeV2),
		0x04,
		0x25,
	}
}
type ThermostatModeSupportedReportV2 struct {
    *report
    BitMask byte
    data []byte
}

func NewThermostatModeSupportedReportV2(data []byte) *ThermostatModeSupportedReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatModeSupportedReportV2{
				BitMask: data[0],
        data: data,
    }
}

