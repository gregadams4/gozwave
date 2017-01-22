package commands
type Initiate struct {
	node byte
	NodeID byte
	StepID byte
}

func NewInitiate() Initiate {
	return Initiate{}
}

func (c *Initiate) SetNode(node int) {
	c.node = byte(node)
}

func (c *Initiate) Set(NodeID byte,StepID byte,) error {
	c.NodeID = NodeID
	c.StepID = StepID

	return nil
}

func (c *Initiate) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(InclusionController),
		0x01,
		c.NodeID,
		c.StepID,
		0x25,
	}
}
type Complete struct {
	node byte
	StepID byte
	Status byte
}

func NewComplete() Complete {
	return Complete{}
}

func (c *Complete) SetNode(node int) {
	c.node = byte(node)
}

func (c *Complete) Set(StepID byte,Status byte,) error {
	c.StepID = StepID
	c.Status = Status

	return nil
}

func (c *Complete) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(InclusionController),
		0x02,
		c.StepID,
		c.Status,
		0x25,
	}
}
