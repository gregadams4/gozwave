package commands
type VersionCommandClassGet struct {
	node byte
	RequestedCommandClass byte
}

func NewVersionCommandClassGet() VersionCommandClassGet {
	return VersionCommandClassGet{}
}

func (c *VersionCommandClassGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *VersionCommandClassGet) Set(RequestedCommandClass byte,) error {
	c.RequestedCommandClass = RequestedCommandClass

	return nil
}

func (c *VersionCommandClassGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Version),
		0x13,
		c.RequestedCommandClass,
		0x25,
	}
}
type VersionCommandClassReport struct {
    *report
    RequestedCommandClass byte
    CommandClassVersion byte
    data []byte
}

func NewVersionCommandClassReport(data []byte) *VersionCommandClassReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &VersionCommandClassReport{
				RequestedCommandClass: data[0],
				CommandClassVersion: data[1],
        data: data,
    }
}

type VersionGet struct {
	node byte
}

func NewVersionGet() VersionGet {
	return VersionGet{}
}

func (c *VersionGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *VersionGet) Set() error {

	return nil
}

func (c *VersionGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Version),
		0x11,
		0x25,
	}
}
type VersionReport struct {
    *report
    ZWaveLibraryType byte
    ZWaveProtocolVersion byte
    ZWaveProtocolSubVersion byte
    ApplicationVersion byte
    ApplicationSubVersion byte
    data []byte
}

func NewVersionReport(data []byte) *VersionReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &VersionReport{
				ZWaveLibraryType: data[0],
				ZWaveProtocolVersion: data[1],
				ZWaveProtocolSubVersion: data[2],
				ApplicationVersion: data[3],
				ApplicationSubVersion: data[4],
        data: data,
    }
}

