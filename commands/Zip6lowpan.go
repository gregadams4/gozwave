package commands
type LowpanFirstFragment struct {
	node byte
	Properties1 byte
	DatagramSize2 byte
	DatagramTag byte
	Payload byte
}

func NewLowpanFirstFragment() LowpanFirstFragment {
	return LowpanFirstFragment{}
}

func (c *LowpanFirstFragment) SetNode(node int) {
	c.node = byte(node)
}

func (c *LowpanFirstFragment) Set(Properties1 byte,DatagramSize2 byte,DatagramTag byte,Payload byte,) error {
	c.Properties1 = Properties1
	c.DatagramSize2 = DatagramSize2
	c.DatagramTag = DatagramTag
	c.Payload = Payload

	return nil
}

func (c *LowpanFirstFragment) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(Zip6lowpan),
		0xC0,
		c.Properties1,
		c.DatagramSize2,
		c.DatagramTag,
		c.Payload,
		0x25,
	}
}
type LowpanSubsequentFragment struct {
	node byte
	Properties1 byte
	DatagramSize2 byte
	DatagramTag byte
	DatagramOffset byte
	Payload byte
}

func NewLowpanSubsequentFragment() LowpanSubsequentFragment {
	return LowpanSubsequentFragment{}
}

func (c *LowpanSubsequentFragment) SetNode(node int) {
	c.node = byte(node)
}

func (c *LowpanSubsequentFragment) Set(Properties1 byte,DatagramSize2 byte,DatagramTag byte,DatagramOffset byte,Payload byte,) error {
	c.Properties1 = Properties1
	c.DatagramSize2 = DatagramSize2
	c.DatagramTag = DatagramTag
	c.DatagramOffset = DatagramOffset
	c.Payload = Payload

	return nil
}

func (c *LowpanSubsequentFragment) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(Zip6lowpan),
		0xE0,
		c.Properties1,
		c.DatagramSize2,
		c.DatagramTag,
		c.DatagramOffset,
		c.Payload,
		0x25,
	}
}
