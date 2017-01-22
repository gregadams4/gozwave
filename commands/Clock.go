package commands
type ClockGet struct {
	node byte
}

func NewClockGet() ClockGet {
	return ClockGet{}
}

func (c *ClockGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ClockGet) Set() error {

	return nil
}

func (c *ClockGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Clock),
		0x05,
		0x25,
	}
}
type ClockReport struct {
    *report
    Level byte
    Minute byte
    data []byte
}

func NewClockReport(data []byte) *ClockReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &ClockReport{
				Level: data[0],
				Minute: data[1],
        data: data,
    }
}

type ClockSet struct {
	node byte
	Level byte
	Minute byte
}

func NewClockSet() ClockSet {
	return ClockSet{}
}

func (c *ClockSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ClockSet) Set(Level byte,Minute byte,) error {
	c.Level = Level
	c.Minute = Minute

	return nil
}

func (c *ClockSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(Clock),
		0x04,
		c.Level,
		c.Minute,
		0x25,
	}
}
