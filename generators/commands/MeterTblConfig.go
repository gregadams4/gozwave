package commands
type MeterTblTablePointAdmNoSet struct {
	node byte
	Properties1 byte
	MeterPointAdmNumberCharacter byte
}

func NewMeterTblTablePointAdmNoSet() MeterTblTablePointAdmNoSet {
	return MeterTblTablePointAdmNoSet{}
}

func (c *MeterTblTablePointAdmNoSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblTablePointAdmNoSet) Set(Properties1 byte,MeterPointAdmNumberCharacter byte,) error {
	c.Properties1 = Properties1
	c.MeterPointAdmNumberCharacter = MeterPointAdmNumberCharacter

	return nil
}

func (c *MeterTblTablePointAdmNoSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(MeterTblConfig),
		0x01,
		c.Properties1,
		c.MeterPointAdmNumberCharacter,
		0x25,
	}
}
