package commands
type CentralSceneSupportedGet struct {
	node byte
}

func NewCentralSceneSupportedGet() CentralSceneSupportedGet {
	return CentralSceneSupportedGet{}
}

func (c *CentralSceneSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *CentralSceneSupportedGet) Set() error {

	return nil
}

func (c *CentralSceneSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(CentralScene),
		0x01,
		0x25,
	}
}
type CentralSceneSupportedReport struct {
    *report
    SupportedScenes byte
    data []byte
}

func NewCentralSceneSupportedReport(data []byte) *CentralSceneSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &CentralSceneSupportedReport{
				SupportedScenes: data[0],
        data: data,
    }
}

type CentralSceneNotification struct {
    *report
    SequenceNumber byte
    Properties1 byte
    SceneNumber byte
    data []byte
}

func NewCentralSceneNotification(data []byte) *CentralSceneNotification {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &CentralSceneNotification{
				SequenceNumber: data[0],
				Properties1: data[1],
				SceneNumber: data[2],
        data: data,
    }
}

