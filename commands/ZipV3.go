package commands
type CommandZipPacketV3 struct {
	node byte
	Properties1 byte
	Properties2 byte
	SeqNo byte
	Properties3 byte
	Properties4 byte
	HeaderLength byte
	Headerextension byte
	ZWavecommand byte
}

func NewCommandZipPacketV3() CommandZipPacketV3 {
	return CommandZipPacketV3{}
}

func (c *CommandZipPacketV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandZipPacketV3) Set(Properties1 byte,Properties2 byte,SeqNo byte,Properties3 byte,Properties4 byte,HeaderLength byte,Headerextension byte,ZWavecommand byte,) error {
	c.Properties1 = Properties1
	c.Properties2 = Properties2
	c.SeqNo = SeqNo
	c.Properties3 = Properties3
	c.Properties4 = Properties4
	c.HeaderLength = HeaderLength
	c.Headerextension = Headerextension
	c.ZWavecommand = ZWavecommand

	return nil
}

func (c *CommandZipPacketV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(10),
		byte(ZipV3),
		0x02,
		c.Properties1,
		c.Properties2,
		c.SeqNo,
		c.Properties3,
		c.Properties4,
		c.HeaderLength,
		c.Headerextension,
		c.ZWavecommand,
		0x25,
	}
}
