package commands
type MultiCmdEncap struct {
	node byte
	NumberofCommands byte
	Encapsulated_Command []MultiCmdEncapEncapsulated_Command
}
type MultiCmdEncapEncapsulated_Command struct {
	CommandLength byte
	CommandClass byte
	Command byte
	Data byte
}

func NewMultiCmdEncap() MultiCmdEncap {
	return MultiCmdEncap{}
}

func (c *MultiCmdEncap) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiCmdEncap) Set(NumberofCommands byte,Encapsulated_Command []MultiCmdEncapEncapsulated_Command,) error {
	c.NumberofCommands = NumberofCommands
	c.Encapsulated_Command = Encapsulated_Command

	return nil
}

func (c *MultiCmdEncap) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(7),
		byte(MultiCmd),
		0x01,
		c.NumberofCommands,
	}
	for _, v := range c.Encapsulated_Command {
		data = append(data, v.CommandLength)
		data = append(data, v.CommandClass)
		data = append(data, v.Command)
		data = append(data, v.Data)
	}
	data = append(data, 0x25)
	return data
	
}
