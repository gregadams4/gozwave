package commands
type MultiChannelAssociationGetV3 struct {
	node byte
	GroupingIdentifier byte
}

func NewMultiChannelAssociationGetV3() MultiChannelAssociationGetV3 {
	return MultiChannelAssociationGetV3{}
}

func (c *MultiChannelAssociationGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelAssociationGetV3) Set(GroupingIdentifier byte,) error {
	c.GroupingIdentifier = GroupingIdentifier

	return nil
}

func (c *MultiChannelAssociationGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MultiChannelAssociationV3),
		0x02,
		c.GroupingIdentifier,
		0x25,
	}
}
type MultiChannelAssociationGroupingsGetV3 struct {
	node byte
}

func NewMultiChannelAssociationGroupingsGetV3() MultiChannelAssociationGroupingsGetV3 {
	return MultiChannelAssociationGroupingsGetV3{}
}

func (c *MultiChannelAssociationGroupingsGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelAssociationGroupingsGetV3) Set() error {

	return nil
}

func (c *MultiChannelAssociationGroupingsGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MultiChannelAssociationV3),
		0x05,
		0x25,
	}
}
type MultiChannelAssociationGroupingsReportV3 struct {
    *report
    SupportedGroupings byte
    data []byte
}

func NewMultiChannelAssociationGroupingsReportV3(data []byte) *MultiChannelAssociationGroupingsReportV3 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelAssociationGroupingsReportV3{
				SupportedGroupings: data[0],
        data: data,
    }
}

type MultiChannelAssociationRemoveV3 struct {
	node byte
	GroupingIdentifier byte
	NodeID byte
	Marker byte
	Vg []MultiChannelAssociationRemoveV3vg
}
type MultiChannelAssociationRemoveV3vg struct {
	MultiChannelNodeID byte
	Properties1 byte
}

func NewMultiChannelAssociationRemoveV3() MultiChannelAssociationRemoveV3 {
	return MultiChannelAssociationRemoveV3{}
}

func (c *MultiChannelAssociationRemoveV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelAssociationRemoveV3) Set(GroupingIdentifier byte,NodeID byte,Marker byte,Vg []MultiChannelAssociationRemoveV3vg,) error {
	c.GroupingIdentifier = GroupingIdentifier
	c.NodeID = NodeID
	c.Marker = Marker
	c.Vg = Vg

	return nil
}

func (c *MultiChannelAssociationRemoveV3) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(7),
		byte(MultiChannelAssociationV3),
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
type MultiChannelAssociationReportV3 struct {
    *report
    GroupingIdentifier byte
    MaxNodesSupported byte
    ReportstoFollow byte
    NodeID byte
    Marker byte
    data []byte
}

func NewMultiChannelAssociationReportV3(data []byte) *MultiChannelAssociationReportV3 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &MultiChannelAssociationReportV3{
				GroupingIdentifier: data[0],
				MaxNodesSupported: data[1],
				ReportstoFollow: data[2],
				NodeID: data[3],
				Marker: data[4],
        data: data,
    }
}

type MultiChannelAssociationSetV3 struct {
	node byte
	GroupingIdentifier byte
	NodeID byte
	Marker byte
	Vg []MultiChannelAssociationSetV3vg
}
type MultiChannelAssociationSetV3vg struct {
	MultiChannelNodeID byte
	Properties1 byte
}

func NewMultiChannelAssociationSetV3() MultiChannelAssociationSetV3 {
	return MultiChannelAssociationSetV3{}
}

func (c *MultiChannelAssociationSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *MultiChannelAssociationSetV3) Set(GroupingIdentifier byte,NodeID byte,Marker byte,Vg []MultiChannelAssociationSetV3vg,) error {
	c.GroupingIdentifier = GroupingIdentifier
	c.NodeID = NodeID
	c.Marker = Marker
	c.Vg = Vg

	return nil
}

func (c *MultiChannelAssociationSetV3) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(7),
		byte(MultiChannelAssociationV3),
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
