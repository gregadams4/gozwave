package commands

import (
	"encoding/binary"
	"log"
)

type SwitchColorSupportedGet struct {
	node byte
}

func NewSwitchColorSupportedGet() SwitchColorSupportedGet {
	return SwitchColorSupportedGet{}
}

func (c *SwitchColorSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorSupportedGet) Set() error {

	return nil
}

func (c *SwitchColorSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SwitchColor),
		0x01,
		0x25,
	}
}

type SwitchColorSupportedReport struct {
	*report
	ColorComponentmask uint16
	data               []byte
}

func NewSwitchColorSupportedReport(data []byte) *SwitchColorSupportedReport {
	if len(data) < 2 {
		//may want to change this to just return nil
		for i := len(data); i < 2; i++ {
			data = append(data, 0x00)
		}
	}

	return &SwitchColorSupportedReport{
		ColorComponentmask: binary.BigEndian.Uint16(data[0:2]),
		data:               data,
	}
}

type SwitchColorGet struct {
	node             byte
	ColorComponentID byte
}

func NewSwitchColorGet() SwitchColorGet {
	return SwitchColorGet{}
}

func (c *SwitchColorGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorGet) Set(ColorComponentID byte) error {
	c.ColorComponentID = ColorComponentID

	return nil
}

func (c *SwitchColorGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SwitchColor),
		0x03,
		c.ColorComponentID,
		0x25,
	}
}

type SwitchColorReport struct {
	*report
	ColorComponentID byte
	Value            byte
	data             []byte
}

func NewSwitchColorReport(data []byte) *SwitchColorReport {
	if len(data) < 2 {
		//may want to change this to just return nil
		for i := len(data); i < 2; i++ {
			data = append(data, 0x00)
		}
	}

	return &SwitchColorReport{
		ColorComponentID: data[0],
		Value:            data[1],
		data:             data,
	}
}

type SwitchColorSet struct {
	node        byte
	Properties1 byte
	Vg1         []SwitchColorSetvg1
}
type SwitchColorSetvg1 struct {
	ColorComponentID byte
	Value            byte
}

func NewSwitchColorSet() SwitchColorSet {
	return SwitchColorSet{}
}

func (c *SwitchColorSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorSet) Set(Properties1 byte, Vg1 []SwitchColorSetvg1) error {
	c.Properties1 = Properties1
	c.Vg1 = Vg1

	return nil
}

func (c *SwitchColorSet) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(5),
		byte(SwitchColor),
		0x05,
		c.Properties1,
	}
	for _, v := range c.Vg1 {
		data = append(data, v.ColorComponentID)
		data = append(data, v.Value)
	}
	data = append(data, 0x25)
	log.Println(data)
	return data

}

type SwitchColorStartLevelChange struct {
	node             byte
	Properties1      byte
	ColorComponentID byte
	StartLevel       byte
}

func NewSwitchColorStartLevelChange() SwitchColorStartLevelChange {
	return SwitchColorStartLevelChange{}
}

func (c *SwitchColorStartLevelChange) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorStartLevelChange) Set(Properties1 byte, ColorComponentID byte, StartLevel byte) error {
	c.Properties1 = Properties1
	c.ColorComponentID = ColorComponentID
	c.StartLevel = StartLevel

	return nil
}

func (c *SwitchColorStartLevelChange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(SwitchColor),
		0x06,
		c.Properties1,
		c.ColorComponentID,
		c.StartLevel,
		0x25,
	}
}

type SwitchColorStopLevelChange struct {
	node             byte
	ColorComponentID byte
}

func NewSwitchColorStopLevelChange() SwitchColorStopLevelChange {
	return SwitchColorStopLevelChange{}
}

func (c *SwitchColorStopLevelChange) SetNode(node int) {
	c.node = byte(node)
}

func (c *SwitchColorStopLevelChange) Set(ColorComponentID byte) error {
	c.ColorComponentID = ColorComponentID

	return nil
}

func (c *SwitchColorStopLevelChange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SwitchColor),
		0x07,
		c.ColorComponentID,
		0x25,
	}
}
