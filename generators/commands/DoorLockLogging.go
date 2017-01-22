package commands
import "encoding/binary"
type DoorLockLoggingRecordsSupportedGet struct {
	node byte
}

func NewDoorLockLoggingRecordsSupportedGet() DoorLockLoggingRecordsSupportedGet {
	return DoorLockLoggingRecordsSupportedGet{}
}

func (c *DoorLockLoggingRecordsSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockLoggingRecordsSupportedGet) Set() error {

	return nil
}

func (c *DoorLockLoggingRecordsSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(DoorLockLogging),
		0x01,
		0x25,
	}
}
type DoorLockLoggingRecordsSupportedReport struct {
    *report
    Maxrecordsstored byte
    data []byte
}

func NewDoorLockLoggingRecordsSupportedReport(data []byte) *DoorLockLoggingRecordsSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &DoorLockLoggingRecordsSupportedReport{
				Maxrecordsstored: data[0],
        data: data,
    }
}

type RecordGet struct {
	node byte
	Recordnumber byte
}

func NewRecordGet() RecordGet {
	return RecordGet{}
}

func (c *RecordGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *RecordGet) Set(Recordnumber byte,) error {
	c.Recordnumber = Recordnumber

	return nil
}

func (c *RecordGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(DoorLockLogging),
		0x03,
		c.Recordnumber,
		0x25,
	}
}
type RecordReport struct {
    *report
    Recordnumber byte
    Year uint16
    Month byte
    Day byte
    Properties1 byte
    MinuteLocalTime byte
    SecondLocalTime byte
    EventType byte
    UserIdentifier byte
    UserCodeLength byte
    USER_CODE byte
    data []byte
}

func NewRecordReport(data []byte) *RecordReport {
    if len(data) < 12 {
				//may want to change this to just return nil
				for i := len(data); i < 12; i++ {
            data = append(data, 0x00)
        }
    }

    return &RecordReport{
				Recordnumber: data[0],
				Year: binary.BigEndian.Uint16(data[1:3]),
				Month: data[3],
				Day: data[4],
				Properties1: data[5],
				MinuteLocalTime: data[6],
				SecondLocalTime: data[7],
				EventType: data[8],
				UserIdentifier: data[9],
				UserCodeLength: data[10],
				USER_CODE: data[11],
        data: data,
    }
}

