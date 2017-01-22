package commands
type SensorMultilevelGetV4 struct {
	node byte
}

func NewSensorMultilevelGetV4() SensorMultilevelGetV4 {
	return SensorMultilevelGetV4{}
}

func (c *SensorMultilevelGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelGetV4) Set() error {

	return nil
}

func (c *SensorMultilevelGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorMultilevelV4),
		0x04,
		0x25,
	}
}
type SensorMultilevelReportV4 struct {
    *report
    SensorType byte
    Level byte
    SensorValue byte
    data []byte
}

func NewSensorMultilevelReportV4(data []byte) *SensorMultilevelReportV4 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelReportV4{
				SensorType: data[0],
				Level: data[1],
				SensorValue: data[2],
        data: data,
    }
}

