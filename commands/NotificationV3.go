package commands
type NotificationGetV3 struct {
    *report
    V1AlarmType byte
    NotificationType byte
    Event byte
    data []byte
}

func NewNotificationGetV3(data []byte) *NotificationGetV3 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationGetV3{
				V1AlarmType: data[0],
				NotificationType: data[1],
				Event: data[2],
        data: data,
    }
}

type NotificationReportV3 struct {
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

func NewNotificationReportV3(data []byte) *NotificationReportV3 {
    if len(data) < 9 {
				//may want to change this to just return nil
				for i := len(data); i < 9; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationReportV3{
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

type NotificationSetV3 struct {
    *report
    NotificationType byte
    NotificationStatus byte
    data []byte
}

func NewNotificationSetV3(data []byte) *NotificationSetV3 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSetV3{
				NotificationType: data[0],
				NotificationStatus: data[1],
        data: data,
    }
}

type NotificationSupportedGetV3 struct {
    *report
    data []byte
}

func NewNotificationSupportedGetV3(data []byte) *NotificationSupportedGetV3 {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSupportedGetV3{
        data: data,
    }
}

type NotificationSupportedReportV3 struct {
    *report
    Properties1 byte
    BitMask byte
    data []byte
}

func NewNotificationSupportedReportV3(data []byte) *NotificationSupportedReportV3 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &NotificationSupportedReportV3{
				Properties1: data[0],
				BitMask: data[1],
        data: data,
    }
}

type EventSupportedGetV3 struct {
	node byte
	NotificationType byte
}

func NewEventSupportedGetV3() EventSupportedGetV3 {
	return EventSupportedGetV3{}
}

func (c *EventSupportedGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *EventSupportedGetV3) Set(NotificationType byte,) error {
	c.NotificationType = NotificationType

	return nil
}

func (c *EventSupportedGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NotificationV3),
		0x01,
		c.NotificationType,
		0x25,
	}
}
type EventSupportedReportV3 struct {
    *report
    NotificationType byte
    Properties1 byte
    BitMask byte
    data []byte
}

func NewEventSupportedReportV3(data []byte) *EventSupportedReportV3 {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &EventSupportedReportV3{
				NotificationType: data[0],
				Properties1: data[1],
				BitMask: data[2],
        data: data,
    }
}

