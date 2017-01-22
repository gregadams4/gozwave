package commands
import "encoding/binary"
type TimeParametersGet struct {
	node byte
}

func NewTimeParametersGet() TimeParametersGet {
	return TimeParametersGet{}
}

func (c *TimeParametersGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *TimeParametersGet) Set() error {

	return nil
}

func (c *TimeParametersGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(TimeParameters),
		0x02,
		0x25,
	}
}
type TimeParametersReport struct {
    *report
    Year uint16
    Month byte
    Day byte
    HourUTC byte
    MinuteUTC byte
    SecondUTC byte
    data []byte
}

func NewTimeParametersReport(data []byte) *TimeParametersReport {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &TimeParametersReport{
				Year: binary.BigEndian.Uint16(data[0:2]),
				Month: data[2],
				Day: data[3],
				HourUTC: data[4],
				MinuteUTC: data[5],
				SecondUTC: data[6],
        data: data,
    }
}

type TimeParametersSet struct {
	node byte
	Year uint16
	Month byte
	Day byte
	HourUTC byte
	MinuteUTC byte
	SecondUTC byte
}

func NewTimeParametersSet() TimeParametersSet {
	return TimeParametersSet{}
}

func (c *TimeParametersSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *TimeParametersSet) Set(Year uint16,Month byte,Day byte,HourUTC byte,MinuteUTC byte,SecondUTC byte,) error {
	c.Year = Year
	c.Month = Month
	c.Day = Day
	c.HourUTC = HourUTC
	c.MinuteUTC = MinuteUTC
	c.SecondUTC = SecondUTC

	return nil
}

func (c *TimeParametersSet) Encode() []byte {
	YearBytes := []byte{}
	binary.BigEndian.PutUint16(YearBytes, c.Year)
	return []byte{
		0x13,
		c.node,
		byte(8),
		byte(TimeParameters),
		0x01,
		YearBytes[0],
		YearBytes[1],
		c.Month,
		c.Day,
		c.HourUTC,
		c.MinuteUTC,
		c.SecondUTC,
		0x25,
	}
}
