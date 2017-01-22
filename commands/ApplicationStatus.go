package commands
type ApplicationBusy struct {
	node byte
	Status byte
	WaitTime byte
}

func NewApplicationBusy() ApplicationBusy {
	return ApplicationBusy{}
}

func (c *ApplicationBusy) SetNode(node int) {
	c.node = byte(node)
}

func (c *ApplicationBusy) Set(Status byte,WaitTime byte,) error {
	c.Status = Status
	c.WaitTime = WaitTime

	return nil
}

func (c *ApplicationBusy) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ApplicationStatus),
		0x01,
		c.Status,
		c.WaitTime,
		0x25,
	}
}
type ApplicationRejectedRequest struct {
	node byte
	Status byte
}

func NewApplicationRejectedRequest() ApplicationRejectedRequest {
	return ApplicationRejectedRequest{}
}

func (c *ApplicationRejectedRequest) SetNode(node int) {
	c.node = byte(node)
}

func (c *ApplicationRejectedRequest) Set(Status byte,) error {
	c.Status = Status

	return nil
}

func (c *ApplicationRejectedRequest) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ApplicationStatus),
		0x02,
		c.Status,
		0x25,
	}
}
