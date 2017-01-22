package commands
type CommandZipPacket struct {
	node byte
	Properties1 byte
	Properties2 byte
	SeqNo byte
	Properties3 byte
	Properties4 byte
	Headerextension byte
	ZWavecommand byte
}

func NewCommandZipPacket() CommandZipPacket {
	return CommandZipPacket{}
}

func (c *CommandZipPacket) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandZipPacket) Set(Properties1 byte,Properties2 byte,SeqNo byte,Properties3 byte,Properties4 byte,Headerextension byte,ZWavecommand byte,) error {
	c.Properties1 = Properties1
	c.Properties2 = Properties2
	c.SeqNo = SeqNo
	c.Properties3 = Properties3
	c.Properties4 = Properties4
	c.Headerextension = Headerextension
	c.ZWavecommand = ZWavecommand

	return nil
}

func (c *CommandZipPacket) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(9),
		byte(Zip),
		0x02,
		c.Properties1,
		c.Properties2,
		c.SeqNo,
		c.Properties3,
		c.Properties4,
		c.Headerextension,
		c.ZWavecommand,
		0x25,
	}
}
