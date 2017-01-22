package commands
type SwitchAllGet struct {
	node byte
}

func NewSwitchAllGet() SwitchAllGet {
	return SwitchAllGet{}
}

func (c *SwitchAllGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchAllGet) Set() error {

	return nil
}

func (c *SwitchAllGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchAll),
		0x02,
		0x25,
	}
}
type SwitchAllOff struct {
	node byte
}

func NewSwitchAllOff() SwitchAllOff {
	return SwitchAllOff{}
}

func (c *SwitchAllOff) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchAllOff) Set() error {

	return nil
}

func (c *SwitchAllOff) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchAll),
		0x05,
		0x25,
	}
}
type SwitchAllOn struct {
	node byte
}

func NewSwitchAllOn() SwitchAllOn {
	return SwitchAllOn{}
}

func (c *SwitchAllOn) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchAllOn) Set() error {

	return nil
}

func (c *SwitchAllOn) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchAll),
		0x04,
		0x25,
	}
}
type SwitchAllReport struct {
    *report
    Mode byte
    data []byte
}

func NewSwitchAllReport(data []byte) *SwitchAllReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchAllReport{
				Mode: data[0],
        data: data,
    }
}

type SwitchAllSet struct {
	node byte
	Mode byte
}

func NewSwitchAllSet() SwitchAllSet {
	return SwitchAllSet{}
}

func (c *SwitchAllSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchAllSet) Set(Mode byte,) error {
	c.Mode = Mode

	return nil
}

func (c *SwitchAllSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SwitchAll),
		0x01,
		c.Mode,
		0x25,
	}
}
