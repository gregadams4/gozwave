package commands
type HrvStatusGet struct {
	node byte
	StatusParameter byte
}

func NewHrvStatusGet() HrvStatusGet {
	return HrvStatusGet{}
}

func (c *HrvStatusGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HrvStatusGet) Set(StatusParameter byte,) error {
	c.StatusParameter = StatusParameter

	return nil
}

func (c *HrvStatusGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(HrvStatus),
		0x01,
		c.StatusParameter,
		0x25,
	}
}
type HrvStatusReport struct {
    *report
    StatusParameter byte
    Properties1 byte
    Value byte
    data []byte
}

func NewHrvStatusReport(data []byte) *HrvStatusReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &HrvStatusReport{
				StatusParameter: data[0],
				Properties1: data[1],
				Value: data[2],
        data: data,
    }
}

type HrvStatusSupportedGet struct {
	node byte
}

func NewHrvStatusSupportedGet() HrvStatusSupportedGet {
	return HrvStatusSupportedGet{}
}

func (c *HrvStatusSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HrvStatusSupportedGet) Set() error {

	return nil
}

func (c *HrvStatusSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(HrvStatus),
		0x03,
		0x25,
	}
}
type HrvStatusSupportedReport struct {
    *report
    BitMask byte
    data []byte
}

func NewHrvStatusSupportedReport(data []byte) *HrvStatusSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &HrvStatusSupportedReport{
				BitMask: data[0],
        data: data,
    }
}

