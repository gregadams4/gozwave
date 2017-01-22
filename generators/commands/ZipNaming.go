package commands
type ZipNamingNameSet struct {
	node byte
	Name byte
}

func NewZipNamingNameSet() ZipNamingNameSet {
	return ZipNamingNameSet{}
}

func (c *ZipNamingNameSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ZipNamingNameSet) Set(Name byte,) error {
	c.Name = Name

	return nil
}

func (c *ZipNamingNameSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ZipNaming),
		0x01,
		c.Name,
		0x25,
	}
}
type ZipNamingNameGet struct {
	node byte
}

func NewZipNamingNameGet() ZipNamingNameGet {
	return ZipNamingNameGet{}
}

func (c *ZipNamingNameGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ZipNamingNameGet) Set() error {

	return nil
}

func (c *ZipNamingNameGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ZipNaming),
		0x02,
		0x25,
	}
}
type ZipNamingNameReport struct {
    *report
    Name byte
    data []byte
}

func NewZipNamingNameReport(data []byte) *ZipNamingNameReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ZipNamingNameReport{
				Name: data[0],
        data: data,
    }
}

type ZipNamingLocationSet struct {
	node byte
	Location byte
}

func NewZipNamingLocationSet() ZipNamingLocationSet {
	return ZipNamingLocationSet{}
}

func (c *ZipNamingLocationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ZipNamingLocationSet) Set(Location byte,) error {
	c.Location = Location

	return nil
}

func (c *ZipNamingLocationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ZipNaming),
		0x04,
		c.Location,
		0x25,
	}
}
type ZipNamingLocationGet struct {
	node byte
}

func NewZipNamingLocationGet() ZipNamingLocationGet {
	return ZipNamingLocationGet{}
}

func (c *ZipNamingLocationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ZipNamingLocationGet) Set() error {

	return nil
}

func (c *ZipNamingLocationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ZipNaming),
		0x05,
		0x25,
	}
}
type ZipNamingLocationReport struct {
    *report
    Location byte
    data []byte
}

func NewZipNamingLocationReport(data []byte) *ZipNamingLocationReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ZipNamingLocationReport{
				Location: data[0],
        data: data,
    }
}

