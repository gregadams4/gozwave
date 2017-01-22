package commands
type RemoteAssociationConfigurationGet struct {
	node byte
	LocalGroupingidentifier byte
}

func NewRemoteAssociationConfigurationGet() RemoteAssociationConfigurationGet {
	return RemoteAssociationConfigurationGet{}
}

func (c *RemoteAssociationConfigurationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *RemoteAssociationConfigurationGet) Set(LocalGroupingidentifier byte,) error {
	c.LocalGroupingidentifier = LocalGroupingidentifier

	return nil
}

func (c *RemoteAssociationConfigurationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(RemoteAssociation),
		0x02,
		c.LocalGroupingidentifier,
		0x25,
	}
}
type RemoteAssociationConfigurationReport struct {
    *report
    LocalGroupingidentifier byte
    RemoteNodeID byte
    RemoteGroupingidentifier byte
    data []byte
}

func NewRemoteAssociationConfigurationReport(data []byte) *RemoteAssociationConfigurationReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &RemoteAssociationConfigurationReport{
				LocalGroupingidentifier: data[0],
				RemoteNodeID: data[1],
				RemoteGroupingidentifier: data[2],
        data: data,
    }
}

type RemoteAssociationConfigurationSet struct {
	node byte
	LocalGroupingidentifier byte
	RemoteNodeID byte
	RemoteGroupingidentifier byte
}

func NewRemoteAssociationConfigurationSet() RemoteAssociationConfigurationSet {
	return RemoteAssociationConfigurationSet{}
}

func (c *RemoteAssociationConfigurationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *RemoteAssociationConfigurationSet) Set(LocalGroupingidentifier byte,RemoteNodeID byte,RemoteGroupingidentifier byte,) error {
	c.LocalGroupingidentifier = LocalGroupingidentifier
	c.RemoteNodeID = RemoteNodeID
	c.RemoteGroupingidentifier = RemoteGroupingidentifier

	return nil
}

func (c *RemoteAssociationConfigurationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(RemoteAssociation),
		0x01,
		c.LocalGroupingidentifier,
		c.RemoteNodeID,
		c.RemoteGroupingidentifier,
		0x25,
	}
}
