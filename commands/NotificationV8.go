package commands
type NotificationGetV8 struct {
    *report
    V1AlarmType byte
    NotificationType byte
    Event byte
    data []byte
}

func NewNotificationGetV8(data []byte) *NotificationGetV8 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationGetV8{
				V1AlarmType: data[0],
				NotificationType: data[1],
				Event: data[2],
        data: data,
    }
}

type NotificationReportV8 struct {
    *report
    V1AlarmType byte
    V1AlarmLevel byte
    Reserved byte
    NotificationStatus byte
    NotificationType byte
    Event byte
    Properties1 byte
    EventParameter byte
    SequenceNumber byte
    data []byte
}

func NewNotificationReportV8(data []byte) *NotificationReportV8 {
    if len(data) < 9 {
				//may want to change this to just return nil
				for i := len(data); i < 9; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationReportV8{
				V1AlarmType: data[0],
				V1AlarmLevel: data[1],
				Reserved: data[2],
				NotificationStatus: data[3],
				NotificationType: data[4],
				Event: data[5],
				Properties1: data[6],
				EventParameter: data[7],
				SequenceNumber: data[8],
        data: data,
    }
}

type NotificationSetV8 struct {
    *report
    NotificationType byte
    NotificationStatus byte
    data []byte
}

func NewNotificationSetV8(data []byte) *NotificationSetV8 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSetV8{
				NotificationType: data[0],
				NotificationStatus: data[1],
        data: data,
    }
}

type NotificationSupportedGetV8 struct {
    *report
    data []byte
}

func NewNotificationSupportedGetV8(data []byte) *NotificationSupportedGetV8 {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSupportedGetV8{
        data: data,
    }
}

type NotificationSupportedReportV8 struct {
    *report
    Properties1 byte
    BitMask byte
    data []byte
}

func NewNotificationSupportedReportV8(data []byte) *NotificationSupportedReportV8 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSupportedReportV8{
				Properties1: data[0],
				BitMask: data[1],
        data: data,
    }
}

type EventSupportedGetV8 struct {
	node byte
	NotificationType byte
}

func NewEventSupportedGetV8() EventSupportedGetV8 {
	return EventSupportedGetV8{}
}

func (c *EventSupportedGetV8) SetNode(node int) {
	c.node = byte(node)
}

func (c *EventSupportedGetV8) Set(NotificationType byte,) error {
	c.NotificationType = NotificationType

	return nil
}

func (c *EventSupportedGetV8) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NotificationV8),
		0x01,
		c.NotificationType,
		0x25,
	}
}
type EventSupportedReportV8 struct {
    *report
    NotificationType byte
    Properties1 byte
    BitMask byte
    data []byte
}

func NewEventSupportedReportV8(data []byte) *EventSupportedReportV8 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &EventSupportedReportV8{
				NotificationType: data[0],
				Properties1: data[1],
				BitMask: data[2],
        data: data,
    }
}

