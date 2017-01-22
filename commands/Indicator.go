package commands
type IndicatorGet struct {
	node byte
}

func NewIndicatorGet() IndicatorGet {
	return IndicatorGet{}
}

func (c *IndicatorGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IndicatorGet) Set() error {

	return nil
}

func (c *IndicatorGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Indicator),
		0x02,
		0x25,
	}
}
type IndicatorReport struct {
    *report
    Value byte
    data []byte
}

func NewIndicatorReport(data []byte) *IndicatorReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &IndicatorReport{
				Value: data[0],
        data: data,
    }
}

type IndicatorSet struct {
	node byte
	Value byte
}

func NewIndicatorSet() IndicatorSet {
	return IndicatorSet{}
}

func (c *IndicatorSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IndicatorSet) Set(Value byte,) error {
	c.Value = Value

	return nil
}

func (c *IndicatorSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Indicator),
		0x01,
		c.Value,
		0x25,
	}
}
