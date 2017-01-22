package commands
type MeterTblPushConfigurationGet struct {
	node byte
}

func NewMeterTblPushConfigurationGet() MeterTblPushConfigurationGet {
	return MeterTblPushConfigurationGet{}
}

func (c *MeterTblPushConfigurationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblPushConfigurationGet) Set() error {

	return nil
}

func (c *MeterTblPushConfigurationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterTblPush),
		0x02,
		0x25,
	}
}
type MeterTblPushConfigurationReport struct {
    *report
    Properties1 byte
    PushDataset byte
    IntervalMonths byte
    IntervalDays byte
    IntervalHours byte
    IntervalMinutes byte
    PushNodeID byte
    data []byte
}

func NewMeterTblPushConfigurationReport(data []byte) *MeterTblPushConfigurationReport {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblPushConfigurationReport{
				Properties1: data[0],
				PushDataset: data[1],
				IntervalMonths: data[2],
				IntervalDays: data[3],
				IntervalHours: data[4],
				IntervalMinutes: data[5],
				PushNodeID: data[6],
        data: data,
    }
}

type MeterTblPushConfigurationSet struct {
	node byte
	Properties1 byte
	PushDataset byte
	IntervalMonths byte
	IntervalDays byte
	IntervalHours byte
	IntervalMinutes byte
	PushNodeID byte
}

func NewMeterTblPushConfigurationSet() MeterTblPushConfigurationSet {
	return MeterTblPushConfigurationSet{}
}

func (c *MeterTblPushConfigurationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblPushConfigurationSet) Set(Properties1 byte,PushDataset byte,IntervalMonths byte,IntervalDays byte,IntervalHours byte,IntervalMinutes byte,PushNodeID byte,) error {
	c.Properties1 = Properties1
	c.PushDataset = PushDataset
	c.IntervalMonths = IntervalMonths
	c.IntervalDays = IntervalDays
	c.IntervalHours = IntervalHours
	c.IntervalMinutes = IntervalMinutes
	c.PushNodeID = PushNodeID

	return nil
}

func (c *MeterTblPushConfigurationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(9),
		byte(MeterTblPush),
		0x01,
		c.Properties1,
		c.PushDataset,
		c.IntervalMonths,
		c.IntervalDays,
		c.IntervalHours,
		c.IntervalMinutes,
		c.PushNodeID,
		0x25,
	}
}
