package commands
type ScheduleChangedGet struct {
	node byte
}

func NewScheduleChangedGet() ScheduleChangedGet {
	return ScheduleChangedGet{}
}

func (c *ScheduleChangedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleChangedGet) Set() error {

	return nil
}

func (c *ScheduleChangedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ClimateControlSchedule),
		0x04,
		0x25,
	}
}
type ScheduleChangedReport struct {
    *report
    ChangeCounter byte
    data []byte
}

func NewScheduleChangedReport(data []byte) *ScheduleChangedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleChangedReport{
				ChangeCounter: data[0],
        data: data,
    }
}

type ScheduleGet struct {
	node byte
	Properties1 byte
}

func NewScheduleGet() ScheduleGet {
	return ScheduleGet{}
}

func (c *ScheduleGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleGet) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *ScheduleGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ClimateControlSchedule),
		0x02,
		c.Properties1,
		0x25,
	}
}
type ScheduleOverrideGet struct {
	node byte
}

func NewScheduleOverrideGet() ScheduleOverrideGet {
	return ScheduleOverrideGet{}
}

func (c *ScheduleOverrideGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleOverrideGet) Set() error {

	return nil
}

func (c *ScheduleOverrideGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ClimateControlSchedule),
		0x07,
		0x25,
	}
}
type ScheduleOverrideReport struct {
    *report
    Properties1 byte
    OverrideState byte
    data []byte
}

func NewScheduleOverrideReport(data []byte) *ScheduleOverrideReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleOverrideReport{
				Properties1: data[0],
				OverrideState: data[1],
        data: data,
    }
}

type ScheduleOverrideSet struct {
	node byte
	Properties1 byte
	OverrideState byte
}

func NewScheduleOverrideSet() ScheduleOverrideSet {
	return ScheduleOverrideSet{}
}

func (c *ScheduleOverrideSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleOverrideSet) Set(Properties1 byte,OverrideState byte,) error {
	c.Properties1 = Properties1
	c.OverrideState = OverrideState

	return nil
}

func (c *ScheduleOverrideSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ClimateControlSchedule),
		0x06,
		c.Properties1,
		c.OverrideState,
		0x25,
	}
}
type ScheduleReport struct {
    *report
    Properties1 byte
    Switchpoint0 byte
    Switchpoint1 byte
    Switchpoint2 byte
    Switchpoint3 byte
    Switchpoint4 byte
    Switchpoint5 byte
    Switchpoint6 byte
    Switchpoint7 byte
    Switchpoint8 byte
    data []byte
}

func NewScheduleReport(data []byte) *ScheduleReport {
    if len(data) < 10 {
				//may want to change this to just return nil
				for i := len(data); i < 10; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScheduleReport{
				Properties1: data[0],
				Switchpoint0: data[1],
				Switchpoint1: data[2],
				Switchpoint2: data[3],
				Switchpoint3: data[4],
				Switchpoint4: data[5],
				Switchpoint5: data[6],
				Switchpoint6: data[7],
				Switchpoint7: data[8],
				Switchpoint8: data[9],
        data: data,
    }
}

type ScheduleSet struct {
	node byte
	Properties1 byte
	Switchpoint0 byte
	Switchpoint1 byte
	Switchpoint2 byte
	Switchpoint3 byte
	Switchpoint4 byte
	Switchpoint5 byte
	Switchpoint6 byte
	Switchpoint7 byte
	Switchpoint8 byte
}

func NewScheduleSet() ScheduleSet {
	return ScheduleSet{}
}

func (c *ScheduleSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScheduleSet) Set(Properties1 byte,Switchpoint0 byte,Switchpoint1 byte,Switchpoint2 byte,Switchpoint3 byte,Switchpoint4 byte,Switchpoint5 byte,Switchpoint6 byte,Switchpoint7 byte,Switchpoint8 byte,) error {
	c.Properties1 = Properties1
	c.Switchpoint0 = Switchpoint0
	c.Switchpoint1 = Switchpoint1
	c.Switchpoint2 = Switchpoint2
	c.Switchpoint3 = Switchpoint3
	c.Switchpoint4 = Switchpoint4
	c.Switchpoint5 = Switchpoint5
	c.Switchpoint6 = Switchpoint6
	c.Switchpoint7 = Switchpoint7
	c.Switchpoint8 = Switchpoint8

	return nil
}

func (c *ScheduleSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(12),
		byte(ClimateControlSchedule),
		0x01,
		c.Properties1,
		c.Switchpoint0,
		c.Switchpoint1,
		c.Switchpoint2,
		c.Switchpoint3,
		c.Switchpoint4,
		c.Switchpoint5,
		c.Switchpoint6,
		c.Switchpoint7,
		c.Switchpoint8,
		0x25,
	}
}
