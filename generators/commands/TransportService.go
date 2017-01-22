package commands
import "encoding/binary"
type CommandFirstFragment struct {
	node byte
	Properties1 byte
	datagram_size_2 byte
	Properties2 byte
	Payload byte
	Checksum uint16
}

func NewCommandFirstFragment() CommandFirstFragment {
	return CommandFirstFragment{}
}

func (c *CommandFirstFragment) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandFirstFragment) Set(Properties1 byte,datagram_size_2 byte,Properties2 byte,Payload byte,Checksum uint16,) error {
	c.Properties1 = Properties1
	c.datagram_size_2 = datagram_size_2
	c.Properties2 = Properties2
	c.Payload = Payload
	c.Checksum = Checksum

	return nil
}

func (c *CommandFirstFragment) Encode() []byte {
	ChecksumBytes := []byte{}
	binary.BigEndian.PutUint16(ChecksumBytes, c.Checksum)
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(TransportService),
		0xC0,
		c.Properties1,
		c.datagram_size_2,
		c.Properties2,
		c.Payload,
		ChecksumBytes[0],
		ChecksumBytes[1],
		0x25,
	}
}
type CommandSubsequentFragment struct {
	node byte
	Properties1 byte
	datagram_size_2 byte
	Properties2 byte
	datagram_offset_2 byte
	Payload byte
	Checksum uint16
}

func NewCommandSubsequentFragment() CommandSubsequentFragment {
	return CommandSubsequentFragment{}
}

func (c *CommandSubsequentFragment) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandSubsequentFragment) Set(Properties1 byte,datagram_size_2 byte,Properties2 byte,datagram_offset_2 byte,Payload byte,Checksum uint16,) error {
	c.Properties1 = Properties1
	c.datagram_size_2 = datagram_size_2
	c.Properties2 = Properties2
	c.datagram_offset_2 = datagram_offset_2
	c.Payload = Payload
	c.Checksum = Checksum

	return nil
}

func (c *CommandSubsequentFragment) Encode() []byte {
	ChecksumBytes := []byte{}
	binary.BigEndian.PutUint16(ChecksumBytes, c.Checksum)
	return []byte{
		0x13,
		c.node,
		byte(8),
		byte(TransportService),
		0xE0,
		c.Properties1,
		c.datagram_size_2,
		c.Properties2,
		c.datagram_offset_2,
		c.Payload,
		ChecksumBytes[0],
		ChecksumBytes[1],
		0x25,
	}
}
