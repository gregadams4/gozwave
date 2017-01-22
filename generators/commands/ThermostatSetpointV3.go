package commands
type ThermostatSetpointGetV3 struct {
	node byte
	Level byte
}

func NewThermostatSetpointGetV3() ThermostatSetpointGetV3 {
	return ThermostatSetpointGetV3{}
}

func (c *ThermostatSetpointGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatSetpointGetV3) Set(Level byte,) error {
	c.Level = Level

	return nil
}

func (c *ThermostatSetpointGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatSetpointV3),
		0x02,
		c.Level,
		0x25,
	}
}
type ThermostatSetpointReportV3 struct {
    *report
    Level byte
    Level2 byte
    Value byte
    data []byte
}

func NewThermostatSetpointReportV3(data []byte) *ThermostatSetpointReportV3 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatSetpointReportV3{
				Level: data[0],
				Level2: data[1],
				Value: data[2],
        data: data,
    }
}

type ThermostatSetpointSetV3 struct {
	node byte
	Level byte
	Level2 byte
	Value byte
}

func NewThermostatSetpointSetV3() ThermostatSetpointSetV3 {
	return ThermostatSetpointSetV3{}
}

func (c *ThermostatSetpointSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatSetpointSetV3) Set(Level byte,Level2 byte,Value byte,) error {
	c.Level = Level
	c.Level2 = Level2
	c.Value = Value

	return nil
}

func (c *ThermostatSetpointSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ThermostatSetpointV3),
		0x01,
		c.Level,
		c.Level2,
		c.Value,
		0x25,
	}
}
type ThermostatSetpointSupportedGetV3 struct {
	node byte
}

func NewThermostatSetpointSupportedGetV3() ThermostatSetpointSupportedGetV3 {
	return ThermostatSetpointSupportedGetV3{}
}

func (c *ThermostatSetpointSupportedGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatSetpointSupportedGetV3) Set() error {

	return nil
}

func (c *ThermostatSetpointSupportedGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatSetpointV3),
		0x04,
		0x25,
	}
}
type ThermostatSetpointSupportedReportV3 struct {
    *report
    BitMask byte
    data []byte
}

func NewThermostatSetpointSupportedReportV3(data []byte) *ThermostatSetpointSupportedReportV3 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatSetpointSupportedReportV3{
				BitMask: data[0],
        data: data,
    }
}

type ThermostatSetpointCapabilitiesGetV3 struct {
	node byte
	Properties1 byte
}

func NewThermostatSetpointCapabilitiesGetV3() ThermostatSetpointCapabilitiesGetV3 {
	return ThermostatSetpointCapabilitiesGetV3{}
}

func (c *ThermostatSetpointCapabilitiesGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatSetpointCapabilitiesGetV3) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *ThermostatSetpointCapabilitiesGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatSetpointV3),
		0x09,
		c.Properties1,
		0x25,
	}
}
type ThermostatSetpointCapabilitiesReportV3 struct {
    *report
    Properties1 byte
    Properties2 byte
    MinValue byte
    Properties3 byte
    MaxValue byte
    data []byte
}

func NewThermostatSetpointCapabilitiesReportV3(data []byte) *ThermostatSetpointCapabilitiesReportV3 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatSetpointCapabilitiesReportV3{
				Properties1: data[0],
				Properties2: data[1],
				MinValue: data[2],
				Properties3: data[3],
				MaxValue: data[4],
        data: data,
    }
}

