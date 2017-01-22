package commands
type EnergyProductionGet struct {
	node byte
	ParameterNumber byte
}

func NewEnergyProductionGet() EnergyProductionGet {
	return EnergyProductionGet{}
}

func (c *EnergyProductionGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *EnergyProductionGet) Set(ParameterNumber byte,) error {
	c.ParameterNumber = ParameterNumber

	return nil
}

func (c *EnergyProductionGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(EnergyProduction),
		0x02,
		c.ParameterNumber,
		0x25,
	}
}
type EnergyProductionReport struct {
    *report
    ParameterNumber byte
    Level byte
    Value byte
    data []byte
}

func NewEnergyProductionReport(data []byte) *EnergyProductionReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &EnergyProductionReport{
				ParameterNumber: data[0],
				Level: data[1],
				Value: data[2],
        data: data,
    }
}

