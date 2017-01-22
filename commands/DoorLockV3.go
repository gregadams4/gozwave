package commands
type DoorLockConfigurationGetV3 struct {
	node byte
}

func NewDoorLockConfigurationGetV3() DoorLockConfigurationGetV3 {
	return DoorLockConfigurationGetV3{}
}

func (c *DoorLockConfigurationGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockConfigurationGetV3) Set() error {

	return nil
}

func (c *DoorLockConfigurationGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(DoorLockV3),
		0x05,
		0x25,
	}
}
type DoorLockConfigurationReportV3 struct {
    *report
    OperationType byte
    Properties1 byte
    LockTimeoutMinutes byte
    LockTimeoutSeconds byte
    data []byte
}

func NewDoorLockConfigurationReportV3(data []byte) *DoorLockConfigurationReportV3 {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &DoorLockConfigurationReportV3{
				OperationType: data[0],
				Properties1: data[1],
				LockTimeoutMinutes: data[2],
				LockTimeoutSeconds: data[3],
        data: data,
    }
}

type DoorLockConfigurationSetV3 struct {
	node byte
	OperationType byte
	Properties1 byte
	LockTimeoutMinutes byte
	LockTimeoutSeconds byte
}

func NewDoorLockConfigurationSetV3() DoorLockConfigurationSetV3 {
	return DoorLockConfigurationSetV3{}
}

func (c *DoorLockConfigurationSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockConfigurationSetV3) Set(OperationType byte,Properties1 byte,LockTimeoutMinutes byte,LockTimeoutSeconds byte,) error {
	c.OperationType = OperationType
	c.Properties1 = Properties1
	c.LockTimeoutMinutes = LockTimeoutMinutes
	c.LockTimeoutSeconds = LockTimeoutSeconds

	return nil
}

func (c *DoorLockConfigurationSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(DoorLockV3),
		0x04,
		c.OperationType,
		c.Properties1,
		c.LockTimeoutMinutes,
		c.LockTimeoutSeconds,
		0x25,
	}
}
type DoorLockOperationGetV3 struct {
	node byte
}

func NewDoorLockOperationGetV3() DoorLockOperationGetV3 {
	return DoorLockOperationGetV3{}
}

func (c *DoorLockOperationGetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockOperationGetV3) Set() error {

	return nil
}

func (c *DoorLockOperationGetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(DoorLockV3),
		0x02,
		0x25,
	}
}
type DoorLockOperationReportV3 struct {
    *report
    CurrentDoorLockMode byte
    Properties1 byte
    DoorCondition byte
    LockTimeoutMinutes byte
    LockTimeoutSeconds byte
    TargetDoorLockMode byte
    Duration byte
    data []byte
}

func NewDoorLockOperationReportV3(data []byte) *DoorLockOperationReportV3 {
    if len(data) < 7 {
				//may want to change this to just return nil
				for i := len(data); i < 7; i++ {
            data = append(data, 0x00)
        }
    }

    return &DoorLockOperationReportV3{
				CurrentDoorLockMode: data[0],
				Properties1: data[1],
				DoorCondition: data[2],
				LockTimeoutMinutes: data[3],
				LockTimeoutSeconds: data[4],
				TargetDoorLockMode: data[5],
				Duration: data[6],
        data: data,
    }
}

type DoorLockOperationSetV3 struct {
	node byte
	DoorLockMode byte
}

func NewDoorLockOperationSetV3() DoorLockOperationSetV3 {
	return DoorLockOperationSetV3{}
}

func (c *DoorLockOperationSetV3) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockOperationSetV3) Set(DoorLockMode byte,) error {
	c.DoorLockMode = DoorLockMode

	return nil
}

func (c *DoorLockOperationSetV3) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(DoorLockV3),
		0x01,
		c.DoorLockMode,
		0x25,
	}
}
