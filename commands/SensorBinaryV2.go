package commands
type SensorBinaryGetV2 struct {
	node byte
	SensorType byte
}

func NewSensorBinaryGetV2() SensorBinaryGetV2 {
	return SensorBinaryGetV2{}
}

func (c *SensorBinaryGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorBinaryGetV2) Set(SensorType byte,) error {
	c.SensorType = SensorType

	return nil
}

func (c *SensorBinaryGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SensorBinaryV2),
		0x02,
		c.SensorType,
		0x25,
	}
}
type SensorBinaryReportV2 struct {
    *report
    SensorValue byte
    SensorType byte
    data []byte
}

func NewSensorBinaryReportV2(data []byte) *SensorBinaryReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorBinaryReportV2{
				SensorValue: data[0],
				SensorType: data[1],
        data: data,
    }
}

type SensorBinarySupportedGetSensorV2 struct {
	node byte
}

func NewSensorBinarySupportedGetSensorV2() SensorBinarySupportedGetSensorV2 {
	return SensorBinarySupportedGetSensorV2{}
}

func (c *SensorBinarySupportedGetSensorV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorBinarySupportedGetSensorV2) Set() error {

	return nil
}

func (c *SensorBinarySupportedGetSensorV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorBinaryV2),
		0x01,
		0x25,
	}
}
type SensorBinarySupportedSensorReportV2 struct {
    *report
    BitMask byte
    data []byte
}

func NewSensorBinarySupportedSensorReportV2(data []byte) *SensorBinarySupportedSensorReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorBinarySupportedSensorReportV2{
				BitMask: data[0],
        data: data,
    }
}

