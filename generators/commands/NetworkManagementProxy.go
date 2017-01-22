package commands
type NodeInfoCachedGet struct {
	node byte
	SeqNo byte
	Properties1 byte
	NodeID byte
}

func NewNodeInfoCachedGet() NodeInfoCachedGet {
	return NodeInfoCachedGet{}
}

func (c *NodeInfoCachedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeInfoCachedGet) Set(SeqNo byte,Properties1 byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.Properties1 = Properties1
	c.NodeID = NodeID

	return nil
}

func (c *NodeInfoCachedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementProxy),
		0x03,
		c.SeqNo,
		c.Properties1,
		c.NodeID,
		0x25,
	}
}
type NodeInfoCachedReport struct {
    *report
    SeqNo byte
    Properties1 byte
    Properties2 byte
    Properties3 byte
    Reserved byte
    BasicDeviceClass byte
    GenericDeviceClass byte
    SpecificDeviceClass byte
    NonSecureCommandClass byte
    SecurityScheme0MARK byte
    SecurityScheme0CommandClass byte
    data []byte
}

func NewNodeInfoCachedReport(data []byte) *NodeInfoCachedReport {
    if len(data) < 11 {
				//may want to change this to just return nil
				for i := len(data); i < 11; i++ {
            data = append(data, 0x00)
        }
    }

    return &NodeInfoCachedReport{
				SeqNo: data[0],
				Properties1: data[1],
				Properties2: data[2],
				Properties3: data[3],
				Reserved: data[4],
				BasicDeviceClass: data[5],
				GenericDeviceClass: data[6],
				SpecificDeviceClass: data[7],
				NonSecureCommandClass: data[8],
				SecurityScheme0MARK: data[9],
				SecurityScheme0CommandClass: data[10],
        data: data,
    }
}

type NodeListGet struct {
	node byte
	SeqNo byte
}

func NewNodeListGet() NodeListGet {
	return NodeListGet{}
}

func (c *NodeListGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeListGet) Set(SeqNo byte,) error {
	c.SeqNo = SeqNo

	return nil
}

func (c *NodeListGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NetworkManagementProxy),
		0x01,
		c.SeqNo,
		0x25,
	}
}
type NodeListReport struct {
    *report
    SeqNo byte
    Status byte
    NodeListControllerID byte
    NodeListData byte
    data []byte
}

func NewNodeListReport(data []byte) *NodeListReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &NodeListReport{
				SeqNo: data[0],
				Status: data[1],
				NodeListControllerID: data[2],
				NodeListData: data[3],
        data: data,
    }
}

