package commands
type WakeUpIntervalGet struct {
	node byte
}

func NewWakeUpIntervalGet() WakeUpIntervalGet {
	return WakeUpIntervalGet{}
}

func (c *WakeUpIntervalGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *WakeUpIntervalGet) Set() error {

	return nil
}

func (c *WakeUpIntervalGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(WakeUp),
		0x05,
		0x25,
	}
}
type WakeUpIntervalReport struct {
    *report
    Seconds byte
    NodeID byte
    data []byte
}

func NewWakeUpIntervalReport(data []byte) *WakeUpIntervalReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &WakeUpIntervalReport{
				Seconds: data[0],
				NodeID: data[1],
        data: data,
    }
}

type WakeUpIntervalSet struct {
	node byte
	Seconds byte
	NodeID byte
}

func NewWakeUpIntervalSet() WakeUpIntervalSet {
	return WakeUpIntervalSet{}
}

func (c *WakeUpIntervalSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *WakeUpIntervalSet) Set(Seconds byte,NodeID byte,) error {
	c.Seconds = Seconds
	c.NodeID = NodeID

	return nil
}

func (c *WakeUpIntervalSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(WakeUp),
		0x04,
		c.Seconds,
		c.NodeID,
		0x25,
	}
}
type WakeUpNoMoreInformation struct {
	node byte
}

func NewWakeUpNoMoreInformation() WakeUpNoMoreInformation {
	return WakeUpNoMoreInformation{}
}

func (c *WakeUpNoMoreInformation) SetNode(node int) {
	c.node = byte(node)
}

func (c *WakeUpNoMoreInformation) Set() error {

	return nil
}

func (c *WakeUpNoMoreInformation) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(WakeUp),
		0x08,
		0x25,
	}
}
type WakeUpNotification struct {
    *report
    data []byte
}

func NewWakeUpNotification(data []byte) *WakeUpNotification {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &WakeUpNotification{
        data: data,
    }
}

