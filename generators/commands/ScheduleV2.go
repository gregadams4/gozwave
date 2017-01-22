package commands
import "encoding/binary"
type ScheduleSupportedGetV2 struct {
	node byte
	ScheduleIDBlock byte
}

func NewScheduleSupportedGetV2() ScheduleSupportedGetV2 {
	return ScheduleSupportedGetV2{}
}

func (c *ScheduleSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleSupportedGetV2) Set(ScheduleIDBlock byte,) error {
	c.ScheduleIDBlock = ScheduleIDBlock

	return nil
}

func (c *ScheduleSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ScheduleV2),
		0x01,
		c.ScheduleIDBlock,
		0x25,
	}
}
type ScheduleSupportedReportV2 struct {
    *report
    NumberofSupportedScheduleID byte
    Properties1 byte
    NumberofsupportedCC byte
    Properties3 byte
    ScheduleIDBlock byte
    NumberofSupportedScheduleBlocks byte
    data []byte
}

func NewScheduleSupportedReportV2(data []byte) *ScheduleSupportedReportV2 {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleSupportedReportV2{
				NumberofSupportedScheduleID: data[0],
				Properties1: data[1],
				NumberofsupportedCC: data[2],
				Properties3: data[3],
				ScheduleIDBlock: data[4],
				NumberofSupportedScheduleBlocks: data[5],
        data: data,
    }
}

type CommandScheduleSetV2 struct {
	node byte
	ScheduleID byte
	ScheduleIDBlock byte
	StartYear byte
	Properties1 byte
	Properties2 byte
	Properties3 byte
	Properties4 byte
	Properties5 byte
	DurationByte uint16
	ReportstoFollow byte
	NumberofCmdtoFollow byte
	Vg1 []CommandScheduleSetV2vg1
}
type CommandScheduleSetV2vg1 struct {
	CmdLength byte
	CmdByte byte
}

func NewCommandScheduleSetV2() CommandScheduleSetV2 {
	return CommandScheduleSetV2{}
}

func (c *CommandScheduleSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandScheduleSetV2) Set(ScheduleID byte,ScheduleIDBlock byte,StartYear byte,Properties1 byte,Properties2 byte,Properties3 byte,Properties4 byte,Properties5 byte,DurationByte uint16,ReportstoFollow byte,NumberofCmdtoFollow byte,Vg1 []CommandScheduleSetV2vg1,) error {
	c.ScheduleID = ScheduleID
	c.ScheduleIDBlock = ScheduleIDBlock
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

func (c *CommandScheduleSetV2) Encode() []byte {
	DurationByteBytes := []byte{}
	binary.BigEndian.PutUint16(DurationByteBytes, c.DurationByte)
	var data = []byte{
		0x13,
		c.node,
		byte(15),
		byte(ScheduleV2),
		0x03,
		c.ScheduleID,
		c.ScheduleIDBlock,
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
type CommandScheduleGetV2 struct {
	node byte
	ScheduleID byte
	ScheduleIDBlock byte
}

func NewCommandScheduleGetV2() CommandScheduleGetV2 {
	return CommandScheduleGetV2{}
}

func (c *CommandScheduleGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandScheduleGetV2) Set(ScheduleID byte,ScheduleIDBlock byte,) error {
	c.ScheduleID = ScheduleID
	c.ScheduleIDBlock = ScheduleIDBlock

	return nil
}

func (c *CommandScheduleGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleV2),
		0x04,
		c.ScheduleID,
		c.ScheduleIDBlock,
		0x25,
	}
}
type CommandScheduleReportV2 struct {
    *report
    ScheduleID byte
    ScheduleIDBlock byte
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

func NewCommandScheduleReportV2(data []byte) *CommandScheduleReportV2 {
    if len(data) < 12 {
				//may want to change this to just return nil
				for i := len(data); i < 12; i++ {
            data = append(data, 0x00)
        }
    }

    return &CommandScheduleReportV2{
				ScheduleID: data[0],
				ScheduleIDBlock: data[1],
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

type ScheduleRemoveV2 struct {
	node byte
	ScheduleID byte
	ScheduleIDBlock byte
}

func NewScheduleRemoveV2() ScheduleRemoveV2 {
	return ScheduleRemoveV2{}
}

func (c *ScheduleRemoveV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleRemoveV2) Set(ScheduleID byte,ScheduleIDBlock byte,) error {
	c.ScheduleID = ScheduleID
	c.ScheduleIDBlock = ScheduleIDBlock

	return nil
}

func (c *ScheduleRemoveV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleV2),
		0x06,
		c.ScheduleID,
		c.ScheduleIDBlock,
		0x25,
	}
}
type ScheduleStateSetV2 struct {
	node byte
	ScheduleID byte
	ScheduleState byte
	ScheduleIDBlock byte
}

func NewScheduleStateSetV2() ScheduleStateSetV2 {
	return ScheduleStateSetV2{}
}

func (c *ScheduleStateSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleStateSetV2) Set(ScheduleID byte,ScheduleState byte,ScheduleIDBlock byte,) error {
	c.ScheduleID = ScheduleID
	c.ScheduleState = ScheduleState
	c.ScheduleIDBlock = ScheduleIDBlock

	return nil
}

func (c *ScheduleStateSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ScheduleV2),
		0x07,
		c.ScheduleID,
		c.ScheduleState,
		c.ScheduleIDBlock,
		0x25,
	}
}
type ScheduleStateGetV2 struct {
	node byte
	ScheduleIDBlock byte
}

func NewScheduleStateGetV2() ScheduleStateGetV2 {
	return ScheduleStateGetV2{}
}

func (c *ScheduleStateGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleStateGetV2) Set(ScheduleIDBlock byte,) error {
	c.ScheduleIDBlock = ScheduleIDBlock

	return nil
}

func (c *ScheduleStateGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ScheduleV2),
		0x08,
		c.ScheduleIDBlock,
		0x25,
	}
}
type ScheduleStateReportV2 struct {
    *report
    NumberofSupportedScheduleID byte
    Properties1 byte
    Properties2 byte
    Properties3 byte
    ScheduleIDBlock byte
    data []byte
}

func NewScheduleStateReportV2(data []byte) *ScheduleStateReportV2 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleStateReportV2{
				NumberofSupportedScheduleID: data[0],
				Properties1: data[1],
				Properties2: data[2],
				Properties3: data[3],
				ScheduleIDBlock: data[4],
        data: data,
    }
}

