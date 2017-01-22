package commands
type SupervisionGet struct {
	node byte
	Properties1 byte
	EncapsulatedCommandLength byte
	EncapsulatedCommand byte
}

func NewSupervisionGet() SupervisionGet {
	return SupervisionGet{}
}

func (c *SupervisionGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SupervisionGet) Set(Properties1 byte,EncapsulatedCommandLength byte,EncapsulatedCommand byte,) error {
	c.Properties1 = Properties1
	c.EncapsulatedCommandLength = EncapsulatedCommandLength
	c.EncapsulatedCommand = EncapsulatedCommand

	return nil
}

func (c *SupervisionGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(Supervision),
		0x01,
		c.Properties1,
		c.EncapsulatedCommandLength,
		c.EncapsulatedCommand,
		0x25,
	}
}
type SupervisionReport struct {
    *report
    Properties1 byte
    Status byte
    Duration byte
    data []byte
}

func NewSupervisionReport(data []byte) *SupervisionReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SupervisionReport{
				Properties1: data[0],
				Status: data[1],
				Duration: data[2],
        data: data,
    }
}

