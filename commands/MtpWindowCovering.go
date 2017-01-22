package commands
type MoveToPositionGet struct {
	node byte
}

func NewMoveToPositionGet() MoveToPositionGet {
	return MoveToPositionGet{}
}

func (c *MoveToPositionGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MoveToPositionGet) Set() error {

	return nil
}

func (c *MoveToPositionGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MtpWindowCovering),
		0x02,
		0x25,
	}
}
type MoveToPositionReport struct {
    *report
    Value byte
    data []byte
}

func NewMoveToPositionReport(data []byte) *MoveToPositionReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &MoveToPositionReport{
				Value: data[0],
        data: data,
    }
}

type MoveToPositionSet struct {
	node byte
	Value byte
}

func NewMoveToPositionSet() MoveToPositionSet {
	return MoveToPositionSet{}
}

func (c *MoveToPositionSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MoveToPositionSet) Set(Value byte,) error {
	c.Value = Value

	return nil
}

func (c *MoveToPositionSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MtpWindowCovering),
		0x01,
		c.Value,
		0x25,
	}
}
