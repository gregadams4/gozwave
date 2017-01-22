package commands
type BatteryGet struct {
	node byte
}

func NewBatteryGet() BatteryGet {
	return BatteryGet{}
}

func (c *BatteryGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *BatteryGet) Set() error {

	return nil
}

func (c *BatteryGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Battery),
		0x02,
		0x25,
	}
}
type BatteryReport struct {
    *report
    BatteryLevel byte
    data []byte
}

func NewBatteryReport(data []byte) *BatteryReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &BatteryReport{
				BatteryLevel: data[0],
        data: data,
    }
}

