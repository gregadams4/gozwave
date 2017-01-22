package commands
type NodeNamingNodeLocationReport struct {
    *report
    Level byte
    Nodelocationchar byte
    data []byte
}

func NewNodeNamingNodeLocationReport(data []byte) *NodeNamingNodeLocationReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &NodeNamingNodeLocationReport{
				Level: data[0],
				Nodelocationchar: data[1],
        data: data,
    }
}

type NodeNamingNodeLocationSet struct {
	node byte
	Level byte
	Nodelocationchar byte
}

func NewNodeNamingNodeLocationSet() NodeNamingNodeLocationSet {
	return NodeNamingNodeLocationSet{}
}

func (c *NodeNamingNodeLocationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeNamingNodeLocationSet) Set(Level byte,Nodelocationchar byte,) error {
	c.Level = Level
	c.Nodelocationchar = Nodelocationchar

	return nil
}

func (c *NodeNamingNodeLocationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NodeNaming),
		0x04,
		c.Level,
		c.Nodelocationchar,
		0x25,
	}
}
type NodeNamingNodeLocationGet struct {
	node byte
}

func NewNodeNamingNodeLocationGet() NodeNamingNodeLocationGet {
	return NodeNamingNodeLocationGet{}
}

func (c *NodeNamingNodeLocationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeNamingNodeLocationGet) Set() error {

	return nil
}

func (c *NodeNamingNodeLocationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(NodeNaming),
		0x05,
		0x25,
	}
}
type NodeNamingNodeNameGet struct {
	node byte
}

func NewNodeNamingNodeNameGet() NodeNamingNodeNameGet {
	return NodeNamingNodeNameGet{}
}

func (c *NodeNamingNodeNameGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeNamingNodeNameGet) Set() error {

	return nil
}

func (c *NodeNamingNodeNameGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(NodeNaming),
		0x02,
		0x25,
	}
}
type NodeNamingNodeNameReport struct {
    *report
    Level byte
    Nodenamechar byte
    data []byte
}

func NewNodeNamingNodeNameReport(data []byte) *NodeNamingNodeNameReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &NodeNamingNodeNameReport{
				Level: data[0],
				Nodenamechar: data[1],
        data: data,
    }
}

type NodeNamingNodeNameSet struct {
	node byte
	Level byte
	Nodenamechar byte
}

func NewNodeNamingNodeNameSet() NodeNamingNodeNameSet {
	return NodeNamingNodeNameSet{}
}

func (c *NodeNamingNodeNameSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *NodeNamingNodeNameSet) Set(Level byte,Nodenamechar byte,) error {
	c.Level = Level
	c.Nodenamechar = Nodenamechar

	return nil
}

func (c *NodeNamingNodeNameSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(NodeNaming),
		0x01,
		c.Level,
		c.Nodenamechar,
		0x25,
	}
}
