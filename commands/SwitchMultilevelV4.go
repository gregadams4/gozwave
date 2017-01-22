package commands
type SwitchMultilevelGetV4 struct {
	node byte
}

func NewSwitchMultilevelGetV4() SwitchMultilevelGetV4 {
	return SwitchMultilevelGetV4{}
}

func (c *SwitchMultilevelGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelGetV4) Set() error {

	return nil
}

func (c *SwitchMultilevelGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchMultilevelV4),
		0x02,
		0x25,
	}
}
type SwitchMultilevelReportV4 struct {
    *report
    CurrentValue byte
    TargetValue byte
    Duration byte
    data []byte
}

func NewSwitchMultilevelReportV4(data []byte) *SwitchMultilevelReportV4 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchMultilevelReportV4{
				CurrentValue: data[0],
				TargetValue: data[1],
				Duration: data[2],
        data: data,
    }
}

type SwitchMultilevelSetV4 struct {
	node byte
	Value byte
	DimmingDuration byte
}

func NewSwitchMultilevelSetV4() SwitchMultilevelSetV4 {
	return SwitchMultilevelSetV4{}
}

func (c *SwitchMultilevelSetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelSetV4) Set(Value byte,DimmingDuration byte,) error {
	c.Value = Value
	c.DimmingDuration = DimmingDuration

	return nil
}

func (c *SwitchMultilevelSetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SwitchMultilevelV4),
		0x01,
		c.Value,
		c.DimmingDuration,
		0x25,
	}
}
type SwitchMultilevelStartLevelChangeV4 struct {
	node byte
	Properties1 byte
	StartLevel byte
	DimmingDuration byte
	StepSize byte
}

func NewSwitchMultilevelStartLevelChangeV4() SwitchMultilevelStartLevelChangeV4 {
	return SwitchMultilevelStartLevelChangeV4{}
}

func (c *SwitchMultilevelStartLevelChangeV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelStartLevelChangeV4) Set(Properties1 byte,StartLevel byte,DimmingDuration byte,StepSize byte,) error {
	c.Properties1 = Properties1
	c.StartLevel = StartLevel
	c.DimmingDuration = DimmingDuration
	c.StepSize = StepSize

	return nil
}

func (c *SwitchMultilevelStartLevelChangeV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(SwitchMultilevelV4),
		0x04,
		c.Properties1,
		c.StartLevel,
		c.DimmingDuration,
		c.StepSize,
		0x25,
	}
}
type SwitchMultilevelStopLevelChangeV4 struct {
	node byte
}

func NewSwitchMultilevelStopLevelChangeV4() SwitchMultilevelStopLevelChangeV4 {
	return SwitchMultilevelStopLevelChangeV4{}
}

func (c *SwitchMultilevelStopLevelChangeV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelStopLevelChangeV4) Set() error {

	return nil
}

func (c *SwitchMultilevelStopLevelChangeV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchMultilevelV4),
		0x05,
		0x25,
	}
}
type SwitchMultilevelSupportedGetV4 struct {
	node byte
}

func NewSwitchMultilevelSupportedGetV4() SwitchMultilevelSupportedGetV4 {
	return SwitchMultilevelSupportedGetV4{}
}

func (c *SwitchMultilevelSupportedGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchMultilevelSupportedGetV4) Set() error {

	return nil
}

func (c *SwitchMultilevelSupportedGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchMultilevelV4),
		0x06,
		0x25,
	}
}
type SwitchMultilevelSupportedReportV4 struct {
    *report
    Properties1 byte
    Properties2 byte
    data []byte
}

func NewSwitchMultilevelSupportedReportV4(data []byte) *SwitchMultilevelSupportedReportV4 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchMultilevelSupportedReportV4{
				Properties1: data[0],
				Properties2: data[1],
        data: data,
    }
}

