package commands
type ScreenMdGetV2 struct {
	node byte
	NumberofReports byte
	NodeID byte
}

func NewScreenMdGetV2() ScreenMdGetV2 {
	return ScreenMdGetV2{}
}

func (c *ScreenMdGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScreenMdGetV2) Set(NumberofReports byte,NodeID byte,) error {
	c.NumberofReports = NumberofReports
	c.NodeID = NodeID

	return nil
}

func (c *ScreenMdGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScreenMdV2),
		0x01,
		c.NumberofReports,
		c.NodeID,
		0x25,
	}
}
type ScreenMdReportV2 struct {
    *report
    Properties1 byte
    Properties2 byte
    data []byte
}

func NewScreenMdReportV2(data []byte) *ScreenMdReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScreenMdReportV2{
				Properties1: data[0],
				Properties2: data[1],
        data: data,
    }
}

