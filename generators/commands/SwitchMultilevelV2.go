package commands
type SwitchMultilevelGetV2 struct {
	node byte
}

func NewSwitchMultilevelGetV2() SwitchMultilevelGetV2 {
	return SwitchMultilevelGetV2{}
}

func (c *SwitchMultilevelGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelGetV2) Set() error {

	return nil
}

func (c *SwitchMultilevelGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchMultilevelV2),
		0x02,
		0x25,
	}
}
type SwitchMultilevelReportV2 struct {
    *report
    Value byte
    data []byte
}

func NewSwitchMultilevelReportV2(data []byte) *SwitchMultilevelReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchMultilevelReportV2{
				Value: data[0],
        data: data,
    }
}

type SwitchMultilevelSetV2 struct {
	node byte
	Value byte
	DimmingDuration byte
}

func NewSwitchMultilevelSetV2() SwitchMultilevelSetV2 {
	return SwitchMultilevelSetV2{}
}

func (c *SwitchMultilevelSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelSetV2) Set(Value byte,DimmingDuration byte,) error {
	c.Value = Value
	c.DimmingDuration = DimmingDuration

	return nil
}

func (c *SwitchMultilevelSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SwitchMultilevelV2),
		0x01,
		c.Value,
		c.DimmingDuration,
		0x25,
	}
}
type SwitchMultilevelStartLevelChangeV2 struct {
	node byte
	Properties1 byte
	StartLevel byte
	DimmingDuration byte
}

func NewSwitchMultilevelStartLevelChangeV2() SwitchMultilevelStartLevelChangeV2 {
	return SwitchMultilevelStartLevelChangeV2{}
}

func (c *SwitchMultilevelStartLevelChangeV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelStartLevelChangeV2) Set(Properties1 byte,StartLevel byte,DimmingDuration byte,) error {
	c.Properties1 = Properties1
	c.StartLevel = StartLevel
	c.DimmingDuration = DimmingDuration

	return nil
}

func (c *SwitchMultilevelStartLevelChangeV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(SwitchMultilevelV2),
		0x04,
		c.Properties1,
		c.StartLevel,
		c.DimmingDuration,
		0x25,
	}
}
type SwitchMultilevelStopLevelChangeV2 struct {
	node byte
}

func NewSwitchMultilevelStopLevelChangeV2() SwitchMultilevelStopLevelChangeV2 {
	return SwitchMultilevelStopLevelChangeV2{}
}

func (c *SwitchMultilevelStopLevelChangeV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelStopLevelChangeV2) Set() error {

	return nil
}

func (c *SwitchMultilevelStopLevelChangeV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchMultilevelV2),
		0x05,
		0x25,
	}
}
