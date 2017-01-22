package commands
type ThermostatFanStateGetV2 struct {
	node byte
}

func NewThermostatFanStateGetV2() ThermostatFanStateGetV2 {
	return ThermostatFanStateGetV2{}
}

func (c *ThermostatFanStateGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanStateGetV2) Set() error {

	return nil
}

func (c *ThermostatFanStateGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatFanStateV2),
		0x02,
		0x25,
	}
}
type ThermostatFanStateReportV2 struct {
    *report
    Level byte
    data []byte
}

func NewThermostatFanStateReportV2(data []byte) *ThermostatFanStateReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatFanStateReportV2{
				Level: data[0],
        data: data,
    }
}

