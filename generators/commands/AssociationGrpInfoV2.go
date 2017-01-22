package commands
type AssociationGroupNameGetV2 struct {
	node byte
	GroupingIdentifier byte
}

func NewAssociationGroupNameGetV2() AssociationGroupNameGetV2 {
	return AssociationGroupNameGetV2{}
}

func (c *AssociationGroupNameGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGroupNameGetV2) Set(GroupingIdentifier byte,) error {
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *AssociationGroupNameGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(AssociationGrpInfoV2),
		0x01,
		c.GroupingIdentifier,
		0x25,
	}
}
type AssociationGroupNameReportV2 struct {
    *report
    GroupingIdentifier byte
    LengthofName byte
    Name byte
    data []byte
}

func NewAssociationGroupNameReportV2(data []byte) *AssociationGroupNameReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationGroupNameReportV2{
				GroupingIdentifier: data[0],
				LengthofName: data[1],
				Name: data[2],
        data: data,
    }
}

type AssociationGroupInfoGetV2 struct {
	node byte
	Properties1 byte
	GroupingIdentifier byte
}

func NewAssociationGroupInfoGetV2() AssociationGroupInfoGetV2 {
	return AssociationGroupInfoGetV2{}
}

func (c *AssociationGroupInfoGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGroupInfoGetV2) Set(Properties1 byte,GroupingIdentifier byte,) error {
	c.Properties1 = Properties1
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *AssociationGroupInfoGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(AssociationGrpInfoV2),
		0x03,
		c.Properties1,
		c.GroupingIdentifier,
		0x25,
	}
}
type AssociationGroupInfoReportV2 struct {
    *report
    Properties1 byte
    data []byte
}

func NewAssociationGroupInfoReportV2(data []byte) *AssociationGroupInfoReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationGroupInfoReportV2{
				Properties1: data[0],
        data: data,
    }
}

type AssociationGroupCommandListGetV2 struct {
	node byte
	Properties1 byte
	GroupingIdentifier byte
}

func NewAssociationGroupCommandListGetV2() AssociationGroupCommandListGetV2 {
	return AssociationGroupCommandListGetV2{}
}

func (c *AssociationGroupCommandListGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGroupCommandListGetV2) Set(Properties1 byte,GroupingIdentifier byte,) error {
	c.Properties1 = Properties1
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *AssociationGroupCommandListGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(AssociationGrpInfoV2),
		0x05,
		c.Properties1,
		c.GroupingIdentifier,
		0x25,
	}
}
type AssociationGroupCommandListReportV2 struct {
    *report
    GroupingIdentifier byte
    ListLength byte
    Command byte
    data []byte
}

func NewAssociationGroupCommandListReportV2(data []byte) *AssociationGroupCommandListReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationGroupCommandListReportV2{
				GroupingIdentifier: data[0],
				ListLength: data[1],
				Command: data[2],
        data: data,
    }
}

