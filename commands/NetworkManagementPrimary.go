package commands
type ControllerChange struct {
	node byte
	SeqNo byte
	Reserved byte
	Mode byte
	txOptions byte
}

func NewControllerChange() ControllerChange {
	return ControllerChange{}
}

func (c *ControllerChange) SetNode(node int) {
	c.node = byte(node)
}

func (c *ControllerChange) Set(SeqNo byte,Reserved byte,Mode byte,txOptions byte,) error {
	c.SeqNo = SeqNo
	c.Reserved = Reserved
	c.Mode = Mode
	c.txOptions = txOptions

	return nil
}

func (c *ControllerChange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(NetworkManagementPrimary),
		0x01,
		c.SeqNo,
		c.Reserved,
		c.Mode,
		c.txOptions,
		0x25,
	}
}
type ControllerChangeStatus struct {
	node byte
	SeqNo byte
	Status byte
	Reserved byte
	NewNodeID byte
	NodeInfoLength byte
	Properties1 byte
	Properties2 byte
	BasicDeviceClass byte
	GenericDeviceClass byte
	SpecificDeviceClass byte
	CommandClass byte
}

func NewControllerChangeStatus() ControllerChangeStatus {
	return ControllerChangeStatus{}
}

func (c *ControllerChangeStatus) SetNode(node int) {
	c.node = byte(node)
}

func (c *ControllerChangeStatus) Set(SeqNo byte,Status byte,Reserved byte,NewNodeID byte,NodeInfoLength byte,Properties1 byte,Properties2 byte,BasicDeviceClass byte,GenericDeviceClass byte,SpecificDeviceClass byte,CommandClass byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status
	c.Reserved = Reserved
	c.NewNodeID = NewNodeID
	c.NodeInfoLength = NodeInfoLength
	c.Properties1 = Properties1
	c.Properties2 = Properties2
	c.BasicDeviceClass = BasicDeviceClass
	c.GenericDeviceClass = GenericDeviceClass
	c.SpecificDeviceClass = SpecificDeviceClass
	c.CommandClass = CommandClass

	return nil
}

func (c *ControllerChangeStatus) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(13),
		byte(NetworkManagementPrimary),
		0x02,
		c.SeqNo,
		c.Status,
		c.Reserved,
		c.NewNodeID,
		c.NodeInfoLength,
		c.Properties1,
		c.Properties2,
		c.BasicDeviceClass,
		c.GenericDeviceClass,
		c.SpecificDeviceClass,
		c.CommandClass,
		0x25,
	}
}
