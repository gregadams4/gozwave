package commands
import "encoding/binary"
type ZwaveplusInfoGetV2 struct {
	node byte
}

func NewZwaveplusInfoGetV2() ZwaveplusInfoGetV2 {
	return ZwaveplusInfoGetV2{}
}

func (c *ZwaveplusInfoGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ZwaveplusInfoGetV2) Set() error {

	return nil
}

func (c *ZwaveplusInfoGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ZwaveplusInfoV2),
		0x01,
		0x25,
	}
}
type ZwaveplusInfoReportV2 struct {
    *report
    ZWavePlusVersion byte
    RoleType byte
    NodeType byte
    InstallerIconType uint16
    UserIconType uint16
    data []byte
}

func NewZwaveplusInfoReportV2(data []byte) *ZwaveplusInfoReportV2 {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &ZwaveplusInfoReportV2{
				ZWavePlusVersion: data[0],
				RoleType: data[1],
				NodeType: data[2],
				InstallerIconType: binary.BigEndian.Uint16(data[3:5]),
				UserIconType: binary.BigEndian.Uint16(data[5:7]),
        data: data,
    }
}

