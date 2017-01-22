package commands
type ProprietaryGet struct {
	node byte
	Data byte
}

func NewProprietaryGet() ProprietaryGet {
	return ProprietaryGet{}
}

func (c *ProprietaryGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ProprietaryGet) Set(Data byte,) error {
	c.Data = Data

	return nil
}

func (c *ProprietaryGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Proprietary),
		0x02,
		c.Data,
		0x25,
	}
}
type ProprietaryReport struct {
    *report
    Data byte
    data []byte
}

func NewProprietaryReport(data []byte) *ProprietaryReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ProprietaryReport{
				Data: data[0],
        data: data,
    }
}

type ProprietarySet struct {
	node byte
	Data byte
}

func NewProprietarySet() ProprietarySet {
	return ProprietarySet{}
}

func (c *ProprietarySet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ProprietarySet) Set(Data byte,) error {
	c.Data = Data

	return nil
}

func (c *ProprietarySet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Proprietary),
		0x01,
		c.Data,
		0x25,
	}
}
