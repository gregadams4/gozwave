package commands
import "encoding/binary"
type DcpListRemove struct {
	node byte
	Year uint16
	Month byte
	Day byte
	HourLocalTime byte
	MinuteLocalTime byte
	SecondLocalTime byte
}

func NewDcpListRemove() DcpListRemove {
	return DcpListRemove{}
}

func (c *DcpListRemove) SetNode(node int) {
	c.node = byte(node)
}

func (c *DcpListRemove) Set(Year uint16,Month byte,Day byte,HourLocalTime byte,MinuteLocalTime byte,SecondLocalTime byte,) error {
	c.Year = Year
	c.Month = Month
	c.Day = Day
	c.HourLocalTime = HourLocalTime
	c.MinuteLocalTime = MinuteLocalTime
	c.SecondLocalTime = SecondLocalTime

	return nil
}

func (c *DcpListRemove) Encode() []byte {
	YearBytes := []byte{}
	binary.BigEndian.PutUint16(YearBytes, c.Year)
	return []byte{
		0x13,
		c.node,
		byte(8),
		byte(DcpConfig),
		0x04,
		YearBytes[0],
		YearBytes[1],
		c.Month,
		c.Day,
		c.HourLocalTime,
		c.MinuteLocalTime,
		c.SecondLocalTime,
		0x25,
	}
}
type DcpListSet struct {
	node byte
	Year uint16
	Month byte
	Day byte
	HourLocalTime byte
	MinuteLocalTime byte
	SecondLocalTime byte
	DCPRateID byte
	Properties1 byte
	StartYear uint16
	StartMonth byte
	StartDay byte
	StartHourLocalTime byte
	StartMinuteLocalTime byte
	StartSecondLocalTime byte
	DurationHourTime byte
	DurationMinuteTime byte
	DurationSecondTime byte
	EventPriority byte
	Loadshedding byte
	StartAssociationGroup byte
	StopAssociationGroup byte
	Randomizationinterval byte
	Vg1 []DcpListSetvg1
}
type DcpListSetvg1 struct {
	GenericDeviceClass byte
	SpecificDeviceClass byte
}

func NewDcpListSet() DcpListSet {
	return DcpListSet{}
}

func (c *DcpListSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DcpListSet) Set(Year uint16,Month byte,Day byte,HourLocalTime byte,MinuteLocalTime byte,SecondLocalTime byte,DCPRateID byte,Properties1 byte,StartYear uint16,StartMonth byte,StartDay byte,StartHourLocalTime byte,StartMinuteLocalTime byte,StartSecondLocalTime byte,DurationHourTime byte,DurationMinuteTime byte,DurationSecondTime byte,EventPriority byte,Loadshedding byte,StartAssociationGroup byte,StopAssociationGroup byte,Randomizationinterval byte,Vg1 []DcpListSetvg1,) error {
	c.Year = Year
	c.Month = Month
	c.Day = Day
	c.HourLocalTime = HourLocalTime
	c.MinuteLocalTime = MinuteLocalTime
	c.SecondLocalTime = SecondLocalTime
	c.DCPRateID = DCPRateID
	c.Properties1 = Properties1
	c.StartYear = StartYear
	c.StartMonth = StartMonth
	c.StartDay = StartDay
	c.StartHourLocalTime = StartHourLocalTime
	c.StartMinuteLocalTime = StartMinuteLocalTime
	c.StartSecondLocalTime = StartSecondLocalTime
	c.DurationHourTime = DurationHourTime
	c.DurationMinuteTime = DurationMinuteTime
	c.DurationSecondTime = DurationSecondTime
	c.EventPriority = EventPriority
	c.Loadshedding = Loadshedding
	c.StartAssociationGroup = StartAssociationGroup
	c.StopAssociationGroup = StopAssociationGroup
	c.Randomizationinterval = Randomizationinterval
	c.Vg1 = Vg1

	return nil
}

func (c *DcpListSet) Encode() []byte {
	YearBytes := []byte{}
	binary.BigEndian.PutUint16(YearBytes, c.Year)
	StartYearBytes := []byte{}
	binary.BigEndian.PutUint16(StartYearBytes, c.StartYear)
	var data = []byte{
		0x13,
		c.node,
		byte(26),
		byte(DcpConfig),
		0x03,
		YearBytes[0],
		YearBytes[1],
		c.Month,
		c.Day,
		c.HourLocalTime,
		c.MinuteLocalTime,
		c.SecondLocalTime,
		c.DCPRateID,
		c.Properties1,
		StartYearBytes[0],
		StartYearBytes[1],
		c.StartMonth,
		c.StartDay,
		c.StartHourLocalTime,
		c.StartMinuteLocalTime,
		c.StartSecondLocalTime,
		c.DurationHourTime,
		c.DurationMinuteTime,
		c.DurationSecondTime,
		c.EventPriority,
		c.Loadshedding,
		c.StartAssociationGroup,
		c.StopAssociationGroup,
		c.Randomizationinterval,
	}
	for _, v := range c.Vg1 {
		data = append(data, v.GenericDeviceClass)
		data = append(data, v.SpecificDeviceClass)
	}
	data = append(data, 0x25)
	return data
	
}
type DcpListSupportedGet struct {
	node byte
}

func NewDcpListSupportedGet() DcpListSupportedGet {
	return DcpListSupportedGet{}
}

func (c *DcpListSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DcpListSupportedGet) Set() error {

	return nil
}

func (c *DcpListSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(DcpConfig),
		0x01,
		0x25,
	}
}
type DcpListSupportedReport struct {
    *report
    DCPListSize byte
    FreeDCPListentries byte
    data []byte
}

func NewDcpListSupportedReport(data []byte) *DcpListSupportedReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &DcpListSupportedReport{
				DCPListSize: data[0],
				FreeDCPListentries: data[1],
        data: data,
    }
}

