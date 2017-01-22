package commands
type ThermostatSetpointGet struct {
	node byte
	Level byte
}

func NewThermostatSetpointGet() ThermostatSetpointGet {
	return ThermostatSetpointGet{}
}

func (c *ThermostatSetpointGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatSetpointGet) Set(Level byte,) error {
	c.Level = Level

	return nil
}

func (c *ThermostatSetpointGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatSetpoint),
		0x02,
		c.Level,
		0x25,
	}
}
type ThermostatSetpointReport struct {
    *report
    Level byte
    Level2 byte
    Value byte
    data []byte
}

func NewThermostatSetpointReport(data []byte) *ThermostatSetpointReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatSetpointReport{
				Level: data[0],
				Level2: data[1],
				Value: data[2],
        data: data,
    }
}

type ThermostatSetpointSet struct {
	node byte
	Level byte
	Level2 byte
	Value byte
}

func NewThermostatSetpointSet() ThermostatSetpointSet {
	return ThermostatSetpointSet{}
}

func (c *ThermostatSetpointSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatSetpointSet) Set(Level byte,Level2 byte,Value byte,) error {
	c.Level = Level
	c.Level2 = Level2
	c.Value = Value

	return nil
}

func (c *ThermostatSetpointSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ThermostatSetpoint),
		0x01,
		c.Level,
		c.Level2,
		c.Value,
		0x25,
	}
}
type ThermostatSetpointSupportedGet struct {
	node byte
}

func NewThermostatSetpointSupportedGet() ThermostatSetpointSupportedGet {
	return ThermostatSetpointSupportedGet{}
}

func (c *ThermostatSetpointSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatSetpointSupportedGet) Set() error {

	return nil
}

func (c *ThermostatSetpointSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatSetpoint),
		0x04,
		0x25,
	}
}
type ThermostatSetpointSupportedReport struct {
    *report
    BitMask byte
    data []byte
}

func NewThermostatSetpointSupportedReport(data []byte) *ThermostatSetpointSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatSetpointSupportedReport{
				BitMask: data[0],
        data: data,
    }
}

