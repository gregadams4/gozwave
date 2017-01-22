package commands
type SensorMultilevelGetV7 struct {
	node byte
	SensorType byte
	Properties1 byte
}

func NewSensorMultilevelGetV7() SensorMultilevelGetV7 {
	return SensorMultilevelGetV7{}
}

func (c *SensorMultilevelGetV7) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelGetV7) Set(SensorType byte,Properties1 byte,) error {
	c.SensorType = SensorType
	c.Properties1 = Properties1

	return nil
}

func (c *SensorMultilevelGetV7) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SensorMultilevelV7),
		0x04,
		c.SensorType,
		c.Properties1,
		0x25,
	}
}
type SensorMultilevelReportV7 struct {
    *report
    SensorType byte
    Level byte
    SensorValue byte
    data []byte
}

func NewSensorMultilevelReportV7(data []byte) *SensorMultilevelReportV7 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelReportV7{
				SensorType: data[0],
				Level: data[1],
				SensorValue: data[2],
        data: data,
    }
}

type SensorMultilevelSupportedGetSensorV7 struct {
	node byte
}

func NewSensorMultilevelSupportedGetSensorV7() SensorMultilevelSupportedGetSensorV7 {
	return SensorMultilevelSupportedGetSensorV7{}
}

func (c *SensorMultilevelSupportedGetSensorV7) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelSupportedGetSensorV7) Set() error {

	return nil
}

func (c *SensorMultilevelSupportedGetSensorV7) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorMultilevelV7),
		0x01,
		0x25,
	}
}
type SensorMultilevelSupportedSensorReportV7 struct {
    *report
    BitMask byte
    data []byte
}

func NewSensorMultilevelSupportedSensorReportV7(data []byte) *SensorMultilevelSupportedSensorReportV7 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelSupportedSensorReportV7{
				BitMask: data[0],
        data: data,
    }
}

type SensorMultilevelSupportedGetScaleV7 struct {
	node byte
	SensorType byte
}

func NewSensorMultilevelSupportedGetScaleV7() SensorMultilevelSupportedGetScaleV7 {
	return SensorMultilevelSupportedGetScaleV7{}
}

func (c *SensorMultilevelSupportedGetScaleV7) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelSupportedGetScaleV7) Set(SensorType byte,) error {
	c.SensorType = SensorType

	return nil
}

func (c *SensorMultilevelSupportedGetScaleV7) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SensorMultilevelV7),
		0x03,
		c.SensorType,
		0x25,
	}
}
type SensorMultilevelSupportedScaleReportV7 struct {
    *report
    SensorType byte
    Properties1 byte
    data []byte
}

func NewSensorMultilevelSupportedScaleReportV7(data []byte) *SensorMultilevelSupportedScaleReportV7 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelSupportedScaleReportV7{
				SensorType: data[0],
				Properties1: data[1],
        data: data,
    }
}

