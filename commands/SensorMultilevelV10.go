package commands
type SensorMultilevelGetV10 struct {
	node byte
	SensorType byte
	Properties1 byte
}

func NewSensorMultilevelGetV10() SensorMultilevelGetV10 {
	return SensorMultilevelGetV10{}
}

func (c *SensorMultilevelGetV10) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelGetV10) Set(SensorType byte,Properties1 byte,) error {
	c.SensorType = SensorType
	c.Properties1 = Properties1

	return nil
}

func (c *SensorMultilevelGetV10) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SensorMultilevelV10),
		0x04,
		c.SensorType,
		c.Properties1,
		0x25,
	}
}
type SensorMultilevelReportV10 struct {
    *report
    SensorType byte
    Level byte
    SensorValue byte
    data []byte
}

func NewSensorMultilevelReportV10(data []byte) *SensorMultilevelReportV10 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelReportV10{
				SensorType: data[0],
				Level: data[1],
				SensorValue: data[2],
        data: data,
    }
}

type SensorMultilevelSupportedGetSensorV10 struct {
	node byte
}

func NewSensorMultilevelSupportedGetSensorV10() SensorMultilevelSupportedGetSensorV10 {
	return SensorMultilevelSupportedGetSensorV10{}
}

func (c *SensorMultilevelSupportedGetSensorV10) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelSupportedGetSensorV10) Set() error {

	return nil
}

func (c *SensorMultilevelSupportedGetSensorV10) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorMultilevelV10),
		0x01,
		0x25,
	}
}
type SensorMultilevelSupportedSensorReportV10 struct {
    *report
    BitMask byte
    data []byte
}

func NewSensorMultilevelSupportedSensorReportV10(data []byte) *SensorMultilevelSupportedSensorReportV10 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelSupportedSensorReportV10{
				BitMask: data[0],
        data: data,
    }
}

type SensorMultilevelSupportedGetScaleV10 struct {
	node byte
	SensorType byte
}

func NewSensorMultilevelSupportedGetScaleV10() SensorMultilevelSupportedGetScaleV10 {
	return SensorMultilevelSupportedGetScaleV10{}
}

func (c *SensorMultilevelSupportedGetScaleV10) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelSupportedGetScaleV10) Set(SensorType byte,) error {
	c.SensorType = SensorType

	return nil
}

func (c *SensorMultilevelSupportedGetScaleV10) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SensorMultilevelV10),
		0x03,
		c.SensorType,
		0x25,
	}
}
type SensorMultilevelSupportedScaleReportV10 struct {
    *report
    SensorType byte
    Properties1 byte
    data []byte
}

func NewSensorMultilevelSupportedScaleReportV10(data []byte) *SensorMultilevelSupportedScaleReportV10 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelSupportedScaleReportV10{
				SensorType: data[0],
				Properties1: data[1],
        data: data,
    }
}

