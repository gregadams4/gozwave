package commands
type MultiInstanceAssociationGet struct {
	node byte
	GroupingIdentifier byte
}

func NewMultiInstanceAssociationGet() MultiInstanceAssociationGet {
	return MultiInstanceAssociationGet{}
}

func (c *MultiInstanceAssociationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiInstanceAssociationGet) Set(GroupingIdentifier byte,) error {
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *MultiInstanceAssociationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MultiInstanceAssociation),
		0x02,
		c.GroupingIdentifier,
		0x25,
	}
}
type MultiInstanceAssociationGroupingsGet struct {
	node byte
}

func NewMultiInstanceAssociationGroupingsGet() MultiInstanceAssociationGroupingsGet {
	return MultiInstanceAssociationGroupingsGet{}
}

func (c *MultiInstanceAssociationGroupingsGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiInstanceAssociationGroupingsGet) Set() error {

	return nil
}

func (c *MultiInstanceAssociationGroupingsGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MultiInstanceAssociation),
		0x05,
		0x25,
	}
}
type MultiInstanceAssociationGroupingsReport struct {
    *report
    SupportedGroupings byte
    data []byte
}

func NewMultiInstanceAssociationGroupingsReport(data []byte) *MultiInstanceAssociationGroupingsReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiInstanceAssociationGroupingsReport{
				SupportedGroupings: data[0],
        data: data,
    }
}

type MultiInstanceAssociationRemove struct {
	node byte
	Groupingidentifier byte
	NodeID byte
	Marker byte
	Vg []MultiInstanceAssociationRemovevg
}
type MultiInstanceAssociationRemovevg struct {
	MultiInstanceNodeID byte
	Instance byte
}

func NewMultiInstanceAssociationRemove() MultiInstanceAssociationRemove {
	return MultiInstanceAssociationRemove{}
}

func (c *MultiInstanceAssociationRemove) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiInstanceAssociationRemove) Set(Groupingidentifier byte,NodeID byte,Marker byte,Vg []MultiInstanceAssociationRemovevg,) error {
	c.Groupingidentifier = Groupingidentifier
	c.NodeID = NodeID
	c.Marker = Marker
	c.Vg = Vg

	return nil
}

func (c *MultiInstanceAssociationRemove) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(7),
		byte(MultiInstanceAssociation),
		0x04,
		c.Groupingidentifier,
		c.NodeID,
		c.Marker,
	}
	for _, v := range c.Vg {
		data = append(data, v.MultiInstanceNodeID)
		data = append(data, v.Instance)
	}
	data = append(data, 0x25)
	return data
	
}
type MultiInstanceAssociationReport struct {
    *report
    GroupingIdentifier byte
    MaxNodesSupported byte
    ReportstoFollow byte
    NodeID byte
    Marker byte
    data []byte
}

func NewMultiInstanceAssociationReport(data []byte) *MultiInstanceAssociationReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiInstanceAssociationReport{
				GroupingIdentifier: data[0],
				MaxNodesSupported: data[1],
				ReportstoFollow: data[2],
				NodeID: data[3],
				Marker: data[4],
        data: data,
    }
}

type MultiInstanceAssociationSet struct {
	node byte
	Groupingidentifier byte
	NodeID byte
	Marker byte
	Vg []MultiInstanceAssociationSetvg
}
type MultiInstanceAssociationSetvg struct {
	MultiInstanceNodeID byte
	Instance byte
}

func NewMultiInstanceAssociationSet() MultiInstanceAssociationSet {
	return MultiInstanceAssociationSet{}
}

func (c *MultiInstanceAssociationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiInstanceAssociationSet) Set(Groupingidentifier byte,NodeID byte,Marker byte,Vg []MultiInstanceAssociationSetvg,) error {
	c.Groupingidentifier = Groupingidentifier
	c.NodeID = NodeID
	c.Marker = Marker
	c.Vg = Vg

	return nil
}

func (c *MultiInstanceAssociationSet) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(7),
		byte(MultiInstanceAssociation),
		0x01,
		c.Groupingidentifier,
		c.NodeID,
		c.Marker,
	}
	for _, v := range c.Vg {
		data = append(data, v.MultiInstanceNodeID)
		data = append(data, v.Instance)
	}
	data = append(data, 0x25)
	return data
	
}
