package commands
type HumidityControlOperatingStateGet struct {
	node byte
}

func NewHumidityControlOperatingStateGet() HumidityControlOperatingStateGet {
	return HumidityControlOperatingStateGet{}
}

func (c *HumidityControlOperatingStateGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *HumidityControlOperatingStateGet) Set() error {

	return nil
}

func (c *HumidityControlOperatingStateGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(HumidityControlOperatingState),
		0x01,
		0x25,
	}
}
type HumidityControlOperatingStateReport struct {
    *report
    Properties1 byte
    data []byte
}

func NewHumidityControlOperatingStateReport(data []byte) *HumidityControlOperatingStateReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &HumidityControlOperatingStateReport{
				Properties1: data[0],
        data: data,
    }
}

