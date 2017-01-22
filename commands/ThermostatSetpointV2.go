package commands
type ThermostatSetpointGetV2 struct {
	node byte
	Level byte
}

func NewThermostatSetpointGetV2() ThermostatSetpointGetV2 {
	return ThermostatSetpointGetV2{}
}

func (c *ThermostatSetpointGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatSetpointGetV2) Set(Level byte,) error {
	c.Level = Level

	return nil
}

func (c *ThermostatSetpointGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatSetpointV2),
		0x02,
		c.Level,
		0x25,
	}
}
type ThermostatSetpointReportV2 struct {
    *report
    Level byte
    Level2 byte
    Value byte
    data []byte
}

func NewThermostatSetpointReportV2(data []byte) *ThermostatSetpointReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatSetpointReportV2{
				Level: data[0],
				Level2: data[1],
				Value: data[2],
        data: data,
    }
}

type ThermostatSetpointSetV2 struct {
	node byte
	Level byte
	Level2 byte
	Value byte
}

func NewThermostatSetpointSetV2() ThermostatSetpointSetV2 {
	return ThermostatSetpointSetV2{}
}

func (c *ThermostatSetpointSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatSetpointSetV2) Set(Level byte,Level2 byte,Value byte,) error {
	c.Level = Level
	c.Level2 = Level2
	c.Value = Value

	return nil
}

func (c *ThermostatSetpointSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ThermostatSetpointV2),
		0x01,
		c.Level,
		c.Level2,
		c.Value,
		0x25,
	}
}
type ThermostatSetpointSupportedGetV2 struct {
	node byte
}

func NewThermostatSetpointSupportedGetV2() ThermostatSetpointSupportedGetV2 {
	return ThermostatSetpointSupportedGetV2{}
}

func (c *ThermostatSetpointSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatSetpointSupportedGetV2) Set() error {

	return nil
}

func (c *ThermostatSetpointSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatSetpointV2),
		0x04,
		0x25,
	}
}
type ThermostatSetpointSupportedReportV2 struct {
    *report
    BitMask byte
    data []byte
}

func NewThermostatSetpointSupportedReportV2(data []byte) *ThermostatSetpointSupportedReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatSetpointSupportedReportV2{
				BitMask: data[0],
        data: data,
    }
}

