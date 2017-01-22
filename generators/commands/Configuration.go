package commands
type ConfigurationGet struct {
	node byte
	ParameterNumber byte
}

func NewConfigurationGet() ConfigurationGet {
	return ConfigurationGet{}
}

func (c *ConfigurationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationGet) Set(ParameterNumber byte,) error {
	c.ParameterNumber = ParameterNumber

	return nil
}

func (c *ConfigurationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Configuration),
		0x05,
		c.ParameterNumber,
		0x25,
	}
}
type ConfigurationReport struct {
    *report
    ParameterNumber byte
    Level byte
    ConfigurationValue byte
    data []byte
}

func NewConfigurationReport(data []byte) *ConfigurationReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationReport{
				ParameterNumber: data[0],
				Level: data[1],
				ConfigurationValue: data[2],
        data: data,
    }
}

type ConfigurationSet struct {
	node byte
	ParameterNumber byte
	Level byte
	ConfigurationValue byte
}

func NewConfigurationSet() ConfigurationSet {
	return ConfigurationSet{}
}

func (c *ConfigurationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationSet) Set(ParameterNumber byte,Level byte,ConfigurationValue byte,) error {
	c.ParameterNumber = ParameterNumber
	c.Level = Level
	c.ConfigurationValue = ConfigurationValue

	return nil
}

func (c *ConfigurationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(Configuration),
		0x04,
		c.ParameterNumber,
		c.Level,
		c.ConfigurationValue,
		0x25,
	}
}
