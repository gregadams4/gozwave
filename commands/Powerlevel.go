package commands
import "encoding/binary"
type PowerlevelGet struct {
	node byte
}

func NewPowerlevelGet() PowerlevelGet {
	return PowerlevelGet{}
}

func (c *PowerlevelGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *PowerlevelGet) Set() error {

	return nil
}

func (c *PowerlevelGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Powerlevel),
		0x02,
		0x25,
	}
}
type PowerlevelReport struct {
    *report
    Powerlevel byte
    Timeout byte
    data []byte
}

func NewPowerlevelReport(data []byte) *PowerlevelReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &PowerlevelReport{
				Powerlevel: data[0],
				Timeout: data[1],
        data: data,
    }
}

type PowerlevelSet struct {
	node byte
	Powerlevel byte
	Timeout byte
}

func NewPowerlevelSet() PowerlevelSet {
	return PowerlevelSet{}
}

func (c *PowerlevelSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *PowerlevelSet) Set(Powerlevel byte,Timeout byte,) error {
	c.Powerlevel = Powerlevel
	c.Timeout = Timeout

	return nil
}

func (c *PowerlevelSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(Powerlevel),
		0x01,
		c.Powerlevel,
		c.Timeout,
		0x25,
	}
}
type PowerlevelTestNodeGet struct {
	node byte
}

func NewPowerlevelTestNodeGet() PowerlevelTestNodeGet {
	return PowerlevelTestNodeGet{}
}

func (c *PowerlevelTestNodeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *PowerlevelTestNodeGet) Set() error {

	return nil
}

func (c *PowerlevelTestNodeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Powerlevel),
		0x05,
		0x25,
	}
}
type PowerlevelTestNodeReport struct {
    *report
    TestNodeID byte
    Statusofoperation byte
    TestFrameCount uint16
    data []byte
}

func NewPowerlevelTestNodeReport(data []byte) *PowerlevelTestNodeReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &PowerlevelTestNodeReport{
				TestNodeID: data[0],
				Statusofoperation: data[1],
				TestFrameCount: binary.BigEndian.Uint16(data[2:4]),
        data: data,
    }
}

type PowerlevelTestNodeSet struct {
	node byte
	TestNodeID byte
	Powerlevel byte
	Testframecount uint16
}

func NewPowerlevelTestNodeSet() PowerlevelTestNodeSet {
	return PowerlevelTestNodeSet{}
}

func (c *PowerlevelTestNodeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *PowerlevelTestNodeSet) Set(TestNodeID byte,Powerlevel byte,Testframecount uint16,) error {
	c.TestNodeID = TestNodeID
	c.Powerlevel = Powerlevel
	c.Testframecount = Testframecount

	return nil
}

func (c *PowerlevelTestNodeSet) Encode() []byte {
	TestframecountBytes := []byte{}
	binary.BigEndian.PutUint16(TestframecountBytes, c.Testframecount)
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(Powerlevel),
		0x04,
		c.TestNodeID,
		c.Powerlevel,
		TestframecountBytes[0],
		TestframecountBytes[1],
		0x25,
	}
}
