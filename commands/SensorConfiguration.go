package commands
type SensorTriggerLevelGet struct {
	node byte
}

func NewSensorTriggerLevelGet() SensorTriggerLevelGet {
	return SensorTriggerLevelGet{}
}

func (c *SensorTriggerLevelGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorTriggerLevelGet) Set() error {

	return nil
}

func (c *SensorTriggerLevelGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorConfiguration),
		0x02,
		0x25,
	}
}
type SensorTriggerLevelReport struct {
    *report
    SensorType byte
    Properties1 byte
    TriggerValue byte
    data []byte
}

func NewSensorTriggerLevelReport(data []byte) *SensorTriggerLevelReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorTriggerLevelReport{
				SensorType: data[0],
				Properties1: data[1],
				TriggerValue: data[2],
        data: data,
    }
}

type SensorTriggerLevelSet struct {
	node byte
	Properties1 byte
	SensorType byte
	Properties2 byte
	TriggerValue byte
}

func NewSensorTriggerLevelSet() SensorTriggerLevelSet {
	return SensorTriggerLevelSet{}
}

func (c *SensorTriggerLevelSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorTriggerLevelSet) Set(Properties1 byte,SensorType byte,Properties2 byte,TriggerValue byte,) error {
	c.Properties1 = Properties1
	c.SensorType = SensorType
	c.Properties2 = Properties2
	c.TriggerValue = TriggerValue

	return nil
}

func (c *SensorTriggerLevelSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(SensorConfiguration),
		0x01,
		c.Properties1,
		c.SensorType,
		c.Properties2,
		c.TriggerValue,
		0x25,
	}
}
