package commands
type MultiChannelCapabilityGetV4 struct {
	node byte
	Properties1 byte
}

func NewMultiChannelCapabilityGetV4() MultiChannelCapabilityGetV4 {
	return MultiChannelCapabilityGetV4{}
}

func (c *MultiChannelCapabilityGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelCapabilityGetV4) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *MultiChannelCapabilityGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MultiChannelV4),
		0x09,
		c.Properties1,
		0x25,
	}
}
type MultiChannelCapabilityReportV4 struct {
    *report
    Properties1 byte
    GenericDeviceClass byte
    SpecificDeviceClass byte
    CommandClass byte
    data []byte
}

func NewMultiChannelCapabilityReportV4(data []byte) *MultiChannelCapabilityReportV4 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelCapabilityReportV4{
				Properties1: data[0],
				GenericDeviceClass: data[1],
				SpecificDeviceClass: data[2],
				CommandClass: data[3],
        data: data,
    }
}

type MultiChannelCmdEncapV4 struct {
	node byte
	Properties1 byte
	Properties2 byte
	CommandClass byte
	Command byte
	Parameter byte
}

func NewMultiChannelCmdEncapV4() MultiChannelCmdEncapV4 {
	return MultiChannelCmdEncapV4{}
}

func (c *MultiChannelCmdEncapV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelCmdEncapV4) Set(Properties1 byte,Properties2 byte,CommandClass byte,Command byte,Parameter byte,) error {
	c.Properties1 = Properties1
	c.Properties2 = Properties2
	c.CommandClass = CommandClass
	c.Command = Command
	c.Parameter = Parameter

	return nil
}

func (c *MultiChannelCmdEncapV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(MultiChannelV4),
		0x0D,
		c.Properties1,
		c.Properties2,
		c.CommandClass,
		c.Command,
		c.Parameter,
		0x25,
	}
}
type MultiChannelEndPointFindV4 struct {
	node byte
	GenericDeviceClass byte
	SpecificDeviceClass byte
}

func NewMultiChannelEndPointFindV4() MultiChannelEndPointFindV4 {
	return MultiChannelEndPointFindV4{}
}

func (c *MultiChannelEndPointFindV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelEndPointFindV4) Set(GenericDeviceClass byte,SpecificDeviceClass byte,) error {
	c.GenericDeviceClass = GenericDeviceClass
	c.SpecificDeviceClass = SpecificDeviceClass

	return nil
}

func (c *MultiChannelEndPointFindV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(MultiChannelV4),
		0x0B,
		c.GenericDeviceClass,
		c.SpecificDeviceClass,
		0x25,
	}
}
type MultiChannelEndPointFindReportV4 struct {
    *report
    ReportstoFollow byte
    GenericDeviceClass byte
    SpecificDeviceClass byte
    data []byte
}

func NewMultiChannelEndPointFindReportV4(data []byte) *MultiChannelEndPointFindReportV4 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelEndPointFindReportV4{
				ReportstoFollow: data[0],
				GenericDeviceClass: data[1],
				SpecificDeviceClass: data[2],
        data: data,
    }
}

type MultiChannelEndPointGetV4 struct {
	node byte
}

func NewMultiChannelEndPointGetV4() MultiChannelEndPointGetV4 {
	return MultiChannelEndPointGetV4{}
}

func (c *MultiChannelEndPointGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelEndPointGetV4) Set() error {

	return nil
}

func (c *MultiChannelEndPointGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MultiChannelV4),
		0x07,
		0x25,
	}
}
type MultiChannelEndPointReportV4 struct {
    *report
    Properties1 byte
    Properties2 byte
    Properties3 byte
    data []byte
}

func NewMultiChannelEndPointReportV4(data []byte) *MultiChannelEndPointReportV4 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelEndPointReportV4{
				Properties1: data[0],
				Properties2: data[1],
				Properties3: data[2],
        data: data,
    }
}

type MultiInstanceCmdEncapV4 struct {
	node byte
	Properties1 byte
	CommandClass byte
	Command byte
	Parameter byte
}

func NewMultiInstanceCmdEncapV4() MultiInstanceCmdEncapV4 {
	return MultiInstanceCmdEncapV4{}
}

func (c *MultiInstanceCmdEncapV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiInstanceCmdEncapV4) Set(Properties1 byte,CommandClass byte,Command byte,Parameter byte,) error {
	c.Properties1 = Properties1
	c.CommandClass = CommandClass
	c.Command = Command
	c.Parameter = Parameter

	return nil
}

func (c *MultiInstanceCmdEncapV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(MultiChannelV4),
		0x06,
		c.Properties1,
		c.CommandClass,
		c.Command,
		c.Parameter,
		0x25,
	}
}
type MultiInstanceGetV4 struct {
	node byte
	CommandClass byte
}

func NewMultiInstanceGetV4() MultiInstanceGetV4 {
	return MultiInstanceGetV4{}
}

func (c *MultiInstanceGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiInstanceGetV4) Set(CommandClass byte,) error {
	c.CommandClass = CommandClass

	return nil
}

func (c *MultiInstanceGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MultiChannelV4),
		0x04,
		c.CommandClass,
		0x25,
	}
}
type MultiInstanceReportV4 struct {
    *report
    CommandClass byte
    Properties1 byte
    data []byte
}

func NewMultiInstanceReportV4(data []byte) *MultiInstanceReportV4 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiInstanceReportV4{
				CommandClass: data[0],
				Properties1: data[1],
        data: data,
    }
}

type MultiChannelAggregatedMembersGetV4 struct {
	node byte
	Properties1 byte
}

func NewMultiChannelAggregatedMembersGetV4() MultiChannelAggregatedMembersGetV4 {
	return MultiChannelAggregatedMembersGetV4{}
}

func (c *MultiChannelAggregatedMembersGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelAggregatedMembersGetV4) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *MultiChannelAggregatedMembersGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MultiChannelV4),
		0x0E,
		c.Properties1,
		0x25,
	}
}
type MultiChannelAggregatedMembersReportV4 struct {
    *report
    Properties1 byte
    NumberofBitMasks byte
    AggregatedMembersBitMask byte
    data []byte
}

func NewMultiChannelAggregatedMembersReportV4(data []byte) *MultiChannelAggregatedMembersReportV4 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelAggregatedMembersReportV4{
				Properties1: data[0],
				NumberofBitMasks: data[1],
				AggregatedMembersBitMask: data[2],
        data: data,
    }
}

