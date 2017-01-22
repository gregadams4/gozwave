package commands
type SwitchBinaryGet struct {
	node byte
}

func NewSwitchBinaryGet() SwitchBinaryGet {
	return SwitchBinaryGet{}
}

func (c *SwitchBinaryGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchBinaryGet) Set() error {

	return nil
}

func (c *SwitchBinaryGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchBinary),
		0x02,
		0x25,
	}
}
type SwitchBinaryReport struct {
    *report
    Value byte
    data []byte
}

func NewSwitchBinaryReport(data []byte) *SwitchBinaryReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchBinaryReport{
				Value: data[0],
        data: data,
    }
}

type SwitchBinarySet struct {
	node byte
	SwitchValue byte
}

func NewSwitchBinarySet() SwitchBinarySet {
	return SwitchBinarySet{}
}

func (c *SwitchBinarySet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchBinarySet) Set(SwitchValue byte,) error {
	c.SwitchValue = SwitchValue

	return nil
}

func (c *SwitchBinarySet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SwitchBinary),
		0x01,
		c.SwitchValue,
		0x25,
	}
}
