package commands
type ThermostatHeatingStatusReport struct {
    *report
    Status byte
    data []byte
}

func NewThermostatHeatingStatusReport(data []byte) *ThermostatHeatingStatusReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatHeatingStatusReport{
				Status: data[0],
        data: data,
    }
}

type ThermostatHeatingModeGet struct {
	node byte
}

func NewThermostatHeatingModeGet() ThermostatHeatingModeGet {
	return ThermostatHeatingModeGet{}
}

func (c *ThermostatHeatingModeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatHeatingModeGet) Set() error {

	return nil
}

func (c *ThermostatHeatingModeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatHeating),
		0x02,
		0x25,
	}
}
type ThermostatHeatingModeReport struct {
    *report
    Mode byte
    data []byte
}

func NewThermostatHeatingModeReport(data []byte) *ThermostatHeatingModeReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatHeatingModeReport{
				Mode: data[0],
        data: data,
    }
}

type ThermostatHeatingModeSet struct {
	node byte
	Mode byte
}

func NewThermostatHeatingModeSet() ThermostatHeatingModeSet {
	return ThermostatHeatingModeSet{}
}

func (c *ThermostatHeatingModeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatHeatingModeSet) Set(Mode byte,) error {
	c.Mode = Mode

	return nil
}

func (c *ThermostatHeatingModeSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatHeating),
		0x01,
		c.Mode,
		0x25,
	}
}
type ThermostatHeatingRelayStatusGet struct {
	node byte
}

func NewThermostatHeatingRelayStatusGet() ThermostatHeatingRelayStatusGet {
	return ThermostatHeatingRelayStatusGet{}
}

func (c *ThermostatHeatingRelayStatusGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatHeatingRelayStatusGet) Set() error {

	return nil
}

func (c *ThermostatHeatingRelayStatusGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatHeating),
		0x09,
		0x25,
	}
}
type ThermostatHeatingRelayStatusReport struct {
    *report
    RelayStatus byte
    data []byte
}

func NewThermostatHeatingRelayStatusReport(data []byte) *ThermostatHeatingRelayStatusReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatHeatingRelayStatusReport{
				RelayStatus: data[0],
        data: data,
    }
}

type ThermostatHeatingSetpointGet struct {
	node byte
	SetpointNr byte
}

func NewThermostatHeatingSetpointGet() ThermostatHeatingSetpointGet {
	return ThermostatHeatingSetpointGet{}
}

func (c *ThermostatHeatingSetpointGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatHeatingSetpointGet) Set(SetpointNr byte,) error {
	c.SetpointNr = SetpointNr

	return nil
}

func (c *ThermostatHeatingSetpointGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatHeating),
		0x05,
		c.SetpointNr,
		0x25,
	}
}
type ThermostatHeatingSetpointReport struct {
    *report
    SetpointNr byte
    Properties1 byte
    Value byte
    data []byte
}

func NewThermostatHeatingSetpointReport(data []byte) *ThermostatHeatingSetpointReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatHeatingSetpointReport{
				SetpointNr: data[0],
				Properties1: data[1],
				Value: data[2],
        data: data,
    }
}

type ThermostatHeatingSetpointSet struct {
	node byte
	SetpointNr byte
	Properties1 byte
	Value byte
}

func NewThermostatHeatingSetpointSet() ThermostatHeatingSetpointSet {
	return ThermostatHeatingSetpointSet{}
}

func (c *ThermostatHeatingSetpointSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatHeatingSetpointSet) Set(SetpointNr byte,Properties1 byte,Value byte,) error {
	c.SetpointNr = SetpointNr
	c.Properties1 = Properties1
	c.Value = Value

	return nil
}

func (c *ThermostatHeatingSetpointSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ThermostatHeating),
		0x04,
		c.SetpointNr,
		c.Properties1,
		c.Value,
		0x25,
	}
}
type ThermostatHeatingStatusGet struct {
	node byte
}

func NewThermostatHeatingStatusGet() ThermostatHeatingStatusGet {
	return ThermostatHeatingStatusGet{}
}

func (c *ThermostatHeatingStatusGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatHeatingStatusGet) Set() error {

	return nil
}

func (c *ThermostatHeatingStatusGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatHeating),
		0x0C,
		0x25,
	}
}
type ThermostatHeatingStatusSet struct {
	node byte
	Status byte
}

func NewThermostatHeatingStatusSet() ThermostatHeatingStatusSet {
	return ThermostatHeatingStatusSet{}
}

func (c *ThermostatHeatingStatusSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatHeatingStatusSet) Set(Status byte,) error {
	c.Status = Status

	return nil
}

func (c *ThermostatHeatingStatusSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatHeating),
		0x0B,
		c.Status,
		0x25,
	}
}
type ThermostatHeatingTimedOffSet struct {
	node byte
	Minutes byte
	Hours byte
}

func NewThermostatHeatingTimedOffSet() ThermostatHeatingTimedOffSet {
	return ThermostatHeatingTimedOffSet{}
}

func (c *ThermostatHeatingTimedOffSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatHeatingTimedOffSet) Set(Minutes byte,Hours byte,) error {
	c.Minutes = Minutes
	c.Hours = Hours

	return nil
}

func (c *ThermostatHeatingTimedOffSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ThermostatHeating),
		0x11,
		c.Minutes,
		c.Hours,
		0x25,
	}
}
