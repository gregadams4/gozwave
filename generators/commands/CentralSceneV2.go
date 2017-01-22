package commands
type CentralSceneSupportedGetV2 struct {
	node byte
}

func NewCentralSceneSupportedGetV2() CentralSceneSupportedGetV2 {
	return CentralSceneSupportedGetV2{}
}

func (c *CentralSceneSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *CentralSceneSupportedGetV2) Set() error {

	return nil
}

func (c *CentralSceneSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(CentralSceneV2),
		0x01,
		0x25,
	}
}
type CentralSceneSupportedReportV2 struct {
    *report
    SupportedScenes byte
    Properties1 byte
    data []byte
}

func NewCentralSceneSupportedReportV2(data []byte) *CentralSceneSupportedReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &CentralSceneSupportedReportV2{
				SupportedScenes: data[0],
				Properties1: data[1],
        data: data,
    }
}

type CentralSceneNotificationV2 struct {
    *report
    SequenceNumber byte
    Properties1 byte
    SceneNumber byte
    data []byte
}

func NewCentralSceneNotificationV2(data []byte) *CentralSceneNotificationV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &CentralSceneNotificationV2{
				SequenceNumber: data[0],
				Properties1: data[1],
				SceneNumber: data[2],
        data: data,
    }
}

