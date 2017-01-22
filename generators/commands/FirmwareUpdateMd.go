package commands
import "encoding/binary"
type FirmwareMdGet struct {
	node byte
}

func NewFirmwareMdGet() FirmwareMdGet {
	return FirmwareMdGet{}
}

func (c *FirmwareMdGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareMdGet) Set() error {

	return nil
}

func (c *FirmwareMdGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(FirmwareUpdateMd),
		0x01,
		0x25,
	}
}
type FirmwareMdReport struct {
    *report
    ManufacturerID uint16
    FirmwareID uint16
    Checksum uint16
    data []byte
}

func NewFirmwareMdReport(data []byte) *FirmwareMdReport {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareMdReport{
				ManufacturerID: binary.BigEndian.Uint16(data[0:2]),
				FirmwareID: binary.BigEndian.Uint16(data[2:4]),
				Checksum: binary.BigEndian.Uint16(data[4:6]),
        data: data,
    }
}

type FirmwareUpdateMdGet struct {
	node byte
	NumberofReports byte
	Properties1 byte
	Reportnumber2 byte
}

func NewFirmwareUpdateMdGet() FirmwareUpdateMdGet {
	return FirmwareUpdateMdGet{}
}

func (c *FirmwareUpdateMdGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateMdGet) Set(NumberofReports byte,Properties1 byte,Reportnumber2 byte,) error {
	c.NumberofReports = NumberofReports
	c.Properties1 = Properties1
	c.Reportnumber2 = Reportnumber2

	return nil
}

func (c *FirmwareUpdateMdGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(FirmwareUpdateMd),
		0x05,
		c.NumberofReports,
		c.Properties1,
		c.Reportnumber2,
		0x25,
	}
}
type FirmwareUpdateMdReport struct {
    *report
    Properties1 byte
    Reportnumber2 byte
    Data byte
    data []byte
}

func NewFirmwareUpdateMdReport(data []byte) *FirmwareUpdateMdReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdReport{
				Properties1: data[0],
				Reportnumber2: data[1],
				Data: data[2],
        data: data,
    }
}

type FirmwareUpdateMdRequestGet struct {
	node byte
	ManufacturerID uint16
	FirmwareID uint16
	Checksum uint16
}

func NewFirmwareUpdateMdRequestGet() FirmwareUpdateMdRequestGet {
	return FirmwareUpdateMdRequestGet{}
}

func (c *FirmwareUpdateMdRequestGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateMdRequestGet) Set(ManufacturerID uint16,FirmwareID uint16,Checksum uint16,) error {
	c.ManufacturerID = ManufacturerID
	c.FirmwareID = FirmwareID
	c.Checksum = Checksum

	return nil
}

func (c *FirmwareUpdateMdRequestGet) Encode() []byte {
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
		byte(FirmwareUpdateMd),
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
type FirmwareUpdateMdRequestReport struct {
    *report
    Status byte
    data []byte
}

func NewFirmwareUpdateMdRequestReport(data []byte) *FirmwareUpdateMdRequestReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdRequestReport{
				Status: data[0],
        data: data,
    }
}

type FirmwareUpdateMdStatusReport struct {
    *report
    Status byte
    data []byte
}

func NewFirmwareUpdateMdStatusReport(data []byte) *FirmwareUpdateMdStatusReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdStatusReport{
				Status: data[0],
        data: data,
    }
}

