package commands
type BasicGet struct {
	node byte
}

func NewBasicGet() BasicGet {
	return BasicGet{}
}

func (c *BasicGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *BasicGet) Set() error {

	return nil
}

func (c *BasicGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Basic),
		0x02,
		0x25,
	}
}
type BasicReport struct {
    *report
    Value byte
    data []byte
}

func NewBasicReport(data []byte) *BasicReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &BasicReport{
				Value: data[0],
        data: data,
    }
}

type BasicSet struct {
	node byte
	Value byte
}

func NewBasicSet() BasicSet {
	return BasicSet{}
}

func (c *BasicSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *BasicSet) Set(Value byte,) error {
	c.Value = Value

	return nil
}

func (c *BasicSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Basic),
		0x01,
		c.Value,
		0x25,
	}
}
