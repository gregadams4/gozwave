package commands
type AvContentSearchMdGet struct {
	node byte
}

func NewAvContentSearchMdGet() AvContentSearchMdGet {
	return AvContentSearchMdGet{}
}

func (c *AvContentSearchMdGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AvContentSearchMdGet) Set() error {

	return nil
}

func (c *AvContentSearchMdGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(AvContentSearchMd),
		0x01,
		0x25,
	}
}
type AvContentSearchMdReport struct {
    *report
    data []byte
}

func NewAvContentSearchMdReport(data []byte) *AvContentSearchMdReport {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &AvContentSearchMdReport{
        data: data,
    }
}

