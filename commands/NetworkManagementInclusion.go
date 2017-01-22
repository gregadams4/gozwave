package commands
type FailedNodeRemove struct {
	node byte
	SeqNo byte
	NodeID byte
}

func NewFailedNodeRemove() FailedNodeRemove {
	return FailedNodeRemove{}
}

func (c *FailedNodeRemove) SetNode(node int) {
	c.node = byte(node)
}

func (c *FailedNodeRemove) Set(SeqNo byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.NodeID = NodeID

	return nil
}

func (c *FailedNodeRemove) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementInclusion),
		0x07,
		c.SeqNo,
		c.NodeID,
		0x25,
	}
}
type FailedNodeRemoveStatus struct {
	node byte
	SeqNo byte
	Status byte
	NodeID byte
}

func NewFailedNodeRemoveStatus() FailedNodeRemoveStatus {
	return FailedNodeRemoveStatus{}
}

func (c *FailedNodeRemoveStatus) SetNode(node int) {
	c.node = byte(node)
}

func (c *FailedNodeRemoveStatus) Set(SeqNo byte,Status byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status
	c.NodeID = NodeID

	return nil
}

func (c *FailedNodeRemoveStatus) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementInclusion),
		0x08,
		c.SeqNo,
		c.Status,
		c.NodeID,
		0x25,
	}
}
type NodeAdd struct {
	node byte
	SeqNo byte
	Reserved byte
	Mode byte
	txOptions byte
}

func NewNodeAdd() NodeAdd {
	return NodeAdd{}
}

func (c *NodeAdd) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeAdd) Set(SeqNo byte,Reserved byte,Mode byte,txOptions byte,) error {
	c.SeqNo = SeqNo
	c.Reserved = Reserved
	c.Mode = Mode
	c.txOptions = txOptions

	return nil
}

func (c *NodeAdd) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(NetworkManagementInclusion),
		0x01,
		c.SeqNo,
		c.Reserved,
		c.Mode,
		c.txOptions,
		0x25,
	}
}
type NodeAddStatus struct {
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

func NewNodeAddStatus() NodeAddStatus {
	return NodeAddStatus{}
}

func (c *NodeAddStatus) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeAddStatus) Set(SeqNo byte,Status byte,Reserved byte,NewNodeID byte,NodeInfoLength byte,Properties1 byte,Properties2 byte,BasicDeviceClass byte,GenericDeviceClass byte,SpecificDeviceClass byte,CommandClass byte,) error {
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

func (c *NodeAddStatus) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(13),
		byte(NetworkManagementInclusion),
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
type NodeRemove struct {
	node byte
	SeqNo byte
	Reserved byte
	Mode byte
}

func NewNodeRemove() NodeRemove {
	return NodeRemove{}
}

func (c *NodeRemove) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeRemove) Set(SeqNo byte,Reserved byte,Mode byte,) error {
	c.SeqNo = SeqNo
	c.Reserved = Reserved
	c.Mode = Mode

	return nil
}

func (c *NodeRemove) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementInclusion),
		0x03,
		c.SeqNo,
		c.Reserved,
		c.Mode,
		0x25,
	}
}
type NodeRemoveStatus struct {
	node byte
	SeqNo byte
	Status byte
	NodeID byte
}

func NewNodeRemoveStatus() NodeRemoveStatus {
	return NodeRemoveStatus{}
}

func (c *NodeRemoveStatus) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeRemoveStatus) Set(SeqNo byte,Status byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status
	c.NodeID = NodeID

	return nil
}

func (c *NodeRemoveStatus) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementInclusion),
		0x04,
		c.SeqNo,
		c.Status,
		c.NodeID,
		0x25,
	}
}
type FailedNodeReplace struct {
	node byte
	SeqNo byte
	NodeID byte
	txOptions byte
	Mode byte
}

func NewFailedNodeReplace() FailedNodeReplace {
	return FailedNodeReplace{}
}

func (c *FailedNodeReplace) SetNode(node int) {
	c.node = byte(node)
}

func (c *FailedNodeReplace) Set(SeqNo byte,NodeID byte,txOptions byte,Mode byte,) error {
	c.SeqNo = SeqNo
	c.NodeID = NodeID
	c.txOptions = txOptions
	c.Mode = Mode

	return nil
}

func (c *FailedNodeReplace) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(NetworkManagementInclusion),
		0x09,
		c.SeqNo,
		c.NodeID,
		c.txOptions,
		c.Mode,
		0x25,
	}
}
type FailedNodeReplaceStatus struct {
	node byte
	SeqNo byte
	Status byte
	NodeID byte
}

func NewFailedNodeReplaceStatus() FailedNodeReplaceStatus {
	return FailedNodeReplaceStatus{}
}

func (c *FailedNodeReplaceStatus) SetNode(node int) {
	c.node = byte(node)
}

func (c *FailedNodeReplaceStatus) Set(SeqNo byte,Status byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status
	c.NodeID = NodeID

	return nil
}

func (c *FailedNodeReplaceStatus) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementInclusion),
		0x0A,
		c.SeqNo,
		c.Status,
		c.NodeID,
		0x25,
	}
}
type NodeNeighborUpdateRequest struct {
	node byte
	SeqNo byte
	NodeID byte
}

func NewNodeNeighborUpdateRequest() NodeNeighborUpdateRequest {
	return NodeNeighborUpdateRequest{}
}

func (c *NodeNeighborUpdateRequest) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeNeighborUpdateRequest) Set(SeqNo byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.NodeID = NodeID

	return nil
}

func (c *NodeNeighborUpdateRequest) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementInclusion),
		0x0B,
		c.SeqNo,
		c.NodeID,
		0x25,
	}
}
type NodeNeighborUpdateStatus struct {
	node byte
	SeqNo byte
	Status byte
}

func NewNodeNeighborUpdateStatus() NodeNeighborUpdateStatus {
	return NodeNeighborUpdateStatus{}
}

func (c *NodeNeighborUpdateStatus) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeNeighborUpdateStatus) Set(SeqNo byte,Status byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status

	return nil
}

func (c *NodeNeighborUpdateStatus) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementInclusion),
		0x0C,
		c.SeqNo,
		c.Status,
		0x25,
	}
}
type ReturnRouteAssign struct {
	node byte
	SeqNo byte
	SourceNodeID byte
	DestinationNodeID byte
}

func NewReturnRouteAssign() ReturnRouteAssign {
	return ReturnRouteAssign{}
}

func (c *ReturnRouteAssign) SetNode(node int) {
	c.node = byte(node)
}

func (c *ReturnRouteAssign) Set(SeqNo byte,SourceNodeID byte,DestinationNodeID byte,) error {
	c.SeqNo = SeqNo
	c.SourceNodeID = SourceNodeID
	c.DestinationNodeID = DestinationNodeID

	return nil
}

func (c *ReturnRouteAssign) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementInclusion),
		0x0D,
		c.SeqNo,
		c.SourceNodeID,
		c.DestinationNodeID,
		0x25,
	}
}
type ReturnRouteAssignComplete struct {
	node byte
	SeqNo byte
	Status byte
}

func NewReturnRouteAssignComplete() ReturnRouteAssignComplete {
	return ReturnRouteAssignComplete{}
}

func (c *ReturnRouteAssignComplete) SetNode(node int) {
	c.node = byte(node)
}

func (c *ReturnRouteAssignComplete) Set(SeqNo byte,Status byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status

	return nil
}

func (c *ReturnRouteAssignComplete) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementInclusion),
		0x0E,
		c.SeqNo,
		c.Status,
		0x25,
	}
}
type ReturnRouteDelete struct {
	node byte
	SeqNo byte
	NodeID byte
}

func NewReturnRouteDelete() ReturnRouteDelete {
	return ReturnRouteDelete{}
}

func (c *ReturnRouteDelete) SetNode(node int) {
	c.node = byte(node)
}

func (c *ReturnRouteDelete) Set(SeqNo byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.NodeID = NodeID

	return nil
}

func (c *ReturnRouteDelete) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementInclusion),
		0x0F,
		c.SeqNo,
		c.NodeID,
		0x25,
	}
}
type ReturnRouteDeleteComplete struct {
	node byte
	SeqNo byte
	Status byte
}

func NewReturnRouteDeleteComplete() ReturnRouteDeleteComplete {
	return ReturnRouteDeleteComplete{}
}

func (c *ReturnRouteDeleteComplete) SetNode(node int) {
	c.node = byte(node)
}

func (c *ReturnRouteDeleteComplete) Set(SeqNo byte,Status byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status

	return nil
}

func (c *ReturnRouteDeleteComplete) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementInclusion),
		0x10,
		c.SeqNo,
		c.Status,
		0x25,
	}
}
