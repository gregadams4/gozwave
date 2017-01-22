package commands
type ThermostatFanStateGet struct {
	node byte
}

func NewThermostatFanStateGet() ThermostatFanStateGet {
	return ThermostatFanStateGet{}
}

func (c *ThermostatFanStateGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatFanStateGet) Set() error {

	return nil
}

func (c *ThermostatFanStateGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatFanState),
		0x02,
		0x25,
	}
}
type ThermostatFanStateReport struct {
    *report
    Level byte
    data []byte
}

func NewThermostatFanStateReport(data []byte) *ThermostatFanStateReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatFanStateReport{
				Level: data[0],
        data: data,
    }
}

