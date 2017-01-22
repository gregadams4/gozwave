package commands
type MultiChannelCapabilityGetV2 struct {
	node byte
	Properties1 byte
}

func NewMultiChannelCapabilityGetV2() MultiChannelCapabilityGetV2 {
	return MultiChannelCapabilityGetV2{}
}

func (c *MultiChannelCapabilityGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelCapabilityGetV2) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *MultiChannelCapabilityGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MultiChannelV2),
		0x09,
		c.Properties1,
		0x25,
	}
}
type MultiChannelCapabilityReportV2 struct {
    *report
    Properties1 byte
    GenericDeviceClass byte
    SpecificDeviceClass byte
    CommandClass byte
    data []byte
}

func NewMultiChannelCapabilityReportV2(data []byte) *MultiChannelCapabilityReportV2 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelCapabilityReportV2{
				Properties1: data[0],
				GenericDeviceClass: data[1],
				SpecificDeviceClass: data[2],
				CommandClass: data[3],
        data: data,
    }
}

type MultiChannelCmdEncapV2 struct {
	node byte
	Properties1 byte
	Properties2 byte
	CommandClass byte
	Command byte
	Parameter byte
}

func NewMultiChannelCmdEncapV2() MultiChannelCmdEncapV2 {
	return MultiChannelCmdEncapV2{}
}

func (c *MultiChannelCmdEncapV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelCmdEncapV2) Set(Properties1 byte,Properties2 byte,CommandClass byte,Command byte,Parameter byte,) error {
	c.Properties1 = Properties1
	c.Properties2 = Properties2
	c.CommandClass = CommandClass
	c.Command = Command
	c.Parameter = Parameter

	return nil
}

func (c *MultiChannelCmdEncapV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(MultiChannelV2),
		0x0D,
		c.Properties1,
		c.Properties2,
		c.CommandClass,
		c.Command,
		c.Parameter,
		0x25,
	}
}
type MultiChannelEndPointFindV2 struct {
	node byte
	GenericDeviceClass byte
	SpecificDeviceClass byte
}

func NewMultiChannelEndPointFindV2() MultiChannelEndPointFindV2 {
	return MultiChannelEndPointFindV2{}
}

func (c *MultiChannelEndPointFindV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelEndPointFindV2) Set(GenericDeviceClass byte,SpecificDeviceClass byte,) error {
	c.GenericDeviceClass = GenericDeviceClass
	c.SpecificDeviceClass = SpecificDeviceClass

	return nil
}

func (c *MultiChannelEndPointFindV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(MultiChannelV2),
		0x0B,
		c.GenericDeviceClass,
		c.SpecificDeviceClass,
		0x25,
	}
}
type MultiChannelEndPointFindReportV2 struct {
    *report
    ReportstoFollow byte
    GenericDeviceClass byte
    SpecificDeviceClass byte
    data []byte
}

func NewMultiChannelEndPointFindReportV2(data []byte) *MultiChannelEndPointFindReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelEndPointFindReportV2{
				ReportstoFollow: data[0],
				GenericDeviceClass: data[1],
				SpecificDeviceClass: data[2],
        data: data,
    }
}

type MultiChannelEndPointGetV2 struct {
	node byte
}

func NewMultiChannelEndPointGetV2() MultiChannelEndPointGetV2 {
	return MultiChannelEndPointGetV2{}
}

func (c *MultiChannelEndPointGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelEndPointGetV2) Set() error {

	return nil
}

func (c *MultiChannelEndPointGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MultiChannelV2),
		0x07,
		0x25,
	}
}
type MultiChannelEndPointReportV2 struct {
    *report
    Properties1 byte
    Properties2 byte
    data []byte
}

func NewMultiChannelEndPointReportV2(data []byte) *MultiChannelEndPointReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelEndPointReportV2{
				Properties1: data[0],
				Properties2: data[1],
        data: data,
    }
}

type MultiInstanceCmdEncapV2 struct {
	node byte
	Properties1 byte
	CommandClass byte
	Command byte
	Parameter byte
}

func NewMultiInstanceCmdEncapV2() MultiInstanceCmdEncapV2 {
	return MultiInstanceCmdEncapV2{}
}

func (c *MultiInstanceCmdEncapV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiInstanceCmdEncapV2) Set(Properties1 byte,CommandClass byte,Command byte,Parameter byte,) error {
	c.Properties1 = Properties1
	c.CommandClass = CommandClass
	c.Command = Command
	c.Parameter = Parameter

	return nil
}

func (c *MultiInstanceCmdEncapV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(MultiChannelV2),
		0x06,
		c.Properties1,
		c.CommandClass,
		c.Command,
		c.Parameter,
		0x25,
	}
}
type MultiInstanceGetV2 struct {
	node byte
	CommandClass byte
}

func NewMultiInstanceGetV2() MultiInstanceGetV2 {
	return MultiInstanceGetV2{}
}

func (c *MultiInstanceGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiInstanceGetV2) Set(CommandClass byte,) error {
	c.CommandClass = CommandClass

	return nil
}

func (c *MultiInstanceGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MultiChannelV2),
		0x04,
		c.CommandClass,
		0x25,
	}
}
type MultiInstanceReportV2 struct {
    *report
    CommandClass byte
    Properties1 byte
    data []byte
}

func NewMultiInstanceReportV2(data []byte) *MultiInstanceReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiInstanceReportV2{
				CommandClass: data[0],
				Properties1: data[1],
        data: data,
    }
}

