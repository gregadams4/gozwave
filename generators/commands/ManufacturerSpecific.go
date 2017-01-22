package commands
import "encoding/binary"
type ManufacturerSpecificGet struct {
	node byte
}

func NewManufacturerSpecificGet() ManufacturerSpecificGet {
	return ManufacturerSpecificGet{}
}

func (c *ManufacturerSpecificGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ManufacturerSpecificGet) Set() error {

	return nil
}

func (c *ManufacturerSpecificGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ManufacturerSpecific),
		0x04,
		0x25,
	}
}
type ManufacturerSpecificReport struct {
    *report
    ManufacturerID uint16
    ProductTypeID uint16
    ProductID uint16
    data []byte
}

func NewManufacturerSpecificReport(data []byte) *ManufacturerSpecificReport {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &ManufacturerSpecificReport{
				ManufacturerID: binary.BigEndian.Uint16(data[0:2]),
				ProductTypeID: binary.BigEndian.Uint16(data[2:4]),
				ProductID: binary.BigEndian.Uint16(data[4:6]),
        data: data,
    }
}

