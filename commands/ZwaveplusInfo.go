package commands
type ZwaveplusInfoGet struct {
	node byte
}

func NewZwaveplusInfoGet() ZwaveplusInfoGet {
	return ZwaveplusInfoGet{}
}

func (c *ZwaveplusInfoGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ZwaveplusInfoGet) Set() error {

	return nil
}

func (c *ZwaveplusInfoGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ZwaveplusInfo),
		0x01,
		0x25,
	}
}
type ZwaveplusInfoReport struct {
    *report
    ZWavePlusVersion byte
    RoleType byte
    NodeType byte
    data []byte
}

func NewZwaveplusInfoReport(data []byte) *ZwaveplusInfoReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &ZwaveplusInfoReport{
				ZWavePlusVersion: data[0],
				RoleType: data[1],
				NodeType: data[2],
        data: data,
    }
}

