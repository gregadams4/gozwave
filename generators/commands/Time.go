package commands
import "encoding/binary"
type DateGet struct {
	node byte
}

func NewDateGet() DateGet {
	return DateGet{}
}

func (c *DateGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DateGet) Set() error {

	return nil
}

func (c *DateGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Time),
		0x03,
		0x25,
	}
}
type DateReport struct {
    *report
    Year uint16
    Month byte
    Day byte
    data []byte
}

func NewDateReport(data []byte) *DateReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &DateReport{
				Year: binary.BigEndian.Uint16(data[0:2]),
				Month: data[2],
				Day: data[3],
        data: data,
    }
}

type TimeGet struct {
	node byte
}

func NewTimeGet() TimeGet {
	return TimeGet{}
}

func (c *TimeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *TimeGet) Set() error {

	return nil
}

func (c *TimeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Time),
		0x01,
		0x25,
	}
}
type TimeReport struct {
    *report
    HourLocalTime byte
    MinuteLocalTime byte
    SecondLocalTime byte
    data []byte
}

func NewTimeReport(data []byte) *TimeReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &TimeReport{
				HourLocalTime: data[0],
				MinuteLocalTime: data[1],
				SecondLocalTime: data[2],
        data: data,
    }
}

