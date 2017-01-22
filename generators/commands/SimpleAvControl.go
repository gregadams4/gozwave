package commands
import "encoding/binary"
type SimpleAvControlGet struct {
	node byte
}

func NewSimpleAvControlGet() SimpleAvControlGet {
	return SimpleAvControlGet{}
}

func (c *SimpleAvControlGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SimpleAvControlGet) Set() error {

	return nil
}

func (c *SimpleAvControlGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(SimpleAvControl),
		0x02,
		0x25,
	}
}
type SimpleAvControlReport struct {
    *report
    Numberofreports byte
    data []byte
}

func NewSimpleAvControlReport(data []byte) *SimpleAvControlReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SimpleAvControlReport{
				Numberofreports: data[0],
        data: data,
    }
}

type SimpleAvControlSet struct {
	node byte
	SequenceNumber byte
	Properties1 byte
	ItemID uint16
	Vg []SimpleAvControlSetvg
}
type SimpleAvControlSetvg struct {
	Command uint16
}

func NewSimpleAvControlSet() SimpleAvControlSet {
	return SimpleAvControlSet{}
}

func (c *SimpleAvControlSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SimpleAvControlSet) Set(SequenceNumber byte,Properties1 byte,ItemID uint16,Vg []SimpleAvControlSetvg,) error {
	c.SequenceNumber = SequenceNumber
	c.Properties1 = Properties1
	c.ItemID = ItemID
	c.Vg = Vg

	return nil
}

func (c *SimpleAvControlSet) Encode() []byte {
	ItemIDBytes := []byte{}
	binary.BigEndian.PutUint16(ItemIDBytes, c.ItemID)
	var data = []byte{
		0x13,
		c.node,
		byte(6),
		byte(SimpleAvControl),
		0x01,
		c.SequenceNumber,
		c.Properties1,
		ItemIDBytes[0],
		ItemIDBytes[1],
	}
	for _, v := range c.Vg {
		CommandBytes := []byte{}
		binary.BigEndian.PutUint16(CommandBytes, v.Command)
		data = append(data, CommandBytes...)
	}
	data = append(data, 0x25)
	return data
	
}
type SimpleAvControlSupportedGet struct {
	node byte
	ReportNo byte
}

func NewSimpleAvControlSupportedGet() SimpleAvControlSupportedGet {
	return SimpleAvControlSupportedGet{}
}

func (c *SimpleAvControlSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SimpleAvControlSupportedGet) Set(ReportNo byte,) error {
	c.ReportNo = ReportNo

	return nil
}

func (c *SimpleAvControlSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(SimpleAvControl),
		0x04,
		c.ReportNo,
		0x25,
	}
}
type SimpleAvControlSupportedReport struct {
    *report
    ReportNo byte
    BitMask byte
    data []byte
}

func NewSimpleAvControlSupportedReport(data []byte) *SimpleAvControlSupportedReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &SimpleAvControlSupportedReport{
				ReportNo: data[0],
				BitMask: data[1],
        data: data,
    }
}

