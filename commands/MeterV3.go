package commands
import "encoding/binary"
type MeterGetV3 struct {
	node byte
	Properties1 byte
}

func NewMeterGetV3() MeterGetV3 {
	return MeterGetV3{}
}

func (c *MeterGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterGetV3) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *MeterGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MeterV3),
		0x01,
		c.Properties1,
		0x25,
	}
}
type MeterReportV3 struct {
    *report
    Properties1 byte
    Properties2 byte
    MeterValue byte
    DeltaTime uint16
    PreviousMeterValue byte
    data []byte
}

func NewMeterReportV3(data []byte) *MeterReportV3 {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterReportV3{
				Properties1: data[0],
				Properties2: data[1],
				MeterValue: data[2],
				DeltaTime: binary.BigEndian.Uint16(data[3:5]),
				PreviousMeterValue: data[5],
        data: data,
    }
}

type MeterResetV3 struct {
	node byte
}

func NewMeterResetV3() MeterResetV3 {
	return MeterResetV3{}
}

func (c *MeterResetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterResetV3) Set() error {

	return nil
}

func (c *MeterResetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterV3),
		0x05,
		0x25,
	}
}
type MeterSupportedGetV3 struct {
	node byte
}

func NewMeterSupportedGetV3() MeterSupportedGetV3 {
	return MeterSupportedGetV3{}
}

func (c *MeterSupportedGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterSupportedGetV3) Set() error {

	return nil
}

func (c *MeterSupportedGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterV3),
		0x03,
		0x25,
	}
}
type MeterSupportedReportV3 struct {
    *report
    Properties1 byte
    ScaleSupported byte
    data []byte
}

func NewMeterSupportedReportV3(data []byte) *MeterSupportedReportV3 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterSupportedReportV3{
				Properties1: data[0],
				ScaleSupported: data[1],
        data: data,
    }
}

