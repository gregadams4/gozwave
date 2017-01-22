package commands
type SensorMultilevelGetV8 struct {
	node byte
	SensorType byte
	Properties1 byte
}

func NewSensorMultilevelGetV8() SensorMultilevelGetV8 {
	return SensorMultilevelGetV8{}
}

func (c *SensorMultilevelGetV8) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelGetV8) Set(SensorType byte,Properties1 byte,) error {
	c.SensorType = SensorType
	c.Properties1 = Properties1

	return nil
}

func (c *SensorMultilevelGetV8) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SensorMultilevelV8),
		0x04,
		c.SensorType,
		c.Properties1,
		0x25,
	}
}
type SensorMultilevelReportV8 struct {
    *report
    SensorType byte
    Level byte
    SensorValue byte
    data []byte
}

func NewSensorMultilevelReportV8(data []byte) *SensorMultilevelReportV8 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelReportV8{
				SensorType: data[0],
				Level: data[1],
				SensorValue: data[2],
        data: data,
    }
}

type SensorMultilevelSupportedGetSensorV8 struct {
	node byte
}

func NewSensorMultilevelSupportedGetSensorV8() SensorMultilevelSupportedGetSensorV8 {
	return SensorMultilevelSupportedGetSensorV8{}
}

func (c *SensorMultilevelSupportedGetSensorV8) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelSupportedGetSensorV8) Set() error {

	return nil
}

func (c *SensorMultilevelSupportedGetSensorV8) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorMultilevelV8),
		0x01,
		0x25,
	}
}
type SensorMultilevelSupportedSensorReportV8 struct {
    *report
    BitMask byte
    data []byte
}

func NewSensorMultilevelSupportedSensorReportV8(data []byte) *SensorMultilevelSupportedSensorReportV8 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelSupportedSensorReportV8{
				BitMask: data[0],
        data: data,
    }
}

type SensorMultilevelSupportedGetScaleV8 struct {
	node byte
	SensorType byte
}

func NewSensorMultilevelSupportedGetScaleV8() SensorMultilevelSupportedGetScaleV8 {
	return SensorMultilevelSupportedGetScaleV8{}
}

func (c *SensorMultilevelSupportedGetScaleV8) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelSupportedGetScaleV8) Set(SensorType byte,) error {
	c.SensorType = SensorType

	return nil
}

func (c *SensorMultilevelSupportedGetScaleV8) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SensorMultilevelV8),
		0x03,
		c.SensorType,
		0x25,
	}
}
type SensorMultilevelSupportedScaleReportV8 struct {
    *report
    SensorType byte
    Properties1 byte
    data []byte
}

func NewSensorMultilevelSupportedScaleReportV8(data []byte) *SensorMultilevelSupportedScaleReportV8 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelSupportedScaleReportV8{
				SensorType: data[0],
				Properties1: data[1],
        data: data,
    }
}

