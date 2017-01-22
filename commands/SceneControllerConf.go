package commands
type SceneControllerConfGet struct {
	node byte
	GroupID byte
}

func NewSceneControllerConfGet() SceneControllerConfGet {
	return SceneControllerConfGet{}
}

func (c *SceneControllerConfGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SceneControllerConfGet) Set(GroupID byte,) error {
	c.GroupID = GroupID

	return nil
}

func (c *SceneControllerConfGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SceneControllerConf),
		0x02,
		c.GroupID,
		0x25,
	}
}
type SceneControllerConfReport struct {
    *report
    GroupID byte
    SceneID byte
    DimmingDuration byte
    data []byte
}

func NewSceneControllerConfReport(data []byte) *SceneControllerConfReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SceneControllerConfReport{
				GroupID: data[0],
				SceneID: data[1],
				DimmingDuration: data[2],
        data: data,
    }
}

type SceneControllerConfSet struct {
	node byte
	GroupID byte
	SceneID byte
	DimmingDuration byte
}

func NewSceneControllerConfSet() SceneControllerConfSet {
	return SceneControllerConfSet{}
}

func (c *SceneControllerConfSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SceneControllerConfSet) Set(GroupID byte,SceneID byte,DimmingDuration byte,) error {
	c.GroupID = GroupID
	c.SceneID = SceneID
	c.DimmingDuration = DimmingDuration

	return nil
}

func (c *SceneControllerConfSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(SceneControllerConf),
		0x01,
		c.GroupID,
		c.SceneID,
		c.DimmingDuration,
		0x25,
	}
}
