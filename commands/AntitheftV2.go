package commands
import "encoding/binary"
type AntitheftSetV2 struct {
	node byte
	Properties1 byte
	MagicCode byte
	ManufacturerID uint16
	AntitheftHintNumberBytes byte
	AntitheftHintByte byte
}

func NewAntitheftSetV2() AntitheftSetV2 {
	return AntitheftSetV2{}
}

func (c *AntitheftSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AntitheftSetV2) Set(Properties1 byte,MagicCode byte,ManufacturerID uint16,AntitheftHintNumberBytes byte,AntitheftHintByte byte,) error {
	c.Properties1 = Properties1
	c.MagicCode = MagicCode
	c.ManufacturerID = ManufacturerID
	c.AntitheftHintNumberBytes = AntitheftHintNumberBytes
	c.AntitheftHintByte = AntitheftHintByte

	return nil
}

func (c *AntitheftSetV2) Encode() []byte {
	ManufacturerIDBytes := []byte{}
	binary.BigEndian.PutUint16(ManufacturerIDBytes, c.ManufacturerID)
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(AntitheftV2),
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
type AntitheftGetV2 struct {
	node byte
}

func NewAntitheftGetV2() AntitheftGetV2 {
	return AntitheftGetV2{}
}

func (c *AntitheftGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AntitheftGetV2) Set() error {

	return nil
}

func (c *AntitheftGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(AntitheftV2),
		0x02,
		0x25,
	}
}
type AntitheftReportV2 struct {
    *report
    AntitheftProtectionStatus byte
    ManufacturerID uint16
    AntitheftHintNumberBytes byte
    AntitheftHintByte byte
    data []byte
}

func NewAntitheftReportV2(data []byte) *AntitheftReportV2 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &AntitheftReportV2{
				AntitheftProtectionStatus: data[0],
				ManufacturerID: binary.BigEndian.Uint16(data[1:3]),
				AntitheftHintNumberBytes: data[3],
				AntitheftHintByte: data[4],
        data: data,
    }
}

