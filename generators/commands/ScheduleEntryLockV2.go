package commands
type ScheduleEntryLockEnableAllSetV2 struct {
	node byte
	Enabled byte
}

func NewScheduleEntryLockEnableAllSetV2() ScheduleEntryLockEnableAllSetV2 {
	return ScheduleEntryLockEnableAllSetV2{}
}

func (c *ScheduleEntryLockEnableAllSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockEnableAllSetV2) Set(Enabled byte,) error {
	c.Enabled = Enabled

	return nil
}

func (c *ScheduleEntryLockEnableAllSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ScheduleEntryLockV2),
		0x02,
		c.Enabled,
		0x25,
	}
}
type ScheduleEntryLockEnableSetV2 struct {
	node byte
	UserIdentifier byte
	Enabled byte
}

func NewScheduleEntryLockEnableSetV2() ScheduleEntryLockEnableSetV2 {
	return ScheduleEntryLockEnableSetV2{}
}

func (c *ScheduleEntryLockEnableSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockEnableSetV2) Set(UserIdentifier byte,Enabled byte,) error {
	c.UserIdentifier = UserIdentifier
	c.Enabled = Enabled

	return nil
}

func (c *ScheduleEntryLockEnableSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleEntryLockV2),
		0x01,
		c.UserIdentifier,
		c.Enabled,
		0x25,
	}
}
type ScheduleEntryLockTimeOffsetGetV2 struct {
	node byte
}

func NewScheduleEntryLockTimeOffsetGetV2() ScheduleEntryLockTimeOffsetGetV2 {
	return ScheduleEntryLockTimeOffsetGetV2{}
}

func (c *ScheduleEntryLockTimeOffsetGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockTimeOffsetGetV2) Set() error {

	return nil
}

func (c *ScheduleEntryLockTimeOffsetGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ScheduleEntryLockV2),
		0x0B,
		0x25,
	}
}
type ScheduleEntryLockTimeOffsetReportV2 struct {
    *report
    Level byte
    MinuteTZO byte
    Level2 byte
    data []byte
}

func NewScheduleEntryLockTimeOffsetReportV2(data []byte) *ScheduleEntryLockTimeOffsetReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleEntryLockTimeOffsetReportV2{
				Level: data[0],
				MinuteTZO: data[1],
				Level2: data[2],
        data: data,
    }
}

type ScheduleEntryLockTimeOffsetSetV2 struct {
	node byte
	Level byte
	MinuteTZO byte
	Level2 byte
}

func NewScheduleEntryLockTimeOffsetSetV2() ScheduleEntryLockTimeOffsetSetV2 {
	return ScheduleEntryLockTimeOffsetSetV2{}
}

func (c *ScheduleEntryLockTimeOffsetSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockTimeOffsetSetV2) Set(Level byte,MinuteTZO byte,Level2 byte,) error {
	c.Level = Level
	c.MinuteTZO = MinuteTZO
	c.Level2 = Level2

	return nil
}

func (c *ScheduleEntryLockTimeOffsetSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ScheduleEntryLockV2),
		0x0D,
		c.Level,
		c.MinuteTZO,
		c.Level2,
		0x25,
	}
}
type ScheduleEntryLockWeekDayGetV2 struct {
	node byte
	UserIdentifier byte
	ScheduleSlotID byte
}

func NewScheduleEntryLockWeekDayGetV2() ScheduleEntryLockWeekDayGetV2 {
	return ScheduleEntryLockWeekDayGetV2{}
}

func (c *ScheduleEntryLockWeekDayGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockWeekDayGetV2) Set(UserIdentifier byte,ScheduleSlotID byte,) error {
	c.UserIdentifier = UserIdentifier
	c.ScheduleSlotID = ScheduleSlotID

	return nil
}

func (c *ScheduleEntryLockWeekDayGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleEntryLockV2),
		0x04,
		c.UserIdentifier,
		c.ScheduleSlotID,
		0x25,
	}
}
type ScheduleEntryLockWeekDayReportV2 struct {
    *report
    UserIdentifier byte
    ScheduleSlotID byte
    DayofWeek byte
    StartHour byte
    StartMinute byte
    StopHour byte
    StopMinute byte
    data []byte
}

func NewScheduleEntryLockWeekDayReportV2(data []byte) *ScheduleEntryLockWeekDayReportV2 {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleEntryLockWeekDayReportV2{
				UserIdentifier: data[0],
				ScheduleSlotID: data[1],
				DayofWeek: data[2],
				StartHour: data[3],
				StartMinute: data[4],
				StopHour: data[5],
				StopMinute: data[6],
        data: data,
    }
}

type ScheduleEntryLockWeekDaySetV2 struct {
	node byte
	SetAction byte
	UserIdentifier byte
	ScheduleSlotID byte
	DayofWeek byte
	StartHour byte
	StartMinute byte
	StopHour byte
	StopMinute byte
}

func NewScheduleEntryLockWeekDaySetV2() ScheduleEntryLockWeekDaySetV2 {
	return ScheduleEntryLockWeekDaySetV2{}
}

func (c *ScheduleEntryLockWeekDaySetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockWeekDaySetV2) Set(SetAction byte,UserIdentifier byte,ScheduleSlotID byte,DayofWeek byte,StartHour byte,StartMinute byte,StopHour byte,StopMinute byte,) error {
	c.SetAction = SetAction
	c.UserIdentifier = UserIdentifier
	c.ScheduleSlotID = ScheduleSlotID
	c.DayofWeek = DayofWeek
	c.StartHour = StartHour
	c.StartMinute = StartMinute
	c.StopHour = StopHour
	c.StopMinute = StopMinute

	return nil
}

func (c *ScheduleEntryLockWeekDaySetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(10),
		byte(ScheduleEntryLockV2),
		0x03,
		c.SetAction,
		c.UserIdentifier,
		c.ScheduleSlotID,
		c.DayofWeek,
		c.StartHour,
		c.StartMinute,
		c.StopHour,
		c.StopMinute,
		0x25,
	}
}
type ScheduleEntryLockYearDayGetV2 struct {
	node byte
	UserIdentifier byte
	ScheduleSlotID byte
}

func NewScheduleEntryLockYearDayGetV2() ScheduleEntryLockYearDayGetV2 {
	return ScheduleEntryLockYearDayGetV2{}
}

func (c *ScheduleEntryLockYearDayGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockYearDayGetV2) Set(UserIdentifier byte,ScheduleSlotID byte,) error {
	c.UserIdentifier = UserIdentifier
	c.ScheduleSlotID = ScheduleSlotID

	return nil
}

func (c *ScheduleEntryLockYearDayGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleEntryLockV2),
		0x07,
		c.UserIdentifier,
		c.ScheduleSlotID,
		0x25,
	}
}
type ScheduleEntryLockYearDayReportV2 struct {
    *report
    UserIdentifier byte
    ScheduleSlotID byte
    StartYear byte
    StartMonth byte
    StartDay byte
    StartHour byte
    StartMinute byte
    StopYear byte
    StopMonth byte
    StopDay byte
    StopHour byte
    StopMinute byte
    data []byte
}

func NewScheduleEntryLockYearDayReportV2(data []byte) *ScheduleEntryLockYearDayReportV2 {
    if len(data) < 12 {
				//may want to change this to just return nil
				for i := len(data); i < 12; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleEntryLockYearDayReportV2{
				UserIdentifier: data[0],
				ScheduleSlotID: data[1],
				StartYear: data[2],
				StartMonth: data[3],
				StartDay: data[4],
				StartHour: data[5],
				StartMinute: data[6],
				StopYear: data[7],
				StopMonth: data[8],
				StopDay: data[9],
				StopHour: data[10],
				StopMinute: data[11],
        data: data,
    }
}

type ScheduleEntryLockYearDaySetV2 struct {
	node byte
	SetAction byte
	UserIdentifier byte
	ScheduleSlotID byte
	StartYear byte
	StartMonth byte
	StartDay byte
	StartHour byte
	StartMinute byte
	StopYear byte
	StopMonth byte
	StopDay byte
	StopHour byte
	StopMinute byte
}

func NewScheduleEntryLockYearDaySetV2() ScheduleEntryLockYearDaySetV2 {
	return ScheduleEntryLockYearDaySetV2{}
}

func (c *ScheduleEntryLockYearDaySetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockYearDaySetV2) Set(SetAction byte,UserIdentifier byte,ScheduleSlotID byte,StartYear byte,StartMonth byte,StartDay byte,StartHour byte,StartMinute byte,StopYear byte,StopMonth byte,StopDay byte,StopHour byte,StopMinute byte,) error {
	c.SetAction = SetAction
	c.UserIdentifier = UserIdentifier
	c.ScheduleSlotID = ScheduleSlotID
	c.StartYear = StartYear
	c.StartMonth = StartMonth
	c.StartDay = StartDay
	c.StartHour = StartHour
	c.StartMinute = StartMinute
	c.StopYear = StopYear
	c.StopMonth = StopMonth
	c.StopDay = StopDay
	c.StopHour = StopHour
	c.StopMinute = StopMinute

	return nil
}

func (c *ScheduleEntryLockYearDaySetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(15),
		byte(ScheduleEntryLockV2),
		0x06,
		c.SetAction,
		c.UserIdentifier,
		c.ScheduleSlotID,
		c.StartYear,
		c.StartMonth,
		c.StartDay,
		c.StartHour,
		c.StartMinute,
		c.StopYear,
		c.StopMonth,
		c.StopDay,
		c.StopHour,
		c.StopMinute,
		0x25,
	}
}
type ScheduleEntryTypeSupportedGetV2 struct {
	node byte
}

func NewScheduleEntryTypeSupportedGetV2() ScheduleEntryTypeSupportedGetV2 {
	return ScheduleEntryTypeSupportedGetV2{}
}

func (c *ScheduleEntryTypeSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryTypeSupportedGetV2) Set() error {

	return nil
}

func (c *ScheduleEntryTypeSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ScheduleEntryLockV2),
		0x09,
		0x25,
	}
}
type ScheduleEntryTypeSupportedReportV2 struct {
    *report
    NumberofSlotsWeekDay byte
    NumberofSlotsYearDay byte
    data []byte
}

func NewScheduleEntryTypeSupportedReportV2(data []byte) *ScheduleEntryTypeSupportedReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleEntryTypeSupportedReportV2{
				NumberofSlotsWeekDay: data[0],
				NumberofSlotsYearDay: data[1],
        data: data,
    }
}

