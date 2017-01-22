package commands
type ProtectionGet struct {
	node byte
}

func NewProtectionGet() ProtectionGet {
	return ProtectionGet{}
}

func (c *ProtectionGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ProtectionGet) Set() error {

	return nil
}

func (c *ProtectionGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Protection),
		0x02,
		0x25,
	}
}
type ProtectionReport struct {
    *report
    ProtectionState byte
    data []byte
}

func NewProtectionReport(data []byte) *ProtectionReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ProtectionReport{
				ProtectionState: data[0],
        data: data,
    }
}

type ProtectionSet struct {
	node byte
	ProtectionState byte
}

func NewProtectionSet() ProtectionSet {
	return ProtectionSet{}
}

func (c *ProtectionSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ProtectionSet) Set(ProtectionState byte,) error {
	c.ProtectionState = ProtectionState

	return nil
}

func (c *ProtectionSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Protection),
		0x01,
		c.ProtectionState,
		0x25,
	}
}
