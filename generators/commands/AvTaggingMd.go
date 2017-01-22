package commands
type AvTaggingMdGet struct {
	node byte
}

func NewAvTaggingMdGet() AvTaggingMdGet {
	return AvTaggingMdGet{}
}

func (c *AvTaggingMdGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AvTaggingMdGet) Set() error {

	return nil
}

func (c *AvTaggingMdGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(AvTaggingMd),
		0x01,
		0x25,
	}
}
type AvTaggingMdReport struct {
    *report
    data []byte
}

func NewAvTaggingMdReport(data []byte) *AvTaggingMdReport {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &AvTaggingMdReport{
        data: data,
    }
}

