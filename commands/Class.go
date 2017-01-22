package commands
type AcceptLost struct {
	node byte
}

func NewAcceptLost() AcceptLost {
	return AcceptLost{}
}

func (c *AcceptLost) SetNode(node int) {
	c.node = byte(node)
}

func (c *AcceptLost) Set() error {

	return nil
}

func (c *AcceptLost) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x17,
		0x25,
	}
}
type AssignId struct {
	node byte
}

func NewAssignId() AssignId {
	return AssignId{}
}

func (c *AssignId) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssignId) Set() error {

	return nil
}

func (c *AssignId) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x03,
		0x25,
	}
}
type AssignReturnRoute struct {
	node byte
}

func NewAssignReturnRoute() AssignReturnRoute {
	return AssignReturnRoute{}
}

func (c *AssignReturnRoute) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssignReturnRoute) Set() error {

	return nil
}

func (c *AssignReturnRoute) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x0C,
		0x25,
	}
}
type CmdAssignSucReturnRoute struct {
	node byte
}

func NewCmdAssignSucReturnRoute() CmdAssignSucReturnRoute {
	return CmdAssignSucReturnRoute{}
}

func (c *CmdAssignSucReturnRoute) SetNode(node int) {
	c.node = byte(node)
}

func (c *CmdAssignSucReturnRoute) Set() error {

	return nil
}

func (c *CmdAssignSucReturnRoute) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x14,
		0x25,
	}
}
type CmdAutomaticControllerUpdateStart struct {
	node byte
}

func NewCmdAutomaticControllerUpdateStart() CmdAutomaticControllerUpdateStart {
	return CmdAutomaticControllerUpdateStart{}
}

func (c *CmdAutomaticControllerUpdateStart) SetNode(node int) {
	c.node = byte(node)
}

func (c *CmdAutomaticControllerUpdateStart) Set() error {

	return nil
}

func (c *CmdAutomaticControllerUpdateStart) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x10,
		0x25,
	}
}
type CmdNodesExist struct {
	node byte
}

func NewCmdNodesExist() CmdNodesExist {
	return CmdNodesExist{}
}

func (c *CmdNodesExist) SetNode(node int) {
	c.node = byte(node)
}

func (c *CmdNodesExist) Set() error {

	return nil
}

func (c *CmdNodesExist) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x1F,
		0x25,
	}
}
type CmdNodesExistReply struct {
	node byte
}

func NewCmdNodesExistReply() CmdNodesExistReply {
	return CmdNodesExistReply{}
}

func (c *CmdNodesExistReply) SetNode(node int) {
	c.node = byte(node)
}

func (c *CmdNodesExistReply) Set() error {

	return nil
}

func (c *CmdNodesExistReply) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x20,
		0x25,
	}
}
type CmdSetNwiMode struct {
	node byte
}

func NewCmdSetNwiMode() CmdSetNwiMode {
	return CmdSetNwiMode{}
}

func (c *CmdSetNwiMode) SetNode(node int) {
	c.node = byte(node)
}

func (c *CmdSetNwiMode) Set() error {

	return nil
}

func (c *CmdSetNwiMode) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x22,
		0x25,
	}
}
type CommandComplete struct {
	node byte
}

func NewCommandComplete() CommandComplete {
	return CommandComplete{}
}

func (c *CommandComplete) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandComplete) Set() error {

	return nil
}

func (c *CommandComplete) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x07,
		0x25,
	}
}
type FindNodesInRange struct {
	node byte
}

func NewFindNodesInRange() FindNodesInRange {
	return FindNodesInRange{}
}

func (c *FindNodesInRange) SetNode(node int) {
	c.node = byte(node)
}

func (c *FindNodesInRange) Set() error {

	return nil
}

func (c *FindNodesInRange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x04,
		0x25,
	}
}
type GetNodesInRange struct {
	node byte
}

func NewGetNodesInRange() GetNodesInRange {
	return GetNodesInRange{}
}

func (c *GetNodesInRange) SetNode(node int) {
	c.node = byte(node)
}

func (c *GetNodesInRange) Set() error {

	return nil
}

func (c *GetNodesInRange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x05,
		0x25,
	}
}
type Lost struct {
	node byte
}

func NewLost() Lost {
	return Lost{}
}

func (c *Lost) SetNode(node int) {
	c.node = byte(node)
}

func (c *Lost) Set() error {

	return nil
}

func (c *Lost) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x16,
		0x25,
	}
}
type NewNodeRegistered struct {
	node byte
}

func NewNewNodeRegistered() NewNodeRegistered {
	return NewNodeRegistered{}
}

func (c *NewNodeRegistered) SetNode(node int) {
	c.node = byte(node)
}

func (c *NewNodeRegistered) Set() error {

	return nil
}

func (c *NewNodeRegistered) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x0D,
		0x25,
	}
}
type NewRangeRegistered struct {
	node byte
}

func NewNewRangeRegistered() NewRangeRegistered {
	return NewRangeRegistered{}
}

func (c *NewRangeRegistered) SetNode(node int) {
	c.node = byte(node)
}

func (c *NewRangeRegistered) Set() error {

	return nil
}

func (c *NewRangeRegistered) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x0E,
		0x25,
	}
}
type NodeInfo struct {
	node byte
	Capability byte
	Security byte
	Properties1 byte
	BasicDeviceClass byte
	GenericDeviceClass byte
	SpecificDeviceClass byte
	CommandClasses byte
}

func NewNodeInfo() NodeInfo {
	return NodeInfo{}
}

func (c *NodeInfo) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeInfo) Set(Capability byte,Security byte,Properties1 byte,BasicDeviceClass byte,GenericDeviceClass byte,SpecificDeviceClass byte,CommandClasses byte,) error {
	c.Capability = Capability
	c.Security = Security
	c.Properties1 = Properties1
	c.BasicDeviceClass = BasicDeviceClass
	c.GenericDeviceClass = GenericDeviceClass
	c.SpecificDeviceClass = SpecificDeviceClass
	c.CommandClasses = CommandClasses

	return nil
}

func (c *NodeInfo) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(9),
		byte(Class),
		0x01,
		c.Capability,
		c.Security,
		c.Properties1,
		c.BasicDeviceClass,
		c.GenericDeviceClass,
		c.SpecificDeviceClass,
		c.CommandClasses,
		0x25,
	}
}
type NodeRangeInfo struct {
	node byte
}

func NewNodeRangeInfo() NodeRangeInfo {
	return NodeRangeInfo{}
}

func (c *NodeRangeInfo) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeRangeInfo) Set() error {

	return nil
}

func (c *NodeRangeInfo) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x06,
		0x25,
	}
}
type ZwaveCmdNop struct {
	node byte
}

func NewZwaveCmdNop() ZwaveCmdNop {
	return ZwaveCmdNop{}
}

func (c *ZwaveCmdNop) SetNode(node int) {
	c.node = byte(node)
}

func (c *ZwaveCmdNop) Set() error {

	return nil
}

func (c *ZwaveCmdNop) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x00,
		0x25,
	}
}
type CmdNopPower struct {
	node byte
}

func NewCmdNopPower() CmdNopPower {
	return CmdNopPower{}
}

func (c *CmdNopPower) SetNode(node int) {
	c.node = byte(node)
}

func (c *CmdNopPower) Set() error {

	return nil
}

func (c *CmdNopPower) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x18,
		0x25,
	}
}
type RequestNodeInfo struct {
	node byte
}

func NewRequestNodeInfo() RequestNodeInfo {
	return RequestNodeInfo{}
}

func (c *RequestNodeInfo) SetNode(node int) {
	c.node = byte(node)
}

func (c *RequestNodeInfo) Set() error {

	return nil
}

func (c *RequestNodeInfo) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x02,
		0x25,
	}
}
type ZwaveCmdReserveNodeIds struct {
	node byte
}

func NewZwaveCmdReserveNodeIds() ZwaveCmdReserveNodeIds {
	return ZwaveCmdReserveNodeIds{}
}

func (c *ZwaveCmdReserveNodeIds) SetNode(node int) {
	c.node = byte(node)
}

func (c *ZwaveCmdReserveNodeIds) Set() error {

	return nil
}

func (c *ZwaveCmdReserveNodeIds) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x19,
		0x25,
	}
}
type CmdReservedIds struct {
	node byte
}

func NewCmdReservedIds() CmdReservedIds {
	return CmdReservedIds{}
}

func (c *CmdReservedIds) SetNode(node int) {
	c.node = byte(node)
}

func (c *CmdReservedIds) Set() error {

	return nil
}

func (c *CmdReservedIds) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x1A,
		0x25,
	}
}
type CmdSetSuc struct {
	node byte
}

func NewCmdSetSuc() CmdSetSuc {
	return CmdSetSuc{}
}

func (c *CmdSetSuc) SetNode(node int) {
	c.node = byte(node)
}

func (c *CmdSetSuc) Set() error {

	return nil
}

func (c *CmdSetSuc) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x12,
		0x25,
	}
}
type CmdSetSucAck struct {
	node byte
}

func NewCmdSetSucAck() CmdSetSucAck {
	return CmdSetSucAck{}
}

func (c *CmdSetSucAck) SetNode(node int) {
	c.node = byte(node)
}

func (c *CmdSetSucAck) Set() error {

	return nil
}

func (c *CmdSetSucAck) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x13,
		0x25,
	}
}
type CmdStaticRouteRequest struct {
	node byte
}

func NewCmdStaticRouteRequest() CmdStaticRouteRequest {
	return CmdStaticRouteRequest{}
}

func (c *CmdStaticRouteRequest) SetNode(node int) {
	c.node = byte(node)
}

func (c *CmdStaticRouteRequest) Set() error {

	return nil
}

func (c *CmdStaticRouteRequest) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x15,
		0x25,
	}
}
type CmdSucNodeId struct {
	node byte
}

func NewCmdSucNodeId() CmdSucNodeId {
	return CmdSucNodeId{}
}

func (c *CmdSucNodeId) SetNode(node int) {
	c.node = byte(node)
}

func (c *CmdSucNodeId) Set() error {

	return nil
}

func (c *CmdSucNodeId) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x11,
		0x25,
	}
}
type TransferEnd struct {
	node byte
}

func NewTransferEnd() TransferEnd {
	return TransferEnd{}
}

func (c *TransferEnd) SetNode(node int) {
	c.node = byte(node)
}

func (c *TransferEnd) Set() error {

	return nil
}

func (c *TransferEnd) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x0B,
		0x25,
	}
}
type TransferNewPrimaryComplete struct {
	node byte
}

func NewTransferNewPrimaryComplete() TransferNewPrimaryComplete {
	return TransferNewPrimaryComplete{}
}

func (c *TransferNewPrimaryComplete) SetNode(node int) {
	c.node = byte(node)
}

func (c *TransferNewPrimaryComplete) Set() error {

	return nil
}

func (c *TransferNewPrimaryComplete) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x0F,
		0x25,
	}
}
type TransferNodeInfo struct {
	node byte
}

func NewTransferNodeInfo() TransferNodeInfo {
	return TransferNodeInfo{}
}

func (c *TransferNodeInfo) SetNode(node int) {
	c.node = byte(node)
}

func (c *TransferNodeInfo) Set() error {

	return nil
}

func (c *TransferNodeInfo) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x09,
		0x25,
	}
}
type TransferPresentation struct {
	node byte
}

func NewTransferPresentation() TransferPresentation {
	return TransferPresentation{}
}

func (c *TransferPresentation) SetNode(node int) {
	c.node = byte(node)
}

func (c *TransferPresentation) Set() error {

	return nil
}

func (c *TransferPresentation) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x08,
		0x25,
	}
}
type TransferRangeInfo struct {
	node byte
}

func NewTransferRangeInfo() TransferRangeInfo {
	return TransferRangeInfo{}
}

func (c *TransferRangeInfo) SetNode(node int) {
	c.node = byte(node)
}

func (c *TransferRangeInfo) Set() error {

	return nil
}

func (c *TransferRangeInfo) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x0A,
		0x25,
	}
}
type ExcludeRequest struct {
	node byte
}

func NewExcludeRequest() ExcludeRequest {
	return ExcludeRequest{}
}

func (c *ExcludeRequest) SetNode(node int) {
	c.node = byte(node)
}

func (c *ExcludeRequest) Set() error {

	return nil
}

func (c *ExcludeRequest) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x23,
		0x25,
	}
}
type AssignReturnRoutePriority struct {
	node byte
}

func NewAssignReturnRoutePriority() AssignReturnRoutePriority {
	return AssignReturnRoutePriority{}
}

func (c *AssignReturnRoutePriority) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssignReturnRoutePriority) Set() error {

	return nil
}

func (c *AssignReturnRoutePriority) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x24,
		0x25,
	}
}
type AssignSucReturnRoutePriority struct {
	node byte
}

func NewAssignSucReturnRoutePriority() AssignSucReturnRoutePriority {
	return AssignSucReturnRoutePriority{}
}

func (c *AssignSucReturnRoutePriority) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssignSucReturnRoutePriority) Set() error {

	return nil
}

func (c *AssignSucReturnRoutePriority) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Class),
		0x25,
		0x25,
	}
}
