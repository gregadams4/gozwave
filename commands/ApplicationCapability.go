package commands
type CommandCommandClassNotSupported struct {
	node byte
	Properties1 byte
	OffendingCommandClass byte
	OffendingCommand byte
}

func NewCommandCommandClassNotSupported() CommandCommandClassNotSupported {
	return CommandCommandClassNotSupported{}
}

func (c *CommandCommandClassNotSupported) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandCommandClassNotSupported) Set(Properties1 byte,OffendingCommandClass byte,OffendingCommand byte,) error {
	c.Properties1 = Properties1
	c.OffendingCommandClass = OffendingCommandClass
	c.OffendingCommand = OffendingCommand

	return nil
}

func (c *CommandCommandClassNotSupported) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ApplicationCapability),
		0x01,
		c.Properties1,
		c.OffendingCommandClass,
		c.OffendingCommand,
		0x25,
	}
}
