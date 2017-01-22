package commands
type CentralSceneSupportedGetV3 struct {
	node byte
}

func NewCentralSceneSupportedGetV3() CentralSceneSupportedGetV3 {
	return CentralSceneSupportedGetV3{}
}

func (c *CentralSceneSupportedGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *CentralSceneSupportedGetV3) Set() error {

	return nil
}

func (c *CentralSceneSupportedGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(CentralSceneV3),
		0x01,
		0x25,
	}
}
type CentralSceneSupportedReportV3 struct {
    *report
    SupportedScenes byte
    Properties1 byte
    data []byte
}

func NewCentralSceneSupportedReportV3(data []byte) *CentralSceneSupportedReportV3 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &CentralSceneSupportedReportV3{
				SupportedScenes: data[0],
				Properties1: data[1],
        data: data,
    }
}

type CentralSceneNotificationV3 struct {
    *report
    SequenceNumber byte
    Properties1 byte
    SceneNumber byte
    data []byte
}

func NewCentralSceneNotificationV3(data []byte) *CentralSceneNotificationV3 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &CentralSceneNotificationV3{
				SequenceNumber: data[0],
				Properties1: data[1],
				SceneNumber: data[2],
        data: data,
    }
}

type CentralSceneConfigurationSetV3 struct {
	node byte
	Properties1 byte
}

func NewCentralSceneConfigurationSetV3() CentralSceneConfigurationSetV3 {
	return CentralSceneConfigurationSetV3{}
}

func (c *CentralSceneConfigurationSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *CentralSceneConfigurationSetV3) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *CentralSceneConfigurationSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(CentralSceneV3),
		0x04,
		c.Properties1,
		0x25,
	}
}
type CentralSceneConfigurationGetV3 struct {
	node byte
}

func NewCentralSceneConfigurationGetV3() CentralSceneConfigurationGetV3 {
	return CentralSceneConfigurationGetV3{}
}

func (c *CentralSceneConfigurationGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *CentralSceneConfigurationGetV3) Set() error {

	return nil
}

func (c *CentralSceneConfigurationGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(CentralSceneV3),
		0x05,
		0x25,
	}
}
type CentralSceneConfigurationReportV3 struct {
    *report
    Properties1 byte
    data []byte
}

func NewCentralSceneConfigurationReportV3(data []byte) *CentralSceneConfigurationReportV3 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &CentralSceneConfigurationReportV3{
				Properties1: data[0],
        data: data,
    }
}

