package commands
type SensorMultilevelGetV5 struct {
	node byte
	SensorType byte
	Properties1 byte
}

func NewSensorMultilevelGetV5() SensorMultilevelGetV5 {
	return SensorMultilevelGetV5{}
}

func (c *SensorMultilevelGetV5) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelGetV5) Set(SensorType byte,Properties1 byte,) error {
	c.SensorType = SensorType
	c.Properties1 = Properties1

	return nil
}

func (c *SensorMultilevelGetV5) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SensorMultilevelV5),
		0x04,
		c.SensorType,
		c.Properties1,
		0x25,
	}
}
type SensorMultilevelReportV5 struct {
    *report
    SensorType byte
    Level byte
    SensorValue byte
    data []byte
}

func NewSensorMultilevelReportV5(data []byte) *SensorMultilevelReportV5 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelReportV5{
				SensorType: data[0],
				Level: data[1],
				SensorValue: data[2],
        data: data,
    }
}

type SensorMultilevelSupportedGetSensorV5 struct {
	node byte
}

func NewSensorMultilevelSupportedGetSensorV5() SensorMultilevelSupportedGetSensorV5 {
	return SensorMultilevelSupportedGetSensorV5{}
}

func (c *SensorMultilevelSupportedGetSensorV5) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelSupportedGetSensorV5) Set() error {

	return nil
}

func (c *SensorMultilevelSupportedGetSensorV5) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorMultilevelV5),
		0x01,
		0x25,
	}
}
type SensorMultilevelSupportedSensorReportV5 struct {
    *report
    BitMask byte
    data []byte
}

func NewSensorMultilevelSupportedSensorReportV5(data []byte) *SensorMultilevelSupportedSensorReportV5 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelSupportedSensorReportV5{
				BitMask: data[0],
        data: data,
    }
}

type SensorMultilevelSupportedGetScaleV5 struct {
	node byte
	SensorType byte
}

func NewSensorMultilevelSupportedGetScaleV5() SensorMultilevelSupportedGetScaleV5 {
	return SensorMultilevelSupportedGetScaleV5{}
}

func (c *SensorMultilevelSupportedGetScaleV5) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelSupportedGetScaleV5) Set(SensorType byte,) error {
	c.SensorType = SensorType

	return nil
}

func (c *SensorMultilevelSupportedGetScaleV5) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SensorMultilevelV5),
		0x03,
		c.SensorType,
		0x25,
	}
}
type SensorMultilevelSupportedScaleReportV5 struct {
    *report
    SensorType byte
    Properties1 byte
    data []byte
}

func NewSensorMultilevelSupportedScaleReportV5(data []byte) *SensorMultilevelSupportedScaleReportV5 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelSupportedScaleReportV5{
				SensorType: data[0],
				Properties1: data[1],
        data: data,
    }
}

