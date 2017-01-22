package commands
import "encoding/binary"
type AntitheftSet struct {
	node byte
	Properties1 byte
	MagicCode byte
	ManufacturerID uint16
	AntitheftHintNumberBytes byte
	AntitheftHintByte byte
}

func NewAntitheftSet() AntitheftSet {
	return AntitheftSet{}
}

func (c *AntitheftSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AntitheftSet) Set(Properties1 byte,MagicCode byte,ManufacturerID uint16,AntitheftHintNumberBytes byte,AntitheftHintByte byte,) error {
	c.Properties1 = Properties1
	c.MagicCode = MagicCode
	c.ManufacturerID = ManufacturerID
	c.AntitheftHintNumberBytes = AntitheftHintNumberBytes
	c.AntitheftHintByte = AntitheftHintByte

	return nil
}

func (c *AntitheftSet) Encode() []byte {
	ManufacturerIDBytes := []byte{}
	binary.BigEndian.PutUint16(ManufacturerIDBytes, c.ManufacturerID)
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(Antitheft),
		0x01,
		c.Properties1,
		c.MagicCode,
		ManufacturerIDBytes[0],
		ManufacturerIDBytes[1],
		c.AntitheftHintNumberBytes,
		c.AntitheftHintByte,
		0x25,
	}
}
type AntitheftGet struct {
	node byte
}

func NewAntitheftGet() AntitheftGet {
	return AntitheftGet{}
}

func (c *AntitheftGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AntitheftGet) Set() error {

	return nil
}

func (c *AntitheftGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Antitheft),
		0x02,
		0x25,
	}
}
type AntitheftReport struct {
    *report
    AntitheftProtectionStatus byte
    ManufacturerID uint16
    AntitheftHintNumberBytes byte
    AntitheftHintByte byte
    data []byte
}

func NewAntitheftReport(data []byte) *AntitheftReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &AntitheftReport{
				AntitheftProtectionStatus: data[0],
				ManufacturerID: binary.BigEndian.Uint16(data[1:3]),
				AntitheftHintNumberBytes: data[3],
				AntitheftHintByte: data[4],
        data: data,
    }
}

