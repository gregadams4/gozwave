package commands
type VersionCommandClassGetV2 struct {
	node byte
	RequestedCommandClass byte
}

func NewVersionCommandClassGetV2() VersionCommandClassGetV2 {
	return VersionCommandClassGetV2{}
}

func (c *VersionCommandClassGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *VersionCommandClassGetV2) Set(RequestedCommandClass byte,) error {
	c.RequestedCommandClass = RequestedCommandClass

	return nil
}

func (c *VersionCommandClassGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(VersionV2),
		0x13,
		c.RequestedCommandClass,
		0x25,
	}
}
type VersionCommandClassReportV2 struct {
    *report
    RequestedCommandClass byte
    CommandClassVersion byte
    data []byte
}

func NewVersionCommandClassReportV2(data []byte) *VersionCommandClassReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &VersionCommandClassReportV2{
				RequestedCommandClass: data[0],
				CommandClassVersion: data[1],
        data: data,
    }
}

type VersionGetV2 struct {
	node byte
}

func NewVersionGetV2() VersionGetV2 {
	return VersionGetV2{}
}

func (c *VersionGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *VersionGetV2) Set() error {

	return nil
}

func (c *VersionGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(VersionV2),
		0x11,
		0x25,
	}
}
type VersionReportV2 struct {
    *report
    ZWaveLibraryType byte
    ZWaveProtocolVersion byte
    ZWaveProtocolSubVersion byte
    Firmware0Version byte
    Firmware0SubVersion byte
    HardwareVersion byte
    Numberoffirmwaretargets byte
    data []byte
}

func NewVersionReportV2(data []byte) *VersionReportV2 {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &VersionReportV2{
				ZWaveLibraryType: data[0],
				ZWaveProtocolVersion: data[1],
				ZWaveProtocolSubVersion: data[2],
				Firmware0Version: data[3],
				Firmware0SubVersion: data[4],
				HardwareVersion: data[5],
				Numberoffirmwaretargets: data[6],
        data: data,
    }
}

