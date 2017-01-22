package commands
import "encoding/binary"
type GatewayModeSet struct {
	node byte
	Mode byte
}

func NewGatewayModeSet() GatewayModeSet {
	return GatewayModeSet{}
}

func (c *GatewayModeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *GatewayModeSet) Set(Mode byte,) error {
	c.Mode = Mode

	return nil
}

func (c *GatewayModeSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ZipGateway),
		0x01,
		c.Mode,
		0x25,
	}
}
type GatewayModeGet struct {
	node byte
}

func NewGatewayModeGet() GatewayModeGet {
	return GatewayModeGet{}
}

func (c *GatewayModeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *GatewayModeGet) Set() error {

	return nil
}

func (c *GatewayModeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ZipGateway),
		0x02,
		0x25,
	}
}
type GatewayModeReport struct {
    *report
    Mode byte
    data []byte
}

func NewGatewayModeReport(data []byte) *GatewayModeReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &GatewayModeReport{
				Mode: data[0],
        data: data,
    }
}

type GatewayPeerSet struct {
	node byte
	PeerProfile byte
	IPv6Address byte
	Port uint16
	Properties1 byte
	PeerName byte
}

func NewGatewayPeerSet() GatewayPeerSet {
	return GatewayPeerSet{}
}

func (c *GatewayPeerSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *GatewayPeerSet) Set(PeerProfile byte,IPv6Address byte,Port uint16,Properties1 byte,PeerName byte,) error {
	c.PeerProfile = PeerProfile
	c.IPv6Address = IPv6Address
	c.Port = Port
	c.Properties1 = Properties1
	c.PeerName = PeerName

	return nil
}

func (c *GatewayPeerSet) Encode() []byte {
	PortBytes := []byte{}
	binary.BigEndian.PutUint16(PortBytes, c.Port)
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(ZipGateway),
		0x04,
		c.PeerProfile,
		c.IPv6Address,
		PortBytes[0],
		PortBytes[1],
		c.Properties1,
		c.PeerName,
		0x25,
	}
}
type GatewayPeerGet struct {
	node byte
	PeerProfile byte
}

func NewGatewayPeerGet() GatewayPeerGet {
	return GatewayPeerGet{}
}

func (c *GatewayPeerGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *GatewayPeerGet) Set(PeerProfile byte,) error {
	c.PeerProfile = PeerProfile

	return nil
}

func (c *GatewayPeerGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ZipGateway),
		0x05,
		c.PeerProfile,
		0x25,
	}
}
type GatewayPeerReport struct {
    *report
    PeerProfile byte
    PeerCount byte
    IPv6Address byte
    Port uint16
    Properties1 byte
    PeerName byte
    data []byte
}

func NewGatewayPeerReport(data []byte) *GatewayPeerReport {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &GatewayPeerReport{
				PeerProfile: data[0],
				PeerCount: data[1],
				IPv6Address: data[2],
				Port: binary.BigEndian.Uint16(data[3:5]),
				Properties1: data[5],
				PeerName: data[6],
        data: data,
    }
}

type GatewayLockSet struct {
	node byte
	Properties1 byte
}

func NewGatewayLockSet() GatewayLockSet {
	return GatewayLockSet{}
}

func (c *GatewayLockSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *GatewayLockSet) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *GatewayLockSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ZipGateway),
		0x07,
		c.Properties1,
		0x25,
	}
}
type UnsolicitedDestinationSet struct {
	node byte
	UnsolicitedIPv6Destination byte
	UnsolicitedDestinationPort uint16
}

func NewUnsolicitedDestinationSet() UnsolicitedDestinationSet {
	return UnsolicitedDestinationSet{}
}

func (c *UnsolicitedDestinationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *UnsolicitedDestinationSet) Set(UnsolicitedIPv6Destination byte,UnsolicitedDestinationPort uint16,) error {
	c.UnsolicitedIPv6Destination = UnsolicitedIPv6Destination
	c.UnsolicitedDestinationPort = UnsolicitedDestinationPort

	return nil
}

func (c *UnsolicitedDestinationSet) Encode() []byte {
	UnsolicitedDestinationPortBytes := []byte{}
	binary.BigEndian.PutUint16(UnsolicitedDestinationPortBytes, c.UnsolicitedDestinationPort)
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ZipGateway),
		0x08,
		c.UnsolicitedIPv6Destination,
		UnsolicitedDestinationPortBytes[0],
		UnsolicitedDestinationPortBytes[1],
		0x25,
	}
}
type UnsolicitedDestinationGet struct {
	node byte
}

func NewUnsolicitedDestinationGet() UnsolicitedDestinationGet {
	return UnsolicitedDestinationGet{}
}

func (c *UnsolicitedDestinationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *UnsolicitedDestinationGet) Set() error {

	return nil
}

func (c *UnsolicitedDestinationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ZipGateway),
		0x09,
		0x25,
	}
}
type UnsolicitedDestinationReport struct {
    *report
    UnsolicitedIPv6Destination byte
    UnsolicitedDestinationPort uint16
    data []byte
}

func NewUnsolicitedDestinationReport(data []byte) *UnsolicitedDestinationReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &UnsolicitedDestinationReport{
				UnsolicitedIPv6Destination: data[0],
				UnsolicitedDestinationPort: binary.BigEndian.Uint16(data[1:3]),
        data: data,
    }
}

type CommandApplicationNodeInfoSet struct {
	node byte
	NonSecureCommandClass byte
	SecurityScheme0MARK byte
	SecurityScheme0CommandClass byte
}

func NewCommandApplicationNodeInfoSet() CommandApplicationNodeInfoSet {
	return CommandApplicationNodeInfoSet{}
}

func (c *CommandApplicationNodeInfoSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandApplicationNodeInfoSet) Set(NonSecureCommandClass byte,SecurityScheme0MARK byte,SecurityScheme0CommandClass byte,) error {
	c.NonSecureCommandClass = NonSecureCommandClass
	c.SecurityScheme0MARK = SecurityScheme0MARK
	c.SecurityScheme0CommandClass = SecurityScheme0CommandClass

	return nil
}

func (c *CommandApplicationNodeInfoSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ZipGateway),
		0x0B,
		c.NonSecureCommandClass,
		c.SecurityScheme0MARK,
		c.SecurityScheme0CommandClass,
		0x25,
	}
}
type CommandApplicationNodeInfoGet struct {
	node byte
}

func NewCommandApplicationNodeInfoGet() CommandApplicationNodeInfoGet {
	return CommandApplicationNodeInfoGet{}
}

func (c *CommandApplicationNodeInfoGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandApplicationNodeInfoGet) Set() error {

	return nil
}

func (c *CommandApplicationNodeInfoGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ZipGateway),
		0x0C,
		0x25,
	}
}
type CommandApplicationNodeInfoReport struct {
    *report
    NonSecureCommandClass byte
    SecurityScheme0MARK byte
    SecurityScheme0CommandClass byte
    data []byte
}

func NewCommandApplicationNodeInfoReport(data []byte) *CommandApplicationNodeInfoReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &CommandApplicationNodeInfoReport{
				NonSecureCommandClass: data[0],
				SecurityScheme0MARK: data[1],
				SecurityScheme0CommandClass: data[2],
        data: data,
    }
}

