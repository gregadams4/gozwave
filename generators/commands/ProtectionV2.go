package commands
import "encoding/binary"
type ProtectionEcGetV2 struct {
	node byte
}

func NewProtectionEcGetV2() ProtectionEcGetV2 {
	return ProtectionEcGetV2{}
}

func (c *ProtectionEcGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ProtectionEcGetV2) Set() error {

	return nil
}

func (c *ProtectionEcGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ProtectionV2),
		0x07,
		0x25,
	}
}
type ProtectionEcReportV2 struct {
    *report
    NodeID byte
    data []byte
}

func NewProtectionEcReportV2(data []byte) *ProtectionEcReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ProtectionEcReportV2{
				NodeID: data[0],
        data: data,
    }
}

type ProtectionEcSetV2 struct {
	node byte
	NodeID byte
}

func NewProtectionEcSetV2() ProtectionEcSetV2 {
	return ProtectionEcSetV2{}
}

func (c *ProtectionEcSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ProtectionEcSetV2) Set(NodeID byte,) error {
	c.NodeID = NodeID

	return nil
}

func (c *ProtectionEcSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ProtectionV2),
		0x06,
		c.NodeID,
		0x25,
	}
}
type ProtectionGetV2 struct {
	node byte
}

func NewProtectionGetV2() ProtectionGetV2 {
	return ProtectionGetV2{}
}

func (c *ProtectionGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ProtectionGetV2) Set() error {

	return nil
}

func (c *ProtectionGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ProtectionV2),
		0x02,
		0x25,
	}
}
type ProtectionReportV2 struct {
    *report
    Level byte
    Level2 byte
    data []byte
}

func NewProtectionReportV2(data []byte) *ProtectionReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &ProtectionReportV2{
				Level: data[0],
				Level2: data[1],
        data: data,
    }
}

type ProtectionSetV2 struct {
	node byte
	Level byte
	Level2 byte
}

func NewProtectionSetV2() ProtectionSetV2 {
	return ProtectionSetV2{}
}

func (c *ProtectionSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ProtectionSetV2) Set(Level byte,Level2 byte,) error {
	c.Level = Level
	c.Level2 = Level2

	return nil
}

func (c *ProtectionSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ProtectionV2),
		0x01,
		c.Level,
		c.Level2,
		0x25,
	}
}
type ProtectionSupportedGetV2 struct {
	node byte
}

func NewProtectionSupportedGetV2() ProtectionSupportedGetV2 {
	return ProtectionSupportedGetV2{}
}

func (c *ProtectionSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ProtectionSupportedGetV2) Set() error {

	return nil
}

func (c *ProtectionSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ProtectionV2),
		0x04,
		0x25,
	}
}
type ProtectionSupportedReportV2 struct {
    *report
    Level byte
    LocalProtectionState uint16
    RFProtectionState uint16
    data []byte
}

func NewProtectionSupportedReportV2(data []byte) *ProtectionSupportedReportV2 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &ProtectionSupportedReportV2{
				Level: data[0],
				LocalProtectionState: binary.BigEndian.Uint16(data[1:3]),
				RFProtectionState: binary.BigEndian.Uint16(data[3:5]),
        data: data,
    }
}

type ProtectionTimeoutGetV2 struct {
	node byte
}

func NewProtectionTimeoutGetV2() ProtectionTimeoutGetV2 {
	return ProtectionTimeoutGetV2{}
}

func (c *ProtectionTimeoutGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ProtectionTimeoutGetV2) Set() error {

	return nil
}

func (c *ProtectionTimeoutGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ProtectionV2),
		0x0A,
		0x25,
	}
}
type ProtectionTimeoutReportV2 struct {
    *report
    Timeout byte
    data []byte
}

func NewProtectionTimeoutReportV2(data []byte) *ProtectionTimeoutReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ProtectionTimeoutReportV2{
				Timeout: data[0],
        data: data,
    }
}

type ProtectionTimeoutSetV2 struct {
	node byte
	Timeout byte
}

func NewProtectionTimeoutSetV2() ProtectionTimeoutSetV2 {
	return ProtectionTimeoutSetV2{}
}

func (c *ProtectionTimeoutSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ProtectionTimeoutSetV2) Set(Timeout byte,) error {
	c.Timeout = Timeout

	return nil
}

func (c *ProtectionTimeoutSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ProtectionV2),
		0x09,
		c.Timeout,
		0x25,
	}
}
