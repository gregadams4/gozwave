package commands
import "encoding/binary"
type IpConfigurationGet struct {
	node byte
}

func NewIpConfigurationGet() IpConfigurationGet {
	return IpConfigurationGet{}
}

func (c *IpConfigurationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IpConfigurationGet) Set() error {

	return nil
}

func (c *IpConfigurationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(IpConfiguration),
		0x02,
		0x25,
	}
}
type IpConfigurationRelease struct {
	node byte
}

func NewIpConfigurationRelease() IpConfigurationRelease {
	return IpConfigurationRelease{}
}

func (c *IpConfigurationRelease) SetNode(node int) {
	c.node = byte(node)
}

func (c *IpConfigurationRelease) Set() error {

	return nil
}

func (c *IpConfigurationRelease) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(IpConfiguration),
		0x04,
		0x25,
	}
}
type IpConfigurationRenew struct {
	node byte
}

func NewIpConfigurationRenew() IpConfigurationRenew {
	return IpConfigurationRenew{}
}

func (c *IpConfigurationRenew) SetNode(node int) {
	c.node = byte(node)
}

func (c *IpConfigurationRenew) Set() error {

	return nil
}

func (c *IpConfigurationRenew) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(IpConfiguration),
		0x05,
		0x25,
	}
}
type IpConfigurationReport struct {
    *report
    Properties1 byte
    IPAddress uint64
    SubnetMask uint64
    Gateway uint64
    DNS1 uint64
    DNS2 uint64
    LeaseTime uint64
    data []byte
}

func NewIpConfigurationReport(data []byte) *IpConfigurationReport {
    if len(data) < 25 {
				//may want to change this to just return nil
				for i := len(data); i < 25; i++ {
            data = append(data, 0x00)
        }
    }

    return &IpConfigurationReport{
				Properties1: data[0],
				IPAddress: binary.BigEndian.Uint64(data[1:5]),
				SubnetMask: binary.BigEndian.Uint64(data[5:9]),
				Gateway: binary.BigEndian.Uint64(data[9:13]),
				DNS1: binary.BigEndian.Uint64(data[13:17]),
				DNS2: binary.BigEndian.Uint64(data[17:21]),
				LeaseTime: binary.BigEndian.Uint64(data[21:25]),
        data: data,
    }
}

type IpConfigurationSet struct {
	node byte
	Properties1 byte
	IPAddress uint64
	SubnetMask uint64
	Gateway uint64
	DNS1 uint64
	DNS2 uint64
}

func NewIpConfigurationSet() IpConfigurationSet {
	return IpConfigurationSet{}
}

func (c *IpConfigurationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IpConfigurationSet) Set(Properties1 byte,IPAddress uint64,SubnetMask uint64,Gateway uint64,DNS1 uint64,DNS2 uint64,) error {
	c.Properties1 = Properties1
	c.IPAddress = IPAddress
	c.SubnetMask = SubnetMask
	c.Gateway = Gateway
	c.DNS1 = DNS1
	c.DNS2 = DNS2

	return nil
}

func (c *IpConfigurationSet) Encode() []byte {
	IPAddressBytes := []byte{}
	binary.BigEndian.PutUint64(IPAddressBytes, c.IPAddress)
	SubnetMaskBytes := []byte{}
	binary.BigEndian.PutUint64(SubnetMaskBytes, c.SubnetMask)
	GatewayBytes := []byte{}
	binary.BigEndian.PutUint64(GatewayBytes, c.Gateway)
	DNS1Bytes := []byte{}
	binary.BigEndian.PutUint64(DNS1Bytes, c.DNS1)
	DNS2Bytes := []byte{}
	binary.BigEndian.PutUint64(DNS2Bytes, c.DNS2)
	return []byte{
		0x13,
		c.node,
		byte(8),
		byte(IpConfiguration),
		0x01,
		c.Properties1,
		IPAddressBytes[0],
		IPAddressBytes[1],
		IPAddressBytes[2],
		IPAddressBytes[3],
		SubnetMaskBytes[0],
		SubnetMaskBytes[1],
		SubnetMaskBytes[2],
		SubnetMaskBytes[3],
		GatewayBytes[0],
		GatewayBytes[1],
		GatewayBytes[2],
		GatewayBytes[3],
		DNS1Bytes[0],
		DNS1Bytes[1],
		DNS1Bytes[2],
		DNS1Bytes[3],
		DNS2Bytes[0],
		DNS2Bytes[1],
		DNS2Bytes[2],
		DNS2Bytes[3],
		0x25,
	}
}
