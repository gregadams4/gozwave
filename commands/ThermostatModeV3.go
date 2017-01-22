package commands
type ThermostatModeGetV3 struct {
	node byte
}

func NewThermostatModeGetV3() ThermostatModeGetV3 {
	return ThermostatModeGetV3{}
}

func (c *ThermostatModeGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatModeGetV3) Set() error {

	return nil
}

func (c *ThermostatModeGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatModeV3),
		0x02,
		0x25,
	}
}
type ThermostatModeReportV3 struct {
    *report
    Level byte
    ManufacturerData byte
    data []byte
}

func NewThermostatModeReportV3(data []byte) *ThermostatModeReportV3 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatModeReportV3{
				Level: data[0],
				ManufacturerData: data[1],
        data: data,
    }
}

type ThermostatModeSetV3 struct {
	node byte
	Level byte
	ManufacturerData byte
}

func NewThermostatModeSetV3() ThermostatModeSetV3 {
	return ThermostatModeSetV3{}
}

func (c *ThermostatModeSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatModeSetV3) Set(Level byte,ManufacturerData byte,) error {
	c.Level = Level
	c.ManufacturerData = ManufacturerData

	return nil
}

func (c *ThermostatModeSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ThermostatModeV3),
		0x01,
		c.Level,
		c.ManufacturerData,
		0x25,
	}
}
type ThermostatModeSupportedGetV3 struct {
	node byte
}

func NewThermostatModeSupportedGetV3() ThermostatModeSupportedGetV3 {
	return ThermostatModeSupportedGetV3{}
}

func (c *ThermostatModeSupportedGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatModeSupportedGetV3) Set() error {

	return nil
}

func (c *ThermostatModeSupportedGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatModeV3),
		0x04,
		0x25,
	}
}
type ThermostatModeSupportedReportV3 struct {
    *report
    BitMask byte
    data []byte
}

func NewThermostatModeSupportedReportV3(data []byte) *ThermostatModeSupportedReportV3 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatModeSupportedReportV3{
				BitMask: data[0],
        data: data,
    }
}

