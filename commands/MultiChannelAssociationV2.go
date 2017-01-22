package commands
type MultiChannelAssociationGetV2 struct {
	node byte
	GroupingIdentifier byte
}

func NewMultiChannelAssociationGetV2() MultiChannelAssociationGetV2 {
	return MultiChannelAssociationGetV2{}
}

func (c *MultiChannelAssociationGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelAssociationGetV2) Set(GroupingIdentifier byte,) error {
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *MultiChannelAssociationGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MultiChannelAssociationV2),
		0x02,
		c.GroupingIdentifier,
		0x25,
	}
}
type MultiChannelAssociationGroupingsGetV2 struct {
	node byte
}

func NewMultiChannelAssociationGroupingsGetV2() MultiChannelAssociationGroupingsGetV2 {
	return MultiChannelAssociationGroupingsGetV2{}
}

func (c *MultiChannelAssociationGroupingsGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelAssociationGroupingsGetV2) Set() error {

	return nil
}

func (c *MultiChannelAssociationGroupingsGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MultiChannelAssociationV2),
		0x05,
		0x25,
	}
}
type MultiChannelAssociationGroupingsReportV2 struct {
    *report
    SupportedGroupings byte
    data []byte
}

func NewMultiChannelAssociationGroupingsReportV2(data []byte) *MultiChannelAssociationGroupingsReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelAssociationGroupingsReportV2{
				SupportedGroupings: data[0],
        data: data,
    }
}

type MultiChannelAssociationRemoveV2 struct {
	node byte
	GroupingIdentifier byte
	NodeID byte
	Marker byte
	Vg []MultiChannelAssociationRemoveV2vg
}
type MultiChannelAssociationRemoveV2vg struct {
	MultiChannelNodeID byte
	Properties1 byte
}

func NewMultiChannelAssociationRemoveV2() MultiChannelAssociationRemoveV2 {
	return MultiChannelAssociationRemoveV2{}
}

func (c *MultiChannelAssociationRemoveV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelAssociationRemoveV2) Set(GroupingIdentifier byte,NodeID byte,Marker byte,Vg []MultiChannelAssociationRemoveV2vg,) error {
	c.GroupingIdentifier = GroupingIdentifier
	c.NodeID = NodeID
	c.Marker = Marker
	c.Vg = Vg

	return nil
}

func (c *MultiChannelAssociationRemoveV2) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(7),
		byte(MultiChannelAssociationV2),
		0x04,
		c.GroupingIdentifier,
		c.NodeID,
		c.Marker,
	}
	for _, v := range c.Vg {
		data = append(data, v.MultiChannelNodeID)
		data = append(data, v.Properties1)
	}
	data = append(data, 0x25)
	return data
	
}
type MultiChannelAssociationReportV2 struct {
    *report
    GroupingIdentifier byte
    MaxNodesSupported byte
    ReportstoFollow byte
    NodeID byte
    Marker byte
    data []byte
}

func NewMultiChannelAssociationReportV2(data []byte) *MultiChannelAssociationReportV2 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelAssociationReportV2{
				GroupingIdentifier: data[0],
				MaxNodesSupported: data[1],
				ReportstoFollow: data[2],
				NodeID: data[3],
				Marker: data[4],
        data: data,
    }
}

type MultiChannelAssociationSetV2 struct {
	node byte
	GroupingIdentifier byte
	NodeID byte
	Marker byte
	Vg []MultiChannelAssociationSetV2vg
}
type MultiChannelAssociationSetV2vg struct {
	MultiChannelNodeID byte
	Properties1 byte
}

func NewMultiChannelAssociationSetV2() MultiChannelAssociationSetV2 {
	return MultiChannelAssociationSetV2{}
}

func (c *MultiChannelAssociationSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelAssociationSetV2) Set(GroupingIdentifier byte,NodeID byte,Marker byte,Vg []MultiChannelAssociationSetV2vg,) error {
	c.GroupingIdentifier = GroupingIdentifier
	c.NodeID = NodeID
	c.Marker = Marker
	c.Vg = Vg

	return nil
}

func (c *MultiChannelAssociationSetV2) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(7),
		byte(MultiChannelAssociationV2),
		0x01,
		c.GroupingIdentifier,
		c.NodeID,
		c.Marker,
	}
	for _, v := range c.Vg {
		data = append(data, v.MultiChannelNodeID)
		data = append(data, v.Properties1)
	}
	data = append(data, 0x25)
	return data
	
}
