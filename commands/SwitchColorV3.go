package commands
import "encoding/binary"
type SwitchColorSupportedGetV3 struct {
	node byte
}

func NewSwitchColorSupportedGetV3() SwitchColorSupportedGetV3 {
	return SwitchColorSupportedGetV3{}
}

func (c *SwitchColorSupportedGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorSupportedGetV3) Set() error {

	return nil
}

func (c *SwitchColorSupportedGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchColorV3),
		0x01,
		0x25,
	}
}
type SwitchColorSupportedReportV3 struct {
    *report
    ColorComponentmask uint16
    data []byte
}

func NewSwitchColorSupportedReportV3(data []byte) *SwitchColorSupportedReportV3 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchColorSupportedReportV3{
				ColorComponentmask: binary.BigEndian.Uint16(data[0:2]),
        data: data,
    }
}

type SwitchColorGetV3 struct {
	node byte
	ColorComponentID byte
}

func NewSwitchColorGetV3() SwitchColorGetV3 {
	return SwitchColorGetV3{}
}

func (c *SwitchColorGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorGetV3) Set(ColorComponentID byte,) error {
	c.ColorComponentID = ColorComponentID

	return nil
}

func (c *SwitchColorGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SwitchColorV3),
		0x03,
		c.ColorComponentID,
		0x25,
	}
}
type SwitchColorReportV3 struct {
    *report
    ColorComponentID byte
    CurrentValue byte
    TargetValue byte
    Duration byte
    data []byte
}

func NewSwitchColorReportV3(data []byte) *SwitchColorReportV3 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &SwitchColorReportV3{
				ColorComponentID: data[0],
				CurrentValue: data[1],
				TargetValue: data[2],
				Duration: data[3],
        data: data,
    }
}

type SwitchColorSetV3 struct {
	node byte
	Properties1 byte
	Duration byte
	Vg1 []SwitchColorSetV3vg1
}
type SwitchColorSetV3vg1 struct {
	ColorComponentID byte
	Value byte
}

func NewSwitchColorSetV3() SwitchColorSetV3 {
	return SwitchColorSetV3{}
}

func (c *SwitchColorSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorSetV3) Set(Properties1 byte,Duration byte,Vg1 []SwitchColorSetV3vg1,) error {
	c.Properties1 = Properties1
	c.Duration = Duration
	c.Vg1 = Vg1

	return nil
}

func (c *SwitchColorSetV3) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(6),
		byte(SwitchColorV3),
		0x05,
		c.Properties1,
		c.Duration,
	}
	for _, v := range c.Vg1 {
		data = append(data, v.ColorComponentID)
		data = append(data, v.Value)
	}
	data = append(data, 0x25)
	return data
	
}
type SwitchColorStartLevelChangeV3 struct {
	node byte
	Properties1 byte
	ColorComponentID byte
	StartLevel byte
	Duration byte
}

func NewSwitchColorStartLevelChangeV3() SwitchColorStartLevelChangeV3 {
	return SwitchColorStartLevelChangeV3{}
}

func (c *SwitchColorStartLevelChangeV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorStartLevelChangeV3) Set(Properties1 byte,ColorComponentID byte,StartLevel byte,Duration byte,) error {
	c.Properties1 = Properties1
	c.ColorComponentID = ColorComponentID
	c.StartLevel = StartLevel
	c.Duration = Duration

	return nil
}

func (c *SwitchColorStartLevelChangeV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(SwitchColorV3),
		0x06,
		c.Properties1,
		c.ColorComponentID,
		c.StartLevel,
		c.Duration,
		0x25,
	}
}
type SwitchColorStopLevelChangeV3 struct {
	node byte
	ColorComponentID byte
}

func NewSwitchColorStopLevelChangeV3() SwitchColorStopLevelChangeV3 {
	return SwitchColorStopLevelChangeV3{}
}

func (c *SwitchColorStopLevelChangeV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorStopLevelChangeV3) Set(ColorComponentID byte,) error {
	c.ColorComponentID = ColorComponentID

	return nil
}

func (c *SwitchColorStopLevelChangeV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SwitchColorV3),
		0x07,
		c.ColorComponentID,
		0x25,
	}
}
