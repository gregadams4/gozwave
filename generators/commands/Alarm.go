package commands
type AlarmGet struct {
	node byte
	AlarmType byte
}

func NewAlarmGet() AlarmGet {
	return AlarmGet{}
}

func (c *AlarmGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AlarmGet) Set(AlarmType byte,) error {
	c.AlarmType = AlarmType

	return nil
}

func (c *AlarmGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Alarm),
		0x04,
		c.AlarmType,
		0x25,
	}
}
type AlarmReport struct {
    *report
    AlarmType byte
    AlarmLevel byte
    data []byte
}

func NewAlarmReport(data []byte) *AlarmReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &AlarmReport{
				AlarmType: data[0],
				AlarmLevel: data[1],
        data: data,
    }
}

