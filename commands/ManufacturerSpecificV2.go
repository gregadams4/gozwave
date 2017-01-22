package commands
import "encoding/binary"
type ManufacturerSpecificGetV2 struct {
	node byte
}

func NewManufacturerSpecificGetV2() ManufacturerSpecificGetV2 {
	return ManufacturerSpecificGetV2{}
}

func (c *ManufacturerSpecificGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ManufacturerSpecificGetV2) Set() error {

	return nil
}

func (c *ManufacturerSpecificGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ManufacturerSpecificV2),
		0x04,
		0x25,
	}
}
type ManufacturerSpecificReportV2 struct {
    *report
    ManufacturerID uint16
    ProductTypeID uint16
    ProductID uint16
    data []byte
}

func NewManufacturerSpecificReportV2(data []byte) *ManufacturerSpecificReportV2 {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &ManufacturerSpecificReportV2{
				ManufacturerID: binary.BigEndian.Uint16(data[0:2]),
				ProductTypeID: binary.BigEndian.Uint16(data[2:4]),
				ProductID: binary.BigEndian.Uint16(data[4:6]),
        data: data,
    }
}

type DeviceSpecificGetV2 struct {
	node byte
	Properties1 byte
}

func NewDeviceSpecificGetV2() DeviceSpecificGetV2 {
	return DeviceSpecificGetV2{}
}

func (c *DeviceSpecificGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *DeviceSpecificGetV2) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *DeviceSpecificGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ManufacturerSpecificV2),
		0x06,
		c.Properties1,
		0x25,
	}
}
type DeviceSpecificReportV2 struct {
    *report
    Properties1 byte
    Properties2 byte
    DeviceIDData byte
    data []byte
}

func NewDeviceSpecificReportV2(data []byte) *DeviceSpecificReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &DeviceSpecificReportV2{
				Properties1: data[0],
				Properties2: data[1],
				DeviceIDData: data[2],
        data: data,
    }
}

