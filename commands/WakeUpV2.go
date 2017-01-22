package commands
type WakeUpIntervalCapabilitiesGetV2 struct {
	node byte
}

func NewWakeUpIntervalCapabilitiesGetV2() WakeUpIntervalCapabilitiesGetV2 {
	return WakeUpIntervalCapabilitiesGetV2{}
}

func (c *WakeUpIntervalCapabilitiesGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *WakeUpIntervalCapabilitiesGetV2) Set() error {

	return nil
}

func (c *WakeUpIntervalCapabilitiesGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(WakeUpV2),
		0x09,
		0x25,
	}
}
type WakeUpIntervalCapabilitiesReportV2 struct {
    *report
    MinimumWakeUpIntervalSeconds byte
    MaximumWakeUpIntervalSeconds byte
    DefaultWakeUpIntervalSeconds byte
    WakeUpIntervalStepSeconds byte
    data []byte
}

func NewWakeUpIntervalCapabilitiesReportV2(data []byte) *WakeUpIntervalCapabilitiesReportV2 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &WakeUpIntervalCapabilitiesReportV2{
				MinimumWakeUpIntervalSeconds: data[0],
				MaximumWakeUpIntervalSeconds: data[1],
				DefaultWakeUpIntervalSeconds: data[2],
				WakeUpIntervalStepSeconds: data[3],
        data: data,
    }
}

type WakeUpIntervalGetV2 struct {
	node byte
}

func NewWakeUpIntervalGetV2() WakeUpIntervalGetV2 {
	return WakeUpIntervalGetV2{}
}

func (c *WakeUpIntervalGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *WakeUpIntervalGetV2) Set() error {

	return nil
}

func (c *WakeUpIntervalGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(WakeUpV2),
		0x05,
		0x25,
	}
}
type WakeUpIntervalReportV2 struct {
    *report
    Seconds byte
    NodeID byte
    data []byte
}

func NewWakeUpIntervalReportV2(data []byte) *WakeUpIntervalReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &WakeUpIntervalReportV2{
				Seconds: data[0],
				NodeID: data[1],
        data: data,
    }
}

type WakeUpIntervalSetV2 struct {
	node byte
	Seconds byte
	NodeID byte
}

func NewWakeUpIntervalSetV2() WakeUpIntervalSetV2 {
	return WakeUpIntervalSetV2{}
}

func (c *WakeUpIntervalSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *WakeUpIntervalSetV2) Set(Seconds byte,NodeID byte,) error {
	c.Seconds = Seconds
	c.NodeID = NodeID

	return nil
}

func (c *WakeUpIntervalSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(WakeUpV2),
		0x04,
		c.Seconds,
		c.NodeID,
		0x25,
	}
}
type WakeUpNoMoreInformationV2 struct {
	node byte
}

func NewWakeUpNoMoreInformationV2() WakeUpNoMoreInformationV2 {
	return WakeUpNoMoreInformationV2{}
}

func (c *WakeUpNoMoreInformationV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *WakeUpNoMoreInformationV2) Set() error {

	return nil
}

func (c *WakeUpNoMoreInformationV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(WakeUpV2),
		0x08,
		0x25,
	}
}
type WakeUpNotificationV2 struct {
    *report
    data []byte
}

func NewWakeUpNotificationV2(data []byte) *WakeUpNotificationV2 {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &WakeUpNotificationV2{
        data: data,
    }
}

