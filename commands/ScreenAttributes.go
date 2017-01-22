package commands
type ScreenAttributesGet struct {
	node byte
}

func NewScreenAttributesGet() ScreenAttributesGet {
	return ScreenAttributesGet{}
}

func (c *ScreenAttributesGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScreenAttributesGet) Set() error {

	return nil
}

func (c *ScreenAttributesGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ScreenAttributes),
		0x01,
		0x25,
	}
}
type ScreenAttributesReport struct {
    *report
    Properties1 byte
    NumberofCharactersperLine byte
    SizeofLineBuffer byte
    NumericalPresentationofaCharacter byte
    data []byte
}

func NewScreenAttributesReport(data []byte) *ScreenAttributesReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScreenAttributesReport{
				Properties1: data[0],
				NumberofCharactersperLine: data[1],
				SizeofLineBuffer: data[2],
				NumericalPresentationofaCharacter: data[3],
        data: data,
    }
}

