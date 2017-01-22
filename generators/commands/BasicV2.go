package commands
type BasicGetV2 struct {
	node byte
}

func NewBasicGetV2() BasicGetV2 {
	return BasicGetV2{}
}

func (c *BasicGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *BasicGetV2) Set() error {

	return nil
}

func (c *BasicGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(BasicV2),
		0x02,
		0x25,
	}
}
type BasicReportV2 struct {
    *report
    CurrentValue byte
    TargetValue byte
    Duration byte
    data []byte
}

func NewBasicReportV2(data []byte) *BasicReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &BasicReportV2{
				CurrentValue: data[0],
				TargetValue: data[1],
				Duration: data[2],
        data: data,
    }
}

type BasicSetV2 struct {
	node byte
	Value byte
}

func NewBasicSetV2() BasicSetV2 {
	return BasicSetV2{}
}

func (c *BasicSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *BasicSetV2) Set(Value byte,) error {
	c.Value = Value

	return nil
}

func (c *BasicSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(BasicV2),
		0x01,
		c.Value,
		0x25,
	}
}
