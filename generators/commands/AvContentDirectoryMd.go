package commands
type AvContentBrowseMdByLetterGet struct {
	node byte
}

func NewAvContentBrowseMdByLetterGet() AvContentBrowseMdByLetterGet {
	return AvContentBrowseMdByLetterGet{}
}

func (c *AvContentBrowseMdByLetterGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AvContentBrowseMdByLetterGet) Set() error {

	return nil
}

func (c *AvContentBrowseMdByLetterGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(AvContentDirectoryMd),
		0x03,
		0x25,
	}
}
type AvContentBrowseMdByLetterReport struct {
    *report
    data []byte
}

func NewAvContentBrowseMdByLetterReport(data []byte) *AvContentBrowseMdByLetterReport {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &AvContentBrowseMdByLetterReport{
        data: data,
    }
}

type AvContentBrowseMdChildCountGet struct {
	node byte
}

func NewAvContentBrowseMdChildCountGet() AvContentBrowseMdChildCountGet {
	return AvContentBrowseMdChildCountGet{}
}

func (c *AvContentBrowseMdChildCountGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AvContentBrowseMdChildCountGet) Set() error {

	return nil
}

func (c *AvContentBrowseMdChildCountGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(AvContentDirectoryMd),
		0x05,
		0x25,
	}
}
type AvContentBrowseMdChildCountReport struct {
    *report
    data []byte
}

func NewAvContentBrowseMdChildCountReport(data []byte) *AvContentBrowseMdChildCountReport {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &AvContentBrowseMdChildCountReport{
        data: data,
    }
}

type AvContentBrowseMdGet struct {
	node byte
}

func NewAvContentBrowseMdGet() AvContentBrowseMdGet {
	return AvContentBrowseMdGet{}
}

func (c *AvContentBrowseMdGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AvContentBrowseMdGet) Set() error {

	return nil
}

func (c *AvContentBrowseMdGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(AvContentDirectoryMd),
		0x01,
		0x25,
	}
}
type AvContentBrowseMdReport struct {
    *report
    data []byte
}

func NewAvContentBrowseMdReport(data []byte) *AvContentBrowseMdReport {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &AvContentBrowseMdReport{
        data: data,
    }
}

type AvMatchItemToRendererMdGet struct {
	node byte
}

func NewAvMatchItemToRendererMdGet() AvMatchItemToRendererMdGet {
	return AvMatchItemToRendererMdGet{}
}

func (c *AvMatchItemToRendererMdGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *AvMatchItemToRendererMdGet) Set() error {

	return nil
}

func (c *AvMatchItemToRendererMdGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(AvContentDirectoryMd),
		0x07,
		0x25,
	}
}
type AvMatchItemToRendererMdReport struct {
    *report
    data []byte
}

func NewAvMatchItemToRendererMdReport(data []byte) *AvMatchItemToRendererMdReport {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &AvMatchItemToRendererMdReport{
        data: data,
    }
}

