package commands

import "log"

type CmdConfigurationGet struct {
	node      byte
	parameter int
}

type CmdConfigurationSet struct {
	node byte
	data []byte
}

type ConfigurationReport struct {
	*report
	node      byte
	Parameter byte
	Value     []byte
	data      []byte
}

func NewConfigurationGet() *CmdConfigurationGet {
	return &CmdConfigurationGet{}
}

func (c *CmdConfigurationGet) SetParameter(v int) {
	c.parameter = v
}

func (c *CmdConfigurationGet) SetNode(v int) {
	c.node = byte(v)
}

func (c *CmdConfigurationGet) Encode() []byte {
	if c.parameter > 255 {
		c.parameter = 255
	} else if c.parameter < 0 {
		c.parameter = 0
	}
	return []byte{
		0x13,              // SEND ZW
		c.node,            // NOD ID
		0x03,              // length
		Configuration,     // command class id
		0x05,              // get
		byte(c.parameter), // value
		0x25,              // transmit options
		//0x08,         // transmit options
	}
}

func NewConfigurationSet() *CmdConfigurationSet {
	return &CmdConfigurationSet{}
}

func (c *CmdConfigurationSet) SetData(v []byte) {
	c.data = v
}

func (c *CmdConfigurationSet) SetNode(v int) {
	c.node = byte(v)
}

func (c *CmdConfigurationSet) Encode() []byte {
	data := []byte{
		0x13,                  // SEND ZW
		c.node,                // NOD ID
		byte(len(c.data) + 2), // length
		Configuration,         // command class id
		0x04,                  // set
	}

	data = append(data, c.data...)
	data = append(data, 0x25)
	return data
}

func NewConfigurationReport(data []byte) *ConfigurationReport {
	if len(data) < 1 {
		data = []byte{0x00}
	}

	log.Println(data)

	if len(data) < 3 {
		return &ConfigurationReport{
			Parameter: data[0],
			data:      data,
		}
	}
	return &ConfigurationReport{
		Parameter: data[0],
		Value:     data[2:],
		data:      data,
	}
}
