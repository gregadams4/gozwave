package commands
import "encoding/binary"
type CommandConfigurationGet struct {
	node byte
	Groupingidentifier byte
	NodeID byte
}

func NewCommandConfigurationGet() CommandConfigurationGet {
	return CommandConfigurationGet{}
}

func (c *CommandConfigurationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandConfigurationGet) Set(Groupingidentifier byte,NodeID byte,) error {
	c.Groupingidentifier = Groupingidentifier
	c.NodeID = NodeID

	return nil
}

func (c *CommandConfigurationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(AssociationCommandConfiguration),
		0x04,
		c.Groupingidentifier,
		c.NodeID,
		0x25,
	}
}
type CommandConfigurationReport struct {
    *report
    Groupingidentifier byte
    NodeID byte
    Properties1 byte
    Commandlength byte
    CommandClassidentifier byte
    Commandidentifier byte
    Commandbyte byte
    data []byte
}

func NewCommandConfigurationReport(data []byte) *CommandConfigurationReport {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &CommandConfigurationReport{
				Groupingidentifier: data[0],
				NodeID: data[1],
				Properties1: data[2],
				Commandlength: data[3],
				CommandClassidentifier: data[4],
				Commandidentifier: data[5],
				Commandbyte: data[6],
        data: data,
    }
}

type CommandConfigurationSet struct {
	node byte
	Groupingidentifier byte
	NodeID byte
	Commandlength byte
	CommandClassidentifier byte
	Commandidentifier byte
	Commandbyte byte
}

func NewCommandConfigurationSet() CommandConfigurationSet {
	return CommandConfigurationSet{}
}

func (c *CommandConfigurationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandConfigurationSet) Set(Groupingidentifier byte,NodeID byte,Commandlength byte,CommandClassidentifier byte,Commandidentifier byte,Commandbyte byte,) error {
	c.Groupingidentifier = Groupingidentifier
	c.NodeID = NodeID
	c.Commandlength = Commandlength
	c.CommandClassidentifier = CommandClassidentifier
	c.Commandidentifier = Commandidentifier
	c.Commandbyte = Commandbyte

	return nil
}

func (c *CommandConfigurationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(8),
		byte(AssociationCommandConfiguration),
		0x03,
		c.Groupingidentifier,
		c.NodeID,
		c.Commandlength,
		c.CommandClassidentifier,
		c.Commandidentifier,
		c.Commandbyte,
		0x25,
	}
}
type CommandRecordsSupportedGet struct {
	node byte
}

func NewCommandRecordsSupportedGet() CommandRecordsSupportedGet {
	return CommandRecordsSupportedGet{}
}

func (c *CommandRecordsSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *CommandRecordsSupportedGet) Set() error {

	return nil
}

func (c *CommandRecordsSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(AssociationCommandConfiguration),
		0x01,
		0x25,
	}
}
type CommandRecordsSupportedReport struct {
    *report
    Properties1 byte
    FreeCommandrecords uint16
    MaxCommandrecords uint16
    data []byte
}

func NewCommandRecordsSupportedReport(data []byte) *CommandRecordsSupportedReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &CommandRecordsSupportedReport{
				Properties1: data[0],
				FreeCommandrecords: binary.BigEndian.Uint16(data[1:3]),
				MaxCommandrecords: binary.BigEndian.Uint16(data[3:5]),
        data: data,
    }
}

