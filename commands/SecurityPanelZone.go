package commands
type SecurityPanelZoneNumberSupportedGet struct {
	node byte
}

func NewSecurityPanelZoneNumberSupportedGet() SecurityPanelZoneNumberSupportedGet {
	return SecurityPanelZoneNumberSupportedGet{}
}

func (c *SecurityPanelZoneNumberSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityPanelZoneNumberSupportedGet) Set() error {

	return nil
}

func (c *SecurityPanelZoneNumberSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SecurityPanelZone),
		0x01,
		0x25,
	}
}
type SecurityPanelZoneStateGet struct {
	node byte
	ZoneNumber byte
}

func NewSecurityPanelZoneStateGet() SecurityPanelZoneStateGet {
	return SecurityPanelZoneStateGet{}
}

func (c *SecurityPanelZoneStateGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityPanelZoneStateGet) Set(ZoneNumber byte,) error {
	c.ZoneNumber = ZoneNumber

	return nil
}

func (c *SecurityPanelZoneStateGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SecurityPanelZone),
		0x05,
		c.ZoneNumber,
		0x25,
	}
}
type SecurityPanelZoneStateReport struct {
    *report
    Zonenumber byte
    ZoneState byte
    data []byte
}

func NewSecurityPanelZoneStateReport(data []byte) *SecurityPanelZoneStateReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SecurityPanelZoneStateReport{
				Zonenumber: data[0],
				ZoneState: data[1],
        data: data,
    }
}

type SecurityPanelZoneSupportedReport struct {
    *report
    Parameters1 byte
    data []byte
}

func NewSecurityPanelZoneSupportedReport(data []byte) *SecurityPanelZoneSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SecurityPanelZoneSupportedReport{
				Parameters1: data[0],
        data: data,
    }
}

type SecurityPanelZoneTypeGet struct {
	node byte
	ZoneNumber byte
}

func NewSecurityPanelZoneTypeGet() SecurityPanelZoneTypeGet {
	return SecurityPanelZoneTypeGet{}
}

func (c *SecurityPanelZoneTypeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityPanelZoneTypeGet) Set(ZoneNumber byte,) error {
	c.ZoneNumber = ZoneNumber

	return nil
}

func (c *SecurityPanelZoneTypeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SecurityPanelZone),
		0x03,
		c.ZoneNumber,
		0x25,
	}
}
type SecurityPanelZoneTypeReport struct {
    *report
    ZoneNumber byte
    ZoneType byte
    data []byte
}

func NewSecurityPanelZoneTypeReport(data []byte) *SecurityPanelZoneTypeReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SecurityPanelZoneTypeReport{
				ZoneNumber: data[0],
				ZoneType: data[1],
        data: data,
    }
}

