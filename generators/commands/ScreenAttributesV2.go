package commands
type ScreenAttributesGetV2 struct {
	node byte
}

func NewScreenAttributesGetV2() ScreenAttributesGetV2 {
	return ScreenAttributesGetV2{}
}

func (c *ScreenAttributesGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ScreenAttributesGetV2) Set() error {

	return nil
}

func (c *ScreenAttributesGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ScreenAttributesV2),
		0x01,
		0x25,
	}
}
type ScreenAttributesReportV2 struct {
    *report
    Properties1 byte
    NumberofCharactersperLine byte
    SizeofLineBuffer byte
    NumericalPresentationofaCharacter byte
    ScreenTimeout byte
    data []byte
}

func NewScreenAttributesReportV2(data []byte) *ScreenAttributesReportV2 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &ScreenAttributesReportV2{
				Properties1: data[0],
				NumberofCharactersperLine: data[1],
				SizeofLineBuffer: data[2],
				NumericalPresentationofaCharacter: data[3],
				ScreenTimeout: data[4],
        data: data,
    }
}

