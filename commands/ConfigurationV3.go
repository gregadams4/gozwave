package commands
import "encoding/binary"
type ConfigurationBulkGetV3 struct {
	node byte
	ParameterOffset uint16
	NumberofParameters byte
}

func NewConfigurationBulkGetV3() ConfigurationBulkGetV3 {
	return ConfigurationBulkGetV3{}
}

func (c *ConfigurationBulkGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationBulkGetV3) Set(ParameterOffset uint16,NumberofParameters byte,) error {
	c.ParameterOffset = ParameterOffset
	c.NumberofParameters = NumberofParameters

	return nil
}

func (c *ConfigurationBulkGetV3) Encode() []byte {
	ParameterOffsetBytes := []byte{}
	binary.BigEndian.PutUint16(ParameterOffsetBytes, c.ParameterOffset)
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ConfigurationV3),
		0x08,
		ParameterOffsetBytes[0],
		ParameterOffsetBytes[1],
		c.NumberofParameters,
		0x25,
	}
}
type ConfigurationBulkReportV3 struct {
    *report
    ParameterOffset uint16
    NumberofParameters byte
    Reportstofollow byte
    Properties1 byte
    data []byte
}

func NewConfigurationBulkReportV3(data []byte) *ConfigurationBulkReportV3 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationBulkReportV3{
				ParameterOffset: binary.BigEndian.Uint16(data[0:2]),
				NumberofParameters: data[2],
				Reportstofollow: data[3],
				Properties1: data[4],
        data: data,
    }
}

type ConfigurationBulkSetV3 struct {
	node byte
	ParameterOffset uint16
	NumberofParameters byte
	Properties1 byte
	Vg []ConfigurationBulkSetV3vg
}
type ConfigurationBulkSetV3vg struct {
	Parameter byte
}

func NewConfigurationBulkSetV3() ConfigurationBulkSetV3 {
	return ConfigurationBulkSetV3{}
}

func (c *ConfigurationBulkSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationBulkSetV3) Set(ParameterOffset uint16,NumberofParameters byte,Properties1 byte,Vg []ConfigurationBulkSetV3vg,) error {
	c.ParameterOffset = ParameterOffset
	c.NumberofParameters = NumberofParameters
	c.Properties1 = Properties1
	c.Vg = Vg

	return nil
}

func (c *ConfigurationBulkSetV3) Encode() []byte {
	ParameterOffsetBytes := []byte{}
	binary.BigEndian.PutUint16(ParameterOffsetBytes, c.ParameterOffset)
	var data = []byte{
		0x13,
		c.node,
		byte(6),
		byte(ConfigurationV3),
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
type ConfigurationGetV3 struct {
	node byte
	ParameterNumber byte
}

func NewConfigurationGetV3() ConfigurationGetV3 {
	return ConfigurationGetV3{}
}

func (c *ConfigurationGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationGetV3) Set(ParameterNumber byte,) error {
	c.ParameterNumber = ParameterNumber

	return nil
}

func (c *ConfigurationGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ConfigurationV3),
		0x05,
		c.ParameterNumber,
		0x25,
	}
}
type ConfigurationReportV3 struct {
    *report
    ParameterNumber byte
    Level byte
    ConfigurationValue byte
    data []byte
}

func NewConfigurationReportV3(data []byte) *ConfigurationReportV3 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationReportV3{
				ParameterNumber: data[0],
				Level: data[1],
				ConfigurationValue: data[2],
        data: data,
    }
}

type ConfigurationSetV3 struct {
	node byte
	ParameterNumber byte
	Level byte
	ConfigurationValue byte
}

func NewConfigurationSetV3() ConfigurationSetV3 {
	return ConfigurationSetV3{}
}

func (c *ConfigurationSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationSetV3) Set(ParameterNumber byte,Level byte,ConfigurationValue byte,) error {
	c.ParameterNumber = ParameterNumber
	c.Level = Level
	c.ConfigurationValue = ConfigurationValue

	return nil
}

func (c *ConfigurationSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ConfigurationV3),
		0x04,
		c.ParameterNumber,
		c.Level,
		c.ConfigurationValue,
		0x25,
	}
}
type ConfigurationNameGetV3 struct {
	node byte
	ParameterNumber uint16
}

func NewConfigurationNameGetV3() ConfigurationNameGetV3 {
	return ConfigurationNameGetV3{}
}

func (c *ConfigurationNameGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationNameGetV3) Set(ParameterNumber uint16,) error {
	c.ParameterNumber = ParameterNumber

	return nil
}

func (c *ConfigurationNameGetV3) Encode() []byte {
	ParameterNumberBytes := []byte{}
	binary.BigEndian.PutUint16(ParameterNumberBytes, c.ParameterNumber)
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ConfigurationV3),
		0x0A,
		ParameterNumberBytes[0],
		ParameterNumberBytes[1],
		0x25,
	}
}
type ConfigurationNameReportV3 struct {
    *report
    ParameterNumber uint16
    Reportstofollow byte
    Name byte
    data []byte
}

func NewConfigurationNameReportV3(data []byte) *ConfigurationNameReportV3 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationNameReportV3{
				ParameterNumber: binary.BigEndian.Uint16(data[0:2]),
				Reportstofollow: data[2],
				Name: data[3],
        data: data,
    }
}

type ConfigurationInfoGetV3 struct {
	node byte
	ParameterNumber uint16
}

func NewConfigurationInfoGetV3() ConfigurationInfoGetV3 {
	return ConfigurationInfoGetV3{}
}

func (c *ConfigurationInfoGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationInfoGetV3) Set(ParameterNumber uint16,) error {
	c.ParameterNumber = ParameterNumber

	return nil
}

func (c *ConfigurationInfoGetV3) Encode() []byte {
	ParameterNumberBytes := []byte{}
	binary.BigEndian.PutUint16(ParameterNumberBytes, c.ParameterNumber)
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ConfigurationV3),
		0x0C,
		ParameterNumberBytes[0],
		ParameterNumberBytes[1],
		0x25,
	}
}
type ConfigurationInfoReportV3 struct {
    *report
    ParameterNumber uint16
    Reportstofollow byte
    Info byte
    data []byte
}

func NewConfigurationInfoReportV3(data []byte) *ConfigurationInfoReportV3 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationInfoReportV3{
				ParameterNumber: binary.BigEndian.Uint16(data[0:2]),
				Reportstofollow: data[2],
				Info: data[3],
        data: data,
    }
}

type ConfigurationPropertiesGetV3 struct {
	node byte
	ParameterNumber uint16
}

func NewConfigurationPropertiesGetV3() ConfigurationPropertiesGetV3 {
	return ConfigurationPropertiesGetV3{}
}

func (c *ConfigurationPropertiesGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationPropertiesGetV3) Set(ParameterNumber uint16,) error {
	c.ParameterNumber = ParameterNumber

	return nil
}

func (c *ConfigurationPropertiesGetV3) Encode() []byte {
	ParameterNumberBytes := []byte{}
	binary.BigEndian.PutUint16(ParameterNumberBytes, c.ParameterNumber)
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ConfigurationV3),
		0x0E,
		ParameterNumberBytes[0],
		ParameterNumberBytes[1],
		0x25,
	}
}
type ConfigurationPropertiesReportV3 struct {
    *report
    ParameterNumber uint16
    Properties1 byte
    MinValue byte
    MaxValue byte
    DefaultValue byte
    NextParameterNumber uint16
    data []byte
}

func NewConfigurationPropertiesReportV3(data []byte) *ConfigurationPropertiesReportV3 {
    if len(data) < 8 {
				//may want to change this to just return nil
				for i := len(data); i < 8; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationPropertiesReportV3{
				ParameterNumber: binary.BigEndian.Uint16(data[0:2]),
				Properties1: data[2],
				MinValue: data[3],
				MaxValue: data[4],
				DefaultValue: data[5],
				NextParameterNumber: binary.BigEndian.Uint16(data[6:8]),
        data: data,
    }
}

