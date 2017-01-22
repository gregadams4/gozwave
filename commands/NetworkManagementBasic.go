package commands
type LearnModeSet struct {
	node byte
	SeqNo byte
	Reserved byte
	Mode byte
}

func NewLearnModeSet() LearnModeSet {
	return LearnModeSet{}
}

func (c *LearnModeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *LearnModeSet) Set(SeqNo byte,Reserved byte,Mode byte,) error {
	c.SeqNo = SeqNo
	c.Reserved = Reserved
	c.Mode = Mode

	return nil
}

func (c *LearnModeSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementBasic),
		0x01,
		c.SeqNo,
		c.Reserved,
		c.Mode,
		0x25,
	}
}
type LearnModeSetStatus struct {
	node byte
	SeqNo byte
	Status byte
	Reserved byte
	NewNodeID byte
}

func NewLearnModeSetStatus() LearnModeSetStatus {
	return LearnModeSetStatus{}
}

func (c *LearnModeSetStatus) SetNode(node int) {
	c.node = byte(node)
}

func (c *LearnModeSetStatus) Set(SeqNo byte,Status byte,Reserved byte,NewNodeID byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status
	c.Reserved = Reserved
	c.NewNodeID = NewNodeID

	return nil
}

func (c *LearnModeSetStatus) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(NetworkManagementBasic),
		0x02,
		c.SeqNo,
		c.Status,
		c.Reserved,
		c.NewNodeID,
		0x25,
	}
}
type NodeInformationSend struct {
	node byte
	SeqNo byte
	Reserved byte
	DestinationNodeID byte
	txOptions byte
}

func NewNodeInformationSend() NodeInformationSend {
	return NodeInformationSend{}
}

func (c *NodeInformationSend) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeInformationSend) Set(SeqNo byte,Reserved byte,DestinationNodeID byte,txOptions byte,) error {
	c.SeqNo = SeqNo
	c.Reserved = Reserved
	c.DestinationNodeID = DestinationNodeID
	c.txOptions = txOptions

	return nil
}

func (c *NodeInformationSend) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(NetworkManagementBasic),
		0x05,
		c.SeqNo,
		c.Reserved,
		c.DestinationNodeID,
		c.txOptions,
		0x25,
	}
}
type NetworkUpdateRequest struct {
	node byte
	SeqNo byte
}

func NewNetworkUpdateRequest() NetworkUpdateRequest {
	return NetworkUpdateRequest{}
}

func (c *NetworkUpdateRequest) SetNode(node int) {
	c.node = byte(node)
}

func (c *NetworkUpdateRequest) Set(SeqNo byte,) error {
	c.SeqNo = SeqNo

	return nil
}

func (c *NetworkUpdateRequest) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NetworkManagementBasic),
		0x03,
		c.SeqNo,
		0x25,
	}
}
type NetworkUpdateRequestStatus struct {
	node byte
	SeqNo byte
	Status byte
}

func NewNetworkUpdateRequestStatus() NetworkUpdateRequestStatus {
	return NetworkUpdateRequestStatus{}
}

func (c *NetworkUpdateRequestStatus) SetNode(node int) {
	c.node = byte(node)
}

func (c *NetworkUpdateRequestStatus) Set(SeqNo byte,Status byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status

	return nil
}

func (c *NetworkUpdateRequestStatus) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementBasic),
		0x04,
		c.SeqNo,
		c.Status,
		0x25,
	}
}
type DefaultSet struct {
	node byte
	SeqNo byte
}

func NewDefaultSet() DefaultSet {
	return DefaultSet{}
}

func (c *DefaultSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DefaultSet) Set(SeqNo byte,) error {
	c.SeqNo = SeqNo

	return nil
}

func (c *DefaultSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NetworkManagementBasic),
		0x06,
		c.SeqNo,
		0x25,
	}
}
type DefaultSetComplete struct {
	node byte
	SeqNo byte
	Status byte
}

func NewDefaultSetComplete() DefaultSetComplete {
	return DefaultSetComplete{}
}

func (c *DefaultSetComplete) SetNode(node int) {
	c.node = byte(node)
}

func (c *DefaultSetComplete) Set(SeqNo byte,Status byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status

	return nil
}

func (c *DefaultSetComplete) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementBasic),
		0x07,
		c.SeqNo,
		c.Status,
		0x25,
	}
}
