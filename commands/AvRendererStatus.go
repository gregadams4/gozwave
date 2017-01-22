package commands
type AvRendererStatusGet struct {
	node byte
}

func NewAvRendererStatusGet() AvRendererStatusGet {
	return AvRendererStatusGet{}
}

func (c *AvRendererStatusGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AvRendererStatusGet) Set() error {

	return nil
}

func (c *AvRendererStatusGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(AvRendererStatus),
		0x01,
		0x25,
	}
}
type AvRendererStatusReport struct {
    *report
    data []byte
}

func NewAvRendererStatusReport(data []byte) *AvRendererStatusReport {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &AvRendererStatusReport{
        data: data,
    }
}

