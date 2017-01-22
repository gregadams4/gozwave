package commands
type BindAccept struct {
	node byte
}

func NewBindAccept() BindAccept {
	return BindAccept{}
}

func (c *BindAccept) SetNode(node int) {
	c.node = byte(node)
}

func (c *BindAccept) Set() error {

	return nil
}

func (c *BindAccept) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ZensorNet),
		0x02,
		0x25,
	}
}
type BindComplete struct {
	node byte
}

func NewBindComplete() BindComplete {
	return BindComplete{}
}

func (c *BindComplete) SetNode(node int) {
	c.node = byte(node)
}

func (c *BindComplete) Set() error {

	return nil
}

func (c *BindComplete) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ZensorNet),
		0x03,
		0x25,
	}
}
type BindRequest struct {
	node byte
}

func NewBindRequest() BindRequest {
	return BindRequest{}
}

func (c *BindRequest) SetNode(node int) {
	c.node = byte(node)
}

func (c *BindRequest) Set() error {

	return nil
}

func (c *BindRequest) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ZensorNet),
		0x01,
		0x25,
	}
}
