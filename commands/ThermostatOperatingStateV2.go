package commands
type ThermostatOperatingStateGetV2 struct {
	node byte
}

func NewThermostatOperatingStateGetV2() ThermostatOperatingStateGetV2 {
	return ThermostatOperatingStateGetV2{}
}

func (c *ThermostatOperatingStateGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatOperatingStateGetV2) Set() error {

	return nil
}

func (c *ThermostatOperatingStateGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatOperatingStateV2),
		0x02,
		0x25,
	}
}
type ThermostatOperatingStateReportV2 struct {
    *report
    Properties1 byte
    data []byte
}

func NewThermostatOperatingStateReportV2(data []byte) *ThermostatOperatingStateReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatOperatingStateReportV2{
				Properties1: data[0],
        data: data,
    }
}

type ThermostatOperatingStateLoggingSupportedGetV2 struct {
	node byte
}

func NewThermostatOperatingStateLoggingSupportedGetV2() ThermostatOperatingStateLoggingSupportedGetV2 {
	return ThermostatOperatingStateLoggingSupportedGetV2{}
}

func (c *ThermostatOperatingStateLoggingSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatOperatingStateLoggingSupportedGetV2) Set() error {

	return nil
}

func (c *ThermostatOperatingStateLoggingSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ThermostatOperatingStateV2),
		0x01,
		0x25,
	}
}
type ThermostatOperatingLoggingSupportedReportV2 struct {
    *report
    BitMask byte
    data []byte
}

func NewThermostatOperatingLoggingSupportedReportV2(data []byte) *ThermostatOperatingLoggingSupportedReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatOperatingLoggingSupportedReportV2{
				BitMask: data[0],
        data: data,
    }
}

type ThermostatOperatingStateLoggingGetV2 struct {
	node byte
	BitMask byte
}

func NewThermostatOperatingStateLoggingGetV2() ThermostatOperatingStateLoggingGetV2 {
	return ThermostatOperatingStateLoggingGetV2{}
}

func (c *ThermostatOperatingStateLoggingGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ThermostatOperatingStateLoggingGetV2) Set(BitMask byte,) error {
	c.BitMask = BitMask

	return nil
}

func (c *ThermostatOperatingStateLoggingGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ThermostatOperatingStateV2),
		0x05,
		c.BitMask,
		0x25,
	}
}
type ThermostatOperatingStateLoggingReportV2 struct {
    *report
    ReportstoFollow byte
    data []byte
}

func NewThermostatOperatingStateLoggingReportV2(data []byte) *ThermostatOperatingStateLoggingReportV2 {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ThermostatOperatingStateLoggingReportV2{
				ReportstoFollow: data[0],
        data: data,
    }
}

