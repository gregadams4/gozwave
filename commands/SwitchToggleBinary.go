package commands
type SwitchToggleBinarySet struct {
	node byte
}

func NewSwitchToggleBinarySet() SwitchToggleBinarySet {
	return SwitchToggleBinarySet{}
}

func (c *SwitchToggleBinarySet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchToggleBinarySet) Set() error {

	return nil
}

func (c *SwitchToggleBinarySet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchToggleBinary),
		0x01,
		0x25,
	}
}
type SwitchToggleBinaryGet struct {
	node byte
}

func NewSwitchToggleBinaryGet() SwitchToggleBinaryGet {
	return SwitchToggleBinaryGet{}
}

func (c *SwitchToggleBinaryGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchToggleBinaryGet) Set() error {

	return nil
}

func (c *SwitchToggleBinaryGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchToggleBinary),
		0x02,
		0x25,
	}
}
type SwitchToggleBinaryReport struct {
    *report
    Value byte
    data []byte
}

func NewSwitchToggleBinaryReport(data []byte) *SwitchToggleBinaryReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchToggleBinaryReport{
				Value: data[0],
        data: data,
    }
}

