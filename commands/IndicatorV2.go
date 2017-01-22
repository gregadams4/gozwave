package commands
type IndicatorGetV2 struct {
	node byte
	IndicatorID byte
}

func NewIndicatorGetV2() IndicatorGetV2 {
	return IndicatorGetV2{}
}

func (c *IndicatorGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *IndicatorGetV2) Set(IndicatorID byte,) error {
	c.IndicatorID = IndicatorID

	return nil
}

func (c *IndicatorGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(IndicatorV2),
		0x02,
		c.IndicatorID,
		0x25,
	}
}
type IndicatorReportV2 struct {
    *report
    Indicator0Value byte
    Properties1 byte
    data []byte
}

func NewIndicatorReportV2(data []byte) *IndicatorReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &IndicatorReportV2{
				Indicator0Value: data[0],
				Properties1: data[1],
        data: data,
    }
}

type IndicatorSetV2 struct {
	node byte
	Indicator0Value byte
	Properties1 byte
	Vg1 []IndicatorSetV2vg1
}
type IndicatorSetV2vg1 struct {
	IndicatorID byte
	PropertyID byte
	Value byte
}

func NewIndicatorSetV2() IndicatorSetV2 {
	return IndicatorSetV2{}
}

func (c *IndicatorSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *IndicatorSetV2) Set(Indicator0Value byte,Properties1 byte,Vg1 []IndicatorSetV2vg1,) error {
	c.Indicator0Value = Indicator0Value
	c.Properties1 = Properties1
	c.Vg1 = Vg1

	return nil
}

func (c *IndicatorSetV2) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(7),
		byte(IndicatorV2),
		0x01,
		c.Indicator0Value,
		c.Properties1,
	}
	for _, v := range c.Vg1 {
		data = append(data, v.IndicatorID)
		data = append(data, v.PropertyID)
		data = append(data, v.Value)
	}
	data = append(data, 0x25)
	return data
	
}
type IndicatorSupportedGetV2 struct {
	node byte
	IndicatorID byte
}

func NewIndicatorSupportedGetV2() IndicatorSupportedGetV2 {
	return IndicatorSupportedGetV2{}
}

func (c *IndicatorSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *IndicatorSupportedGetV2) Set(IndicatorID byte,) error {
	c.IndicatorID = IndicatorID

	return nil
}

func (c *IndicatorSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(IndicatorV2),
		0x04,
		c.IndicatorID,
		0x25,
	}
}
type IndicatorSupportedReportV2 struct {
    *report
    IndicatorID byte
    NextIndicatorID byte
    Properties1 byte
    PropertySupportedBitMask byte
    data []byte
}

func NewIndicatorSupportedReportV2(data []byte) *IndicatorSupportedReportV2 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &IndicatorSupportedReportV2{
				IndicatorID: data[0],
				NextIndicatorID: data[1],
				Properties1: data[2],
				PropertySupportedBitMask: data[3],
        data: data,
    }
}

