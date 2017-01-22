package commands
import "encoding/binary"
type CommandFirstSegmentV2 struct {
	node byte
	Properties1 byte
	datagram_size_2 byte
	Properties2 byte
	HeaderExtensionLength byte
	HeaderExtension byte
	Payload byte
	FrameCheckSequence uint16
}

func NewCommandFirstSegmentV2() CommandFirstSegmentV2 {
	return CommandFirstSegmentV2{}
}

func (c *CommandFirstSegmentV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandFirstSegmentV2) Set(Properties1 byte,datagram_size_2 byte,Properties2 byte,HeaderExtensionLength byte,HeaderExtension byte,Payload byte,FrameCheckSequence uint16,) error {
	c.Properties1 = Properties1
	c.datagram_size_2 = datagram_size_2
	c.Properties2 = Properties2
	c.HeaderExtensionLength = HeaderExtensionLength
	c.HeaderExtension = HeaderExtension
	c.Payload = Payload
	c.FrameCheckSequence = FrameCheckSequence

	return nil
}

func (c *CommandFirstSegmentV2) Encode() []byte {
	FrameCheckSequenceBytes := []byte{}
	binary.BigEndian.PutUint16(FrameCheckSequenceBytes, c.FrameCheckSequence)
	return []byte{
		0x13,
		c.node,
		byte(9),
		byte(TransportServiceV2),
		0xC0,
		c.Properties1,
		c.datagram_size_2,
		c.Properties2,
		c.HeaderExtensionLength,
		c.HeaderExtension,
		c.Payload,
		FrameCheckSequenceBytes[0],
		FrameCheckSequenceBytes[1],
		0x25,
	}
}
type CommandSegmentCompleteV2 struct {
	node byte
	Properties1 byte
	Properties2 byte
}

func NewCommandSegmentCompleteV2() CommandSegmentCompleteV2 {
	return CommandSegmentCompleteV2{}
}

func (c *CommandSegmentCompleteV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandSegmentCompleteV2) Set(Properties1 byte,Properties2 byte,) error {
	c.Properties1 = Properties1
	c.Properties2 = Properties2

	return nil
}

func (c *CommandSegmentCompleteV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(TransportServiceV2),
		0xE8,
		c.Properties1,
		c.Properties2,
		0x25,
	}
}
type CommandSegmentRequestV2 struct {
	node byte
	Properties1 byte
	Properties2 byte
	datagram_offset_2 byte
}

func NewCommandSegmentRequestV2() CommandSegmentRequestV2 {
	return CommandSegmentRequestV2{}
}

func (c *CommandSegmentRequestV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandSegmentRequestV2) Set(Properties1 byte,Properties2 byte,datagram_offset_2 byte,) error {
	c.Properties1 = Properties1
	c.Properties2 = Properties2
	c.datagram_offset_2 = datagram_offset_2

	return nil
}

func (c *CommandSegmentRequestV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(TransportServiceV2),
		0xC8,
		c.Properties1,
		c.Properties2,
		c.datagram_offset_2,
		0x25,
	}
}
type CommandSegmentWaitV2 struct {
	node byte
	Properties1 byte
	pending_fragments byte
}

func NewCommandSegmentWaitV2() CommandSegmentWaitV2 {
	return CommandSegmentWaitV2{}
}

func (c *CommandSegmentWaitV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandSegmentWaitV2) Set(Properties1 byte,pending_fragments byte,) error {
	c.Properties1 = Properties1
	c.pending_fragments = pending_fragments

	return nil
}

func (c *CommandSegmentWaitV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(TransportServiceV2),
		0xF0,
		c.Properties1,
		c.pending_fragments,
		0x25,
	}
}
type CommandSubsequentSegmentV2 struct {
	node byte
	Properties1 byte
	datagram_size_2 byte
	Properties2 byte
	datagram_offset_2 byte
	HeaderExtensionLength byte
	HeaderExtension byte
	Payload byte
	FrameCheckSequence uint16
}

func NewCommandSubsequentSegmentV2() CommandSubsequentSegmentV2 {
	return CommandSubsequentSegmentV2{}
}

func (c *CommandSubsequentSegmentV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandSubsequentSegmentV2) Set(Properties1 byte,datagram_size_2 byte,Properties2 byte,datagram_offset_2 byte,HeaderExtensionLength byte,HeaderExtension byte,Payload byte,FrameCheckSequence uint16,) error {
	c.Properties1 = Properties1
	c.datagram_size_2 = datagram_size_2
	c.Properties2 = Properties2
	c.datagram_offset_2 = datagram_offset_2
	c.HeaderExtensionLength = HeaderExtensionLength
	c.HeaderExtension = HeaderExtension
	c.Payload = Payload
	c.FrameCheckSequence = FrameCheckSequence

	return nil
}

func (c *CommandSubsequentSegmentV2) Encode() []byte {
	FrameCheckSequenceBytes := []byte{}
	binary.BigEndian.PutUint16(FrameCheckSequenceBytes, c.FrameCheckSequence)
	return []byte{
		0x13,
		c.node,
		byte(10),
		byte(TransportServiceV2),
		0xE0,
		c.Properties1,
		c.datagram_size_2,
		c.Properties2,
		c.datagram_offset_2,
		c.HeaderExtensionLength,
		c.HeaderExtension,
		c.Payload,
		FrameCheckSequenceBytes[0],
		FrameCheckSequenceBytes[1],
		0x25,
	}
}
