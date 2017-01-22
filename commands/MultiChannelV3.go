package commands
type MultiChannelCapabilityGetV3 struct {
	node byte
	Properties1 byte
}

func NewMultiChannelCapabilityGetV3() MultiChannelCapabilityGetV3 {
	return MultiChannelCapabilityGetV3{}
}

func (c *MultiChannelCapabilityGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelCapabilityGetV3) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *MultiChannelCapabilityGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MultiChannelV3),
		0x09,
		c.Properties1,
		0x25,
	}
}
type MultiChannelCapabilityReportV3 struct {
    *report
    Properties1 byte
    GenericDeviceClass byte
    SpecificDeviceClass byte
    CommandClass byte
    data []byte
}

func NewMultiChannelCapabilityReportV3(data []byte) *MultiChannelCapabilityReportV3 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelCapabilityReportV3{
				Properties1: data[0],
				GenericDeviceClass: data[1],
				SpecificDeviceClass: data[2],
				CommandClass: data[3],
        data: data,
    }
}

type MultiChannelCmdEncapV3 struct {
	node byte
	Properties1 byte
	Properties2 byte
	CommandClass byte
	Command byte
	Parameter byte
}

func NewMultiChannelCmdEncapV3() MultiChannelCmdEncapV3 {
	return MultiChannelCmdEncapV3{}
}

func (c *MultiChannelCmdEncapV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelCmdEncapV3) Set(Properties1 byte,Properties2 byte,CommandClass byte,Command byte,Parameter byte,) error {
	c.Properties1 = Properties1
	c.Properties2 = Properties2
	c.CommandClass = CommandClass
	c.Command = Command
	c.Parameter = Parameter

	return nil
}

func (c *MultiChannelCmdEncapV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(MultiChannelV3),
		0x0D,
		c.Properties1,
		c.Properties2,
		c.CommandClass,
		c.Command,
		c.Parameter,
		0x25,
	}
}
type MultiChannelEndPointFindV3 struct {
	node byte
	GenericDeviceClass byte
	SpecificDeviceClass byte
}

func NewMultiChannelEndPointFindV3() MultiChannelEndPointFindV3 {
	return MultiChannelEndPointFindV3{}
}

func (c *MultiChannelEndPointFindV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelEndPointFindV3) Set(GenericDeviceClass byte,SpecificDeviceClass byte,) error {
	c.GenericDeviceClass = GenericDeviceClass
	c.SpecificDeviceClass = SpecificDeviceClass

	return nil
}

func (c *MultiChannelEndPointFindV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(MultiChannelV3),
		0x0B,
		c.GenericDeviceClass,
		c.SpecificDeviceClass,
		0x25,
	}
}
type MultiChannelEndPointFindReportV3 struct {
    *report
    ReportstoFollow byte
    GenericDeviceClass byte
    SpecificDeviceClass byte
    data []byte
}

func NewMultiChannelEndPointFindReportV3(data []byte) *MultiChannelEndPointFindReportV3 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelEndPointFindReportV3{
				ReportstoFollow: data[0],
				GenericDeviceClass: data[1],
				SpecificDeviceClass: data[2],
        data: data,
    }
}

type MultiChannelEndPointGetV3 struct {
	node byte
}

func NewMultiChannelEndPointGetV3() MultiChannelEndPointGetV3 {
	return MultiChannelEndPointGetV3{}
}

func (c *MultiChannelEndPointGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelEndPointGetV3) Set() error {

	return nil
}

func (c *MultiChannelEndPointGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MultiChannelV3),
		0x07,
		0x25,
	}
}
type MultiChannelEndPointReportV3 struct {
    *report
    Properties1 byte
    Properties2 byte
    data []byte
}

func NewMultiChannelEndPointReportV3(data []byte) *MultiChannelEndPointReportV3 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelEndPointReportV3{
				Properties1: data[0],
				Properties2: data[1],
        data: data,
    }
}

type MultiInstanceCmdEncapV3 struct {
	node byte
	Properties1 byte
	CommandClass byte
	Command byte
	Parameter byte
}

func NewMultiInstanceCmdEncapV3() MultiInstanceCmdEncapV3 {
	return MultiInstanceCmdEncapV3{}
}

func (c *MultiInstanceCmdEncapV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiInstanceCmdEncapV3) Set(Properties1 byte,CommandClass byte,Command byte,Parameter byte,) error {
	c.Properties1 = Properties1
	c.CommandClass = CommandClass
	c.Command = Command
	c.Parameter = Parameter

	return nil
}

func (c *MultiInstanceCmdEncapV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(MultiChannelV3),
		0x06,
		c.Properties1,
		c.CommandClass,
		c.Command,
		c.Parameter,
		0x25,
	}
}
type MultiInstanceGetV3 struct {
	node byte
	CommandClass byte
}

func NewMultiInstanceGetV3() MultiInstanceGetV3 {
	return MultiInstanceGetV3{}
}

func (c *MultiInstanceGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiInstanceGetV3) Set(CommandClass byte,) error {
	c.CommandClass = CommandClass

	return nil
}

func (c *MultiInstanceGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MultiChannelV3),
		0x04,
		c.CommandClass,
		0x25,
	}
}
type MultiInstanceReportV3 struct {
    *report
    CommandClass byte
    Properties1 byte
    data []byte
}

func NewMultiInstanceReportV3(data []byte) *MultiInstanceReportV3 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiInstanceReportV3{
				CommandClass: data[0],
				Properties1: data[1],
        data: data,
    }
}

