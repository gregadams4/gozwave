package commands
import "encoding/binary"
type DmxAddressSet struct {
	node byte
	Properties1 byte
	ChannelID byte
}

func NewDmxAddressSet() DmxAddressSet {
	return DmxAddressSet{}
}

func (c *DmxAddressSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DmxAddressSet) Set(Properties1 byte,ChannelID byte,) error {
	c.Properties1 = Properties1
	c.ChannelID = ChannelID

	return nil
}

func (c *DmxAddressSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(Dmx),
		0x01,
		c.Properties1,
		c.ChannelID,
		0x25,
	}
}
type DmxAddressGet struct {
	node byte
}

func NewDmxAddressGet() DmxAddressGet {
	return DmxAddressGet{}
}

func (c *DmxAddressGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DmxAddressGet) Set() error {

	return nil
}

func (c *DmxAddressGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Dmx),
		0x02,
		0x25,
	}
}
type DmxAddressReport struct {
    *report
    Properties1 byte
    ChannelID byte
    data []byte
}

func NewDmxAddressReport(data []byte) *DmxAddressReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &DmxAddressReport{
				Properties1: data[0],
				ChannelID: data[1],
        data: data,
    }
}

type DmxCapabilityGet struct {
	node byte
	ChannelID byte
}

func NewDmxCapabilityGet() DmxCapabilityGet {
	return DmxCapabilityGet{}
}

func (c *DmxCapabilityGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DmxCapabilityGet) Set(ChannelID byte,) error {
	c.ChannelID = ChannelID

	return nil
}

func (c *DmxCapabilityGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Dmx),
		0x04,
		c.ChannelID,
		0x25,
	}
}
type DmxCapabilityReport struct {
    *report
    ChannelID byte
    PropertyID uint16
    DeviceChannels byte
    MaxChannels byte
    data []byte
}

func NewDmxCapabilityReport(data []byte) *DmxCapabilityReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &DmxCapabilityReport{
				ChannelID: data[0],
				PropertyID: binary.BigEndian.Uint16(data[1:3]),
				DeviceChannels: data[3],
				MaxChannels: data[4],
        data: data,
    }
}

type DmxData struct {
	node byte
	Source byte
	Properties1 byte
	DMXchannel byte
}

func NewDmxData() DmxData {
	return DmxData{}
}

func (c *DmxData) SetNode(node int) {
	c.node = byte(node)
}

func (c *DmxData) Set(Source byte,Properties1 byte,DMXchannel byte,) error {
	c.Source = Source
	c.Properties1 = Properties1
	c.DMXchannel = DMXchannel

	return nil
}

func (c *DmxData) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(Dmx),
		0x06,
		c.Source,
		c.Properties1,
		c.DMXchannel,
		0x25,
	}
}
