package commands
type SensorMultilevelGetV2 struct {
	node byte
}

func NewSensorMultilevelGetV2() SensorMultilevelGetV2 {
	return SensorMultilevelGetV2{}
}

func (c *SensorMultilevelGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelGetV2) Set() error {

	return nil
}

func (c *SensorMultilevelGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorMultilevelV2),
		0x04,
		0x25,
	}
}
type SensorMultilevelReportV2 struct {
    *report
    SensorType byte
    Level byte
    SensorValue byte
    data []byte
}

func NewSensorMultilevelReportV2(data []byte) *SensorMultilevelReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelReportV2{
				SensorType: data[0],
				Level: data[1],
				SensorValue: data[2],
        data: data,
    }
}

