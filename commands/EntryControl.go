package commands
type EntryControlNotification struct {
    *report
    SequenceNumber byte
    Properties1 byte
    EventType byte
    EventDataLength byte
    EventData byte
    data []byte
}

func NewEntryControlNotification(data []byte) *EntryControlNotification {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &EntryControlNotification{
				SequenceNumber: data[0],
				Properties1: data[1],
				EventType: data[2],
				EventDataLength: data[3],
				EventData: data[4],
        data: data,
    }
}

type EntryControlKeySupportedGet struct {
	node byte
}

func NewEntryControlKeySupportedGet() EntryControlKeySupportedGet {
	return EntryControlKeySupportedGet{}
}

func (c *EntryControlKeySupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *EntryControlKeySupportedGet) Set() error {

	return nil
}

func (c *EntryControlKeySupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(EntryControl),
		0x02,
		0x25,
	}
}
type EntryControlKeySupportedReport struct {
    *report
    KeySupportedBitMaskLength byte
    KeySupportedBitMask byte
    data []byte
}

func NewEntryControlKeySupportedReport(data []byte) *EntryControlKeySupportedReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &EntryControlKeySupportedReport{
				KeySupportedBitMaskLength: data[0],
				KeySupportedBitMask: data[1],
        data: data,
    }
}

type EntryControlEventSupportedGet struct {
	node byte
}

func NewEntryControlEventSupportedGet() EntryControlEventSupportedGet {
	return EntryControlEventSupportedGet{}
}

func (c *EntryControlEventSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *EntryControlEventSupportedGet) Set() error {

	return nil
}

func (c *EntryControlEventSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(EntryControl),
		0x04,
		0x25,
	}
}
type EntryControlEventSupportedReport struct {
    *report
    Properties1 byte
    DataTypeSupportedBitMask byte
    Properties2 byte
    EventTypeSupportedBitMask byte
    KeyCachedSizesupportedMinimum byte
    KeyCachedSizesupportedMaximum byte
    KeyCachedTimeoutsupportedMinimum byte
    KeyCachedTimeoutsupportedMaximum byte
    data []byte
}

func NewEntryControlEventSupportedReport(data []byte) *EntryControlEventSupportedReport {
    if len(data) < 8 {
				//may want to change this to just return nil
				for i := len(data); i < 8; i++ {
            data = append(data, 0x00)
        }
    }

    return &EntryControlEventSupportedReport{
				Properties1: data[0],
				DataTypeSupportedBitMask: data[1],
				Properties2: data[2],
				EventTypeSupportedBitMask: data[3],
				KeyCachedSizesupportedMinimum: data[4],
				KeyCachedSizesupportedMaximum: data[5],
				KeyCachedTimeoutsupportedMinimum: data[6],
				KeyCachedTimeoutsupportedMaximum: data[7],
        data: data,
    }
}

type EntryControlConfigurationSet struct {
	node byte
	KeyCacheSize byte
	KeyCacheTimeout byte
}

func NewEntryControlConfigurationSet() EntryControlConfigurationSet {
	return EntryControlConfigurationSet{}
}

func (c *EntryControlConfigurationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *EntryControlConfigurationSet) Set(KeyCacheSize byte,KeyCacheTimeout byte,) error {
	c.KeyCacheSize = KeyCacheSize
	c.KeyCacheTimeout = KeyCacheTimeout

	return nil
}

func (c *EntryControlConfigurationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(EntryControl),
		0x06,
		c.KeyCacheSize,
		c.KeyCacheTimeout,
		0x25,
	}
}
type EntryControlConfigurationGet struct {
	node byte
}

func NewEntryControlConfigurationGet() EntryControlConfigurationGet {
	return EntryControlConfigurationGet{}
}

func (c *EntryControlConfigurationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *EntryControlConfigurationGet) Set() error {

	return nil
}

func (c *EntryControlConfigurationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(EntryControl),
		0x07,
		0x25,
	}
}
type EntryControlConfigurationReport struct {
    *report
    KeyCacheSize byte
    KeyCacheTimeout byte
    data []byte
}

func NewEntryControlConfigurationReport(data []byte) *EntryControlConfigurationReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &EntryControlConfigurationReport{
				KeyCacheSize: data[0],
				KeyCacheTimeout: data[1],
        data: data,
    }
}

