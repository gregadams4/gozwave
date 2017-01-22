package commands
import "encoding/binary"
type ConfigurationBulkGetV4 struct {
	node byte
	ParameterOffset uint16
	NumberofParameters byte
}

func NewConfigurationBulkGetV4() ConfigurationBulkGetV4 {
	return ConfigurationBulkGetV4{}
}

func (c *ConfigurationBulkGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationBulkGetV4) Set(ParameterOffset uint16,NumberofParameters byte,) error {
	c.ParameterOffset = ParameterOffset
	c.NumberofParameters = NumberofParameters

	return nil
}

func (c *ConfigurationBulkGetV4) Encode() []byte {
	ParameterOffsetBytes := []byte{}
	binary.BigEndian.PutUint16(ParameterOffsetBytes, c.ParameterOffset)
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ConfigurationV4),
		0x08,
		ParameterOffsetBytes[0],
		ParameterOffsetBytes[1],
		c.NumberofParameters,
		0x25,
	}
}
type ConfigurationBulkReportV4 struct {
    *report
    ParameterOffset uint16
    NumberofParameters byte
    Reportstofollow byte
    Properties1 byte
    data []byte
}

func NewConfigurationBulkReportV4(data []byte) *ConfigurationBulkReportV4 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationBulkReportV4{
				ParameterOffset: binary.BigEndian.Uint16(data[0:2]),
				NumberofParameters: data[2],
				Reportstofollow: data[3],
				Properties1: data[4],
        data: data,
    }
}

type ConfigurationBulkSetV4 struct {
	node byte
	ParameterOffset uint16
	NumberofParameters byte
	Properties1 byte
	Vg []ConfigurationBulkSetV4vg
}
type ConfigurationBulkSetV4vg struct {
	Parameter byte
}

func NewConfigurationBulkSetV4() ConfigurationBulkSetV4 {
	return ConfigurationBulkSetV4{}
}

func (c *ConfigurationBulkSetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationBulkSetV4) Set(ParameterOffset uint16,NumberofParameters byte,Properties1 byte,Vg []ConfigurationBulkSetV4vg,) error {
	c.ParameterOffset = ParameterOffset
	c.NumberofParameters = NumberofParameters
	c.Properties1 = Properties1
	c.Vg = Vg

	return nil
}

func (c *ConfigurationBulkSetV4) Encode() []byte {
	ParameterOffsetBytes := []byte{}
	binary.BigEndian.PutUint16(ParameterOffsetBytes, c.ParameterOffset)
	var data = []byte{
		0x13,
		c.node,
		byte(6),
		byte(ConfigurationV4),
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
type ConfigurationGetV4 struct {
	node byte
	ParameterNumber byte
}

func NewConfigurationGetV4() ConfigurationGetV4 {
	return ConfigurationGetV4{}
}

func (c *ConfigurationGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationGetV4) Set(ParameterNumber byte,) error {
	c.ParameterNumber = ParameterNumber

	return nil
}

func (c *ConfigurationGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ConfigurationV4),
		0x05,
		c.ParameterNumber,
		0x25,
	}
}
type ConfigurationReportV4 struct {
    *report
    ParameterNumber byte
    Level byte
    ConfigurationValue byte
    data []byte
}

func NewConfigurationReportV4(data []byte) *ConfigurationReportV4 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationReportV4{
				ParameterNumber: data[0],
				Level: data[1],
				ConfigurationValue: data[2],
        data: data,
    }
}

type ConfigurationSetV4 struct {
	node byte
	ParameterNumber byte
	Level byte
	ConfigurationValue byte
}

func NewConfigurationSetV4() ConfigurationSetV4 {
	return ConfigurationSetV4{}
}

func (c *ConfigurationSetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationSetV4) Set(ParameterNumber byte,Level byte,ConfigurationValue byte,) error {
	c.ParameterNumber = ParameterNumber
	c.Level = Level
	c.ConfigurationValue = ConfigurationValue

	return nil
}

func (c *ConfigurationSetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(ConfigurationV4),
		0x04,
		c.ParameterNumber,
		c.Level,
		c.ConfigurationValue,
		0x25,
	}
}
type ConfigurationNameGetV4 struct {
	node byte
	ParameterNumber uint16
}

func NewConfigurationNameGetV4() ConfigurationNameGetV4 {
	return ConfigurationNameGetV4{}
}

func (c *ConfigurationNameGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationNameGetV4) Set(ParameterNumber uint16,) error {
	c.ParameterNumber = ParameterNumber

	return nil
}

func (c *ConfigurationNameGetV4) Encode() []byte {
	ParameterNumberBytes := []byte{}
	binary.BigEndian.PutUint16(ParameterNumberBytes, c.ParameterNumber)
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ConfigurationV4),
		0x0A,
		ParameterNumberBytes[0],
		ParameterNumberBytes[1],
		0x25,
	}
}
type ConfigurationNameReportV4 struct {
    *report
    ParameterNumber uint16
    Reportstofollow byte
    Name byte
    data []byte
}

func NewConfigurationNameReportV4(data []byte) *ConfigurationNameReportV4 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationNameReportV4{
				ParameterNumber: binary.BigEndian.Uint16(data[0:2]),
				Reportstofollow: data[2],
				Name: data[3],
        data: data,
    }
}

type ConfigurationInfoGetV4 struct {
	node byte
	ParameterNumber uint16
}

func NewConfigurationInfoGetV4() ConfigurationInfoGetV4 {
	return ConfigurationInfoGetV4{}
}

func (c *ConfigurationInfoGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationInfoGetV4) Set(ParameterNumber uint16,) error {
	c.ParameterNumber = ParameterNumber

	return nil
}

func (c *ConfigurationInfoGetV4) Encode() []byte {
	ParameterNumberBytes := []byte{}
	binary.BigEndian.PutUint16(ParameterNumberBytes, c.ParameterNumber)
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ConfigurationV4),
		0x0C,
		ParameterNumberBytes[0],
		ParameterNumberBytes[1],
		0x25,
	}
}
type ConfigurationInfoReportV4 struct {
    *report
    ParameterNumber uint16
    Reportstofollow byte
    Info byte
    data []byte
}

func NewConfigurationInfoReportV4(data []byte) *ConfigurationInfoReportV4 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationInfoReportV4{
				ParameterNumber: binary.BigEndian.Uint16(data[0:2]),
				Reportstofollow: data[2],
				Info: data[3],
        data: data,
    }
}

type ConfigurationPropertiesGetV4 struct {
	node byte
	ParameterNumber uint16
}

func NewConfigurationPropertiesGetV4() ConfigurationPropertiesGetV4 {
	return ConfigurationPropertiesGetV4{}
}

func (c *ConfigurationPropertiesGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationPropertiesGetV4) Set(ParameterNumber uint16,) error {
	c.ParameterNumber = ParameterNumber

	return nil
}

func (c *ConfigurationPropertiesGetV4) Encode() []byte {
	ParameterNumberBytes := []byte{}
	binary.BigEndian.PutUint16(ParameterNumberBytes, c.ParameterNumber)
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ConfigurationV4),
		0x0E,
		ParameterNumberBytes[0],
		ParameterNumberBytes[1],
		0x25,
	}
}
type ConfigurationPropertiesReportV4 struct {
    *report
    ParameterNumber uint16
    Properties1 byte
    MinValue byte
    MaxValue byte
    DefaultValue byte
    NextParameterNumber uint16
    Properties2 byte
    data []byte
}

func NewConfigurationPropertiesReportV4(data []byte) *ConfigurationPropertiesReportV4 {
    if len(data) < 9 {
				//may want to change this to just return nil
				for i := len(data); i < 9; i++ {
            data = append(data, 0x00)
        }
    }

    return &ConfigurationPropertiesReportV4{
				ParameterNumber: binary.BigEndian.Uint16(data[0:2]),
				Properties1: data[2],
				MinValue: data[3],
				MaxValue: data[4],
				DefaultValue: data[5],
				NextParameterNumber: binary.BigEndian.Uint16(data[6:8]),
				Properties2: data[8],
        data: data,
    }
}

type ConfigurationDefaultResetV4 struct {
	node byte
}

func NewConfigurationDefaultResetV4() ConfigurationDefaultResetV4 {
	return ConfigurationDefaultResetV4{}
}

func (c *ConfigurationDefaultResetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *ConfigurationDefaultResetV4) Set() error {

	return nil
}

func (c *ConfigurationDefaultResetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ConfigurationV4),
		0x01,
		0x25,
	}
}
