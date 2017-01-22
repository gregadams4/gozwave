package commands
type SceneActuatorConfGet struct {
	node byte
	SceneID byte
}

func NewSceneActuatorConfGet() SceneActuatorConfGet {
	return SceneActuatorConfGet{}
}

func (c *SceneActuatorConfGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SceneActuatorConfGet) Set(SceneID byte,) error {
	c.SceneID = SceneID

	return nil
}

func (c *SceneActuatorConfGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SceneActuatorConf),
		0x02,
		c.SceneID,
		0x25,
	}
}
type SceneActuatorConfReport struct {
    *report
    SceneID byte
    Level byte
    DimmingDuration byte
    data []byte
}

func NewSceneActuatorConfReport(data []byte) *SceneActuatorConfReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SceneActuatorConfReport{
				SceneID: data[0],
				Level: data[1],
				DimmingDuration: data[2],
        data: data,
    }
}

type SceneActuatorConfSet struct {
	node byte
	SceneID byte
	DimmingDuration byte
	Level2 byte
	Level byte
}

func NewSceneActuatorConfSet() SceneActuatorConfSet {
	return SceneActuatorConfSet{}
}

func (c *SceneActuatorConfSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SceneActuatorConfSet) Set(SceneID byte,DimmingDuration byte,Level2 byte,Level byte,) error {
	c.SceneID = SceneID
	c.DimmingDuration = DimmingDuration
	c.Level2 = Level2
	c.Level = Level

	return nil
}

func (c *SceneActuatorConfSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(SceneActuatorConf),
		0x01,
		c.SceneID,
		c.DimmingDuration,
		c.Level2,
		c.Level,
		0x25,
	}
}
