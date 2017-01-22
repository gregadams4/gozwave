package commands
import "encoding/binary"
type FirmwareMdGetV4 struct {
	node byte
}

func NewFirmwareMdGetV4() FirmwareMdGetV4 {
	return FirmwareMdGetV4{}
}

func (c *FirmwareMdGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareMdGetV4) Set() error {

	return nil
}

func (c *FirmwareMdGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(FirmwareUpdateMdV4),
		0x01,
		0x25,
	}
}
type FirmwareMdReportV4 struct {
    *report
    ManufacturerID uint16
    Firmware0ID uint16
    Firmware0Checksum uint16
    FirmwareUpgradable byte
    NumberofFirmwareTargets byte
    MaxFragmentSize uint16
    data []byte
}

func NewFirmwareMdReportV4(data []byte) *FirmwareMdReportV4 {
    if len(data) < 10 {
				//may want to change this to just return nil
				for i := len(data); i < 10; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareMdReportV4{
				ManufacturerID: binary.BigEndian.Uint16(data[0:2]),
				Firmware0ID: binary.BigEndian.Uint16(data[2:4]),
				Firmware0Checksum: binary.BigEndian.Uint16(data[4:6]),
				FirmwareUpgradable: data[6],
				NumberofFirmwareTargets: data[7],
				MaxFragmentSize: binary.BigEndian.Uint16(data[8:10]),
        data: data,
    }
}

type FirmwareUpdateMdGetV4 struct {
	node byte
	NumberofReports byte
	Properties1 byte
	Reportnumber2 byte
}

func NewFirmwareUpdateMdGetV4() FirmwareUpdateMdGetV4 {
	return FirmwareUpdateMdGetV4{}
}

func (c *FirmwareUpdateMdGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateMdGetV4) Set(NumberofReports byte,Properties1 byte,Reportnumber2 byte,) error {
	c.NumberofReports = NumberofReports
	c.Properties1 = Properties1
	c.Reportnumber2 = Reportnumber2

	return nil
}

func (c *FirmwareUpdateMdGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(FirmwareUpdateMdV4),
		0x05,
		c.NumberofReports,
		c.Properties1,
		c.Reportnumber2,
		0x25,
	}
}
type FirmwareUpdateMdReportV4 struct {
    *report
    Properties1 byte
    Reportnumber2 byte
    Data byte
    Checksum uint16
    data []byte
}

func NewFirmwareUpdateMdReportV4(data []byte) *FirmwareUpdateMdReportV4 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdReportV4{
				Properties1: data[0],
				Reportnumber2: data[1],
				Data: data[2],
				Checksum: binary.BigEndian.Uint16(data[3:5]),
        data: data,
    }
}

type FirmwareUpdateMdRequestGetV4 struct {
	node byte
	ManufacturerID uint16
	FirmwareID uint16
	Checksum uint16
	FirmwareTarget byte
	FragmentSize uint16
	Properties1 byte
}

func NewFirmwareUpdateMdRequestGetV4() FirmwareUpdateMdRequestGetV4 {
	return FirmwareUpdateMdRequestGetV4{}
}

func (c *FirmwareUpdateMdRequestGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateMdRequestGetV4) Set(ManufacturerID uint16,FirmwareID uint16,Checksum uint16,FirmwareTarget byte,FragmentSize uint16,Properties1 byte,) error {
	c.ManufacturerID = ManufacturerID
	c.FirmwareID = FirmwareID
	c.Checksum = Checksum
	c.FirmwareTarget = FirmwareTarget
	c.FragmentSize = FragmentSize
	c.Properties1 = Properties1

	return nil
}

func (c *FirmwareUpdateMdRequestGetV4) Encode() []byte {
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
		byte(8),
		byte(FirmwareUpdateMdV4),
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
		c.Properties1,
		0x25,
	}
}
type FirmwareUpdateMdRequestReportV4 struct {
    *report
    Status byte
    data []byte
}

func NewFirmwareUpdateMdRequestReportV4(data []byte) *FirmwareUpdateMdRequestReportV4 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdRequestReportV4{
				Status: data[0],
        data: data,
    }
}

type FirmwareUpdateMdStatusReportV4 struct {
    *report
    Status byte
    WaitTime uint16
    data []byte
}

func NewFirmwareUpdateMdStatusReportV4(data []byte) *FirmwareUpdateMdStatusReportV4 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdStatusReportV4{
				Status: data[0],
				WaitTime: binary.BigEndian.Uint16(data[1:3]),
        data: data,
    }
}

type FirmwareUpdateActivationSetV4 struct {
	node byte
	ManufacturerID uint16
	FirmwareID uint16
	Checksum uint16
	FirmwareTarget byte
}

func NewFirmwareUpdateActivationSetV4() FirmwareUpdateActivationSetV4 {
	return FirmwareUpdateActivationSetV4{}
}

func (c *FirmwareUpdateActivationSetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateActivationSetV4) Set(ManufacturerID uint16,FirmwareID uint16,Checksum uint16,FirmwareTarget byte,) error {
	c.ManufacturerID = ManufacturerID
	c.FirmwareID = FirmwareID
	c.Checksum = Checksum
	c.FirmwareTarget = FirmwareTarget

	return nil
}

func (c *FirmwareUpdateActivationSetV4) Encode() []byte {
	ManufacturerIDBytes := []byte{}
	binary.BigEndian.PutUint16(ManufacturerIDBytes, c.ManufacturerID)
	FirmwareIDBytes := []byte{}
	binary.BigEndian.PutUint16(FirmwareIDBytes, c.FirmwareID)
	ChecksumBytes := []byte{}
	binary.BigEndian.PutUint16(ChecksumBytes, c.Checksum)
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(FirmwareUpdateMdV4),
		0x08,
		ManufacturerIDBytes[0],
		ManufacturerIDBytes[1],
		FirmwareIDBytes[0],
		FirmwareIDBytes[1],
		ChecksumBytes[0],
		ChecksumBytes[1],
		c.FirmwareTarget,
		0x25,
	}
}
type FirmwareUpdateActivationStatusReportV4 struct {
    *report
    ManufacturerID uint16
    FirmwareID byte
    Checksum uint16
    FirmwareTarget byte
    FirmwareUpdateStatus byte
    data []byte
}

func NewFirmwareUpdateActivationStatusReportV4(data []byte) *FirmwareUpdateActivationStatusReportV4 {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateActivationStatusReportV4{
				ManufacturerID: binary.BigEndian.Uint16(data[0:2]),
				FirmwareID: data[2],
				Checksum: binary.BigEndian.Uint16(data[3:5]),
				FirmwareTarget: data[5],
				FirmwareUpdateStatus: data[6],
        data: data,
    }
}

