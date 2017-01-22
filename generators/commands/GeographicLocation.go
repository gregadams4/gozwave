package commands
type GeographicLocationGet struct {
	node byte
}

func NewGeographicLocationGet() GeographicLocationGet {
	return GeographicLocationGet{}
}

func (c *GeographicLocationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *GeographicLocationGet) Set() error {

	return nil
}

func (c *GeographicLocationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(GeographicLocation),
		0x02,
		0x25,
	}
}
type GeographicLocationReport struct {
    *report
    LongitudeDegrees byte
    Level byte
    LatitudeDegrees byte
    Level2 byte
    data []byte
}

func NewGeographicLocationReport(data []byte) *GeographicLocationReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &GeographicLocationReport{
				LongitudeDegrees: data[0],
				Level: data[1],
				LatitudeDegrees: data[2],
				Level2: data[3],
        data: data,
    }
}

type GeographicLocationSet struct {
	node byte
	LongitudeDegrees byte
	Level byte
	LatitudeDegrees byte
	Level2 byte
}

func NewGeographicLocationSet() GeographicLocationSet {
	return GeographicLocationSet{}
}

func (c *GeographicLocationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *GeographicLocationSet) Set(LongitudeDegrees byte,Level byte,LatitudeDegrees byte,Level2 byte,) error {
	c.LongitudeDegrees = LongitudeDegrees
	c.Level = Level
	c.LatitudeDegrees = LatitudeDegrees
	c.Level2 = Level2

	return nil
}

func (c *GeographicLocationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(GeographicLocation),
		0x01,
		c.LongitudeDegrees,
		c.Level,
		c.LatitudeDegrees,
		c.Level2,
		0x25,
	}
}
