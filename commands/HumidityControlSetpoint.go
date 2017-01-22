package commands
type HumidityControlSetpointSet struct {
	node byte
	Properties1 byte
	Properties2 byte
	Value byte
}

func NewHumidityControlSetpointSet() HumidityControlSetpointSet {
	return HumidityControlSetpointSet{}
}

func (c *HumidityControlSetpointSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HumidityControlSetpointSet) Set(Properties1 byte,Properties2 byte,Value byte,) error {
	c.Properties1 = Properties1
	c.Properties2 = Properties2
	c.Value = Value

	return nil
}

func (c *HumidityControlSetpointSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(HumidityControlSetpoint),
		0x01,
		c.Properties1,
		c.Properties2,
		c.Value,
		0x25,
	}
}
type HumidityControlSetpointGet struct {
	node byte
	Properties1 byte
}

func NewHumidityControlSetpointGet() HumidityControlSetpointGet {
	return HumidityControlSetpointGet{}
}

func (c *HumidityControlSetpointGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HumidityControlSetpointGet) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *HumidityControlSetpointGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(HumidityControlSetpoint),
		0x02,
		c.Properties1,
		0x25,
	}
}
type HumidityControlSetpointReport struct {
    *report
    Properties1 byte
    Properties2 byte
    Value byte
    data []byte
}

func NewHumidityControlSetpointReport(data []byte) *HumidityControlSetpointReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &HumidityControlSetpointReport{
				Properties1: data[0],
				Properties2: data[1],
				Value: data[2],
        data: data,
    }
}

type HumidityControlSetpointSupportedGet struct {
	node byte
}

func NewHumidityControlSetpointSupportedGet() HumidityControlSetpointSupportedGet {
	return HumidityControlSetpointSupportedGet{}
}

func (c *HumidityControlSetpointSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HumidityControlSetpointSupportedGet) Set() error {

	return nil
}

func (c *HumidityControlSetpointSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(HumidityControlSetpoint),
		0x04,
		0x25,
	}
}
type HumidityControlSetpointSupportedReport struct {
    *report
    BitMask byte
    data []byte
}

func NewHumidityControlSetpointSupportedReport(data []byte) *HumidityControlSetpointSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &HumidityControlSetpointSupportedReport{
				BitMask: data[0],
        data: data,
    }
}

type HumidityControlSetpointScaleSupportedGet struct {
	node byte
	Properties1 byte
}

func NewHumidityControlSetpointScaleSupportedGet() HumidityControlSetpointScaleSupportedGet {
	return HumidityControlSetpointScaleSupportedGet{}
}

func (c *HumidityControlSetpointScaleSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HumidityControlSetpointScaleSupportedGet) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *HumidityControlSetpointScaleSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(HumidityControlSetpoint),
		0x06,
		c.Properties1,
		0x25,
	}
}
type HumidityControlSetpointScaleSupportedReport struct {
    *report
    Properties1 byte
    data []byte
}

func NewHumidityControlSetpointScaleSupportedReport(data []byte) *HumidityControlSetpointScaleSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &HumidityControlSetpointScaleSupportedReport{
				Properties1: data[0],
        data: data,
    }
}

type HumidityControlSetpointCapabilitiesGet struct {
	node byte
	Properties1 byte
}

func NewHumidityControlSetpointCapabilitiesGet() HumidityControlSetpointCapabilitiesGet {
	return HumidityControlSetpointCapabilitiesGet{}
}

func (c *HumidityControlSetpointCapabilitiesGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HumidityControlSetpointCapabilitiesGet) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *HumidityControlSetpointCapabilitiesGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(HumidityControlSetpoint),
		0x08,
		c.Properties1,
		0x25,
	}
}
type HumidityControlSetpointCapabilitiesReport struct {
    *report
    Properties1 byte
    Properties2 byte
    MinimumValue byte
    Properties3 byte
    MaximumValue byte
    data []byte
}

func NewHumidityControlSetpointCapabilitiesReport(data []byte) *HumidityControlSetpointCapabilitiesReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &HumidityControlSetpointCapabilitiesReport{
				Properties1: data[0],
				Properties2: data[1],
				MinimumValue: data[2],
				Properties3: data[3],
				MaximumValue: data[4],
        data: data,
    }
}

