package commands
type AssociationGroupNameGetV3 struct {
	node byte
	GroupingIdentifier byte
}

func NewAssociationGroupNameGetV3() AssociationGroupNameGetV3 {
	return AssociationGroupNameGetV3{}
}

func (c *AssociationGroupNameGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGroupNameGetV3) Set(GroupingIdentifier byte,) error {
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *AssociationGroupNameGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(AssociationGrpInfoV3),
		0x01,
		c.GroupingIdentifier,
		0x25,
	}
}
type AssociationGroupNameReportV3 struct {
    *report
    GroupingIdentifier byte
    LengthofName byte
    Name byte
    data []byte
}

func NewAssociationGroupNameReportV3(data []byte) *AssociationGroupNameReportV3 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationGroupNameReportV3{
				GroupingIdentifier: data[0],
				LengthofName: data[1],
				Name: data[2],
        data: data,
    }
}

type AssociationGroupInfoGetV3 struct {
	node byte
	Properties1 byte
	GroupingIdentifier byte
}

func NewAssociationGroupInfoGetV3() AssociationGroupInfoGetV3 {
	return AssociationGroupInfoGetV3{}
}

func (c *AssociationGroupInfoGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGroupInfoGetV3) Set(Properties1 byte,GroupingIdentifier byte,) error {
	c.Properties1 = Properties1
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *AssociationGroupInfoGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(AssociationGrpInfoV3),
		0x03,
		c.Properties1,
		c.GroupingIdentifier,
		0x25,
	}
}
type AssociationGroupInfoReportV3 struct {
    *report
    Properties1 byte
    data []byte
}

func NewAssociationGroupInfoReportV3(data []byte) *AssociationGroupInfoReportV3 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationGroupInfoReportV3{
				Properties1: data[0],
        data: data,
    }
}

type AssociationGroupCommandListGetV3 struct {
	node byte
	Properties1 byte
	GroupingIdentifier byte
}

func NewAssociationGroupCommandListGetV3() AssociationGroupCommandListGetV3 {
	return AssociationGroupCommandListGetV3{}
}

func (c *AssociationGroupCommandListGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGroupCommandListGetV3) Set(Properties1 byte,GroupingIdentifier byte,) error {
	c.Properties1 = Properties1
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *AssociationGroupCommandListGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(AssociationGrpInfoV3),
		0x05,
		c.Properties1,
		c.GroupingIdentifier,
		0x25,
	}
}
type AssociationGroupCommandListReportV3 struct {
    *report
    GroupingIdentifier byte
    ListLength byte
    Command byte
    data []byte
}

func NewAssociationGroupCommandListReportV3(data []byte) *AssociationGroupCommandListReportV3 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationGroupCommandListReportV3{
				GroupingIdentifier: data[0],
				ListLength: data[1],
				Command: data[2],
        data: data,
    }
}

