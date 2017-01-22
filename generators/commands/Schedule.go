package commands
import "encoding/binary"
type ScheduleSupportedGet struct {
	node byte
}

func NewScheduleSupportedGet() ScheduleSupportedGet {
	return ScheduleSupportedGet{}
}

func (c *ScheduleSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleSupportedGet) Set() error {

	return nil
}

func (c *ScheduleSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Schedule),
		0x01,
		0x25,
	}
}
type ScheduleSupportedReport struct {
    *report
    NumberofSupportedScheduleID byte
    Properties1 byte
    NumberofsupportedCC byte
    Properties3 byte
    data []byte
}

func NewScheduleSupportedReport(data []byte) *ScheduleSupportedReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleSupportedReport{
				NumberofSupportedScheduleID: data[0],
				Properties1: data[1],
				NumberofsupportedCC: data[2],
				Properties3: data[3],
        data: data,
    }
}

type CommandScheduleSet struct {
	node byte
	ScheduleID byte
	UserIdentifier byte
	StartYear byte
	Properties1 byte
	Properties2 byte
	Properties3 byte
	Properties4 byte
	Properties5 byte
	DurationByte uint16
	ReportstoFollow byte
	NumberofCmdtoFollow byte
	Vg1 []CommandScheduleSetvg1
}
type CommandScheduleSetvg1 struct {
	CmdLength byte
	CmdByte byte
}

func NewCommandScheduleSet() CommandScheduleSet {
	return CommandScheduleSet{}
}

func (c *CommandScheduleSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandScheduleSet) Set(ScheduleID byte,UserIdentifier byte,StartYear byte,Properties1 byte,Properties2 byte,Properties3 byte,Properties4 byte,Properties5 byte,DurationByte uint16,ReportstoFollow byte,NumberofCmdtoFollow byte,Vg1 []CommandScheduleSetvg1,) error {
	c.ScheduleID = ScheduleID
	c.UserIdentifier = UserIdentifier
	c.StartYear = StartYear
	c.Properties1 = Properties1
	c.Properties2 = Properties2
	c.Properties3 = Properties3
	c.Properties4 = Properties4
	c.Properties5 = Properties5
	c.DurationByte = DurationByte
	c.ReportstoFollow = ReportstoFollow
	c.NumberofCmdtoFollow = NumberofCmdtoFollow
	c.Vg1 = Vg1

	return nil
}

func (c *CommandScheduleSet) Encode() []byte {
	DurationByteBytes := []byte{}
	binary.BigEndian.PutUint16(DurationByteBytes, c.DurationByte)
	var data = []byte{
		0x13,
		c.node,
		byte(15),
		byte(Schedule),
		0x03,
		c.ScheduleID,
		c.UserIdentifier,
		c.StartYear,
		c.Properties1,
		c.Properties2,
		c.Properties3,
		c.Properties4,
		c.Properties5,
		DurationByteBytes[0],
		DurationByteBytes[1],
		c.ReportstoFollow,
		c.NumberofCmdtoFollow,
	}
	for _, v := range c.Vg1 {
		data = append(data, v.CmdLength)
		data = append(data, v.CmdByte)
	}
	data = append(data, 0x25)
	return data
	
}
type CommandScheduleGet struct {
	node byte
	ScheduleID byte
}

func NewCommandScheduleGet() CommandScheduleGet {
	return CommandScheduleGet{}
}

func (c *CommandScheduleGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandScheduleGet) Set(ScheduleID byte,) error {
	c.ScheduleID = ScheduleID

	return nil
}

func (c *CommandScheduleGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Schedule),
		0x04,
		c.ScheduleID,
		0x25,
	}
}
type CommandScheduleReport struct {
    *report
    ScheduleID byte
    UserIdentifier byte
    StartYear byte
    Properties1 byte
    Properties2 byte
    Properties3 byte
    Properties4 byte
    Properties5 byte
    DurationByte uint16
    ReportstoFollow byte
    NumberofCmdtoFollow byte
    data []byte
}

func NewCommandScheduleReport(data []byte) *CommandScheduleReport {
    if len(data) < 12 {
				//may want to change this to just return nil
				for i := len(data); i < 12; i++ {
            data = append(data, 0x00)
        }
    }

    return &CommandScheduleReport{
				ScheduleID: data[0],
				UserIdentifier: data[1],
				StartYear: data[2],
				Properties1: data[3],
				Properties2: data[4],
				Properties3: data[5],
				Properties4: data[6],
				Properties5: data[7],
				DurationByte: binary.BigEndian.Uint16(data[8:10]),
				ReportstoFollow: data[10],
				NumberofCmdtoFollow: data[11],
        data: data,
    }
}

type ScheduleRemove struct {
	node byte
	ScheduleID byte
}

func NewScheduleRemove() ScheduleRemove {
	return ScheduleRemove{}
}

func (c *ScheduleRemove) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleRemove) Set(ScheduleID byte,) error {
	c.ScheduleID = ScheduleID

	return nil
}

func (c *ScheduleRemove) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Schedule),
		0x06,
		c.ScheduleID,
		0x25,
	}
}
type ScheduleStateSet struct {
	node byte
	ScheduleID byte
	ScheduleState byte
}

func NewScheduleStateSet() ScheduleStateSet {
	return ScheduleStateSet{}
}

func (c *ScheduleStateSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleStateSet) Set(ScheduleID byte,ScheduleState byte,) error {
	c.ScheduleID = ScheduleID
	c.ScheduleState = ScheduleState

	return nil
}

func (c *ScheduleStateSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(Schedule),
		0x07,
		c.ScheduleID,
		c.ScheduleState,
		0x25,
	}
}
type ScheduleStateGet struct {
	node byte
}

func NewScheduleStateGet() ScheduleStateGet {
	return ScheduleStateGet{}
}

func (c *ScheduleStateGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleStateGet) Set() error {

	return nil
}

func (c *ScheduleStateGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Schedule),
		0x08,
		0x25,
	}
}
type ScheduleStateReport struct {
    *report
    NumberofSupportedScheduleID byte
    Properties1 byte
    Properties2 byte
    Properties3 byte
    data []byte
}

func NewScheduleStateReport(data []byte) *ScheduleStateReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleStateReport{
				NumberofSupportedScheduleID: data[0],
				Properties1: data[1],
				Properties2: data[2],
				Properties3: data[3],
        data: data,
    }
}

