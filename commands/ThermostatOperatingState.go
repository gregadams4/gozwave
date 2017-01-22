package commands
type ThermostatOperatingStateGet struct {
	node byte
}

func NewThermostatOperatingStateGet() ThermostatOperatingStateGet {
	return ThermostatOperatingStateGet{}
}

func (c *ThermostatOperatingStateGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatOperatingStateGet) Set() error {

	return nil
}

func (c *ThermostatOperatingStateGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatOperatingState),
		0x02,
		0x25,
	}
}
type ThermostatOperatingStateReport struct {
    *report
    Level byte
    data []byte
}

func NewThermostatOperatingStateReport(data []byte) *ThermostatOperatingStateReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatOperatingStateReport{
				Level: data[0],
        data: data,
    }
}

