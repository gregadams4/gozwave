package commands
type SwitchBinaryGetV2 struct {
	node byte
}

func NewSwitchBinaryGetV2() SwitchBinaryGetV2 {
	return SwitchBinaryGetV2{}
}

func (c *SwitchBinaryGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchBinaryGetV2) Set() error {

	return nil
}

func (c *SwitchBinaryGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchBinaryV2),
		0x02,
		0x25,
	}
}
type SwitchBinaryReportV2 struct {
    *report
    CurrentValue byte
    TargetValue byte
    Duration byte
    data []byte
}

func NewSwitchBinaryReportV2(data []byte) *SwitchBinaryReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchBinaryReportV2{
				CurrentValue: data[0],
				TargetValue: data[1],
				Duration: data[2],
        data: data,
    }
}

type SwitchBinarySetV2 struct {
	node byte
	TargetValue byte
	Duration byte
}

func NewSwitchBinarySetV2() SwitchBinarySetV2 {
	return SwitchBinarySetV2{}
}

func (c *SwitchBinarySetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchBinarySetV2) Set(TargetValue byte,Duration byte,) error {
	c.TargetValue = TargetValue
	c.Duration = Duration

	return nil
}

func (c *SwitchBinarySetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SwitchBinaryV2),
		0x01,
		c.TargetValue,
		c.Duration,
		0x25,
	}
}
