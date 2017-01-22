package commands
type IpAssociationSet struct {
	node byte
	GroupingIdentifier byte
	IPv6Address byte
	EndPoint byte
}

func NewIpAssociationSet() IpAssociationSet {
	return IpAssociationSet{}
}

func (c *IpAssociationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IpAssociationSet) Set(GroupingIdentifier byte,IPv6Address byte,EndPoint byte,) error {
	c.GroupingIdentifier = GroupingIdentifier
	c.IPv6Address = IPv6Address
	c.EndPoint = EndPoint

	return nil
}

func (c *IpAssociationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(IpAssociation),
		0x01,
		c.GroupingIdentifier,
		c.IPv6Address,
		c.EndPoint,
		0x25,
	}
}
type IpAssociationGet struct {
	node byte
	GroupingIdentifier byte
	Index byte
}

func NewIpAssociationGet() IpAssociationGet {
	return IpAssociationGet{}
}

func (c *IpAssociationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IpAssociationGet) Set(GroupingIdentifier byte,Index byte,) error {
	c.GroupingIdentifier = GroupingIdentifier
	c.Index = Index

	return nil
}

func (c *IpAssociationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(IpAssociation),
		0x02,
		c.GroupingIdentifier,
		c.Index,
		0x25,
	}
}
type IpAssociationReport struct {
    *report
    GroupingIdentifier byte
    Index byte
    ActualNodes byte
    IPv6Address byte
    EndPoint byte
    data []byte
}

func NewIpAssociationReport(data []byte) *IpAssociationReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &IpAssociationReport{
				GroupingIdentifier: data[0],
				Index: data[1],
				ActualNodes: data[2],
				IPv6Address: data[3],
				EndPoint: data[4],
        data: data,
    }
}

type IpAssociationRemove struct {
	node byte
	GroupingIdentifier byte
	IPv6Address byte
	EndPoint byte
}

func NewIpAssociationRemove() IpAssociationRemove {
	return IpAssociationRemove{}
}

func (c *IpAssociationRemove) SetNode(node int) {
	c.node = byte(node)
}

func (c *IpAssociationRemove) Set(GroupingIdentifier byte,IPv6Address byte,EndPoint byte,) error {
	c.GroupingIdentifier = GroupingIdentifier
	c.IPv6Address = IPv6Address
	c.EndPoint = EndPoint

	return nil
}

func (c *IpAssociationRemove) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(IpAssociation),
		0x04,
		c.GroupingIdentifier,
		c.IPv6Address,
		c.EndPoint,
		0x25,
	}
}
