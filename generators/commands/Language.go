package commands
import "encoding/binary"
type LanguageGet struct {
	node byte
}

func NewLanguageGet() LanguageGet {
	return LanguageGet{}
}

func (c *LanguageGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *LanguageGet) Set() error {

	return nil
}

func (c *LanguageGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Language),
		0x02,
		0x25,
	}
}
type LanguageReport struct {
    *report
    Language byte
    Country uint16
    data []byte
}

func NewLanguageReport(data []byte) *LanguageReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &LanguageReport{
				Language: data[0],
				Country: binary.BigEndian.Uint16(data[1:3]),
        data: data,
    }
}

type LanguageSet struct {
	node byte
	Language byte
	Country uint16
}

func NewLanguageSet() LanguageSet {
	return LanguageSet{}
}

func (c *LanguageSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *LanguageSet) Set(Language byte,Country uint16,) error {
	c.Language = Language
	c.Country = Country

	return nil
}

func (c *LanguageSet) Encode() []byte {
	CountryBytes := []byte{}
	binary.BigEndian.PutUint16(CountryBytes, c.Country)
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(Language),
		0x01,
		c.Language,
		CountryBytes[0],
		CountryBytes[1],
		0x25,
	}
}
