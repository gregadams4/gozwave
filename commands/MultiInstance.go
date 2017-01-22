package commands
type MultiInstanceCmdEncap struct {
	node byte
	Instance byte
	CommandClass byte
	Command byte
	Parameter byte
}

func NewMultiInstanceCmdEncap() MultiInstanceCmdEncap {
	return MultiInstanceCmdEncap{}
}

func (c *MultiInstanceCmdEncap) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiInstanceCmdEncap) Set(Instance byte,CommandClass byte,Command byte,Parameter byte,) error {
	c.Instance = Instance
	c.CommandClass = CommandClass
	c.Command = Command
	c.Parameter = Parameter

	return nil
}

func (c *MultiInstanceCmdEncap) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(MultiInstance),
		0x06,
		c.Instance,
		c.CommandClass,
		c.Command,
		c.Parameter,
		0x25,
	}
}
type MultiInstanceGet struct {
	node byte
	CommandClass byte
}

func NewMultiInstanceGet() MultiInstanceGet {
	return MultiInstanceGet{}
}

func (c *MultiInstanceGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiInstanceGet) Set(CommandClass byte,) error {
	c.CommandClass = CommandClass

	return nil
}

func (c *MultiInstanceGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MultiInstance),
		0x04,
		c.CommandClass,
		0x25,
	}
}
type MultiInstanceReport struct {
    *report
    CommandClass byte
    Instances byte
    data []byte
}

func NewMultiInstanceReport(data []byte) *MultiInstanceReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiInstanceReport{
				CommandClass: data[0],
				Instances: data[1],
        data: data,
    }
}

