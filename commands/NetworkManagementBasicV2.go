package commands
type LearnModeSetV2 struct {
	node byte
	SeqNo byte
	Reserved byte
	Mode byte
}

func NewLearnModeSetV2() LearnModeSetV2 {
	return LearnModeSetV2{}
}

func (c *LearnModeSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *LearnModeSetV2) Set(SeqNo byte,Reserved byte,Mode byte,) error {
	c.SeqNo = SeqNo
	c.Reserved = Reserved
	c.Mode = Mode

	return nil
}

func (c *LearnModeSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementBasicV2),
		0x01,
		c.SeqNo,
		c.Reserved,
		c.Mode,
		0x25,
	}
}
type LearnModeSetStatusV2 struct {
	node byte
	SeqNo byte
	Status byte
	Reserved byte
	NewNodeID byte
	GrantedKeys byte
	DSK byte
}

func NewLearnModeSetStatusV2() LearnModeSetStatusV2 {
	return LearnModeSetStatusV2{}
}

func (c *LearnModeSetStatusV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *LearnModeSetStatusV2) Set(SeqNo byte,Status byte,Reserved byte,NewNodeID byte,GrantedKeys byte,DSK byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status
	c.Reserved = Reserved
	c.NewNodeID = NewNodeID
	c.GrantedKeys = GrantedKeys
	c.DSK = DSK

	return nil
}

func (c *LearnModeSetStatusV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(8),
		byte(NetworkManagementBasicV2),
		0x02,
		c.SeqNo,
		c.Status,
		c.Reserved,
		c.NewNodeID,
		c.GrantedKeys,
		c.DSK,
		0x25,
	}
}
type NodeInformationSendV2 struct {
	node byte
	SeqNo byte
	Reserved byte
	DestinationNodeID byte
	txOptions byte
}

func NewNodeInformationSendV2() NodeInformationSendV2 {
	return NodeInformationSendV2{}
}

func (c *NodeInformationSendV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeInformationSendV2) Set(SeqNo byte,Reserved byte,DestinationNodeID byte,txOptions byte,) error {
	c.SeqNo = SeqNo
	c.Reserved = Reserved
	c.DestinationNodeID = DestinationNodeID
	c.txOptions = txOptions

	return nil
}

func (c *NodeInformationSendV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(NetworkManagementBasicV2),
		0x05,
		c.SeqNo,
		c.Reserved,
		c.DestinationNodeID,
		c.txOptions,
		0x25,
	}
}
type NetworkUpdateRequestV2 struct {
	node byte
	SeqNo byte
}

func NewNetworkUpdateRequestV2() NetworkUpdateRequestV2 {
	return NetworkUpdateRequestV2{}
}

func (c *NetworkUpdateRequestV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NetworkUpdateRequestV2) Set(SeqNo byte,) error {
	c.SeqNo = SeqNo

	return nil
}

func (c *NetworkUpdateRequestV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NetworkManagementBasicV2),
		0x03,
		c.SeqNo,
		0x25,
	}
}
type NetworkUpdateRequestStatusV2 struct {
	node byte
	SeqNo byte
	Status byte
}

func NewNetworkUpdateRequestStatusV2() NetworkUpdateRequestStatusV2 {
	return NetworkUpdateRequestStatusV2{}
}

func (c *NetworkUpdateRequestStatusV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NetworkUpdateRequestStatusV2) Set(SeqNo byte,Status byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status

	return nil
}

func (c *NetworkUpdateRequestStatusV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementBasicV2),
		0x04,
		c.SeqNo,
		c.Status,
		0x25,
	}
}
type DefaultSetV2 struct {
	node byte
	SeqNo byte
}

func NewDefaultSetV2() DefaultSetV2 {
	return DefaultSetV2{}
}

func (c *DefaultSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *DefaultSetV2) Set(SeqNo byte,) error {
	c.SeqNo = SeqNo

	return nil
}

func (c *DefaultSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NetworkManagementBasicV2),
		0x06,
		c.SeqNo,
		0x25,
	}
}
type DefaultSetCompleteV2 struct {
	node byte
	SeqNo byte
	Status byte
}

func NewDefaultSetCompleteV2() DefaultSetCompleteV2 {
	return DefaultSetCompleteV2{}
}

func (c *DefaultSetCompleteV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *DefaultSetCompleteV2) Set(SeqNo byte,Status byte,) error {
	c.SeqNo = SeqNo
	c.Status = Status

	return nil
}

func (c *DefaultSetCompleteV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NetworkManagementBasicV2),
		0x07,
		c.SeqNo,
		c.Status,
		0x25,
	}
}
type DskGetV2 struct {
	node byte
	SeqNo byte
}

func NewDskGetV2() DskGetV2 {
	return DskGetV2{}
}

func (c *DskGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *DskGetV2) Set(SeqNo byte,) error {
	c.SeqNo = SeqNo

	return nil
}

func (c *DskGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NetworkManagementBasicV2),
		0x08,
		c.SeqNo,
		0x25,
	}
}
type DskReportV2 struct {
    *report
    SeqNo byte
    DSK byte
    data []byte
}

func NewDskReportV2(data []byte) *DskReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &DskReportV2{
				SeqNo: data[0],
				DSK: data[1],
        data: data,
    }
}

