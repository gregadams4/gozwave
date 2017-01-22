package commands
type GatewayConfigurationSet struct {
	node byte
	LANIPv6Address byte
	LANIPv6PrefixLength byte
	PortalIPv6Prefix byte
	PortalIPv6PrefixLength byte
	DefaultGatewayIPv6Address byte
	PANIPv6Prefix byte
}

func NewGatewayConfigurationSet() GatewayConfigurationSet {
	return GatewayConfigurationSet{}
}

func (c *GatewayConfigurationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *GatewayConfigurationSet) Set(LANIPv6Address byte,LANIPv6PrefixLength byte,PortalIPv6Prefix byte,PortalIPv6PrefixLength byte,DefaultGatewayIPv6Address byte,PANIPv6Prefix byte,) error {
	c.LANIPv6Address = LANIPv6Address
	c.LANIPv6PrefixLength = LANIPv6PrefixLength
	c.PortalIPv6Prefix = PortalIPv6Prefix
	c.PortalIPv6PrefixLength = PortalIPv6PrefixLength
	c.DefaultGatewayIPv6Address = DefaultGatewayIPv6Address
	c.PANIPv6Prefix = PANIPv6Prefix

	return nil
}

func (c *GatewayConfigurationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(8),
		byte(ZipPortal),
		0x01,
		c.LANIPv6Address,
		c.LANIPv6PrefixLength,
		c.PortalIPv6Prefix,
		c.PortalIPv6PrefixLength,
		c.DefaultGatewayIPv6Address,
		c.PANIPv6Prefix,
		0x25,
	}
}
type GatewayConfigurationStatus struct {
	node byte
	Status byte
}

func NewGatewayConfigurationStatus() GatewayConfigurationStatus {
	return GatewayConfigurationStatus{}
}

func (c *GatewayConfigurationStatus) SetNode(node int) {
	c.node = byte(node)
}

func (c *GatewayConfigurationStatus) Set(Status byte,) error {
	c.Status = Status

	return nil
}

func (c *GatewayConfigurationStatus) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ZipPortal),
		0x02,
		c.Status,
		0x25,
	}
}
type GatewayConfigurationGet struct {
	node byte
}

func NewGatewayConfigurationGet() GatewayConfigurationGet {
	return GatewayConfigurationGet{}
}

func (c *GatewayConfigurationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *GatewayConfigurationGet) Set() error {

	return nil
}

func (c *GatewayConfigurationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ZipPortal),
		0x03,
		0x25,
	}
}
type GatewayConfigurationReport struct {
    *report
    LANIPv6Address byte
    LANIPv6PrefixLength byte
    PortalIPv6Prefix byte
    PortalIPv6PrefixLength byte
    DefaultGatewayIPv6Address byte
    PANIPv6Prefix byte
    data []byte
}

func NewGatewayConfigurationReport(data []byte) *GatewayConfigurationReport {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &GatewayConfigurationReport{
				LANIPv6Address: data[0],
				LANIPv6PrefixLength: data[1],
				PortalIPv6Prefix: data[2],
				PortalIPv6PrefixLength: data[3],
				DefaultGatewayIPv6Address: data[4],
				PANIPv6Prefix: data[5],
        data: data,
    }
}

