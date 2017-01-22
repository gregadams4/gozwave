package commands
type ScreenMdGet struct {
	node byte
	NumberofReports byte
	NodeID byte
}

func NewScreenMdGet() ScreenMdGet {
	return ScreenMdGet{}
}

func (c *ScreenMdGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScreenMdGet) Set(NumberofReports byte,NodeID byte,) error {
	c.NumberofReports = NumberofReports
	c.NodeID = NodeID

	return nil
}

func (c *ScreenMdGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScreenMd),
		0x01,
		c.NumberofReports,
		c.NodeID,
		0x25,
	}
}
type ScreenMdReport struct {
    *report
    Properties1 byte
    data []byte
}

func NewScreenMdReport(data []byte) *ScreenMdReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScreenMdReport{
				Properties1: data[0],
        data: data,
    }
}

