package commands
import "encoding/binary"
type MeterGetV4 struct {
	node byte
	Properties1 byte
	Scale2 byte
}

func NewMeterGetV4() MeterGetV4 {
	return MeterGetV4{}
}

func (c *MeterGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterGetV4) Set(Properties1 byte,Scale2 byte,) error {
	c.Properties1 = Properties1
	c.Scale2 = Scale2

	return nil
}

func (c *MeterGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(MeterV4),
		0x01,
		c.Properties1,
		c.Scale2,
		0x25,
	}
}
type MeterReportV4 struct {
    *report
    Properties1 byte
    Properties2 byte
    MeterValue byte
    DeltaTime uint16
    PreviousMeterValue byte
    Scale2 byte
    data []byte
}

func NewMeterReportV4(data []byte) *MeterReportV4 {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterReportV4{
				Properties1: data[0],
				Properties2: data[1],
				MeterValue: data[2],
				DeltaTime: binary.BigEndian.Uint16(data[3:5]),
				PreviousMeterValue: data[5],
				Scale2: data[6],
        data: data,
    }
}

type MeterResetV4 struct {
	node byte
}

func NewMeterResetV4() MeterResetV4 {
	return MeterResetV4{}
}

func (c *MeterResetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterResetV4) Set() error {

	return nil
}

func (c *MeterResetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterV4),
		0x05,
		0x25,
	}
}
type MeterSupportedGetV4 struct {
	node byte
}

func NewMeterSupportedGetV4() MeterSupportedGetV4 {
	return MeterSupportedGetV4{}
}

func (c *MeterSupportedGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterSupportedGetV4) Set() error {

	return nil
}

func (c *MeterSupportedGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterV4),
		0x03,
		0x25,
	}
}
type MeterSupportedReportV4 struct {
    *report
    Properties1 byte
    Properties2 byte
    NumberofScaleSupportedBytestoFollow byte
    ScaleSupported byte
    data []byte
}

func NewMeterSupportedReportV4(data []byte) *MeterSupportedReportV4 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterSupportedReportV4{
				Properties1: data[0],
				Properties2: data[1],
				NumberofScaleSupportedBytestoFollow: data[2],
				ScaleSupported: data[3],
        data: data,
    }
}

