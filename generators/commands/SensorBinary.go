package commands
type SensorBinaryGet struct {
	node byte
}

func NewSensorBinaryGet() SensorBinaryGet {
	return SensorBinaryGet{}
}

func (c *SensorBinaryGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorBinaryGet) Set() error {

	return nil
}

func (c *SensorBinaryGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorBinary),
		0x02,
		0x25,
	}
}
type SensorBinaryReport struct {
    *report
    SensorValue byte
    data []byte
}

func NewSensorBinaryReport(data []byte) *SensorBinaryReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorBinaryReport{
				SensorValue: data[0],
        data: data,
    }
}

