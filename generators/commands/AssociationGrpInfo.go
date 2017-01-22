package commands
type AssociationGroupNameGet struct {
	node byte
	GroupingIdentifier byte
}

func NewAssociationGroupNameGet() AssociationGroupNameGet {
	return AssociationGroupNameGet{}
}

func (c *AssociationGroupNameGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGroupNameGet) Set(GroupingIdentifier byte,) error {
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *AssociationGroupNameGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(AssociationGrpInfo),
		0x01,
		c.GroupingIdentifier,
		0x25,
	}
}
type AssociationGroupNameReport struct {
    *report
    GroupingIdentifier byte
    LengthofName byte
    Name byte
    data []byte
}

func NewAssociationGroupNameReport(data []byte) *AssociationGroupNameReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationGroupNameReport{
				GroupingIdentifier: data[0],
				LengthofName: data[1],
				Name: data[2],
        data: data,
    }
}

type AssociationGroupInfoGet struct {
	node byte
	Properties1 byte
	GroupingIdentifier byte
}

func NewAssociationGroupInfoGet() AssociationGroupInfoGet {
	return AssociationGroupInfoGet{}
}

func (c *AssociationGroupInfoGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGroupInfoGet) Set(Properties1 byte,GroupingIdentifier byte,) error {
	c.Properties1 = Properties1
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *AssociationGroupInfoGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(AssociationGrpInfo),
		0x03,
		c.Properties1,
		c.GroupingIdentifier,
		0x25,
	}
}
type AssociationGroupInfoReport struct {
    *report
    Properties1 byte
    data []byte
}

func NewAssociationGroupInfoReport(data []byte) *AssociationGroupInfoReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationGroupInfoReport{
				Properties1: data[0],
        data: data,
    }
}

type AssociationGroupCommandListGet struct {
	node byte
	Properties1 byte
	GroupingIdentifier byte
}

func NewAssociationGroupCommandListGet() AssociationGroupCommandListGet {
	return AssociationGroupCommandListGet{}
}

func (c *AssociationGroupCommandListGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AssociationGroupCommandListGet) Set(Properties1 byte,GroupingIdentifier byte,) error {
	c.Properties1 = Properties1
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *AssociationGroupCommandListGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(AssociationGrpInfo),
		0x05,
		c.Properties1,
		c.GroupingIdentifier,
		0x25,
	}
}
type AssociationGroupCommandListReport struct {
    *report
    GroupingIdentifier byte
    ListLength byte
    Command byte
    data []byte
}

func NewAssociationGroupCommandListReport(data []byte) *AssociationGroupCommandListReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &AssociationGroupCommandListReport{
				GroupingIdentifier: data[0],
				ListLength: data[1],
				Command: data[2],
        data: data,
    }
}

