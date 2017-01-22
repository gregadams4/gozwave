package commands
type SensorMultilevelGetV6 struct {
	node byte
	SensorType byte
	Properties1 byte
}

func NewSensorMultilevelGetV6() SensorMultilevelGetV6 {
	return SensorMultilevelGetV6{}
}

func (c *SensorMultilevelGetV6) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelGetV6) Set(SensorType byte,Properties1 byte,) error {
	c.SensorType = SensorType
	c.Properties1 = Properties1

	return nil
}

func (c *SensorMultilevelGetV6) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(SensorMultilevelV6),
		0x04,
		c.SensorType,
		c.Properties1,
		0x25,
	}
}
type SensorMultilevelReportV6 struct {
    *report
    SensorType byte
    Level byte
    SensorValue byte
    data []byte
}

func NewSensorMultilevelReportV6(data []byte) *SensorMultilevelReportV6 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelReportV6{
				SensorType: data[0],
				Level: data[1],
				SensorValue: data[2],
        data: data,
    }
}

type SensorMultilevelSupportedGetSensorV6 struct {
	node byte
}

func NewSensorMultilevelSupportedGetSensorV6() SensorMultilevelSupportedGetSensorV6 {
	return SensorMultilevelSupportedGetSensorV6{}
}

func (c *SensorMultilevelSupportedGetSensorV6) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelSupportedGetSensorV6) Set() error {

	return nil
}

func (c *SensorMultilevelSupportedGetSensorV6) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorMultilevelV6),
		0x01,
		0x25,
	}
}
type SensorMultilevelSupportedSensorReportV6 struct {
    *report
    BitMask byte
    data []byte
}

func NewSensorMultilevelSupportedSensorReportV6(data []byte) *SensorMultilevelSupportedSensorReportV6 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelSupportedSensorReportV6{
				BitMask: data[0],
        data: data,
    }
}

type SensorMultilevelSupportedGetScaleV6 struct {
	node byte
	SensorType byte
}

func NewSensorMultilevelSupportedGetScaleV6() SensorMultilevelSupportedGetScaleV6 {
	return SensorMultilevelSupportedGetScaleV6{}
}

func (c *SensorMultilevelSupportedGetScaleV6) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorMultilevelSupportedGetScaleV6) Set(SensorType byte,) error {
	c.SensorType = SensorType

	return nil
}

func (c *SensorMultilevelSupportedGetScaleV6) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SensorMultilevelV6),
		0x03,
		c.SensorType,
		0x25,
	}
}
type SensorMultilevelSupportedScaleReportV6 struct {
    *report
    SensorType byte
    Properties1 byte
    data []byte
}

func NewSensorMultilevelSupportedScaleReportV6(data []byte) *SensorMultilevelSupportedScaleReportV6 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorMultilevelSupportedScaleReportV6{
				SensorType: data[0],
				Properties1: data[1],
        data: data,
    }
}

