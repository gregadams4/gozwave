package commands
type BarrierOperatorSet struct {
	node byte
	TargetValue byte
}

func NewBarrierOperatorSet() BarrierOperatorSet {
	return BarrierOperatorSet{}
}

func (c *BarrierOperatorSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *BarrierOperatorSet) Set(TargetValue byte,) error {
	c.TargetValue = TargetValue

	return nil
}

func (c *BarrierOperatorSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(BarrierOperator),
		0x01,
		c.TargetValue,
		0x25,
	}
}
type BarrierOperatorGet struct {
	node byte
}

func NewBarrierOperatorGet() BarrierOperatorGet {
	return BarrierOperatorGet{}
}

func (c *BarrierOperatorGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *BarrierOperatorGet) Set() error {

	return nil
}

func (c *BarrierOperatorGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(BarrierOperator),
		0x02,
		0x25,
	}
}
type BarrierOperatorReport struct {
    *report
    State byte
    data []byte
}

func NewBarrierOperatorReport(data []byte) *BarrierOperatorReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &BarrierOperatorReport{
				State: data[0],
        data: data,
    }
}

type BarrierOperatorSignalSupportedGet struct {
	node byte
}

func NewBarrierOperatorSignalSupportedGet() BarrierOperatorSignalSupportedGet {
	return BarrierOperatorSignalSupportedGet{}
}

func (c *BarrierOperatorSignalSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *BarrierOperatorSignalSupportedGet) Set() error {

	return nil
}

func (c *BarrierOperatorSignalSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(BarrierOperator),
		0x04,
		0x25,
	}
}
type BarrierOperatorSignalSupportedReport struct {
    *report
    BitMask byte
    data []byte
}

func NewBarrierOperatorSignalSupportedReport(data []byte) *BarrierOperatorSignalSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &BarrierOperatorSignalSupportedReport{
				BitMask: data[0],
        data: data,
    }
}

type BarrierOperatorSignalSet struct {
	node byte
	SubsystemType byte
	SubsystemState byte
}

func NewBarrierOperatorSignalSet() BarrierOperatorSignalSet {
	return BarrierOperatorSignalSet{}
}

func (c *BarrierOperatorSignalSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *BarrierOperatorSignalSet) Set(SubsystemType byte,SubsystemState byte,) error {
	c.SubsystemType = SubsystemType
	c.SubsystemState = SubsystemState

	return nil
}

func (c *BarrierOperatorSignalSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(BarrierOperator),
		0x06,
		c.SubsystemType,
		c.SubsystemState,
		0x25,
	}
}
type BarrierOperatorSignalGet struct {
	node byte
	SubsystemType byte
}

func NewBarrierOperatorSignalGet() BarrierOperatorSignalGet {
	return BarrierOperatorSignalGet{}
}

func (c *BarrierOperatorSignalGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *BarrierOperatorSignalGet) Set(SubsystemType byte,) error {
	c.SubsystemType = SubsystemType

	return nil
}

func (c *BarrierOperatorSignalGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(BarrierOperator),
		0x07,
		c.SubsystemType,
		0x25,
	}
}
type BarrierOperatorSignalReport struct {
    *report
    SubsystemType byte
    SubsystemState byte
    data []byte
}

func NewBarrierOperatorSignalReport(data []byte) *BarrierOperatorSignalReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &BarrierOperatorSignalReport{
				SubsystemType: data[0],
				SubsystemState: data[1],
        data: data,
    }
}

