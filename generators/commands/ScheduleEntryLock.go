package commands
type ScheduleEntryLockEnableAllSet struct {
	node byte
	Enabled byte
}

func NewScheduleEntryLockEnableAllSet() ScheduleEntryLockEnableAllSet {
	return ScheduleEntryLockEnableAllSet{}
}

func (c *ScheduleEntryLockEnableAllSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockEnableAllSet) Set(Enabled byte,) error {
	c.Enabled = Enabled

	return nil
}

func (c *ScheduleEntryLockEnableAllSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ScheduleEntryLock),
		0x02,
		c.Enabled,
		0x25,
	}
}
type ScheduleEntryLockEnableSet struct {
	node byte
	UserIdentifier byte
	Enabled byte
}

func NewScheduleEntryLockEnableSet() ScheduleEntryLockEnableSet {
	return ScheduleEntryLockEnableSet{}
}

func (c *ScheduleEntryLockEnableSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockEnableSet) Set(UserIdentifier byte,Enabled byte,) error {
	c.UserIdentifier = UserIdentifier
	c.Enabled = Enabled

	return nil
}

func (c *ScheduleEntryLockEnableSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleEntryLock),
		0x01,
		c.UserIdentifier,
		c.Enabled,
		0x25,
	}
}
type ScheduleEntryLockWeekDayGet struct {
	node byte
	UserIdentifier byte
	ScheduleSlotID byte
}

func NewScheduleEntryLockWeekDayGet() ScheduleEntryLockWeekDayGet {
	return ScheduleEntryLockWeekDayGet{}
}

func (c *ScheduleEntryLockWeekDayGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockWeekDayGet) Set(UserIdentifier byte,ScheduleSlotID byte,) error {
	c.UserIdentifier = UserIdentifier
	c.ScheduleSlotID = ScheduleSlotID

	return nil
}

func (c *ScheduleEntryLockWeekDayGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleEntryLock),
		0x04,
		c.UserIdentifier,
		c.ScheduleSlotID,
		0x25,
	}
}
type ScheduleEntryLockWeekDayReport struct {
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

func NewScheduleEntryLockWeekDayReport(data []byte) *ScheduleEntryLockWeekDayReport {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleEntryLockWeekDayReport{
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

type ScheduleEntryLockWeekDaySet struct {
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

func NewScheduleEntryLockWeekDaySet() ScheduleEntryLockWeekDaySet {
	return ScheduleEntryLockWeekDaySet{}
}

func (c *ScheduleEntryLockWeekDaySet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockWeekDaySet) Set(SetAction byte,UserIdentifier byte,ScheduleSlotID byte,DayofWeek byte,StartHour byte,StartMinute byte,StopHour byte,StopMinute byte,) error {
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

func (c *ScheduleEntryLockWeekDaySet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(10),
		byte(ScheduleEntryLock),
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
type ScheduleEntryLockYearDayGet struct {
	node byte
	UserIdentifier byte
	ScheduleSlotID byte
}

func NewScheduleEntryLockYearDayGet() ScheduleEntryLockYearDayGet {
	return ScheduleEntryLockYearDayGet{}
}

func (c *ScheduleEntryLockYearDayGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockYearDayGet) Set(UserIdentifier byte,ScheduleSlotID byte,) error {
	c.UserIdentifier = UserIdentifier
	c.ScheduleSlotID = ScheduleSlotID

	return nil
}

func (c *ScheduleEntryLockYearDayGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleEntryLock),
		0x07,
		c.UserIdentifier,
		c.ScheduleSlotID,
		0x25,
	}
}
type ScheduleEntryLockYearDayReport struct {
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

func NewScheduleEntryLockYearDayReport(data []byte) *ScheduleEntryLockYearDayReport {
    if len(data) < 12 {
				//may want to change this to just return nil
				for i := len(data); i < 12; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleEntryLockYearDayReport{
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

type ScheduleEntryLockYearDaySet struct {
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

func NewScheduleEntryLockYearDaySet() ScheduleEntryLockYearDaySet {
	return ScheduleEntryLockYearDaySet{}
}

func (c *ScheduleEntryLockYearDaySet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryLockYearDaySet) Set(SetAction byte,UserIdentifier byte,ScheduleSlotID byte,StartYear byte,StartMonth byte,StartDay byte,StartHour byte,StartMinute byte,StopYear byte,StopMonth byte,StopDay byte,StopHour byte,StopMinute byte,) error {
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

func (c *ScheduleEntryLockYearDaySet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(15),
		byte(ScheduleEntryLock),
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
type ScheduleEntryTypeSupportedGet struct {
	node byte
}

func NewScheduleEntryTypeSupportedGet() ScheduleEntryTypeSupportedGet {
	return ScheduleEntryTypeSupportedGet{}
}

func (c *ScheduleEntryTypeSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleEntryTypeSupportedGet) Set() error {

	return nil
}

func (c *ScheduleEntryTypeSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ScheduleEntryLock),
		0x09,
		0x25,
	}
}
type ScheduleEntryTypeSupportedReport struct {
    *report
    NumberofSlotsWeekDay byte
    NumberofSlotsYearDay byte
    data []byte
}

func NewScheduleEntryTypeSupportedReport(data []byte) *ScheduleEntryTypeSupportedReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleEntryTypeSupportedReport{
				NumberofSlotsWeekDay: data[0],
				NumberofSlotsYearDay: data[1],
        data: data,
    }
}

