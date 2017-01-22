package commands
import "encoding/binary"
type ScheduleSupportedGetV3 struct {
	node byte
	ScheduleIDBlock byte
}

func NewScheduleSupportedGetV3() ScheduleSupportedGetV3 {
	return ScheduleSupportedGetV3{}
}

func (c *ScheduleSupportedGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleSupportedGetV3) Set(ScheduleIDBlock byte,) error {
	c.ScheduleIDBlock = ScheduleIDBlock

	return nil
}

func (c *ScheduleSupportedGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ScheduleV3),
		0x01,
		c.ScheduleIDBlock,
		0x25,
	}
}
type ScheduleSupportedReportV3 struct {
    *report
    NumberofSupportedScheduleID byte
    Properties1 byte
    NumberofsupportedCC byte
    Properties3 byte
    ScheduleIDBlock byte
    NumberofSupportedScheduleBlocks byte
    data []byte
}

func NewScheduleSupportedReportV3(data []byte) *ScheduleSupportedReportV3 {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleSupportedReportV3{
				NumberofSupportedScheduleID: data[0],
				Properties1: data[1],
				NumberofsupportedCC: data[2],
				Properties3: data[3],
				ScheduleIDBlock: data[4],
				NumberofSupportedScheduleBlocks: data[5],
        data: data,
    }
}

type CommandScheduleSetV3 struct {
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
	Vg1 []CommandScheduleSetV3vg1
}
type CommandScheduleSetV3vg1 struct {
	CmdLength byte
	CmdByte byte
}

func NewCommandScheduleSetV3() CommandScheduleSetV3 {
	return CommandScheduleSetV3{}
}

func (c *CommandScheduleSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandScheduleSetV3) Set(ScheduleID byte,ScheduleIDBlock byte,StartYear byte,Properties1 byte,Properties2 byte,Properties3 byte,Properties4 byte,Properties5 byte,DurationByte uint16,ReportstoFollow byte,NumberofCmdtoFollow byte,Vg1 []CommandScheduleSetV3vg1,) error {
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

func (c *CommandScheduleSetV3) Encode() []byte {
	DurationByteBytes := []byte{}
	binary.BigEndian.PutUint16(DurationByteBytes, c.DurationByte)
	var data = []byte{
		0x13,
		c.node,
		byte(15),
		byte(ScheduleV3),
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
type CommandScheduleGetV3 struct {
	node byte
	ScheduleID byte
	ScheduleIDBlock byte
	Properties1 byte
}

func NewCommandScheduleGetV3() CommandScheduleGetV3 {
	return CommandScheduleGetV3{}
}

func (c *CommandScheduleGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandScheduleGetV3) Set(ScheduleID byte,ScheduleIDBlock byte,Properties1 byte,) error {
	c.ScheduleID = ScheduleID
	c.ScheduleIDBlock = ScheduleIDBlock
	c.Properties1 = Properties1

	return nil
}

func (c *CommandScheduleGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ScheduleV3),
		0x04,
		c.ScheduleID,
		c.ScheduleIDBlock,
		c.Properties1,
		0x25,
	}
}
type CommandScheduleReportV3 struct {
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

func NewCommandScheduleReportV3(data []byte) *CommandScheduleReportV3 {
    if len(data) < 12 {
				//may want to change this to just return nil
				for i := len(data); i < 12; i++ {
            data = append(data, 0x00)
        }
    }

    return &CommandScheduleReportV3{
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

type ScheduleRemoveV3 struct {
	node byte
	ScheduleID byte
	ScheduleIDBlock byte
}

func NewScheduleRemoveV3() ScheduleRemoveV3 {
	return ScheduleRemoveV3{}
}

func (c *ScheduleRemoveV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleRemoveV3) Set(ScheduleID byte,ScheduleIDBlock byte,) error {
	c.ScheduleID = ScheduleID
	c.ScheduleIDBlock = ScheduleIDBlock

	return nil
}

func (c *ScheduleRemoveV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ScheduleV3),
		0x06,
		c.ScheduleID,
		c.ScheduleIDBlock,
		0x25,
	}
}
type ScheduleStateSetV3 struct {
	node byte
	ScheduleID byte
	ScheduleState byte
	ScheduleIDBlock byte
}

func NewScheduleStateSetV3() ScheduleStateSetV3 {
	return ScheduleStateSetV3{}
}

func (c *ScheduleStateSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleStateSetV3) Set(ScheduleID byte,ScheduleState byte,ScheduleIDBlock byte,) error {
	c.ScheduleID = ScheduleID
	c.ScheduleState = ScheduleState
	c.ScheduleIDBlock = ScheduleIDBlock

	return nil
}

func (c *ScheduleStateSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ScheduleV3),
		0x07,
		c.ScheduleID,
		c.ScheduleState,
		c.ScheduleIDBlock,
		0x25,
	}
}
type ScheduleStateGetV3 struct {
	node byte
	ScheduleIDBlock byte
}

func NewScheduleStateGetV3() ScheduleStateGetV3 {
	return ScheduleStateGetV3{}
}

func (c *ScheduleStateGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleStateGetV3) Set(ScheduleIDBlock byte,) error {
	c.ScheduleIDBlock = ScheduleIDBlock

	return nil
}

func (c *ScheduleStateGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ScheduleV3),
		0x08,
		c.ScheduleIDBlock,
		0x25,
	}
}
type ScheduleStateReportV3 struct {
    *report
    NumberofSupportedScheduleID byte
    Properties1 byte
    Properties2 byte
    Properties3 byte
    ScheduleIDBlock byte
    data []byte
}

func NewScheduleStateReportV3(data []byte) *ScheduleStateReportV3 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleStateReportV3{
				NumberofSupportedScheduleID: data[0],
				Properties1: data[1],
				Properties2: data[2],
				Properties3: data[3],
				ScheduleIDBlock: data[4],
        data: data,
    }
}

