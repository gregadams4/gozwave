package commands
import "encoding/binary"
type FirmwareMdGetV5 struct {
	node byte
}

func NewFirmwareMdGetV5() FirmwareMdGetV5 {
	return FirmwareMdGetV5{}
}

func (c *FirmwareMdGetV5) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareMdGetV5) Set() error {

	return nil
}

func (c *FirmwareMdGetV5) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(FirmwareUpdateMdV5),
		0x01,
		0x25,
	}
}
type FirmwareMdReportV5 struct {
    *report
    ManufacturerID uint16
    Firmware0ID uint16
    Firmware0Checksum uint16
    FirmwareUpgradable byte
    NumberofFirmwareTargets byte
    MaxFragmentSize uint16
    HardwareVersion byte
    data []byte
}

func NewFirmwareMdReportV5(data []byte) *FirmwareMdReportV5 {
    if len(data) < 11 {
				//may want to change this to just return nil
				for i := len(data); i < 11; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareMdReportV5{
				ManufacturerID: binary.BigEndian.Uint16(data[0:2]),
				Firmware0ID: binary.BigEndian.Uint16(data[2:4]),
				Firmware0Checksum: binary.BigEndian.Uint16(data[4:6]),
				FirmwareUpgradable: data[6],
				NumberofFirmwareTargets: data[7],
				MaxFragmentSize: binary.BigEndian.Uint16(data[8:10]),
				HardwareVersion: data[10],
        data: data,
    }
}

type FirmwareUpdateMdGetV5 struct {
	node byte
	NumberofReports byte
	Properties1 byte
	Reportnumber2 byte
}

func NewFirmwareUpdateMdGetV5() FirmwareUpdateMdGetV5 {
	return FirmwareUpdateMdGetV5{}
}

func (c *FirmwareUpdateMdGetV5) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateMdGetV5) Set(NumberofReports byte,Properties1 byte,Reportnumber2 byte,) error {
	c.NumberofReports = NumberofReports
	c.Properties1 = Properties1
	c.Reportnumber2 = Reportnumber2

	return nil
}

func (c *FirmwareUpdateMdGetV5) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(FirmwareUpdateMdV5),
		0x05,
		c.NumberofReports,
		c.Properties1,
		c.Reportnumber2,
		0x25,
	}
}
type FirmwareUpdateMdReportV5 struct {
    *report
    Properties1 byte
    Reportnumber2 byte
    Data byte
    Checksum uint16
    data []byte
}

func NewFirmwareUpdateMdReportV5(data []byte) *FirmwareUpdateMdReportV5 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdReportV5{
				Properties1: data[0],
				Reportnumber2: data[1],
				Data: data[2],
				Checksum: binary.BigEndian.Uint16(data[3:5]),
        data: data,
    }
}

type FirmwareUpdateMdRequestGetV5 struct {
	node byte
	ManufacturerID uint16
	FirmwareID uint16
	Checksum uint16
	FirmwareTarget byte
	FragmentSize uint16
	Properties1 byte
	HardwareVersion byte
}

func NewFirmwareUpdateMdRequestGetV5() FirmwareUpdateMdRequestGetV5 {
	return FirmwareUpdateMdRequestGetV5{}
}

func (c *FirmwareUpdateMdRequestGetV5) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateMdRequestGetV5) Set(ManufacturerID uint16,FirmwareID uint16,Checksum uint16,FirmwareTarget byte,FragmentSize uint16,Properties1 byte,HardwareVersion byte,) error {
	c.ManufacturerID = ManufacturerID
	c.FirmwareID = FirmwareID
	c.Checksum = Checksum
	c.FirmwareTarget = FirmwareTarget
	c.FragmentSize = FragmentSize
	c.Properties1 = Properties1
	c.HardwareVersion = HardwareVersion

	return nil
}

func (c *FirmwareUpdateMdRequestGetV5) Encode() []byte {
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
		byte(9),
		byte(FirmwareUpdateMdV5),
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
		c.HardwareVersion,
		0x25,
	}
}
type FirmwareUpdateMdRequestReportV5 struct {
    *report
    Status byte
    data []byte
}

func NewFirmwareUpdateMdRequestReportV5(data []byte) *FirmwareUpdateMdRequestReportV5 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdRequestReportV5{
				Status: data[0],
        data: data,
    }
}

type FirmwareUpdateMdStatusReportV5 struct {
    *report
    Status byte
    WaitTime uint16
    data []byte
}

func NewFirmwareUpdateMdStatusReportV5(data []byte) *FirmwareUpdateMdStatusReportV5 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdStatusReportV5{
				Status: data[0],
				WaitTime: binary.BigEndian.Uint16(data[1:3]),
        data: data,
    }
}

type FirmwareUpdateActivationSetV5 struct {
	node byte
	ManufacturerID uint16
	FirmwareID uint16
	Checksum uint16
	FirmwareTarget byte
	HardwareVersion byte
}

func NewFirmwareUpdateActivationSetV5() FirmwareUpdateActivationSetV5 {
	return FirmwareUpdateActivationSetV5{}
}

func (c *FirmwareUpdateActivationSetV5) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateActivationSetV5) Set(ManufacturerID uint16,FirmwareID uint16,Checksum uint16,FirmwareTarget byte,HardwareVersion byte,) error {
	c.ManufacturerID = ManufacturerID
	c.FirmwareID = FirmwareID
	c.Checksum = Checksum
	c.FirmwareTarget = FirmwareTarget
	c.HardwareVersion = HardwareVersion

	return nil
}

func (c *FirmwareUpdateActivationSetV5) Encode() []byte {
	ManufacturerIDBytes := []byte{}
	binary.BigEndian.PutUint16(ManufacturerIDBytes, c.ManufacturerID)
	FirmwareIDBytes := []byte{}
	binary.BigEndian.PutUint16(FirmwareIDBytes, c.FirmwareID)
	ChecksumBytes := []byte{}
	binary.BigEndian.PutUint16(ChecksumBytes, c.Checksum)
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(FirmwareUpdateMdV5),
		0x08,
		ManufacturerIDBytes[0],
		ManufacturerIDBytes[1],
		FirmwareIDBytes[0],
		FirmwareIDBytes[1],
		ChecksumBytes[0],
		ChecksumBytes[1],
		c.FirmwareTarget,
		c.HardwareVersion,
		0x25,
	}
}
type FirmwareUpdateActivationStatusReportV5 struct {
    *report
    ManufacturerID uint16
    FirmwareID byte
    Checksum uint16
    FirmwareTarget byte
    FirmwareUpdateStatus byte
    HardwareVersion byte
    data []byte
}

func NewFirmwareUpdateActivationStatusReportV5(data []byte) *FirmwareUpdateActivationStatusReportV5 {
    if len(data) < 8 {
				//may want to change this to just return nil
				for i := len(data); i < 8; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateActivationStatusReportV5{
				ManufacturerID: binary.BigEndian.Uint16(data[0:2]),
				FirmwareID: data[2],
				Checksum: binary.BigEndian.Uint16(data[3:5]),
				FirmwareTarget: data[5],
				FirmwareUpdateStatus: data[6],
				HardwareVersion: data[7],
        data: data,
    }
}

type FirmwareUpdateMdPrepareGetV5 struct {
	node byte
	ManufacturerID uint16
	FirmwareID uint16
	FirmwareTarget byte
	FragmentSize uint16
	HardwareVersion byte
}

func NewFirmwareUpdateMdPrepareGetV5() FirmwareUpdateMdPrepareGetV5 {
	return FirmwareUpdateMdPrepareGetV5{}
}

func (c *FirmwareUpdateMdPrepareGetV5) SetNode(node int) {
	c.node = byte(node)
}

func (c *FirmwareUpdateMdPrepareGetV5) Set(ManufacturerID uint16,FirmwareID uint16,FirmwareTarget byte,FragmentSize uint16,HardwareVersion byte,) error {
	c.ManufacturerID = ManufacturerID
	c.FirmwareID = FirmwareID
	c.FirmwareTarget = FirmwareTarget
	c.FragmentSize = FragmentSize
	c.HardwareVersion = HardwareVersion

	return nil
}

func (c *FirmwareUpdateMdPrepareGetV5) Encode() []byte {
	ManufacturerIDBytes := []byte{}
	binary.BigEndian.PutUint16(ManufacturerIDBytes, c.ManufacturerID)
	FirmwareIDBytes := []byte{}
	binary.BigEndian.PutUint16(FirmwareIDBytes, c.FirmwareID)
	FragmentSizeBytes := []byte{}
	binary.BigEndian.PutUint16(FragmentSizeBytes, c.FragmentSize)
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(FirmwareUpdateMdV5),
		0x0A,
		ManufacturerIDBytes[0],
		ManufacturerIDBytes[1],
		FirmwareIDBytes[0],
		FirmwareIDBytes[1],
		c.FirmwareTarget,
		FragmentSizeBytes[0],
		FragmentSizeBytes[1],
		c.HardwareVersion,
		0x25,
	}
}
type FirmwareUpdateMdPrepareReportV5 struct {
    *report
    Status byte
    FirmwareChecksum uint16
    data []byte
}

func NewFirmwareUpdateMdPrepareReportV5(data []byte) *FirmwareUpdateMdPrepareReportV5 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &FirmwareUpdateMdPrepareReportV5{
				Status: data[0],
				FirmwareChecksum: binary.BigEndian.Uint16(data[1:3]),
        data: data,
    }
}

