package commands
type DoorLockConfigurationGetV2 struct {
	node byte
}

func NewDoorLockConfigurationGetV2() DoorLockConfigurationGetV2 {
	return DoorLockConfigurationGetV2{}
}

func (c *DoorLockConfigurationGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockConfigurationGetV2) Set() error {

	return nil
}

func (c *DoorLockConfigurationGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(DoorLockV2),
		0x05,
		0x25,
	}
}
type DoorLockConfigurationReportV2 struct {
    *report
    OperationType byte
    Properties1 byte
    LockTimeoutMinutes byte
    LockTimeoutSeconds byte
    data []byte
}

func NewDoorLockConfigurationReportV2(data []byte) *DoorLockConfigurationReportV2 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &DoorLockConfigurationReportV2{
				OperationType: data[0],
				Properties1: data[1],
				LockTimeoutMinutes: data[2],
				LockTimeoutSeconds: data[3],
        data: data,
    }
}

type DoorLockConfigurationSetV2 struct {
	node byte
	OperationType byte
	Properties1 byte
	LockTimeoutMinutes byte
	LockTimeoutSeconds byte
}

func NewDoorLockConfigurationSetV2() DoorLockConfigurationSetV2 {
	return DoorLockConfigurationSetV2{}
}

func (c *DoorLockConfigurationSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockConfigurationSetV2) Set(OperationType byte,Properties1 byte,LockTimeoutMinutes byte,LockTimeoutSeconds byte,) error {
	c.OperationType = OperationType
	c.Properties1 = Properties1
	c.LockTimeoutMinutes = LockTimeoutMinutes
	c.LockTimeoutSeconds = LockTimeoutSeconds

	return nil
}

func (c *DoorLockConfigurationSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(DoorLockV2),
		0x04,
		c.OperationType,
		c.Properties1,
		c.LockTimeoutMinutes,
		c.LockTimeoutSeconds,
		0x25,
	}
}
type DoorLockOperationGetV2 struct {
	node byte
}

func NewDoorLockOperationGetV2() DoorLockOperationGetV2 {
	return DoorLockOperationGetV2{}
}

func (c *DoorLockOperationGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockOperationGetV2) Set() error {

	return nil
}

func (c *DoorLockOperationGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(DoorLockV2),
		0x02,
		0x25,
	}
}
type DoorLockOperationReportV2 struct {
    *report
    DoorLockMode byte
    Properties1 byte
    DoorCondition byte
    LockTimeoutMinutes byte
    LockTimeoutSeconds byte
    data []byte
}

func NewDoorLockOperationReportV2(data []byte) *DoorLockOperationReportV2 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &DoorLockOperationReportV2{
				DoorLockMode: data[0],
				Properties1: data[1],
				DoorCondition: data[2],
				LockTimeoutMinutes: data[3],
				LockTimeoutSeconds: data[4],
        data: data,
    }
}

type DoorLockOperationSetV2 struct {
	node byte
	DoorLockMode byte
}

func NewDoorLockOperationSetV2() DoorLockOperationSetV2 {
	return DoorLockOperationSetV2{}
}

func (c *DoorLockOperationSetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockOperationSetV2) Set(DoorLockMode byte,) error {
	c.DoorLockMode = DoorLockMode

	return nil
}

func (c *DoorLockOperationSetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(DoorLockV2),
		0x01,
		c.DoorLockMode,
		0x25,
	}
}
