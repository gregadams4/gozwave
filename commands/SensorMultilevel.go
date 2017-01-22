package commands
type SensorMultilevelGet struct {
	node byte
}

func NewSensorMultilevelGet() SensorMultilevelGet {
	return SensorMultilevelGet{}
}

func (c *SensorMultilevelGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelGet) Set() error {

	return nil
}

func (c *SensorMultilevelGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorMultilevel),
		0x04,
		0x25,
	}
}
type SensorMultilevelReport struct {
    *report
    SensorType byte
    Level byte
    SensorValue byte
    data []byte
}

func NewSensorMultilevelReport(data []byte) *SensorMultilevelReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelReport{
				SensorType: data[0],
				Level: data[1],
				SensorValue: data[2],
        data: data,
    }
}

