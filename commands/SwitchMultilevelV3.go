package commands
type SwitchMultilevelGetV3 struct {
	node byte
}

func NewSwitchMultilevelGetV3() SwitchMultilevelGetV3 {
	return SwitchMultilevelGetV3{}
}

func (c *SwitchMultilevelGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelGetV3) Set() error {

	return nil
}

func (c *SwitchMultilevelGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchMultilevelV3),
		0x02,
		0x25,
	}
}
type SwitchMultilevelReportV3 struct {
    *report
    Value byte
    data []byte
}

func NewSwitchMultilevelReportV3(data []byte) *SwitchMultilevelReportV3 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchMultilevelReportV3{
				Value: data[0],
        data: data,
    }
}

type SwitchMultilevelSetV3 struct {
	node byte
	Value byte
	DimmingDuration byte
}

func NewSwitchMultilevelSetV3() SwitchMultilevelSetV3 {
	return SwitchMultilevelSetV3{}
}

func (c *SwitchMultilevelSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelSetV3) Set(Value byte,DimmingDuration byte,) error {
	c.Value = Value
	c.DimmingDuration = DimmingDuration

	return nil
}

func (c *SwitchMultilevelSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SwitchMultilevelV3),
		0x01,
		c.Value,
		c.DimmingDuration,
		0x25,
	}
}
type SwitchMultilevelStartLevelChangeV3 struct {
	node byte
	Properties1 byte
	StartLevel byte
	DimmingDuration byte
	StepSize byte
}

func NewSwitchMultilevelStartLevelChangeV3() SwitchMultilevelStartLevelChangeV3 {
	return SwitchMultilevelStartLevelChangeV3{}
}

func (c *SwitchMultilevelStartLevelChangeV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelStartLevelChangeV3) Set(Properties1 byte,StartLevel byte,DimmingDuration byte,StepSize byte,) error {
	c.Properties1 = Properties1
	c.StartLevel = StartLevel
	c.DimmingDuration = DimmingDuration
	c.StepSize = StepSize

	return nil
}

func (c *SwitchMultilevelStartLevelChangeV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(SwitchMultilevelV3),
		0x04,
		c.Properties1,
		c.StartLevel,
		c.DimmingDuration,
		c.StepSize,
		0x25,
	}
}
type SwitchMultilevelStopLevelChangeV3 struct {
	node byte
}

func NewSwitchMultilevelStopLevelChangeV3() SwitchMultilevelStopLevelChangeV3 {
	return SwitchMultilevelStopLevelChangeV3{}
}

func (c *SwitchMultilevelStopLevelChangeV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelStopLevelChangeV3) Set() error {

	return nil
}

func (c *SwitchMultilevelStopLevelChangeV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchMultilevelV3),
		0x05,
		0x25,
	}
}
type SwitchMultilevelSupportedGetV3 struct {
	node byte
}

func NewSwitchMultilevelSupportedGetV3() SwitchMultilevelSupportedGetV3 {
	return SwitchMultilevelSupportedGetV3{}
}

func (c *SwitchMultilevelSupportedGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelSupportedGetV3) Set() error {

	return nil
}

func (c *SwitchMultilevelSupportedGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchMultilevelV3),
		0x06,
		0x25,
	}
}
type SwitchMultilevelSupportedReportV3 struct {
    *report
    Properties1 byte
    Properties2 byte
    data []byte
}

func NewSwitchMultilevelSupportedReportV3(data []byte) *SwitchMultilevelSupportedReportV3 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchMultilevelSupportedReportV3{
				Properties1: data[0],
				Properties2: data[1],
        data: data,
    }
}

