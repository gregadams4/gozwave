package commands
type AssociationGet struct {
	node byte
	GroupingIdentifier byte
}

func NewAssociationGet() AssociationGet {
	return AssociationGet{}
}

func (c *AssociationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGet) Set(GroupingIdentifier byte,) error {
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *AssociationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Association),
		0x02,
		c.GroupingIdentifier,
		0x25,
	}
}
type AssociationGroupingsGet struct {
	node byte
}

func NewAssociationGroupingsGet() AssociationGroupingsGet {
	return AssociationGroupingsGet{}
}

func (c *AssociationGroupingsGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGroupingsGet) Set() error {

	return nil
}

func (c *AssociationGroupingsGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Association),
		0x05,
		0x25,
	}
}
type AssociationGroupingsReport struct {
    *report
    SupportedGroupings byte
    data []byte
}

func NewAssociationGroupingsReport(data []byte) *AssociationGroupingsReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationGroupingsReport{
				SupportedGroupings: data[0],
        data: data,
    }
}

type AssociationRemove struct {
	node byte
	GroupingIdentifier byte
	NodeID byte
}

func NewAssociationRemove() AssociationRemove {
	return AssociationRemove{}
}

func (c *AssociationRemove) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationRemove) Set(GroupingIdentifier byte,NodeID byte,) error {
	c.GroupingIdentifier = GroupingIdentifier
	c.NodeID = NodeID

	return nil
}

func (c *AssociationRemove) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(Association),
		0x04,
		c.GroupingIdentifier,
		c.NodeID,
		0x25,
	}
}
type AssociationReport struct {
    *report
    GroupingIdentifier byte
    MaxNodesSupported byte
    ReportstoFollow byte
    NodeID byte
    data []byte
}

func NewAssociationReport(data []byte) *AssociationReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationReport{
				GroupingIdentifier: data[0],
				MaxNodesSupported: data[1],
				ReportstoFollow: data[2],
				NodeID: data[3],
        data: data,
    }
}

type AssociationSet struct {
	node byte
	GroupingIdentifier byte
	NodeID byte
}

func NewAssociationSet() AssociationSet {
	return AssociationSet{}
}

func (c *AssociationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationSet) Set(GroupingIdentifier byte,NodeID byte,) error {
	c.GroupingIdentifier = GroupingIdentifier
	c.NodeID = NodeID

	return nil
}

func (c *AssociationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(Association),
		0x01,
		c.GroupingIdentifier,
		c.NodeID,
		0x25,
	}
}
