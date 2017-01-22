package commands
type CmdEncapsulation struct {
	node byte
	Data byte
}

func NewCmdEncapsulation() CmdEncapsulation {
	return CmdEncapsulation{}
}

func (c *CmdEncapsulation) SetNode(node int) {
	c.node = byte(node)
}

func (c *CmdEncapsulation) Set(Data byte,) error {
	c.Data = Data

	return nil
}

func (c *CmdEncapsulation) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(PrepaymentEncapsulation),
		0x01,
		c.Data,
		0x25,
	}
}
