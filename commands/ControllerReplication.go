package commands
type CtrlReplicationTransferGroup struct {
	node byte
	SequenceNumber byte
	GroupID byte
	NodeID byte
}

func NewCtrlReplicationTransferGroup() CtrlReplicationTransferGroup {
	return CtrlReplicationTransferGroup{}
}

func (c *CtrlReplicationTransferGroup) SetNode(node int) {
	c.node = byte(node)
}

func (c *CtrlReplicationTransferGroup) Set(SequenceNumber byte,GroupID byte,NodeID byte,) error {
	c.SequenceNumber = SequenceNumber
	c.GroupID = GroupID
	c.NodeID = NodeID

	return nil
}

func (c *CtrlReplicationTransferGroup) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ControllerReplication),
		0x31,
		c.SequenceNumber,
		c.GroupID,
		c.NodeID,
		0x25,
	}
}
type CtrlReplicationTransferGroupName struct {
	node byte
	SequenceNumber byte
	GroupID byte
	GroupName byte
}

func NewCtrlReplicationTransferGroupName() CtrlReplicationTransferGroupName {
	return CtrlReplicationTransferGroupName{}
}

func (c *CtrlReplicationTransferGroupName) SetNode(node int) {
	c.node = byte(node)
}

func (c *CtrlReplicationTransferGroupName) Set(SequenceNumber byte,GroupID byte,GroupName byte,) error {
	c.SequenceNumber = SequenceNumber
	c.GroupID = GroupID
	c.GroupName = GroupName

	return nil
}

func (c *CtrlReplicationTransferGroupName) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ControllerReplication),
		0x32,
		c.SequenceNumber,
		c.GroupID,
		c.GroupName,
		0x25,
	}
}
type CtrlReplicationTransferScene struct {
	node byte
	SequenceNumber byte
	SceneID byte
	NodeID byte
	Level byte
}

func NewCtrlReplicationTransferScene() CtrlReplicationTransferScene {
	return CtrlReplicationTransferScene{}
}

func (c *CtrlReplicationTransferScene) SetNode(node int) {
	c.node = byte(node)
}

func (c *CtrlReplicationTransferScene) Set(SequenceNumber byte,SceneID byte,NodeID byte,Level byte,) error {
	c.SequenceNumber = SequenceNumber
	c.SceneID = SceneID
	c.NodeID = NodeID
	c.Level = Level

	return nil
}

func (c *CtrlReplicationTransferScene) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(ControllerReplication),
		0x33,
		c.SequenceNumber,
		c.SceneID,
		c.NodeID,
		c.Level,
		0x25,
	}
}
type CtrlReplicationTransferSceneName struct {
	node byte
	SequenceNumber byte
	SceneID byte
	SceneName byte
}

func NewCtrlReplicationTransferSceneName() CtrlReplicationTransferSceneName {
	return CtrlReplicationTransferSceneName{}
}

func (c *CtrlReplicationTransferSceneName) SetNode(node int) {
	c.node = byte(node)
}

func (c *CtrlReplicationTransferSceneName) Set(SequenceNumber byte,SceneID byte,SceneName byte,) error {
	c.SequenceNumber = SequenceNumber
	c.SceneID = SceneID
	c.SceneName = SceneName

	return nil
}

func (c *CtrlReplicationTransferSceneName) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ControllerReplication),
		0x34,
		c.SequenceNumber,
		c.SceneID,
		c.SceneName,
		0x25,
	}
}
