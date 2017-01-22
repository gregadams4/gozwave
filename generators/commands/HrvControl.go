package commands
type HrvControlBypassGet struct {
	node byte
}

func NewHrvControlBypassGet() HrvControlBypassGet {
	return HrvControlBypassGet{}
}

func (c *HrvControlBypassGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HrvControlBypassGet) Set() error {

	return nil
}

func (c *HrvControlBypassGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(HrvControl),
		0x05,
		0x25,
	}
}
type HrvControlBypassReport struct {
    *report
    Bypass byte
    data []byte
}

func NewHrvControlBypassReport(data []byte) *HrvControlBypassReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &HrvControlBypassReport{
				Bypass: data[0],
        data: data,
    }
}

type HrvControlBypassSet struct {
	node byte
	Bypass byte
}

func NewHrvControlBypassSet() HrvControlBypassSet {
	return HrvControlBypassSet{}
}

func (c *HrvControlBypassSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HrvControlBypassSet) Set(Bypass byte,) error {
	c.Bypass = Bypass

	return nil
}

func (c *HrvControlBypassSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(HrvControl),
		0x04,
		c.Bypass,
		0x25,
	}
}
type HrvControlModeGet struct {
	node byte
}

func NewHrvControlModeGet() HrvControlModeGet {
	return HrvControlModeGet{}
}

func (c *HrvControlModeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HrvControlModeGet) Set() error {

	return nil
}

func (c *HrvControlModeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(HrvControl),
		0x02,
		0x25,
	}
}
type HrvControlModeReport struct {
    *report
    Properties1 byte
    data []byte
}

func NewHrvControlModeReport(data []byte) *HrvControlModeReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &HrvControlModeReport{
				Properties1: data[0],
        data: data,
    }
}

type HrvControlModeSet struct {
	node byte
	Properties1 byte
}

func NewHrvControlModeSet() HrvControlModeSet {
	return HrvControlModeSet{}
}

func (c *HrvControlModeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HrvControlModeSet) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *HrvControlModeSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(HrvControl),
		0x01,
		c.Properties1,
		0x25,
	}
}
type HrvControlModeSupportedGet struct {
	node byte
}

func NewHrvControlModeSupportedGet() HrvControlModeSupportedGet {
	return HrvControlModeSupportedGet{}
}

func (c *HrvControlModeSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HrvControlModeSupportedGet) Set() error {

	return nil
}

func (c *HrvControlModeSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(HrvControl),
		0x0A,
		0x25,
	}
}
type HrvControlModeSupportedReport struct {
    *report
    Properties1 byte
    BitMask byte
    data []byte
}

func NewHrvControlModeSupportedReport(data []byte) *HrvControlModeSupportedReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &HrvControlModeSupportedReport{
				Properties1: data[0],
				BitMask: data[1],
        data: data,
    }
}

type HrvControlVentilationRateGet struct {
	node byte
}

func NewHrvControlVentilationRateGet() HrvControlVentilationRateGet {
	return HrvControlVentilationRateGet{}
}

func (c *HrvControlVentilationRateGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HrvControlVentilationRateGet) Set() error {

	return nil
}

func (c *HrvControlVentilationRateGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(HrvControl),
		0x08,
		0x25,
	}
}
type HrvControlVentilationRateReport struct {
    *report
    VentilationRate byte
    data []byte
}

func NewHrvControlVentilationRateReport(data []byte) *HrvControlVentilationRateReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &HrvControlVentilationRateReport{
				VentilationRate: data[0],
        data: data,
    }
}

type HrvControlVentilationRateSet struct {
	node byte
	VentilationRate byte
}

func NewHrvControlVentilationRateSet() HrvControlVentilationRateSet {
	return HrvControlVentilationRateSet{}
}

func (c *HrvControlVentilationRateSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HrvControlVentilationRateSet) Set(VentilationRate byte,) error {
	c.VentilationRate = VentilationRate

	return nil
}

func (c *HrvControlVentilationRateSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(HrvControl),
		0x07,
		c.VentilationRate,
		0x25,
	}
}
