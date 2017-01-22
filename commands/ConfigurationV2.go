package commands
import "encoding/binary"
type ConfigurationBulkGetV2 struct {
	node byte
	ParameterOffset uint16
	NumberofParameters byte
}

func NewConfigurationBulkGetV2() ConfigurationBulkGetV2 {
	return ConfigurationBulkGetV2{}
}

func (c *ConfigurationBulkGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationBulkGetV2) Set(ParameterOffset uint16,NumberofParameters byte,) error {
	c.ParameterOffset = ParameterOffset
	c.NumberofParameters = NumberofParameters

	return nil
}

func (c *ConfigurationBulkGetV2) Encode() []byte {
	ParameterOffsetBytes := []byte{}
	binary.BigEndian.PutUint16(ParameterOffsetBytes, c.ParameterOffset)
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ConfigurationV2),
		0x08,
		ParameterOffsetBytes[0],
		ParameterOffsetBytes[1],
		c.NumberofParameters,
		0x25,
	}
}
type ConfigurationBulkReportV2 struct {
    *report
    ParameterOffset uint16
    NumberofParameters byte
    Reportstofollow byte
    Properties1 byte
    data []byte
}

func NewConfigurationBulkReportV2(data []byte) *ConfigurationBulkReportV2 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationBulkReportV2{
				ParameterOffset: binary.BigEndian.Uint16(data[0:2]),
				NumberofParameters: data[2],
				Reportstofollow: data[3],
				Properties1: data[4],
        data: data,
    }
}

type ConfigurationBulkSetV2 struct {
	node byte
	ParameterOffset uint16
	NumberofParameters byte
	Properties1 byte
	Vg []ConfigurationBulkSetV2vg
}
type ConfigurationBulkSetV2vg struct {
	Parameter byte
}

func NewConfigurationBulkSetV2() ConfigurationBulkSetV2 {
	return ConfigurationBulkSetV2{}
}

func (c *ConfigurationBulkSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationBulkSetV2) Set(ParameterOffset uint16,NumberofParameters byte,Properties1 byte,Vg []ConfigurationBulkSetV2vg,) error {
	c.ParameterOffset = ParameterOffset
	c.NumberofParameters = NumberofParameters
	c.Properties1 = Properties1
	c.Vg = Vg

	return nil
}

func (c *ConfigurationBulkSetV2) Encode() []byte {
	ParameterOffsetBytes := []byte{}
	binary.BigEndian.PutUint16(ParameterOffsetBytes, c.ParameterOffset)
	var data = []byte{
		0x13,
		c.node,
		byte(6),
		byte(ConfigurationV2),
		0x07,
		ParameterOffsetBytes[0],
		ParameterOffsetBytes[1],
		c.NumberofParameters,
		c.Properties1,
	}
	for _, v := range c.Vg {
		data = append(data, v.Parameter)
	}
	data = append(data, 0x25)
	return data
	
}
type ConfigurationGetV2 struct {
	node byte
	ParameterNumber byte
}

func NewConfigurationGetV2() ConfigurationGetV2 {
	return ConfigurationGetV2{}
}

func (c *ConfigurationGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationGetV2) Set(ParameterNumber byte,) error {
	c.ParameterNumber = ParameterNumber

	return nil
}

func (c *ConfigurationGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ConfigurationV2),
		0x05,
		c.ParameterNumber,
		0x25,
	}
}
type ConfigurationReportV2 struct {
    *report
    ParameterNumber byte
    Level byte
    ConfigurationValue byte
    data []byte
}

func NewConfigurationReportV2(data []byte) *ConfigurationReportV2 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationReportV2{
				ParameterNumber: data[0],
				Level: data[1],
				ConfigurationValue: data[2],
        data: data,
    }
}

type ConfigurationSetV2 struct {
	node byte
	ParameterNumber byte
	Level byte
	ConfigurationValue byte
}

func NewConfigurationSetV2() ConfigurationSetV2 {
	return ConfigurationSetV2{}
}

func (c *ConfigurationSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationSetV2) Set(ParameterNumber byte,Level byte,ConfigurationValue byte,) error {
	c.ParameterNumber = ParameterNumber
	c.Level = Level
	c.ConfigurationValue = ConfigurationValue

	return nil
}

func (c *ConfigurationSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ConfigurationV2),
		0x04,
		c.ParameterNumber,
		c.Level,
		c.ConfigurationValue,
		0x25,
	}
}
