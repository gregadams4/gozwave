package commands
type NotificationGetV6 struct {
    *report
    V1AlarmType byte
    NotificationType byte
    Event byte
    data []byte
}

func NewNotificationGetV6(data []byte) *NotificationGetV6 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationGetV6{
				V1AlarmType: data[0],
				NotificationType: data[1],
				Event: data[2],
        data: data,
    }
}

type NotificationReportV6 struct {
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

func NewNotificationReportV6(data []byte) *NotificationReportV6 {
    if len(data) < 9 {
				//may want to change this to just return nil
				for i := len(data); i < 9; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationReportV6{
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

type NotificationSetV6 struct {
    *report
    NotificationType byte
    NotificationStatus byte
    data []byte
}

func NewNotificationSetV6(data []byte) *NotificationSetV6 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSetV6{
				NotificationType: data[0],
				NotificationStatus: data[1],
        data: data,
    }
}

type NotificationSupportedGetV6 struct {
    *report
    data []byte
}

func NewNotificationSupportedGetV6(data []byte) *NotificationSupportedGetV6 {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSupportedGetV6{
        data: data,
    }
}

type NotificationSupportedReportV6 struct {
    *report
    Properties1 byte
    BitMask byte
    data []byte
}

func NewNotificationSupportedReportV6(data []byte) *NotificationSupportedReportV6 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSupportedReportV6{
				Properties1: data[0],
				BitMask: data[1],
        data: data,
    }
}

type EventSupportedGetV6 struct {
	node byte
	NotificationType byte
}

func NewEventSupportedGetV6() EventSupportedGetV6 {
	return EventSupportedGetV6{}
}

func (c *EventSupportedGetV6) SetNode(node int) {
	c.node = byte(node)
}

func (c *EventSupportedGetV6) Set(NotificationType byte,) error {
	c.NotificationType = NotificationType

	return nil
}

func (c *EventSupportedGetV6) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NotificationV6),
		0x01,
		c.NotificationType,
		0x25,
	}
}
type EventSupportedReportV6 struct {
    *report
    NotificationType byte
    Properties1 byte
    BitMask byte
    data []byte
}

func NewEventSupportedReportV6(data []byte) *EventSupportedReportV6 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &EventSupportedReportV6{
				NotificationType: data[0],
				Properties1: data[1],
				BitMask: data[2],
        data: data,
    }
}

