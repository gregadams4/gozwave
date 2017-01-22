package commands
type ScheduleEntryLockEnableAllSetV3 struct {
	node byte
	Enabled byte
}

func NewScheduleEntryLockEnableAllSetV3() ScheduleEntryLockEnableAllSetV3 {
	return ScheduleEntryLockEnableAllSetV3{}
}

func (c *ScheduleEntryLockEnableAllSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockEnableAllSetV3) Set(Enabled byte,) error {
	c.Enabled = Enabled

	return nil
}

func (c *ScheduleEntryLockEnableAllSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ScheduleEntryLockV3),
		0x02,
		c.Enabled,
		0x25,
	}
}
type ScheduleEntryLockEnableSetV3 struct {
	node byte
	UserIdentifier byte
	Enabled byte
}

func NewScheduleEntryLockEnableSetV3() ScheduleEntryLockEnableSetV3 {
	return ScheduleEntryLockEnableSetV3{}
}

func (c *ScheduleEntryLockEnableSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockEnableSetV3) Set(UserIdentifier byte,Enabled byte,) error {
	c.UserIdentifier = UserIdentifier
	c.Enabled = Enabled

	return nil
}

func (c *ScheduleEntryLockEnableSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleEntryLockV3),
		0x01,
		c.UserIdentifier,
		c.Enabled,
		0x25,
	}
}
type ScheduleEntryLockTimeOffsetGetV3 struct {
	node byte
}

func NewScheduleEntryLockTimeOffsetGetV3() ScheduleEntryLockTimeOffsetGetV3 {
	return ScheduleEntryLockTimeOffsetGetV3{}
}

func (c *ScheduleEntryLockTimeOffsetGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockTimeOffsetGetV3) Set() error {

	return nil
}

func (c *ScheduleEntryLockTimeOffsetGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ScheduleEntryLockV3),
		0x0B,
		0x25,
	}
}
type ScheduleEntryLockTimeOffsetReportV3 struct {
    *report
    Level byte
    MinuteTZO byte
    Level2 byte
    data []byte
}

func NewScheduleEntryLockTimeOffsetReportV3(data []byte) *ScheduleEntryLockTimeOffsetReportV3 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleEntryLockTimeOffsetReportV3{
				Level: data[0],
				MinuteTZO: data[1],
				Level2: data[2],
        data: data,
    }
}

type ScheduleEntryLockTimeOffsetSetV3 struct {
	node byte
	Level byte
	MinuteTZO byte
	Level2 byte
}

func NewScheduleEntryLockTimeOffsetSetV3() ScheduleEntryLockTimeOffsetSetV3 {
	return ScheduleEntryLockTimeOffsetSetV3{}
}

func (c *ScheduleEntryLockTimeOffsetSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockTimeOffsetSetV3) Set(Level byte,MinuteTZO byte,Level2 byte,) error {
	c.Level = Level
	c.MinuteTZO = MinuteTZO
	c.Level2 = Level2

	return nil
}

func (c *ScheduleEntryLockTimeOffsetSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ScheduleEntryLockV3),
		0x0D,
		c.Level,
		c.MinuteTZO,
		c.Level2,
		0x25,
	}
}
type ScheduleEntryLockWeekDayGetV3 struct {
	node byte
	UserIdentifier byte
	ScheduleSlotID byte
}

func NewScheduleEntryLockWeekDayGetV3() ScheduleEntryLockWeekDayGetV3 {
	return ScheduleEntryLockWeekDayGetV3{}
}

func (c *ScheduleEntryLockWeekDayGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockWeekDayGetV3) Set(UserIdentifier byte,ScheduleSlotID byte,) error {
	c.UserIdentifier = UserIdentifier
	c.ScheduleSlotID = ScheduleSlotID

	return nil
}

func (c *ScheduleEntryLockWeekDayGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleEntryLockV3),
		0x04,
		c.UserIdentifier,
		c.ScheduleSlotID,
		0x25,
	}
}
type ScheduleEntryLockWeekDayReportV3 struct {
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

func NewScheduleEntryLockWeekDayReportV3(data []byte) *ScheduleEntryLockWeekDayReportV3 {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleEntryLockWeekDayReportV3{
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

type ScheduleEntryLockWeekDaySetV3 struct {
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

func NewScheduleEntryLockWeekDaySetV3() ScheduleEntryLockWeekDaySetV3 {
	return ScheduleEntryLockWeekDaySetV3{}
}

func (c *ScheduleEntryLockWeekDaySetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockWeekDaySetV3) Set(SetAction byte,UserIdentifier byte,ScheduleSlotID byte,DayofWeek byte,StartHour byte,StartMinute byte,StopHour byte,StopMinute byte,) error {
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

func (c *ScheduleEntryLockWeekDaySetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(10),
		byte(ScheduleEntryLockV3),
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
type ScheduleEntryLockYearDayGetV3 struct {
	node byte
	UserIdentifier byte
	ScheduleSlotID byte
}

func NewScheduleEntryLockYearDayGetV3() ScheduleEntryLockYearDayGetV3 {
	return ScheduleEntryLockYearDayGetV3{}
}

func (c *ScheduleEntryLockYearDayGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockYearDayGetV3) Set(UserIdentifier byte,ScheduleSlotID byte,) error {
	c.UserIdentifier = UserIdentifier
	c.ScheduleSlotID = ScheduleSlotID

	return nil
}

func (c *ScheduleEntryLockYearDayGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleEntryLockV3),
		0x07,
		c.UserIdentifier,
		c.ScheduleSlotID,
		0x25,
	}
}
type ScheduleEntryLockYearDayReportV3 struct {
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

func NewScheduleEntryLockYearDayReportV3(data []byte) *ScheduleEntryLockYearDayReportV3 {
    if len(data) < 12 {
				//may want to change this to just return nil
				for i := len(data); i < 12; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleEntryLockYearDayReportV3{
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

type ScheduleEntryLockYearDaySetV3 struct {
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

func NewScheduleEntryLockYearDaySetV3() ScheduleEntryLockYearDaySetV3 {
	return ScheduleEntryLockYearDaySetV3{}
}

func (c *ScheduleEntryLockYearDaySetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockYearDaySetV3) Set(SetAction byte,UserIdentifier byte,ScheduleSlotID byte,StartYear byte,StartMonth byte,StartDay byte,StartHour byte,StartMinute byte,StopYear byte,StopMonth byte,StopDay byte,StopHour byte,StopMinute byte,) error {
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

func (c *ScheduleEntryLockYearDaySetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(15),
		byte(ScheduleEntryLockV3),
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
type ScheduleEntryTypeSupportedGetV3 struct {
	node byte
}

func NewScheduleEntryTypeSupportedGetV3() ScheduleEntryTypeSupportedGetV3 {
	return ScheduleEntryTypeSupportedGetV3{}
}

func (c *ScheduleEntryTypeSupportedGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryTypeSupportedGetV3) Set() error {

	return nil
}

func (c *ScheduleEntryTypeSupportedGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ScheduleEntryLockV3),
		0x09,
		0x25,
	}
}
type ScheduleEntryTypeSupportedReportV3 struct {
    *report
    NumberofSlotsWeekDay byte
    NumberofSlotsYearDay byte
    NumberofSlotsDailyRepeating byte
    data []byte
}

func NewScheduleEntryTypeSupportedReportV3(data []byte) *ScheduleEntryTypeSupportedReportV3 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleEntryTypeSupportedReportV3{
				NumberofSlotsWeekDay: data[0],
				NumberofSlotsYearDay: data[1],
				NumberofSlotsDailyRepeating: data[2],
        data: data,
    }
}

type ScheduleEntryLockDailyRepeatingGetV3 struct {
	node byte
	UserIdentifier byte
	ScheduleSlotID byte
}

func NewScheduleEntryLockDailyRepeatingGetV3() ScheduleEntryLockDailyRepeatingGetV3 {
	return ScheduleEntryLockDailyRepeatingGetV3{}
}

func (c *ScheduleEntryLockDailyRepeatingGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockDailyRepeatingGetV3) Set(UserIdentifier byte,ScheduleSlotID byte,) error {
	c.UserIdentifier = UserIdentifier
	c.ScheduleSlotID = ScheduleSlotID

	return nil
}

func (c *ScheduleEntryLockDailyRepeatingGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleEntryLockV3),
		0x0E,
		c.UserIdentifier,
		c.ScheduleSlotID,
		0x25,
	}
}
type ScheduleEntryLockDailyRepeatingReportV3 struct {
    *report
    UserIdentifier byte
    ScheduleSlotID byte
    WeekDayBitmask byte
    StartHour byte
    StartMinute byte
    DurationHour byte
    DurationMinute byte
    data []byte
}

func NewScheduleEntryLockDailyRepeatingReportV3(data []byte) *ScheduleEntryLockDailyRepeatingReportV3 {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleEntryLockDailyRepeatingReportV3{
				UserIdentifier: data[0],
				ScheduleSlotID: data[1],
				WeekDayBitmask: data[2],
				StartHour: data[3],
				StartMinute: data[4],
				DurationHour: data[5],
				DurationMinute: data[6],
        data: data,
    }
}

type ScheduleEntryLockDailyRepeatingSetV3 struct {
	node byte
	SetAction byte
	UserIdentifier byte
	ScheduleSlotID byte
	WeekDayBitmask byte
	StartHour byte
	StartMinute byte
	DurationHour byte
	DurationMinute byte
}

func NewScheduleEntryLockDailyRepeatingSetV3() ScheduleEntryLockDailyRepeatingSetV3 {
	return ScheduleEntryLockDailyRepeatingSetV3{}
}

func (c *ScheduleEntryLockDailyRepeatingSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockDailyRepeatingSetV3) Set(SetAction byte,UserIdentifier byte,ScheduleSlotID byte,WeekDayBitmask byte,StartHour byte,StartMinute byte,DurationHour byte,DurationMinute byte,) error {
	c.SetAction = SetAction
	c.UserIdentifier = UserIdentifier
	c.ScheduleSlotID = ScheduleSlotID
	c.WeekDayBitmask = WeekDayBitmask
	c.StartHour = StartHour
	c.StartMinute = StartMinute
	c.DurationHour = DurationHour
	c.DurationMinute = DurationMinute

	return nil
}

func (c *ScheduleEntryLockDailyRepeatingSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(10),
		byte(ScheduleEntryLockV3),
		0x10,
		c.SetAction,
		c.UserIdentifier,
		c.ScheduleSlotID,
		c.WeekDayBitmask,
		c.StartHour,
		c.StartMinute,
		c.DurationHour,
		c.DurationMinute,
		0x25,
	}
}
