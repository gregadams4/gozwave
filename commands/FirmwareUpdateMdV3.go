package commands
import "encoding/binary"
type FirmwareMdGetV3 struct {
	node byte
}

func NewFirmwareMdGetV3() FirmwareMdGetV3 {
	return FirmwareMdGetV3{}
}

func (c *FirmwareMdGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareMdGetV3) Set() error {

	return nil
}

func (c *FirmwareMdGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(FirmwareUpdateMdV3),
		0x01,
		0x25,
	}
}
type FirmwareMdReportV3 struct {
    *report
    ManufacturerID uint16
    Firmware0ID uint16
    Firmware0Checksum uint16
    FirmwareUpgradable byte
    NumberofFirmwareTargets byte
    MaxFragmentSize uint16
    data []byte
}

func NewFirmwareMdReportV3(data []byte) *FirmwareMdReportV3 {
    if len(data) < 10 {
				//may want to change this to just return nil
				for i := len(data); i < 10; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareMdReportV3{
				ManufacturerID: binary.BigEndian.Uint16(data[0:2]),
				Firmware0ID: binary.BigEndian.Uint16(data[2:4]),
				Firmware0Checksum: binary.BigEndian.Uint16(data[4:6]),
				FirmwareUpgradable: data[6],
				NumberofFirmwareTargets: data[7],
				MaxFragmentSize: binary.BigEndian.Uint16(data[8:10]),
        data: data,
    }
}

type FirmwareUpdateMdGetV3 struct {
	node byte
	NumberofReports byte
	Properties1 byte
	Reportnumber2 byte
}

func NewFirmwareUpdateMdGetV3() FirmwareUpdateMdGetV3 {
	return FirmwareUpdateMdGetV3{}
}

func (c *FirmwareUpdateMdGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateMdGetV3) Set(NumberofReports byte,Properties1 byte,Reportnumber2 byte,) error {
	c.NumberofReports = NumberofReports
	c.Properties1 = Properties1
	c.Reportnumber2 = Reportnumber2

	return nil
}

func (c *FirmwareUpdateMdGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(FirmwareUpdateMdV3),
		0x05,
		c.NumberofReports,
		c.Properties1,
		c.Reportnumber2,
		0x25,
	}
}
type FirmwareUpdateMdReportV3 struct {
    *report
    Properties1 byte
    Reportnumber2 byte
    Data byte
    Checksum uint16
    data []byte
}

func NewFirmwareUpdateMdReportV3(data []byte) *FirmwareUpdateMdReportV3 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdReportV3{
				Properties1: data[0],
				Reportnumber2: data[1],
				Data: data[2],
				Checksum: binary.BigEndian.Uint16(data[3:5]),
        data: data,
    }
}

type FirmwareUpdateMdRequestGetV3 struct {
	node byte
	ManufacturerID uint16
	FirmwareID uint16
	Checksum uint16
	FirmwareTarget byte
	FragmentSize uint16
}

func NewFirmwareUpdateMdRequestGetV3() FirmwareUpdateMdRequestGetV3 {
	return FirmwareUpdateMdRequestGetV3{}
}

func (c *FirmwareUpdateMdRequestGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateMdRequestGetV3) Set(ManufacturerID uint16,FirmwareID uint16,Checksum uint16,FirmwareTarget byte,FragmentSize uint16,) error {
	c.ManufacturerID = ManufacturerID
	c.FirmwareID = FirmwareID
	c.Checksum = Checksum
	c.FirmwareTarget = FirmwareTarget
	c.FragmentSize = FragmentSize

	return nil
}

func (c *FirmwareUpdateMdRequestGetV3) Encode() []byte {
	ManufacturerIDBytes := []byte{}
	binary.BigEndian.PutUint16(ManufacturerIDBytes, c.ManufacturerID)
	FirmwareIDBytes := []byte{}
	binary.BigEndian.PutUint16(FirmwareIDBytes, c.FirmwareID)
	ChecksumBytes := []byte{}
	binary.BigEndian.PutUint16(ChecksumBytes, c.Checksum)
	FragmentSizeBytes := []byte{}
	binary.BigEndian.PutUint16(FragmentSizeBytes, c.FragmentSize)
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(FirmwareUpdateMdV3),
		0x03,
		ManufacturerIDBytes[0],
		ManufacturerIDBytes[1],
		FirmwareIDBytes[0],
		FirmwareIDBytes[1],
		ChecksumBytes[0],
		ChecksumBytes[1],
		c.FirmwareTarget,
		FragmentSizeBytes[0],
		FragmentSizeBytes[1],
		0x25,
	}
}
type FirmwareUpdateMdRequestReportV3 struct {
    *report
    Status byte
    data []byte
}

func NewFirmwareUpdateMdRequestReportV3(data []byte) *FirmwareUpdateMdRequestReportV3 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdRequestReportV3{
				Status: data[0],
        data: data,
    }
}

type FirmwareUpdateMdStatusReportV3 struct {
    *report
    Status byte
    WaitTime uint16
    data []byte
}

func NewFirmwareUpdateMdStatusReportV3(data []byte) *FirmwareUpdateMdStatusReportV3 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdStatusReportV3{
				Status: data[0],
				WaitTime: binary.BigEndian.Uint16(data[1:3]),
        data: data,
    }
}

