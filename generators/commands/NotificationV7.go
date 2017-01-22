package commands
type NotificationGetV7 struct {
    *report
    V1AlarmType byte
    NotificationType byte
    Event byte
    data []byte
}

func NewNotificationGetV7(data []byte) *NotificationGetV7 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationGetV7{
				V1AlarmType: data[0],
				NotificationType: data[1],
				Event: data[2],
        data: data,
    }
}

type NotificationReportV7 struct {
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

func NewNotificationReportV7(data []byte) *NotificationReportV7 {
    if len(data) < 9 {
				//may want to change this to just return nil
				for i := len(data); i < 9; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationReportV7{
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

type NotificationSetV7 struct {
    *report
    NotificationType byte
    NotificationStatus byte
    data []byte
}

func NewNotificationSetV7(data []byte) *NotificationSetV7 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSetV7{
				NotificationType: data[0],
				NotificationStatus: data[1],
        data: data,
    }
}

type NotificationSupportedGetV7 struct {
    *report
    data []byte
}

func NewNotificationSupportedGetV7(data []byte) *NotificationSupportedGetV7 {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSupportedGetV7{
        data: data,
    }
}

type NotificationSupportedReportV7 struct {
    *report
    Properties1 byte
    BitMask byte
    data []byte
}

func NewNotificationSupportedReportV7(data []byte) *NotificationSupportedReportV7 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSupportedReportV7{
				Properties1: data[0],
				BitMask: data[1],
        data: data,
    }
}

type EventSupportedGetV7 struct {
	node byte
	NotificationType byte
}

func NewEventSupportedGetV7() EventSupportedGetV7 {
	return EventSupportedGetV7{}
}

func (c *EventSupportedGetV7) SetNode(node int) {
	c.node = byte(node)
}

func (c *EventSupportedGetV7) Set(NotificationType byte,) error {
	c.NotificationType = NotificationType

	return nil
}

func (c *EventSupportedGetV7) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NotificationV7),
		0x01,
		c.NotificationType,
		0x25,
	}
}
type EventSupportedReportV7 struct {
    *report
    NotificationType byte
    Properties1 byte
    BitMask byte
    data []byte
}

func NewEventSupportedReportV7(data []byte) *EventSupportedReportV7 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &EventSupportedReportV7{
				NotificationType: data[0],
				Properties1: data[1],
				BitMask: data[2],
        data: data,
    }
}

