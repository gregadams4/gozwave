package commands
import "encoding/binary"
type SensorAlarmGet struct {
	node byte
	SensorType byte
}

func NewSensorAlarmGet() SensorAlarmGet {
	return SensorAlarmGet{}
}

func (c *SensorAlarmGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorAlarmGet) Set(SensorType byte,) error {
	c.SensorType = SensorType

	return nil
}

func (c *SensorAlarmGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SensorAlarm),
		0x01,
		c.SensorType,
		0x25,
	}
}
type SensorAlarmReport struct {
    *report
    SourceNodeID byte
    SensorType byte
    SensorState byte
    Seconds uint16
    data []byte
}

func NewSensorAlarmReport(data []byte) *SensorAlarmReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorAlarmReport{
				SourceNodeID: data[0],
				SensorType: data[1],
				SensorState: data[2],
				Seconds: binary.BigEndian.Uint16(data[3:5]),
        data: data,
    }
}

type SensorAlarmSupportedGet struct {
	node byte
}

func NewSensorAlarmSupportedGet() SensorAlarmSupportedGet {
	return SensorAlarmSupportedGet{}
}

func (c *SensorAlarmSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorAlarmSupportedGet) Set() error {

	return nil
}

func (c *SensorAlarmSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SensorAlarm),
		0x03,
		0x25,
	}
}
type SensorAlarmSupportedReport struct {
    *report
    NumberofBitMasks byte
    BitMask byte
    data []byte
}

func NewSensorAlarmSupportedReport(data []byte) *SensorAlarmSupportedReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SensorAlarmSupportedReport{
				NumberofBitMasks: data[0],
				BitMask: data[1],
        data: data,
    }
}

