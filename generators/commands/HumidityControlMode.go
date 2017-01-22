package commands
type HumidityControlModeSet struct {
	node byte
	Properties1 byte
}

func NewHumidityControlModeSet() HumidityControlModeSet {
	return HumidityControlModeSet{}
}

func (c *HumidityControlModeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HumidityControlModeSet) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *HumidityControlModeSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(HumidityControlMode),
		0x01,
		c.Properties1,
		0x25,
	}
}
type HumidityControlModeGet struct {
	node byte
}

func NewHumidityControlModeGet() HumidityControlModeGet {
	return HumidityControlModeGet{}
}

func (c *HumidityControlModeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HumidityControlModeGet) Set() error {

	return nil
}

func (c *HumidityControlModeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(HumidityControlMode),
		0x02,
		0x25,
	}
}
type HumidityControlModeReport struct {
    *report
    Properties1 byte
    data []byte
}

func NewHumidityControlModeReport(data []byte) *HumidityControlModeReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &HumidityControlModeReport{
				Properties1: data[0],
        data: data,
    }
}

type HumidityControlModeSupportedGet struct {
	node byte
}

func NewHumidityControlModeSupportedGet() HumidityControlModeSupportedGet {
	return HumidityControlModeSupportedGet{}
}

func (c *HumidityControlModeSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HumidityControlModeSupportedGet) Set() error {

	return nil
}

func (c *HumidityControlModeSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(HumidityControlMode),
		0x04,
		0x25,
	}
}
type HumidityControlModeSupportedReport struct {
    *report
    BitMask byte
    data []byte
}

func NewHumidityControlModeSupportedReport(data []byte) *HumidityControlModeSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &HumidityControlModeSupportedReport{
				BitMask: data[0],
        data: data,
    }
}

