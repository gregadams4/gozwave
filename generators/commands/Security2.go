package commands
type Security2NonceGet struct {
	node byte
	SequenceNumber byte
}

func NewSecurity2NonceGet() Security2NonceGet {
	return Security2NonceGet{}
}

func (c *Security2NonceGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *Security2NonceGet) Set(SequenceNumber byte,) error {
	c.SequenceNumber = SequenceNumber

	return nil
}

func (c *Security2NonceGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Security2),
		0x01,
		c.SequenceNumber,
		0x25,
	}
}
type Security2NonceReport struct {
    *report
    SequenceNumber byte
    Properties1 byte
    ReceiversEntropyInput byte
    data []byte
}

func NewSecurity2NonceReport(data []byte) *Security2NonceReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &Security2NonceReport{
				SequenceNumber: data[0],
				Properties1: data[1],
				ReceiversEntropyInput: data[2],
        data: data,
    }
}

type Security2MessageEncapsulation struct {
	node byte
	SequenceNumber byte
	Properties1 byte
	CCMCiphertextObject byte
	Vg1 []Security2MessageEncapsulationvg1
}
type Security2MessageEncapsulationvg1 struct {
	ExtensionLength byte
	Properties1 byte
	Extension byte
}

func NewSecurity2MessageEncapsulation() Security2MessageEncapsulation {
	return Security2MessageEncapsulation{}
}

func (c *Security2MessageEncapsulation) SetNode(node int) {
	c.node = byte(node)
}

func (c *Security2MessageEncapsulation) Set(SequenceNumber byte,Properties1 byte,CCMCiphertextObject byte,Vg1 []Security2MessageEncapsulationvg1,) error {
	c.SequenceNumber = SequenceNumber
	c.Properties1 = Properties1
	c.CCMCiphertextObject = CCMCiphertextObject
	c.Vg1 = Vg1

	return nil
}

func (c *Security2MessageEncapsulation) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(8),
		byte(Security2),
		0x03,
		c.SequenceNumber,
		c.Properties1,
		c.CCMCiphertextObject,
	}
	for _, v := range c.Vg1 {
		data = append(data, v.ExtensionLength)
		data = append(data, v.Properties1)
		data = append(data, v.Extension)
	}
	data = append(data, 0x25)
	return data
	
}
type KexGet struct {
	node byte
}

func NewKexGet() KexGet {
	return KexGet{}
}

func (c *KexGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *KexGet) Set() error {

	return nil
}

func (c *KexGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Security2),
		0x04,
		0x25,
	}
}
type KexReport struct {
    *report
    Properties1 byte
    SupportedKEXSchemes byte
    SupportedECDHProfiles byte
    RequestedKeys byte
    data []byte
}

func NewKexReport(data []byte) *KexReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &KexReport{
				Properties1: data[0],
				SupportedKEXSchemes: data[1],
				SupportedECDHProfiles: data[2],
				RequestedKeys: data[3],
        data: data,
    }
}

type KexSet struct {
	node byte
	Properties1 byte
	SelectedKEXScheme byte
	SelectedECDHProfile byte
	GrantedKeys byte
}

func NewKexSet() KexSet {
	return KexSet{}
}

func (c *KexSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *KexSet) Set(Properties1 byte,SelectedKEXScheme byte,SelectedECDHProfile byte,GrantedKeys byte,) error {
	c.Properties1 = Properties1
	c.SelectedKEXScheme = SelectedKEXScheme
	c.SelectedECDHProfile = SelectedECDHProfile
	c.GrantedKeys = GrantedKeys

	return nil
}

func (c *KexSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(Security2),
		0x06,
		c.Properties1,
		c.SelectedKEXScheme,
		c.SelectedECDHProfile,
		c.GrantedKeys,
		0x25,
	}
}
type KexFail struct {
	node byte
	KEXFailType byte
}

func NewKexFail() KexFail {
	return KexFail{}
}

func (c *KexFail) SetNode(node int) {
	c.node = byte(node)
}

func (c *KexFail) Set(KEXFailType byte,) error {
	c.KEXFailType = KEXFailType

	return nil
}

func (c *KexFail) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Security2),
		0x07,
		c.KEXFailType,
		0x25,
	}
}
type PublicKeyReport struct {
    *report
    Properties1 byte
    ECDHPublicKey byte
    data []byte
}

func NewPublicKeyReport(data []byte) *PublicKeyReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &PublicKeyReport{
				Properties1: data[0],
				ECDHPublicKey: data[1],
        data: data,
    }
}

type Security2NetworkKeyGet struct {
	node byte
	RequestedKey byte
}

func NewSecurity2NetworkKeyGet() Security2NetworkKeyGet {
	return Security2NetworkKeyGet{}
}

func (c *Security2NetworkKeyGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *Security2NetworkKeyGet) Set(RequestedKey byte,) error {
	c.RequestedKey = RequestedKey

	return nil
}

func (c *Security2NetworkKeyGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Security2),
		0x09,
		c.RequestedKey,
		0x25,
	}
}
type Security2NetworkKeyReport struct {
    *report
    GrantedKey byte
    NetworkKey byte
    data []byte
}

func NewSecurity2NetworkKeyReport(data []byte) *Security2NetworkKeyReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &Security2NetworkKeyReport{
				GrantedKey: data[0],
				NetworkKey: data[1],
        data: data,
    }
}

type Security2NetworkKeyVerify struct {
	node byte
}

func NewSecurity2NetworkKeyVerify() Security2NetworkKeyVerify {
	return Security2NetworkKeyVerify{}
}

func (c *Security2NetworkKeyVerify) SetNode(node int) {
	c.node = byte(node)
}

func (c *Security2NetworkKeyVerify) Set() error {

	return nil
}

func (c *Security2NetworkKeyVerify) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Security2),
		0x0B,
		0x25,
	}
}
type Security2TransferEnd struct {
	node byte
	Properties1 byte
}

func NewSecurity2TransferEnd() Security2TransferEnd {
	return Security2TransferEnd{}
}

func (c *Security2TransferEnd) SetNode(node int) {
	c.node = byte(node)
}

func (c *Security2TransferEnd) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *Security2TransferEnd) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Security2),
		0x0C,
		c.Properties1,
		0x25,
	}
}
type Security2CommandsSupportedGet struct {
	node byte
}

func NewSecurity2CommandsSupportedGet() Security2CommandsSupportedGet {
	return Security2CommandsSupportedGet{}
}

func (c *Security2CommandsSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *Security2CommandsSupportedGet) Set() error {

	return nil
}

func (c *Security2CommandsSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Security2),
		0x0D,
		0x25,
	}
}
type Security2CommandsSupportedReport struct {
    *report
    CommandClass byte
    data []byte
}

func NewSecurity2CommandsSupportedReport(data []byte) *Security2CommandsSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &Security2CommandsSupportedReport{
				CommandClass: data[0],
        data: data,
    }
}

type Security2CapabilitiesGet struct {
	node byte
}

func NewSecurity2CapabilitiesGet() Security2CapabilitiesGet {
	return Security2CapabilitiesGet{}
}

func (c *Security2CapabilitiesGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *Security2CapabilitiesGet) Set() error {

	return nil
}

func (c *Security2CapabilitiesGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Security2),
		0x0F,
		0x25,
	}
}
type Security2CapabilitiesReport struct {
    *report
    SupportedSPANs byte
    SupportedMPANs byte
    data []byte
}

func NewSecurity2CapabilitiesReport(data []byte) *Security2CapabilitiesReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &Security2CapabilitiesReport{
				SupportedSPANs: data[0],
				SupportedMPANs: data[1],
        data: data,
    }
}

