package commands
type ThermostatFanModeGet struct {
	node byte
}

func NewThermostatFanModeGet() ThermostatFanModeGet {
	return ThermostatFanModeGet{}
}

func (c *ThermostatFanModeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanModeGet) Set() error {

	return nil
}

func (c *ThermostatFanModeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatFanMode),
		0x02,
		0x25,
	}
}
type ThermostatFanModeReport struct {
    *report
    Level byte
    data []byte
}

func NewThermostatFanModeReport(data []byte) *ThermostatFanModeReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatFanModeReport{
				Level: data[0],
        data: data,
    }
}

type ThermostatFanModeSet struct {
	node byte
	Level byte
}

func NewThermostatFanModeSet() ThermostatFanModeSet {
	return ThermostatFanModeSet{}
}

func (c *ThermostatFanModeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanModeSet) Set(Level byte,) error {
	c.Level = Level

	return nil
}

func (c *ThermostatFanModeSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatFanMode),
		0x01,
		c.Level,
		0x25,
	}
}
type ThermostatFanModeSupportedGet struct {
	node byte
}

func NewThermostatFanModeSupportedGet() ThermostatFanModeSupportedGet {
	return ThermostatFanModeSupportedGet{}
}

func (c *ThermostatFanModeSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanModeSupportedGet) Set() error {

	return nil
}

func (c *ThermostatFanModeSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatFanMode),
		0x04,
		0x25,
	}
}
type ThermostatFanModeSupportedReport struct {
    *report
    BitMask byte
    data []byte
}

func NewThermostatFanModeSupportedReport(data []byte) *ThermostatFanModeSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatFanModeSupportedReport{
				BitMask: data[0],
        data: data,
    }
}

