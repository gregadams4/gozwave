package commands
type AlarmGetV2 struct {
	node byte
	AlarmType byte
	ZWaveAlarmType byte
}

func NewAlarmGetV2() AlarmGetV2 {
	return AlarmGetV2{}
}

func (c *AlarmGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AlarmGetV2) Set(AlarmType byte,ZWaveAlarmType byte,) error {
	c.AlarmType = AlarmType
	c.ZWaveAlarmType = ZWaveAlarmType

	return nil
}

func (c *AlarmGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(AlarmV2),
		0x04,
		c.AlarmType,
		c.ZWaveAlarmType,
		0x25,
	}
}
type AlarmReportV2 struct {
    *report
    AlarmType byte
    AlarmLevel byte
    ZensorNetSourceNodeID byte
    ZWaveAlarmStatus byte
    ZWaveAlarmType byte
    ZWaveAlarmEvent byte
    NumberofEventParameters byte
    EventParameter byte
    data []byte
}

func NewAlarmReportV2(data []byte) *AlarmReportV2 {
    if len(data) < 8 {
				//may want to change this to just return nil
				for i := len(data); i < 8; i++ {
            data = append(data, 0x00)
        }
    }

    return &AlarmReportV2{
				AlarmType: data[0],
				AlarmLevel: data[1],
				ZensorNetSourceNodeID: data[2],
				ZWaveAlarmStatus: data[3],
				ZWaveAlarmType: data[4],
				ZWaveAlarmEvent: data[5],
				NumberofEventParameters: data[6],
				EventParameter: data[7],
        data: data,
    }
}

type AlarmSetV2 struct {
	node byte
	ZWaveAlarmType byte
	ZWaveAlarmStatus byte
}

func NewAlarmSetV2() AlarmSetV2 {
	return AlarmSetV2{}
}

func (c *AlarmSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AlarmSetV2) Set(ZWaveAlarmType byte,ZWaveAlarmStatus byte,) error {
	c.ZWaveAlarmType = ZWaveAlarmType
	c.ZWaveAlarmStatus = ZWaveAlarmStatus

	return nil
}

func (c *AlarmSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(AlarmV2),
		0x06,
		c.ZWaveAlarmType,
		c.ZWaveAlarmStatus,
		0x25,
	}
}
type AlarmTypeSupportedGetV2 struct {
	node byte
}

func NewAlarmTypeSupportedGetV2() AlarmTypeSupportedGetV2 {
	return AlarmTypeSupportedGetV2{}
}

func (c *AlarmTypeSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *AlarmTypeSupportedGetV2) Set() error {

	return nil
}

func (c *AlarmTypeSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(AlarmV2),
		0x07,
		0x25,
	}
}
type AlarmTypeSupportedReportV2 struct {
    *report
    Properties1 byte
    BitMask byte
    data []byte
}

func NewAlarmTypeSupportedReportV2(data []byte) *AlarmTypeSupportedReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &AlarmTypeSupportedReportV2{
				Properties1: data[0],
				BitMask: data[1],
        data: data,
    }
}

