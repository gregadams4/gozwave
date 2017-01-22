package commands
import "encoding/binary"
type MeterPulseGet struct {
	node byte
}

func NewMeterPulseGet() MeterPulseGet {
	return MeterPulseGet{}
}

func (c *MeterPulseGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterPulseGet) Set() error {

	return nil
}

func (c *MeterPulseGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterPulse),
		0x04,
		0x25,
	}
}
type MeterPulseReport struct {
    *report
    PulseCount uint64
    data []byte
}

func NewMeterPulseReport(data []byte) *MeterPulseReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterPulseReport{
				PulseCount: binary.BigEndian.Uint64(data[0:4]),
        data: data,
    }
}

