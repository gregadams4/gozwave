package commands
type MeterGet struct {
	node byte
}

func NewMeterGet() MeterGet {
	return MeterGet{}
}

func (c *MeterGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterGet) Set() error {

	return nil
}

func (c *MeterGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Meter),
		0x01,
		0x25,
	}
}
type MeterReport struct {
    *report
    MeterType byte
    Properties1 byte
    MeterValue byte
    data []byte
}

func NewMeterReport(data []byte) *MeterReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterReport{
				MeterType: data[0],
				Properties1: data[1],
				MeterValue: data[2],
        data: data,
    }
}

