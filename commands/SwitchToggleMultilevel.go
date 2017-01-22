package commands
type SwitchToggleMultilevelSet struct {
	node byte
}

func NewSwitchToggleMultilevelSet() SwitchToggleMultilevelSet {
	return SwitchToggleMultilevelSet{}
}

func (c *SwitchToggleMultilevelSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchToggleMultilevelSet) Set() error {

	return nil
}

func (c *SwitchToggleMultilevelSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchToggleMultilevel),
		0x01,
		0x25,
	}
}
type SwitchToggleMultilevelGet struct {
	node byte
}

func NewSwitchToggleMultilevelGet() SwitchToggleMultilevelGet {
	return SwitchToggleMultilevelGet{}
}

func (c *SwitchToggleMultilevelGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchToggleMultilevelGet) Set() error {

	return nil
}

func (c *SwitchToggleMultilevelGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchToggleMultilevel),
		0x02,
		0x25,
	}
}
type SwitchToggleMultilevelReport struct {
    *report
    Value byte
    data []byte
}

func NewSwitchToggleMultilevelReport(data []byte) *SwitchToggleMultilevelReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchToggleMultilevelReport{
				Value: data[0],
        data: data,
    }
}

type SwitchToggleMultilevelStartLevelChange struct {
	node byte
	Level byte
	StartLevel byte
}

func NewSwitchToggleMultilevelStartLevelChange() SwitchToggleMultilevelStartLevelChange {
	return SwitchToggleMultilevelStartLevelChange{}
}

func (c *SwitchToggleMultilevelStartLevelChange) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchToggleMultilevelStartLevelChange) Set(Level byte,StartLevel byte,) error {
	c.Level = Level
	c.StartLevel = StartLevel

	return nil
}

func (c *SwitchToggleMultilevelStartLevelChange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SwitchToggleMultilevel),
		0x04,
		c.Level,
		c.StartLevel,
		0x25,
	}
}
type SwitchToggleMultilevelStopLevelChange struct {
	node byte
}

func NewSwitchToggleMultilevelStopLevelChange() SwitchToggleMultilevelStopLevelChange {
	return SwitchToggleMultilevelStopLevelChange{}
}

func (c *SwitchToggleMultilevelStopLevelChange) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchToggleMultilevelStopLevelChange) Set() error {

	return nil
}

func (c *SwitchToggleMultilevelStopLevelChange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchToggleMultilevel),
		0x05,
		0x25,
	}
}
