package commands
type AssociationGetV2 struct {
	node byte
	GroupingIdentifier byte
}

func NewAssociationGetV2() AssociationGetV2 {
	return AssociationGetV2{}
}

func (c *AssociationGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGetV2) Set(GroupingIdentifier byte,) error {
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *AssociationGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(AssociationV2),
		0x02,
		c.GroupingIdentifier,
		0x25,
	}
}
type AssociationGroupingsGetV2 struct {
	node byte
}

func NewAssociationGroupingsGetV2() AssociationGroupingsGetV2 {
	return AssociationGroupingsGetV2{}
}

func (c *AssociationGroupingsGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGroupingsGetV2) Set() error {

	return nil
}

func (c *AssociationGroupingsGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(AssociationV2),
		0x05,
		0x25,
	}
}
type AssociationGroupingsReportV2 struct {
    *report
    SupportedGroupings byte
    data []byte
}

func NewAssociationGroupingsReportV2(data []byte) *AssociationGroupingsReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationGroupingsReportV2{
				SupportedGroupings: data[0],
        data: data,
    }
}

type AssociationRemoveV2 struct {
	node byte
	GroupingIdentifier byte
	NodeID byte
}

func NewAssociationRemoveV2() AssociationRemoveV2 {
	return AssociationRemoveV2{}
}

func (c *AssociationRemoveV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationRemoveV2) Set(GroupingIdentifier byte,NodeID byte,) error {
	c.GroupingIdentifier = GroupingIdentifier
	c.NodeID = NodeID

	return nil
}

func (c *AssociationRemoveV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(AssociationV2),
		0x04,
		c.GroupingIdentifier,
		c.NodeID,
		0x25,
	}
}
type AssociationReportV2 struct {
    *report
    GroupingIdentifier byte
    MaxNodesSupported byte
    ReportstoFollow byte
    NodeID byte
    data []byte
}

func NewAssociationReportV2(data []byte) *AssociationReportV2 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationReportV2{
				GroupingIdentifier: data[0],
				MaxNodesSupported: data[1],
				ReportstoFollow: data[2],
				NodeID: data[3],
        data: data,
    }
}

type AssociationSetV2 struct {
	node byte
	GroupingIdentifier byte
	NodeID byte
}

func NewAssociationSetV2() AssociationSetV2 {
	return AssociationSetV2{}
}

func (c *AssociationSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationSetV2) Set(GroupingIdentifier byte,NodeID byte,) error {
	c.GroupingIdentifier = GroupingIdentifier
	c.NodeID = NodeID

	return nil
}

func (c *AssociationSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(AssociationV2),
		0x01,
		c.GroupingIdentifier,
		c.NodeID,
		0x25,
	}
}
type AssociationSpecificGroupGetV2 struct {
	node byte
}

func NewAssociationSpecificGroupGetV2() AssociationSpecificGroupGetV2 {
	return AssociationSpecificGroupGetV2{}
}

func (c *AssociationSpecificGroupGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationSpecificGroupGetV2) Set() error {

	return nil
}

func (c *AssociationSpecificGroupGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(AssociationV2),
		0x0B,
		0x25,
	}
}
type AssociationSpecificGroupReportV2 struct {
    *report
    Group byte
    data []byte
}

func NewAssociationSpecificGroupReportV2(data []byte) *AssociationSpecificGroupReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationSpecificGroupReportV2{
				Group: data[0],
        data: data,
    }
}

