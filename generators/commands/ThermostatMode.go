package commands
type ThermostatModeGet struct {
	node byte
}

func NewThermostatModeGet() ThermostatModeGet {
	return ThermostatModeGet{}
}

func (c *ThermostatModeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatModeGet) Set() error {

	return nil
}

func (c *ThermostatModeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatMode),
		0x02,
		0x25,
	}
}
type ThermostatModeReport struct {
    *report
    Level byte
    data []byte
}

func NewThermostatModeReport(data []byte) *ThermostatModeReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatModeReport{
				Level: data[0],
        data: data,
    }
}

type ThermostatModeSet struct {
	node byte
	Level byte
}

func NewThermostatModeSet() ThermostatModeSet {
	return ThermostatModeSet{}
}

func (c *ThermostatModeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatModeSet) Set(Level byte,) error {
	c.Level = Level

	return nil
}

func (c *ThermostatModeSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatMode),
		0x01,
		c.Level,
		0x25,
	}
}
type ThermostatModeSupportedGet struct {
	node byte
}

func NewThermostatModeSupportedGet() ThermostatModeSupportedGet {
	return ThermostatModeSupportedGet{}
}

func (c *ThermostatModeSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatModeSupportedGet) Set() error {

	return nil
}

func (c *ThermostatModeSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatMode),
		0x04,
		0x25,
	}
}
type ThermostatModeSupportedReport struct {
    *report
    BitMask byte
    data []byte
}

func NewThermostatModeSupportedReport(data []byte) *ThermostatModeSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatModeSupportedReport{
				BitMask: data[0],
        data: data,
    }
}

