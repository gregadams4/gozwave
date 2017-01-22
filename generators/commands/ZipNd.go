package commands
type ZipNodeSolicitation struct {
	node byte
	Reserved byte
	NodeID byte
	IPv6Address byte
}

func NewZipNodeSolicitation() ZipNodeSolicitation {
	return ZipNodeSolicitation{}
}

func (c *ZipNodeSolicitation) SetNode(node int) {
	c.node = byte(node)
}

func (c *ZipNodeSolicitation) Set(Reserved byte,NodeID byte,IPv6Address byte,) error {
	c.Reserved = Reserved
	c.NodeID = NodeID
	c.IPv6Address = IPv6Address

	return nil
}

func (c *ZipNodeSolicitation) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ZipNd),
		0x03,
		c.Reserved,
		c.NodeID,
		c.IPv6Address,
		0x25,
	}
}
type ZipInvNodeSolicitation struct {
	node byte
	Properties1 byte
	NodeID byte
}

func NewZipInvNodeSolicitation() ZipInvNodeSolicitation {
	return ZipInvNodeSolicitation{}
}

func (c *ZipInvNodeSolicitation) SetNode(node int) {
	c.node = byte(node)
}

func (c *ZipInvNodeSolicitation) Set(Properties1 byte,NodeID byte,) error {
	c.Properties1 = Properties1
	c.NodeID = NodeID

	return nil
}

func (c *ZipInvNodeSolicitation) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ZipNd),
		0x04,
		c.Properties1,
		c.NodeID,
		0x25,
	}
}
type ZipNodeAdvertisement struct {
	node byte
	Properties1 byte
	NodeID byte
	IPv6Address byte
	HomeID byte
}

func NewZipNodeAdvertisement() ZipNodeAdvertisement {
	return ZipNodeAdvertisement{}
}

func (c *ZipNodeAdvertisement) SetNode(node int) {
	c.node = byte(node)
}

func (c *ZipNodeAdvertisement) Set(Properties1 byte,NodeID byte,IPv6Address byte,HomeID byte,) error {
	c.Properties1 = Properties1
	c.NodeID = NodeID
	c.IPv6Address = IPv6Address
	c.HomeID = HomeID

	return nil
}

func (c *ZipNodeAdvertisement) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(ZipNd),
		0x01,
		c.Properties1,
		c.NodeID,
		c.IPv6Address,
		c.HomeID,
		0x25,
	}
}
