package commands
type WindowCoveringSupportedGet struct {
	node byte
}

func NewWindowCoveringSupportedGet() WindowCoveringSupportedGet {
	return WindowCoveringSupportedGet{}
}

func (c *WindowCoveringSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *WindowCoveringSupportedGet) Set() error {

	return nil
}

func (c *WindowCoveringSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(WindowCovering),
		0x01,
		0x25,
	}
}
type WindowCoveringSupportedReport struct {
    *report
    Properties1 byte
    ParameterMask byte
    data []byte
}

func NewWindowCoveringSupportedReport(data []byte) *WindowCoveringSupportedReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &WindowCoveringSupportedReport{
				Properties1: data[0],
				ParameterMask: data[1],
        data: data,
    }
}

type WindowCoveringGet struct {
	node byte
	ParameterID byte
}

func NewWindowCoveringGet() WindowCoveringGet {
	return WindowCoveringGet{}
}

func (c *WindowCoveringGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *WindowCoveringGet) Set(ParameterID byte,) error {
	c.ParameterID = ParameterID

	return nil
}

func (c *WindowCoveringGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(WindowCovering),
		0x03,
		c.ParameterID,
		0x25,
	}
}
type WindowCoveringReport struct {
    *report
    ParameterID byte
    CurrentValue byte
    TargetValue byte
    Duration byte
    data []byte
}

func NewWindowCoveringReport(data []byte) *WindowCoveringReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &WindowCoveringReport{
				ParameterID: data[0],
				CurrentValue: data[1],
				TargetValue: data[2],
				Duration: data[3],
        data: data,
    }
}

type WindowCoveringSet struct {
	node byte
	Properties1 byte
	Duration byte
	Vg1 []WindowCoveringSetvg1
}
type WindowCoveringSetvg1 struct {
	ParameterID byte
	Value byte
}

func NewWindowCoveringSet() WindowCoveringSet {
	return WindowCoveringSet{}
}

func (c *WindowCoveringSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *WindowCoveringSet) Set(Properties1 byte,Duration byte,Vg1 []WindowCoveringSetvg1,) error {
	c.Properties1 = Properties1
	c.Duration = Duration
	c.Vg1 = Vg1

	return nil
}

func (c *WindowCoveringSet) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(6),
		byte(WindowCovering),
		0x05,
		c.Properties1,
		c.Duration,
	}
	for _, v := range c.Vg1 {
		data = append(data, v.ParameterID)
		data = append(data, v.Value)
	}
	data = append(data, 0x25)
	return data
	
}
type WindowCoveringStartLevelChange struct {
	node byte
	Properties1 byte
	ParameterID byte
	Duration byte
}

func NewWindowCoveringStartLevelChange() WindowCoveringStartLevelChange {
	return WindowCoveringStartLevelChange{}
}

func (c *WindowCoveringStartLevelChange) SetNode(node int) {
	c.node = byte(node)
}

func (c *WindowCoveringStartLevelChange) Set(Properties1 byte,ParameterID byte,Duration byte,) error {
	c.Properties1 = Properties1
	c.ParameterID = ParameterID
	c.Duration = Duration

	return nil
}

func (c *WindowCoveringStartLevelChange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(WindowCovering),
		0x06,
		c.Properties1,
		c.ParameterID,
		c.Duration,
		0x25,
	}
}
type WindowCoveringStopLevelChange struct {
	node byte
	ParameterID byte
}

func NewWindowCoveringStopLevelChange() WindowCoveringStopLevelChange {
	return WindowCoveringStopLevelChange{}
}

func (c *WindowCoveringStopLevelChange) SetNode(node int) {
	c.node = byte(node)
}

func (c *WindowCoveringStopLevelChange) Set(ParameterID byte,) error {
	c.ParameterID = ParameterID

	return nil
}

func (c *WindowCoveringStopLevelChange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(WindowCovering),
		0x07,
		c.ParameterID,
		0x25,
	}
}
