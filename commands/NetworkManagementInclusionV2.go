package commands
type FailedNodeRemoveV2 struct {
	node byte
	SeqNo byte
	NodeID byte
}

func NewFailedNodeRemoveV2() FailedNodeRemoveV2 {
	return FailedNodeRemoveV2{}
}

func (c *FailedNodeRemoveV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *FailedNodeRemoveV2) Set(SeqNo byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.NodeID = NodeID

	return nil
}

func (c *FailedNodeRemoveV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementInclusionV2),
		0x07,
		c.SeqNo,
		c.NodeID,
		0x25,
	}
}
type FailedNodeRemoveStatusV2 struct {
	node byte
	SeqNo byte
	Status byte
	NodeID byte
}

func NewFailedNodeRemoveStatusV2() FailedNodeRemoveStatusV2 {
	return FailedNodeRemoveStatusV2{}
}

func (c *FailedNodeRemoveStatusV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *FailedNodeRemoveStatusV2) Set(SeqNo byte,Status byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status
	c.NodeID = NodeID

	return nil
}

func (c *FailedNodeRemoveStatusV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementInclusionV2),
		0x08,
		c.SeqNo,
		c.Status,
		c.NodeID,
		0x25,
	}
}
type NodeAddV2 struct {
	node byte
	SeqNo byte
	Reserved byte
	Mode byte
	txOptions byte
}

func NewNodeAddV2() NodeAddV2 {
	return NodeAddV2{}
}

func (c *NodeAddV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeAddV2) Set(SeqNo byte,Reserved byte,Mode byte,txOptions byte,) error {
	c.SeqNo = SeqNo
	c.Reserved = Reserved
	c.Mode = Mode
	c.txOptions = txOptions

	return nil
}

func (c *NodeAddV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(NetworkManagementInclusionV2),
		0x01,
		c.SeqNo,
		c.Reserved,
		c.Mode,
		c.txOptions,
		0x25,
	}
}
type NodeAddStatusV2 struct {
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
	NonSecureCommandClass byte
	SecurityMARK byte
	SecureCommandClass byte
	GrantedKeys byte
	KEXFailType byte
}

func NewNodeAddStatusV2() NodeAddStatusV2 {
	return NodeAddStatusV2{}
}

func (c *NodeAddStatusV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeAddStatusV2) Set(SeqNo byte,Status byte,Reserved byte,NewNodeID byte,NodeInfoLength byte,Properties1 byte,Properties2 byte,BasicDeviceClass byte,GenericDeviceClass byte,SpecificDeviceClass byte,NonSecureCommandClass byte,SecurityMARK byte,SecureCommandClass byte,GrantedKeys byte,KEXFailType byte,) error {
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
	c.NonSecureCommandClass = NonSecureCommandClass
	c.SecurityMARK = SecurityMARK
	c.SecureCommandClass = SecureCommandClass
	c.GrantedKeys = GrantedKeys
	c.KEXFailType = KEXFailType

	return nil
}

func (c *NodeAddStatusV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(17),
		byte(NetworkManagementInclusionV2),
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
		c.NonSecureCommandClass,
		c.SecurityMARK,
		c.SecureCommandClass,
		c.GrantedKeys,
		c.KEXFailType,
		0x25,
	}
}
type NodeRemoveV2 struct {
	node byte
	SeqNo byte
	Reserved byte
	Mode byte
}

func NewNodeRemoveV2() NodeRemoveV2 {
	return NodeRemoveV2{}
}

func (c *NodeRemoveV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeRemoveV2) Set(SeqNo byte,Reserved byte,Mode byte,) error {
	c.SeqNo = SeqNo
	c.Reserved = Reserved
	c.Mode = Mode

	return nil
}

func (c *NodeRemoveV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementInclusionV2),
		0x03,
		c.SeqNo,
		c.Reserved,
		c.Mode,
		0x25,
	}
}
type NodeRemoveStatusV2 struct {
	node byte
	SeqNo byte
	Status byte
	NodeID byte
}

func NewNodeRemoveStatusV2() NodeRemoveStatusV2 {
	return NodeRemoveStatusV2{}
}

func (c *NodeRemoveStatusV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeRemoveStatusV2) Set(SeqNo byte,Status byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status
	c.NodeID = NodeID

	return nil
}

func (c *NodeRemoveStatusV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementInclusionV2),
		0x04,
		c.SeqNo,
		c.Status,
		c.NodeID,
		0x25,
	}
}
type FailedNodeReplaceV2 struct {
	node byte
	SeqNo byte
	NodeID byte
	txOptions byte
	Mode byte
}

func NewFailedNodeReplaceV2() FailedNodeReplaceV2 {
	return FailedNodeReplaceV2{}
}

func (c *FailedNodeReplaceV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *FailedNodeReplaceV2) Set(SeqNo byte,NodeID byte,txOptions byte,Mode byte,) error {
	c.SeqNo = SeqNo
	c.NodeID = NodeID
	c.txOptions = txOptions
	c.Mode = Mode

	return nil
}

func (c *FailedNodeReplaceV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(NetworkManagementInclusionV2),
		0x09,
		c.SeqNo,
		c.NodeID,
		c.txOptions,
		c.Mode,
		0x25,
	}
}
type FailedNodeReplaceStatusV2 struct {
	node byte
	SeqNo byte
	Status byte
	NodeID byte
	GrantedKeys byte
	KEXFailType byte
}

func NewFailedNodeReplaceStatusV2() FailedNodeReplaceStatusV2 {
	return FailedNodeReplaceStatusV2{}
}

func (c *FailedNodeReplaceStatusV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *FailedNodeReplaceStatusV2) Set(SeqNo byte,Status byte,NodeID byte,GrantedKeys byte,KEXFailType byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status
	c.NodeID = NodeID
	c.GrantedKeys = GrantedKeys
	c.KEXFailType = KEXFailType

	return nil
}

func (c *FailedNodeReplaceStatusV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(NetworkManagementInclusionV2),
		0x0A,
		c.SeqNo,
		c.Status,
		c.NodeID,
		c.GrantedKeys,
		c.KEXFailType,
		0x25,
	}
}
type NodeNeighborUpdateRequestV2 struct {
	node byte
	SeqNo byte
	NodeID byte
}

func NewNodeNeighborUpdateRequestV2() NodeNeighborUpdateRequestV2 {
	return NodeNeighborUpdateRequestV2{}
}

func (c *NodeNeighborUpdateRequestV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeNeighborUpdateRequestV2) Set(SeqNo byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.NodeID = NodeID

	return nil
}

func (c *NodeNeighborUpdateRequestV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementInclusionV2),
		0x0B,
		c.SeqNo,
		c.NodeID,
		0x25,
	}
}
type NodeNeighborUpdateStatusV2 struct {
	node byte
	SeqNo byte
	Status byte
}

func NewNodeNeighborUpdateStatusV2() NodeNeighborUpdateStatusV2 {
	return NodeNeighborUpdateStatusV2{}
}

func (c *NodeNeighborUpdateStatusV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeNeighborUpdateStatusV2) Set(SeqNo byte,Status byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status

	return nil
}

func (c *NodeNeighborUpdateStatusV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementInclusionV2),
		0x0C,
		c.SeqNo,
		c.Status,
		0x25,
	}
}
type ReturnRouteAssignV2 struct {
	node byte
	SeqNo byte
	SourceNodeID byte
	DestinationNodeID byte
}

func NewReturnRouteAssignV2() ReturnRouteAssignV2 {
	return ReturnRouteAssignV2{}
}

func (c *ReturnRouteAssignV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ReturnRouteAssignV2) Set(SeqNo byte,SourceNodeID byte,DestinationNodeID byte,) error {
	c.SeqNo = SeqNo
	c.SourceNodeID = SourceNodeID
	c.DestinationNodeID = DestinationNodeID

	return nil
}

func (c *ReturnRouteAssignV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementInclusionV2),
		0x0D,
		c.SeqNo,
		c.SourceNodeID,
		c.DestinationNodeID,
		0x25,
	}
}
type ReturnRouteAssignCompleteV2 struct {
	node byte
	SeqNo byte
	Status byte
}

func NewReturnRouteAssignCompleteV2() ReturnRouteAssignCompleteV2 {
	return ReturnRouteAssignCompleteV2{}
}

func (c *ReturnRouteAssignCompleteV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ReturnRouteAssignCompleteV2) Set(SeqNo byte,Status byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status

	return nil
}

func (c *ReturnRouteAssignCompleteV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementInclusionV2),
		0x0E,
		c.SeqNo,
		c.Status,
		0x25,
	}
}
type ReturnRouteDeleteV2 struct {
	node byte
	SeqNo byte
	NodeID byte
}

func NewReturnRouteDeleteV2() ReturnRouteDeleteV2 {
	return ReturnRouteDeleteV2{}
}

func (c *ReturnRouteDeleteV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ReturnRouteDeleteV2) Set(SeqNo byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.NodeID = NodeID

	return nil
}

func (c *ReturnRouteDeleteV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementInclusionV2),
		0x0F,
		c.SeqNo,
		c.NodeID,
		0x25,
	}
}
type ReturnRouteDeleteCompleteV2 struct {
	node byte
	SeqNo byte
	Status byte
}

func NewReturnRouteDeleteCompleteV2() ReturnRouteDeleteCompleteV2 {
	return ReturnRouteDeleteCompleteV2{}
}

func (c *ReturnRouteDeleteCompleteV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ReturnRouteDeleteCompleteV2) Set(SeqNo byte,Status byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status

	return nil
}

func (c *ReturnRouteDeleteCompleteV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementInclusionV2),
		0x10,
		c.SeqNo,
		c.Status,
		0x25,
	}
}
type NodeAddKeysReportV2 struct {
    *report
    SeqNo byte
    Properties1 byte
    RequestedKeys byte
    data []byte
}

func NewNodeAddKeysReportV2(data []byte) *NodeAddKeysReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &NodeAddKeysReportV2{
				SeqNo: data[0],
				Properties1: data[1],
				RequestedKeys: data[2],
        data: data,
    }
}

type NodeAddKeysSetV2 struct {
	node byte
	SeqNo byte
	Properties1 byte
	GrantedKeys byte
}

func NewNodeAddKeysSetV2() NodeAddKeysSetV2 {
	return NodeAddKeysSetV2{}
}

func (c *NodeAddKeysSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeAddKeysSetV2) Set(SeqNo byte,Properties1 byte,GrantedKeys byte,) error {
	c.SeqNo = SeqNo
	c.Properties1 = Properties1
	c.GrantedKeys = GrantedKeys

	return nil
}

func (c *NodeAddKeysSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementInclusionV2),
		0x12,
		c.SeqNo,
		c.Properties1,
		c.GrantedKeys,
		0x25,
	}
}
type NodeAddDskReportV2 struct {
    *report
    SeqNo byte
    Properties1 byte
    DSK byte
    data []byte
}

func NewNodeAddDskReportV2(data []byte) *NodeAddDskReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &NodeAddDskReportV2{
				SeqNo: data[0],
				Properties1: data[1],
				DSK: data[2],
        data: data,
    }
}

type NodeAddDskSetV2 struct {
	node byte
	SeqNo byte
	Properties1 byte
	InputDSK byte
}

func NewNodeAddDskSetV2() NodeAddDskSetV2 {
	return NodeAddDskSetV2{}
}

func (c *NodeAddDskSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeAddDskSetV2) Set(SeqNo byte,Properties1 byte,InputDSK byte,) error {
	c.SeqNo = SeqNo
	c.Properties1 = Properties1
	c.InputDSK = InputDSK

	return nil
}

func (c *NodeAddDskSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementInclusionV2),
		0x14,
		c.SeqNo,
		c.Properties1,
		c.InputDSK,
		0x25,
	}
}
