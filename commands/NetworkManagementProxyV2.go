package commands
type NodeInfoCachedGetV2 struct {
	node byte
	SeqNo byte
	Properties1 byte
	NodeID byte
}

func NewNodeInfoCachedGetV2() NodeInfoCachedGetV2 {
	return NodeInfoCachedGetV2{}
}

func (c *NodeInfoCachedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeInfoCachedGetV2) Set(SeqNo byte,Properties1 byte,NodeID byte,) error {
	c.SeqNo = SeqNo
	c.Properties1 = Properties1
	c.NodeID = NodeID

	return nil
}

func (c *NodeInfoCachedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(NetworkManagementProxyV2),
		0x03,
		c.SeqNo,
		c.Properties1,
		c.NodeID,
		0x25,
	}
}
type NodeInfoCachedReportV2 struct {
    *report
    SeqNo byte
    Properties1 byte
    Properties2 byte
    Properties3 byte
    GrantedKeys byte
    BasicDeviceClass byte
    GenericDeviceClass byte
    SpecificDeviceClass byte
    NonSecureCommandClass byte
    SecurityMARK byte
    SecureCommandClass byte
    data []byte
}

func NewNodeInfoCachedReportV2(data []byte) *NodeInfoCachedReportV2 {
    if len(data) < 11 {
				//may want to change this to just return nil
				for i := len(data); i < 11; i++ {
            data = append(data, 0x00)
        }
    }

    return &NodeInfoCachedReportV2{
				SeqNo: data[0],
				Properties1: data[1],
				Properties2: data[2],
				Properties3: data[3],
				GrantedKeys: data[4],
				BasicDeviceClass: data[5],
				GenericDeviceClass: data[6],
				SpecificDeviceClass: data[7],
				NonSecureCommandClass: data[8],
				SecurityMARK: data[9],
				SecureCommandClass: data[10],
        data: data,
    }
}

type NodeListGetV2 struct {
	node byte
	SeqNo byte
}

func NewNodeListGetV2() NodeListGetV2 {
	return NodeListGetV2{}
}

func (c *NodeListGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeListGetV2) Set(SeqNo byte,) error {
	c.SeqNo = SeqNo

	return nil
}

func (c *NodeListGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NetworkManagementProxyV2),
		0x01,
		c.SeqNo,
		0x25,
	}
}
type NodeListReportV2 struct {
    *report
    SeqNo byte
    Status byte
    NodeListControllerID byte
    NodeListData byte
    data []byte
}

func NewNodeListReportV2(data []byte) *NodeListReportV2 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &NodeListReportV2{
				SeqNo: data[0],
				Status: data[1],
				NodeListControllerID: data[2],
				NodeListData: data[3],
        data: data,
    }
}

