package commands
import "encoding/binary"
type MailboxConfigurationGet struct {
	node byte
}

func NewMailboxConfigurationGet() MailboxConfigurationGet {
	return MailboxConfigurationGet{}
}

func (c *MailboxConfigurationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MailboxConfigurationGet) Set() error {

	return nil
}

func (c *MailboxConfigurationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Mailbox),
		0x01,
		0x25,
	}
}
type MailboxConfigurationSet struct {
	node byte
	Properties1 byte
	ForwardingDestinationIPv6Address byte
	UDPPortNumber uint16
}

func NewMailboxConfigurationSet() MailboxConfigurationSet {
	return MailboxConfigurationSet{}
}

func (c *MailboxConfigurationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MailboxConfigurationSet) Set(Properties1 byte,ForwardingDestinationIPv6Address byte,UDPPortNumber uint16,) error {
	c.Properties1 = Properties1
	c.ForwardingDestinationIPv6Address = ForwardingDestinationIPv6Address
	c.UDPPortNumber = UDPPortNumber

	return nil
}

func (c *MailboxConfigurationSet) Encode() []byte {
	UDPPortNumberBytes := []byte{}
	binary.BigEndian.PutUint16(UDPPortNumberBytes, c.UDPPortNumber)
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(Mailbox),
		0x02,
		c.Properties1,
		c.ForwardingDestinationIPv6Address,
		UDPPortNumberBytes[0],
		UDPPortNumberBytes[1],
		0x25,
	}
}
type MailboxConfigurationReport struct {
    *report
    Properties1 byte
    MailboxCapacity uint16
    ForwardingDestinationIPv6Address byte
    UDPPortNumber uint16
    data []byte
}

func NewMailboxConfigurationReport(data []byte) *MailboxConfigurationReport {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &MailboxConfigurationReport{
				Properties1: data[0],
				MailboxCapacity: binary.BigEndian.Uint16(data[1:3]),
				ForwardingDestinationIPv6Address: data[3],
				UDPPortNumber: binary.BigEndian.Uint16(data[4:6]),
        data: data,
    }
}

type MailboxQueue struct {
	node byte
	SequenceNumber byte
	Properties1 byte
	QueueHandle byte
	MailboxEntry byte
}

func NewMailboxQueue() MailboxQueue {
	return MailboxQueue{}
}

func (c *MailboxQueue) SetNode(node int) {
	c.node = byte(node)
}

func (c *MailboxQueue) Set(SequenceNumber byte,Properties1 byte,QueueHandle byte,MailboxEntry byte,) error {
	c.SequenceNumber = SequenceNumber
	c.Properties1 = Properties1
	c.QueueHandle = QueueHandle
	c.MailboxEntry = MailboxEntry

	return nil
}

func (c *MailboxQueue) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(Mailbox),
		0x04,
		c.SequenceNumber,
		c.Properties1,
		c.QueueHandle,
		c.MailboxEntry,
		0x25,
	}
}
type MailboxWakeupNotification struct {
    *report
    QueueHandle byte
    data []byte
}

func NewMailboxWakeupNotification(data []byte) *MailboxWakeupNotification {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &MailboxWakeupNotification{
				QueueHandle: data[0],
        data: data,
    }
}

type MailboxNodeFailing struct {
	node byte
	QueueHandle byte
}

func NewMailboxNodeFailing() MailboxNodeFailing {
	return MailboxNodeFailing{}
}

func (c *MailboxNodeFailing) SetNode(node int) {
	c.node = byte(node)
}

func (c *MailboxNodeFailing) Set(QueueHandle byte,) error {
	c.QueueHandle = QueueHandle

	return nil
}

func (c *MailboxNodeFailing) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Mailbox),
		0x06,
		c.QueueHandle,
		0x25,
	}
}
