package commands
type SensorMultilevelGetV9 struct {
	node byte
	SensorType byte
	Properties1 byte
}

func NewSensorMultilevelGetV9() SensorMultilevelGetV9 {
	return SensorMultilevelGetV9{}
}

func (c *SensorMultilevelGetV9) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelGetV9) Set(SensorType byte,Properties1 byte,) error {
	c.SensorType = SensorType
	c.Properties1 = Properties1

	return nil
}

func (c *SensorMultilevelGetV9) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SensorMultilevelV9),
		0x04,
		c.SensorType,
		c.Properties1,
		0x25,
	}
}
type SensorMultilevelReportV9 struct {
    *report
    SensorType byte
    Level byte
    SensorValue byte
    data []byte
}

func NewSensorMultilevelReportV9(data []byte) *SensorMultilevelReportV9 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelReportV9{
				SensorType: data[0],
				Level: data[1],
				SensorValue: data[2],
        data: data,
    }
}

type SensorMultilevelSupportedGetSensorV9 struct {
	node byte
}

func NewSensorMultilevelSupportedGetSensorV9() SensorMultilevelSupportedGetSensorV9 {
	return SensorMultilevelSupportedGetSensorV9{}
}

func (c *SensorMultilevelSupportedGetSensorV9) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelSupportedGetSensorV9) Set() error {

	return nil
}

func (c *SensorMultilevelSupportedGetSensorV9) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorMultilevelV9),
		0x01,
		0x25,
	}
}
type SensorMultilevelSupportedSensorReportV9 struct {
    *report
    BitMask byte
    data []byte
}

func NewSensorMultilevelSupportedSensorReportV9(data []byte) *SensorMultilevelSupportedSensorReportV9 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelSupportedSensorReportV9{
				BitMask: data[0],
        data: data,
    }
}

type SensorMultilevelSupportedGetScaleV9 struct {
	node byte
	SensorType byte
}

func NewSensorMultilevelSupportedGetScaleV9() SensorMultilevelSupportedGetScaleV9 {
	return SensorMultilevelSupportedGetScaleV9{}
}

func (c *SensorMultilevelSupportedGetScaleV9) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelSupportedGetScaleV9) Set(SensorType byte,) error {
	c.SensorType = SensorType

	return nil
}

func (c *SensorMultilevelSupportedGetScaleV9) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SensorMultilevelV9),
		0x03,
		c.SensorType,
		0x25,
	}
}
type SensorMultilevelSupportedScaleReportV9 struct {
    *report
    SensorType byte
    Properties1 byte
    data []byte
}

func NewSensorMultilevelSupportedScaleReportV9(data []byte) *SensorMultilevelSupportedScaleReportV9 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelSupportedScaleReportV9{
				SensorType: data[0],
				Properties1: data[1],
        data: data,
    }
}

