package commands
import "encoding/binary"
type DcpEventStatusGet struct {
	node byte
	Year uint16
	Month byte
	Day byte
	HourLocalTime byte
	MinuteLocalTime byte
	SecondLocalTime byte
}

func NewDcpEventStatusGet() DcpEventStatusGet {
	return DcpEventStatusGet{}
}

func (c *DcpEventStatusGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DcpEventStatusGet) Set(Year uint16,Month byte,Day byte,HourLocalTime byte,MinuteLocalTime byte,SecondLocalTime byte,) error {
	c.Year = Year
	c.Month = Month
	c.Day = Day
	c.HourLocalTime = HourLocalTime
	c.MinuteLocalTime = MinuteLocalTime
	c.SecondLocalTime = SecondLocalTime

	return nil
}

func (c *DcpEventStatusGet) Encode() []byte {
	YearBytes := []byte{}
	binary.BigEndian.PutUint16(YearBytes, c.Year)
	return []byte{
		0x13,
		c.node,
		byte(8),
		byte(DcpMonitor),
		0x03,
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
type DcpEventStatusReport struct {
    *report
    Year uint16
    Month byte
    Day byte
    HourLocalTime byte
    MinuteLocalTime byte
    SecondLocalTime byte
    Eventstatus byte
    data []byte
}

func NewDcpEventStatusReport(data []byte) *DcpEventStatusReport {
    if len(data) < 8 {
				//may want to change this to just return nil
				for i := len(data); i < 8; i++ {
            data = append(data, 0x00)
        }
    }

    return &DcpEventStatusReport{
				Year: binary.BigEndian.Uint16(data[0:2]),
				Month: data[2],
				Day: data[3],
				HourLocalTime: data[4],
				MinuteLocalTime: data[5],
				SecondLocalTime: data[6],
				Eventstatus: data[7],
        data: data,
    }
}

type DcpListGet struct {
	node byte
}

func NewDcpListGet() DcpListGet {
	return DcpListGet{}
}

func (c *DcpListGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DcpListGet) Set() error {

	return nil
}

func (c *DcpListGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(DcpMonitor),
		0x01,
		0x25,
	}
}
type DcpListReport struct {
    *report
    ReportstoFollow byte
    Year uint16
    Month byte
    Day byte
    HourLocalTime byte
    MinuteLocalTime byte
    SecondLocalTime byte
    DCPID byte
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
    data []byte
}

func NewDcpListReport(data []byte) *DcpListReport {
    if len(data) < 25 {
				//may want to change this to just return nil
				for i := len(data); i < 25; i++ {
            data = append(data, 0x00)
        }
    }

    return &DcpListReport{
				ReportstoFollow: data[0],
				Year: binary.BigEndian.Uint16(data[1:3]),
				Month: data[3],
				Day: data[4],
				HourLocalTime: data[5],
				MinuteLocalTime: data[6],
				SecondLocalTime: data[7],
				DCPID: data[8],
				Properties1: data[9],
				StartYear: binary.BigEndian.Uint16(data[10:12]),
				StartMonth: data[12],
				StartDay: data[13],
				StartHourLocalTime: data[14],
				StartMinuteLocalTime: data[15],
				StartSecondLocalTime: data[16],
				DurationHourTime: data[17],
				DurationMinuteTime: data[18],
				DurationSecondTime: data[19],
				EventPriority: data[20],
				Loadshedding: data[21],
				StartAssociationGroup: data[22],
				StopAssociationGroup: data[23],
				Randomizationinterval: data[24],
        data: data,
    }
}

