package commands
type DoorLockConfigurationGet struct {
	node byte
}

func NewDoorLockConfigurationGet() DoorLockConfigurationGet {
	return DoorLockConfigurationGet{}
}

func (c *DoorLockConfigurationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockConfigurationGet) Set() error {

	return nil
}

func (c *DoorLockConfigurationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(DoorLock),
		0x05,
		0x25,
	}
}
type DoorLockConfigurationReport struct {
    *report
    OperationType byte
    Properties1 byte
    LockTimeoutMinutes byte
    LockTimeoutSeconds byte
    data []byte
}

func NewDoorLockConfigurationReport(data []byte) *DoorLockConfigurationReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &DoorLockConfigurationReport{
				OperationType: data[0],
				Properties1: data[1],
				LockTimeoutMinutes: data[2],
				LockTimeoutSeconds: data[3],
        data: data,
    }
}

type DoorLockConfigurationSet struct {
	node byte
	OperationType byte
	Properties1 byte
	LockTimeoutMinutes byte
	LockTimeoutSeconds byte
}

func NewDoorLockConfigurationSet() DoorLockConfigurationSet {
	return DoorLockConfigurationSet{}
}

func (c *DoorLockConfigurationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockConfigurationSet) Set(OperationType byte,Properties1 byte,LockTimeoutMinutes byte,LockTimeoutSeconds byte,) error {
	c.OperationType = OperationType
	c.Properties1 = Properties1
	c.LockTimeoutMinutes = LockTimeoutMinutes
	c.LockTimeoutSeconds = LockTimeoutSeconds

	return nil
}

func (c *DoorLockConfigurationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(DoorLock),
		0x04,
		c.OperationType,
		c.Properties1,
		c.LockTimeoutMinutes,
		c.LockTimeoutSeconds,
		0x25,
	}
}
type DoorLockOperationGet struct {
	node byte
}

func NewDoorLockOperationGet() DoorLockOperationGet {
	return DoorLockOperationGet{}
}

func (c *DoorLockOperationGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockOperationGet) Set() error {

	return nil
}

func (c *DoorLockOperationGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(DoorLock),
		0x02,
		0x25,
	}
}
type DoorLockOperationReport struct {
    *report
    DoorLockMode byte
    Properties1 byte
    DoorCondition byte
    LockTimeoutMinutes byte
    LockTimeoutSeconds byte
    data []byte
}

func NewDoorLockOperationReport(data []byte) *DoorLockOperationReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &DoorLockOperationReport{
				DoorLockMode: data[0],
				Properties1: data[1],
				DoorCondition: data[2],
				LockTimeoutMinutes: data[3],
				LockTimeoutSeconds: data[4],
        data: data,
    }
}

type DoorLockOperationSet struct {
	node byte
	DoorLockMode byte
}

func NewDoorLockOperationSet() DoorLockOperationSet {
	return DoorLockOperationSet{}
}

func (c *DoorLockOperationSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *DoorLockOperationSet) Set(DoorLockMode byte,) error {
	c.DoorLockMode = DoorLockMode

	return nil
}

func (c *DoorLockOperationSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(DoorLock),
		0x01,
		c.DoorLockMode,
		0x25,
	}
}
