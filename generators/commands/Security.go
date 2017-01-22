package commands
type NetworkKeySet struct {
	node byte
	NetworkKeybyte byte
}

func NewNetworkKeySet() NetworkKeySet {
	return NetworkKeySet{}
}

func (c *NetworkKeySet) SetNode(node int) {
	c.node = byte(node)
}

func (c *NetworkKeySet) Set(NetworkKeybyte byte,) error {
	c.NetworkKeybyte = NetworkKeybyte

	return nil
}

func (c *NetworkKeySet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Security),
		0x06,
		c.NetworkKeybyte,
		0x25,
	}
}
type NetworkKeyVerify struct {
	node byte
}

func NewNetworkKeyVerify() NetworkKeyVerify {
	return NetworkKeyVerify{}
}

func (c *NetworkKeyVerify) SetNode(node int) {
	c.node = byte(node)
}

func (c *NetworkKeyVerify) Set() error {

	return nil
}

func (c *NetworkKeyVerify) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Security),
		0x07,
		0x25,
	}
}
type SecurityCommandsSupportedGet struct {
	node byte
}

func NewSecurityCommandsSupportedGet() SecurityCommandsSupportedGet {
	return SecurityCommandsSupportedGet{}
}

func (c *SecurityCommandsSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityCommandsSupportedGet) Set() error {

	return nil
}

func (c *SecurityCommandsSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Security),
		0x02,
		0x25,
	}
}
type SecurityCommandsSupportedReport struct {
    *report
    Reportstofollow byte
    CommandClasssupport byte
    COMMAND_CLASS_MARK byte
    CommandClasscontrol byte
    data []byte
}

func NewSecurityCommandsSupportedReport(data []byte) *SecurityCommandsSupportedReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &SecurityCommandsSupportedReport{
				Reportstofollow: data[0],
				CommandClasssupport: data[1],
				COMMAND_CLASS_MARK: data[2],
				CommandClasscontrol: data[3],
        data: data,
    }
}

type SecurityMessageEncapsulation struct {
	node byte
	InitializationVectorbyte byte
	Properties1 byte
	Commandbyte byte
	ReceiversnonceIdentifier byte
	MessageAuthenticationCodebyte byte
}

func NewSecurityMessageEncapsulation() SecurityMessageEncapsulation {
	return SecurityMessageEncapsulation{}
}

func (c *SecurityMessageEncapsulation) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityMessageEncapsulation) Set(InitializationVectorbyte byte,Properties1 byte,Commandbyte byte,ReceiversnonceIdentifier byte,MessageAuthenticationCodebyte byte,) error {
	c.InitializationVectorbyte = InitializationVectorbyte
	c.Properties1 = Properties1
	c.Commandbyte = Commandbyte
	c.ReceiversnonceIdentifier = ReceiversnonceIdentifier
	c.MessageAuthenticationCodebyte = MessageAuthenticationCodebyte

	return nil
}

func (c *SecurityMessageEncapsulation) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(Security),
		0x81,
		c.InitializationVectorbyte,
		c.Properties1,
		c.Commandbyte,
		c.ReceiversnonceIdentifier,
		c.MessageAuthenticationCodebyte,
		0x25,
	}
}
type SecurityMessageEncapsulationNonceGet struct {
	node byte
	InitializationVectorbyte byte
	Properties1 byte
	Commandbyte byte
	ReceiversnonceIdentifier byte
	MessageAuthenticationCodebyte byte
}

func NewSecurityMessageEncapsulationNonceGet() SecurityMessageEncapsulationNonceGet {
	return SecurityMessageEncapsulationNonceGet{}
}

func (c *SecurityMessageEncapsulationNonceGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityMessageEncapsulationNonceGet) Set(InitializationVectorbyte byte,Properties1 byte,Commandbyte byte,ReceiversnonceIdentifier byte,MessageAuthenticationCodebyte byte,) error {
	c.InitializationVectorbyte = InitializationVectorbyte
	c.Properties1 = Properties1
	c.Commandbyte = Commandbyte
	c.ReceiversnonceIdentifier = ReceiversnonceIdentifier
	c.MessageAuthenticationCodebyte = MessageAuthenticationCodebyte

	return nil
}

func (c *SecurityMessageEncapsulationNonceGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(7),
		byte(Security),
		0xC1,
		c.InitializationVectorbyte,
		c.Properties1,
		c.Commandbyte,
		c.ReceiversnonceIdentifier,
		c.MessageAuthenticationCodebyte,
		0x25,
	}
}
type SecurityNonceGet struct {
	node byte
}

func NewSecurityNonceGet() SecurityNonceGet {
	return SecurityNonceGet{}
}

func (c *SecurityNonceGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecurityNonceGet) Set() error {

	return nil
}

func (c *SecurityNonceGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Security),
		0x40,
		0x25,
	}
}
type SecurityNonceReport struct {
    *report
    Noncebyte byte
    data []byte
}

func NewSecurityNonceReport(data []byte) *SecurityNonceReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SecurityNonceReport{
				Noncebyte: data[0],
        data: data,
    }
}

type SecuritySchemeGet struct {
	node byte
	SupportedSecuritySchemes byte
}

func NewSecuritySchemeGet() SecuritySchemeGet {
	return SecuritySchemeGet{}
}

func (c *SecuritySchemeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecuritySchemeGet) Set(SupportedSecuritySchemes byte,) error {
	c.SupportedSecuritySchemes = SupportedSecuritySchemes

	return nil
}

func (c *SecuritySchemeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Security),
		0x04,
		c.SupportedSecuritySchemes,
		0x25,
	}
}
type SecuritySchemeInherit struct {
	node byte
	SupportedSecuritySchemes byte
}

func NewSecuritySchemeInherit() SecuritySchemeInherit {
	return SecuritySchemeInherit{}
}

func (c *SecuritySchemeInherit) SetNode(node int) {
	c.node = byte(node)
}

func (c *SecuritySchemeInherit) Set(SupportedSecuritySchemes byte,) error {
	c.SupportedSecuritySchemes = SupportedSecuritySchemes

	return nil
}

func (c *SecuritySchemeInherit) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Security),
		0x08,
		c.SupportedSecuritySchemes,
		0x25,
	}
}
type SecuritySchemeReport struct {
    *report
    SupportedSecuritySchemes byte
    data []byte
}

func NewSecuritySchemeReport(data []byte) *SecuritySchemeReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &SecuritySchemeReport{
				SupportedSecuritySchemes: data[0],
        data: data,
    }
}

