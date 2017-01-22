package commands

type CmdMeterGet struct {
	node byte
}

func NewMeterGet() *CmdMeterGet {
	return &CmdMeterGet{}
}

func (c *CmdMeterGet) SetNode(v int) {
	c.node = byte(v)
}

func (c *CmdMeterGet) Encode() []byte {
	return []byte{
		0x13,   // SEND ZW
		c.node, // NOD ID
		0x02,   // length
		Meter,  // command class id
		0x03,   // supported get
		// byte(c.parameter), // value
		0x25, // transmit options
		//0x08,         // transmit options
	}
}

type MeterReport struct {
	*report
	node   byte
	Status byte
	data   []byte
}

func NewMeterReport(data []byte) *MeterReport {
	if len(data) < 1 {
		data = []byte{0x00}
	}

	return &MeterReport{
		Status: data[0],
		data:   data,
	}
}
