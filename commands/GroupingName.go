package commands
type GroupingNameGet struct {
	node byte
	Groupingidentifier byte
}

func NewGroupingNameGet() GroupingNameGet {
	return GroupingNameGet{}
}

func (c *GroupingNameGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *GroupingNameGet) Set(Groupingidentifier byte,) error {
	c.Groupingidentifier = Groupingidentifier

	return nil
}

func (c *GroupingNameGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(GroupingName),
		0x02,
		c.Groupingidentifier,
		0x25,
	}
}
type GroupingNameReport struct {
    *report
    Groupingidentifier byte
    Properties1 byte
    GroupingName byte
    data []byte
}

func NewGroupingNameReport(data []byte) *GroupingNameReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &GroupingNameReport{
				Groupingidentifier: data[0],
				Properties1: data[1],
				GroupingName: data[2],
        data: data,
    }
}

type GroupingNameSet struct {
	node byte
	Groupingidentifier byte
	Properties1 byte
	GroupingName byte
}

func NewGroupingNameSet() GroupingNameSet {
	return GroupingNameSet{}
}

func (c *GroupingNameSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *GroupingNameSet) Set(Groupingidentifier byte,Properties1 byte,GroupingName byte,) error {
	c.Groupingidentifier = Groupingidentifier
	c.Properties1 = Properties1
	c.GroupingName = GroupingName

	return nil
}

func (c *GroupingNameSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(GroupingName),
		0x01,
		c.Groupingidentifier,
		c.Properties1,
		c.GroupingName,
		0x25,
	}
}
