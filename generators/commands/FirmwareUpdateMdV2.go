package commands
import "encoding/binary"
type FirmwareMdGetV2 struct {
	node byte
}

func NewFirmwareMdGetV2() FirmwareMdGetV2 {
	return FirmwareMdGetV2{}
}

func (c *FirmwareMdGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareMdGetV2) Set() error {

	return nil
}

func (c *FirmwareMdGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(FirmwareUpdateMdV2),
		0x01,
		0x25,
	}
}
type FirmwareMdReportV2 struct {
    *report
    ManufacturerID uint16
    FirmwareID uint16
    Checksum uint16
    data []byte
}

func NewFirmwareMdReportV2(data []byte) *FirmwareMdReportV2 {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareMdReportV2{
				ManufacturerID: binary.BigEndian.Uint16(data[0:2]),
				FirmwareID: binary.BigEndian.Uint16(data[2:4]),
				Checksum: binary.BigEndian.Uint16(data[4:6]),
        data: data,
    }
}

type FirmwareUpdateMdGetV2 struct {
	node byte
	NumberofReports byte
	Properties1 byte
	Reportnumber2 byte
}

func NewFirmwareUpdateMdGetV2() FirmwareUpdateMdGetV2 {
	return FirmwareUpdateMdGetV2{}
}

func (c *FirmwareUpdateMdGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateMdGetV2) Set(NumberofReports byte,Properties1 byte,Reportnumber2 byte,) error {
	c.NumberofReports = NumberofReports
	c.Properties1 = Properties1
	c.Reportnumber2 = Reportnumber2

	return nil
}

func (c *FirmwareUpdateMdGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(FirmwareUpdateMdV2),
		0x05,
		c.NumberofReports,
		c.Properties1,
		c.Reportnumber2,
		0x25,
	}
}
type FirmwareUpdateMdReportV2 struct {
    *report
    Properties1 byte
    Reportnumber2 byte
    Data byte
    Checksum uint16
    data []byte
}

func NewFirmwareUpdateMdReportV2(data []byte) *FirmwareUpdateMdReportV2 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdReportV2{
				Properties1: data[0],
				Reportnumber2: data[1],
				Data: data[2],
				Checksum: binary.BigEndian.Uint16(data[3:5]),
        data: data,
    }
}

type FirmwareUpdateMdRequestGetV2 struct {
	node byte
	ManufacturerID uint16
	FirmwareID uint16
	Checksum uint16
}

func NewFirmwareUpdateMdRequestGetV2() FirmwareUpdateMdRequestGetV2 {
	return FirmwareUpdateMdRequestGetV2{}
}

func (c *FirmwareUpdateMdRequestGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateMdRequestGetV2) Set(ManufacturerID uint16,FirmwareID uint16,Checksum uint16,) error {
	c.ManufacturerID = ManufacturerID
	c.FirmwareID = FirmwareID
	c.Checksum = Checksum

	return nil
}

func (c *FirmwareUpdateMdRequestGetV2) Encode() []byte {
	ManufacturerIDBytes := []byte{}
	binary.BigEndian.PutUint16(ManufacturerIDBytes, c.ManufacturerID)
	FirmwareIDBytes := []byte{}
	binary.BigEndian.PutUint16(FirmwareIDBytes, c.FirmwareID)
	ChecksumBytes := []byte{}
	binary.BigEndian.PutUint16(ChecksumBytes, c.Checksum)
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(FirmwareUpdateMdV2),
		0x03,
		ManufacturerIDBytes[0],
		ManufacturerIDBytes[1],
		FirmwareIDBytes[0],
		FirmwareIDBytes[1],
		ChecksumBytes[0],
		ChecksumBytes[1],
		0x25,
	}
}
type FirmwareUpdateMdRequestReportV2 struct {
    *report
    Status byte
    data []byte
}

func NewFirmwareUpdateMdRequestReportV2(data []byte) *FirmwareUpdateMdRequestReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdRequestReportV2{
				Status: data[0],
        data: data,
    }
}

type FirmwareUpdateMdStatusReportV2 struct {
    *report
    Status byte
    data []byte
}

func NewFirmwareUpdateMdStatusReportV2(data []byte) *FirmwareUpdateMdStatusReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdStatusReportV2{
				Status: data[0],
        data: data,
    }
}

