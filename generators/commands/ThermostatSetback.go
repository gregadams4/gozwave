package commands
type ThermostatSetbackGet struct {
	node byte
}

func NewThermostatSetbackGet() ThermostatSetbackGet {
	return ThermostatSetbackGet{}
}

func (c *ThermostatSetbackGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatSetbackGet) Set() error {

	return nil
}

func (c *ThermostatSetbackGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatSetback),
		0x02,
		0x25,
	}
}
type ThermostatSetbackReport struct {
    *report
    Properties1 byte
    SetbackState byte
    data []byte
}

func NewThermostatSetbackReport(data []byte) *ThermostatSetbackReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatSetbackReport{
				Properties1: data[0],
				SetbackState: data[1],
        data: data,
    }
}

type ThermostatSetbackSet struct {
	node byte
	Properties1 byte
	SetbackState byte
}

func NewThermostatSetbackSet() ThermostatSetbackSet {
	return ThermostatSetbackSet{}
}

func (c *ThermostatSetbackSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatSetbackSet) Set(Properties1 byte,SetbackState byte,) error {
	c.Properties1 = Properties1
	c.SetbackState = SetbackState

	return nil
}

func (c *ThermostatSetbackSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ThermostatSetback),
		0x01,
		c.Properties1,
		c.SetbackState,
		0x25,
	}
}
