package commands
type SceneActivationSet struct {
	node byte
	SceneID byte
	DimmingDuration byte
}

func NewSceneActivationSet() SceneActivationSet {
	return SceneActivationSet{}
}

func (c *SceneActivationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SceneActivationSet) Set(SceneID byte,DimmingDuration byte,) error {
	c.SceneID = SceneID
	c.DimmingDuration = DimmingDuration

	return nil
}

func (c *SceneActivationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SceneActivation),
		0x01,
		c.SceneID,
		c.DimmingDuration,
		0x25,
	}
}
