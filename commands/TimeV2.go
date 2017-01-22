package commands
import "encoding/binary"
type DateGetV2 struct {
	node byte
}

func NewDateGetV2() DateGetV2 {
	return DateGetV2{}
}

func (c *DateGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *DateGetV2) Set() error {

	return nil
}

func (c *DateGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(TimeV2),
		0x03,
		0x25,
	}
}
type DateReportV2 struct {
    *report
    Year uint16
    Month byte
    Day byte
    data []byte
}

func NewDateReportV2(data []byte) *DateReportV2 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &DateReportV2{
				Year: binary.BigEndian.Uint16(data[0:2]),
				Month: data[2],
				Day: data[3],
        data: data,
    }
}

type TimeGetV2 struct {
	node byte
}

func NewTimeGetV2() TimeGetV2 {
	return TimeGetV2{}
}

func (c *TimeGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *TimeGetV2) Set() error {

	return nil
}

func (c *TimeGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(TimeV2),
		0x01,
		0x25,
	}
}
type TimeOffsetGetV2 struct {
	node byte
}

func NewTimeOffsetGetV2() TimeOffsetGetV2 {
	return TimeOffsetGetV2{}
}

func (c *TimeOffsetGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *TimeOffsetGetV2) Set() error {

	return nil
}

func (c *TimeOffsetGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(TimeV2),
		0x06,
		0x25,
	}
}
type TimeOffsetReportV2 struct {
    *report
    Level byte
    MinuteTZO byte
    Level2 byte
    MonthStartDST byte
    DayStartDST byte
    HourStartDST byte
    MonthEndDST byte
    DayEndDST byte
    HourEndDST byte
    data []byte
}

func NewTimeOffsetReportV2(data []byte) *TimeOffsetReportV2 {
    if len(data) < 9 {
				//may want to change this to just return nil
				for i := len(data); i < 9; i++ {
            data = append(data, 0x00)
        }
    }

    return &TimeOffsetReportV2{
				Level: data[0],
				MinuteTZO: data[1],
				Level2: data[2],
				MonthStartDST: data[3],
				DayStartDST: data[4],
				HourStartDST: data[5],
				MonthEndDST: data[6],
				DayEndDST: data[7],
				HourEndDST: data[8],
        data: data,
    }
}

type TimeOffsetSetV2 struct {
	node byte
	Level byte
	MinuteTZO byte
	Level2 byte
	MonthStartDST byte
	DayStartDST byte
	HourStartDST byte
	MonthEndDST byte
	DayEndDST byte
	HourEndDST byte
}

func NewTimeOffsetSetV2() TimeOffsetSetV2 {
	return TimeOffsetSetV2{}
}

func (c *TimeOffsetSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *TimeOffsetSetV2) Set(Level byte,MinuteTZO byte,Level2 byte,MonthStartDST byte,DayStartDST byte,HourStartDST byte,MonthEndDST byte,DayEndDST byte,HourEndDST byte,) error {
	c.Level = Level
	c.MinuteTZO = MinuteTZO
	c.Level2 = Level2
	c.MonthStartDST = MonthStartDST
	c.DayStartDST = DayStartDST
	c.HourStartDST = HourStartDST
	c.MonthEndDST = MonthEndDST
	c.DayEndDST = DayEndDST
	c.HourEndDST = HourEndDST

	return nil
}

func (c *TimeOffsetSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(11),
		byte(TimeV2),
		0x05,
		c.Level,
		c.MinuteTZO,
		c.Level2,
		c.MonthStartDST,
		c.DayStartDST,
		c.HourStartDST,
		c.MonthEndDST,
		c.DayEndDST,
		c.HourEndDST,
		0x25,
	}
}
type TimeReportV2 struct {
    *report
    HourLocalTime byte
    MinuteLocalTime byte
    SecondLocalTime byte
    data []byte
}

func NewTimeReportV2(data []byte) *TimeReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &TimeReportV2{
				HourLocalTime: data[0],
				MinuteLocalTime: data[1],
				SecondLocalTime: data[2],
        data: data,
    }
}

