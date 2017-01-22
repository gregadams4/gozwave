package commands
type BasicWindowCoveringStartLevelChange struct {
	node byte
	Level byte
}

func NewBasicWindowCoveringStartLevelChange() BasicWindowCoveringStartLevelChange {
	return BasicWindowCoveringStartLevelChange{}
}

func (c *BasicWindowCoveringStartLevelChange) SetNode(node int) {
	c.node = byte(node)
}

func (c *BasicWindowCoveringStartLevelChange) Set(Level byte,) error {
	c.Level = Level

	return nil
}

func (c *BasicWindowCoveringStartLevelChange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(BasicWindowCovering),
		0x01,
		c.Level,
		0x25,
	}
}
type BasicWindowCoveringStopLevelChange struct {
	node byte
}

func NewBasicWindowCoveringStopLevelChange() BasicWindowCoveringStopLevelChange {
	return BasicWindowCoveringStopLevelChange{}
}

func (c *BasicWindowCoveringStopLevelChange) SetNode(node int) {
	c.node = byte(node)
}

func (c *BasicWindowCoveringStopLevelChange) Set() error {

	return nil
}

func (c *BasicWindowCoveringStopLevelChange) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(BasicWindowCovering),
		0x02,
		0x25,
	}
}
