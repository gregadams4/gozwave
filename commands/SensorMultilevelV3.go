package commands
type SensorMultilevelGetV3 struct {
	node byte
}

func NewSensorMultilevelGetV3() SensorMultilevelGetV3 {
	return SensorMultilevelGetV3{}
}

func (c *SensorMultilevelGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelGetV3) Set() error {

	return nil
}

func (c *SensorMultilevelGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorMultilevelV3),
		0x04,
		0x25,
	}
}
type SensorMultilevelReportV3 struct {
    *report
    SensorType byte
    Level byte
    SensorValue byte
    data []byte
}

func NewSensorMultilevelReportV3(data []byte) *SensorMultilevelReportV3 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelReportV3{
				SensorType: data[0],
				Level: data[1],
				SensorValue: data[2],
        data: data,
    }
}

