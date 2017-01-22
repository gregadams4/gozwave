package commands

import "encoding/binary"

type SwitchColorSupportedGetV2 struct {
	node byte
}

func NewSwitchColorSupportedGetV2() SwitchColorSupportedGetV2 {
	return SwitchColorSupportedGetV2{}
}

func (c *SwitchColorSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorSupportedGetV2) Set() error {

	return nil
}

func (c *SwitchColorSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchColorV2),
		0x01,
		0x25,
	}
}

type SwitchColorSupportedReportV2 struct {
	*report
	ColorComponentmask uint16
	data               []byte
}

func NewSwitchColorSupportedReportV2(data []byte) *SwitchColorSupportedReportV2 {
	if len(data) < 2 {
		//may want to change this to just return nil
		for i := len(data); i < 2; i++ {
			data = append(data, 0x00)
		}
	}

	return &SwitchColorSupportedReportV2{
		ColorComponentmask: binary.BigEndian.Uint16(data[0:2]),
		data:               data,
	}
}

type SwitchColorGetV2 struct {
	node             byte
	ColorComponentID byte
}

func NewSwitchColorGetV2() SwitchColorGetV2 {
	return SwitchColorGetV2{}
}

func (c *SwitchColorGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorGetV2) Set(ColorComponentID byte) error {
	c.ColorComponentID = ColorComponentID

	return nil
}

func (c *SwitchColorGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SwitchColorV2),
		0x03,
		c.ColorComponentID,
		0x25,
	}
}

type SwitchColorReportV2 struct {
	*report
	ColorComponentID byte
	Value            byte
	data             []byte
}

func NewSwitchColorReportV2(data []byte) *SwitchColorReportV2 {
	if len(data) < 2 {
		//may want to change this to just return nil
		for i := len(data); i < 2; i++ {
			data = append(data, 0x00)
		}
	}

	return &SwitchColorReportV2{
		ColorComponentID: data[0],
		Value:            data[1],
		data:             data,
	}
}

type SwitchColorSetV2 struct {
	node        byte
	Properties1 byte
	Vg1         []SwitchColorSetV2vg1
	Duration    byte
}
type SwitchColorSetV2vg1 struct {
	ColorComponentID byte
	Value            byte
}

func NewSwitchColorSetV2() SwitchColorSetV2 {
	return SwitchColorSetV2{}
}

func (c *SwitchColorSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorSetV2) Set(Properties1 byte, Duration byte, Vg1 []SwitchColorSetV2vg1) error {
	c.Properties1 = Properties1
	c.Duration = Duration
	c.Vg1 = Vg1

	return nil
}

func (c *SwitchColorSetV2) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(6),
		byte(SwitchColorV2),
		0x05,
		c.Properties1,
	}
	for _, v := range c.Vg1 {
		data = append(data, v.ColorComponentID)
		data = append(data, v.Value)
	}
	data = append(data, c.Duration)
	data = append(data, 0x25)
	return data

}

type SwitchColorStartLevelChangeV2 struct {
	node             byte
	Properties1      byte
	ColorComponentID byte
	StartLevel       byte
}

func NewSwitchColorStartLevelChangeV2() SwitchColorStartLevelChangeV2 {
	return SwitchColorStartLevelChangeV2{}
}

func (c *SwitchColorStartLevelChangeV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorStartLevelChangeV2) Set(Properties1 byte, ColorComponentID byte, StartLevel byte) error {
	c.Properties1 = Properties1
	c.ColorComponentID = ColorComponentID
	c.StartLevel = StartLevel

	return nil
}

func (c *SwitchColorStartLevelChangeV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(SwitchColorV2),
		0x06,
		c.Properties1,
		c.ColorComponentID,
		c.StartLevel,
		0x25,
	}
}

type SwitchColorStopLevelChangeV2 struct {
	node             byte
	ColorComponentID byte
}

func NewSwitchColorStopLevelChangeV2() SwitchColorStopLevelChangeV2 {
	return SwitchColorStopLevelChangeV2{}
}

func (c *SwitchColorStopLevelChangeV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorStopLevelChangeV2) Set(ColorComponentID byte) error {
	c.ColorComponentID = ColorComponentID

	return nil
}

func (c *SwitchColorStopLevelChangeV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SwitchColorV2),
		0x07,
		c.ColorComponentID,
		0x25,
	}
}
