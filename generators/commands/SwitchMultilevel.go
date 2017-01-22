package commands
type SwitchMultilevelGet struct {
	node byte
}

func NewSwitchMultilevelGet() SwitchMultilevelGet {
	return SwitchMultilevelGet{}
}

func (c *SwitchMultilevelGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelGet) Set() error {

	return nil
}

func (c *SwitchMultilevelGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchMultilevel),
		0x02,
		0x25,
	}
}
type SwitchMultilevelReport struct {
    *report
    Value byte
    data []byte
}

func NewSwitchMultilevelReport(data []byte) *SwitchMultilevelReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchMultilevelReport{
				Value: data[0],
        data: data,
    }
}

type SwitchMultilevelSet struct {
	node byte
	Value byte
}

func NewSwitchMultilevelSet() SwitchMultilevelSet {
	return SwitchMultilevelSet{}
}

func (c *SwitchMultilevelSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelSet) Set(Value byte,) error {
	c.Value = Value

	return nil
}

func (c *SwitchMultilevelSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SwitchMultilevel),
		0x01,
		c.Value,
		0x25,
	}
}
type SwitchMultilevelStartLevelChange struct {
	node byte
	Level byte
	StartLevel byte
}

func NewSwitchMultilevelStartLevelChange() SwitchMultilevelStartLevelChange {
	return SwitchMultilevelStartLevelChange{}
}

func (c *SwitchMultilevelStartLevelChange) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelStartLevelChange) Set(Level byte,StartLevel byte,) error {
	c.Level = Level
	c.StartLevel = StartLevel

	return nil
}

func (c *SwitchMultilevelStartLevelChange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SwitchMultilevel),
		0x04,
		c.Level,
		c.StartLevel,
		0x25,
	}
}
type SwitchMultilevelStopLevelChange struct {
	node byte
}

func NewSwitchMultilevelStopLevelChange() SwitchMultilevelStopLevelChange {
	return SwitchMultilevelStopLevelChange{}
}

func (c *SwitchMultilevelStopLevelChange) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelStopLevelChange) Set() error {

	return nil
}

func (c *SwitchMultilevelStopLevelChange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchMultilevel),
		0x05,
		0x25,
	}
}
