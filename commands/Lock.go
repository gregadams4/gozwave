package commands
type LockGet struct {
	node byte
}

func NewLockGet() LockGet {
	return LockGet{}
}

func (c *LockGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *LockGet) Set() error {

	return nil
}

func (c *LockGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Lock),
		0x02,
		0x25,
	}
}
type LockReport struct {
    *report
    LockState byte
    data []byte
}

func NewLockReport(data []byte) *LockReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &LockReport{
				LockState: data[0],
        data: data,
    }
}

type LockSet struct {
	node byte
	LockState byte
}

func NewLockSet() LockSet {
	return LockSet{}
}

func (c *LockSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *LockSet) Set(LockState byte,) error {
	c.LockState = LockState

	return nil
}

func (c *LockSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Lock),
		0x01,
		c.LockState,
		0x25,
	}
}
