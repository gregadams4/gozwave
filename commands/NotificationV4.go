package commands
type NotificationGetV4 struct {
    *report
    V1AlarmType byte
    NotificationType byte
    Event byte
    data []byte
}

func NewNotificationGetV4(data []byte) *NotificationGetV4 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationGetV4{
				V1AlarmType: data[0],
				NotificationType: data[1],
				Event: data[2],
        data: data,
    }
}

type NotificationReportV4 struct {
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

func NewNotificationReportV4(data []byte) *NotificationReportV4 {
    if len(data) < 9 {
				//may want to change this to just return nil
				for i := len(data); i < 9; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationReportV4{
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

type NotificationSetV4 struct {
    *report
    NotificationType byte
    NotificationStatus byte
    data []byte
}

func NewNotificationSetV4(data []byte) *NotificationSetV4 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSetV4{
				NotificationType: data[0],
				NotificationStatus: data[1],
        data: data,
    }
}

type NotificationSupportedGetV4 struct {
    *report
    data []byte
}

func NewNotificationSupportedGetV4(data []byte) *NotificationSupportedGetV4 {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSupportedGetV4{
        data: data,
    }
}

type NotificationSupportedReportV4 struct {
    *report
    Properties1 byte
    BitMask byte
    data []byte
}

func NewNotificationSupportedReportV4(data []byte) *NotificationSupportedReportV4 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSupportedReportV4{
				Properties1: data[0],
				BitMask: data[1],
        data: data,
    }
}

type EventSupportedGetV4 struct {
	node byte
	NotificationType byte
}

func NewEventSupportedGetV4() EventSupportedGetV4 {
	return EventSupportedGetV4{}
}

func (c *EventSupportedGetV4) SetNode(node int) {
	c.node = byte(node)
}

func (c *EventSupportedGetV4) Set(NotificationType byte,) error {
	c.NotificationType = NotificationType

	return nil
}

func (c *EventSupportedGetV4) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NotificationV4),
		0x01,
		c.NotificationType,
		0x25,
	}
}
type EventSupportedReportV4 struct {
    *report
    NotificationType byte
    Properties1 byte
    BitMask byte
    data []byte
}

func NewEventSupportedReportV4(data []byte) *EventSupportedReportV4 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &EventSupportedReportV4{
				NotificationType: data[0],
				Properties1: data[1],
				BitMask: data[2],
        data: data,
    }
}

