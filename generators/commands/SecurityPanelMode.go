package commands
import "encoding/binary"
type SecurityPanelModeGet struct {
	node byte
}

func NewSecurityPanelModeGet() SecurityPanelModeGet {
	return SecurityPanelModeGet{}
}

func (c *SecurityPanelModeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityPanelModeGet) Set() error {

	return nil
}

func (c *SecurityPanelModeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SecurityPanelMode),
		0x03,
		0x25,
	}
}
type SecurityPanelModeReport struct {
    *report
    MODE byte
    data []byte
}

func NewSecurityPanelModeReport(data []byte) *SecurityPanelModeReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SecurityPanelModeReport{
				MODE: data[0],
        data: data,
    }
}

type SecurityPanelModeSet struct {
	node byte
	MODE byte
}

func NewSecurityPanelModeSet() SecurityPanelModeSet {
	return SecurityPanelModeSet{}
}

func (c *SecurityPanelModeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityPanelModeSet) Set(MODE byte,) error {
	c.MODE = MODE

	return nil
}

func (c *SecurityPanelModeSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SecurityPanelMode),
		0x05,
		c.MODE,
		0x25,
	}
}
type SecurityPanelModeSupportedGet struct {
	node byte
}

func NewSecurityPanelModeSupportedGet() SecurityPanelModeSupportedGet {
	return SecurityPanelModeSupportedGet{}
}

func (c *SecurityPanelModeSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityPanelModeSupportedGet) Set() error {

	return nil
}

func (c *SecurityPanelModeSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SecurityPanelMode),
		0x01,
		0x25,
	}
}
type SecurityPanelModeSupportedReport struct {
    *report
    SupportedModeBitMask uint16
    data []byte
}

func NewSecurityPanelModeSupportedReport(data []byte) *SecurityPanelModeSupportedReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SecurityPanelModeSupportedReport{
				SupportedModeBitMask: binary.BigEndian.Uint16(data[0:2]),
        data: data,
    }
}

