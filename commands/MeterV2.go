package commands
import "encoding/binary"
type MeterGetV2 struct {
	node byte
	Properties1 byte
}

func NewMeterGetV2() MeterGetV2 {
	return MeterGetV2{}
}

func (c *MeterGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterGetV2) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *MeterGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MeterV2),
		0x01,
		c.Properties1,
		0x25,
	}
}
type MeterReportV2 struct {
    *report
    Properties1 byte
    Properties2 byte
    MeterValue byte
    DeltaTime uint16
    PreviousMeterValue byte
    data []byte
}

func NewMeterReportV2(data []byte) *MeterReportV2 {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterReportV2{
				Properties1: data[0],
				Properties2: data[1],
				MeterValue: data[2],
				DeltaTime: binary.BigEndian.Uint16(data[3:5]),
				PreviousMeterValue: data[5],
        data: data,
    }
}

type MeterResetV2 struct {
	node byte
}

func NewMeterResetV2() MeterResetV2 {
	return MeterResetV2{}
}

func (c *MeterResetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterResetV2) Set() error {

	return nil
}

func (c *MeterResetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterV2),
		0x05,
		0x25,
	}
}
type MeterSupportedGetV2 struct {
	node byte
}

func NewMeterSupportedGetV2() MeterSupportedGetV2 {
	return MeterSupportedGetV2{}
}

func (c *MeterSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterSupportedGetV2) Set() error {

	return nil
}

func (c *MeterSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterV2),
		0x03,
		0x25,
	}
}
type MeterSupportedReportV2 struct {
    *report
    Properties1 byte
    Properties2 byte
    data []byte
}

func NewMeterSupportedReportV2(data []byte) *MeterSupportedReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterSupportedReportV2{
				Properties1: data[0],
				Properties2: data[1],
        data: data,
    }
}

