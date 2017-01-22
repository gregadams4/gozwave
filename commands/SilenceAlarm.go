package commands
import "encoding/binary"
type SensorAlarmSet struct {
	node byte
	Mode byte
	Seconds uint16
	NumberofBitMasks byte
	BitMask byte
}

func NewSensorAlarmSet() SensorAlarmSet {
	return SensorAlarmSet{}
}

func (c *SensorAlarmSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *SensorAlarmSet) Set(Mode byte,Seconds uint16,NumberofBitMasks byte,BitMask byte,) error {
	c.Mode = Mode
	c.Seconds = Seconds
	c.NumberofBitMasks = NumberofBitMasks
	c.BitMask = BitMask

	return nil
}

func (c *SensorAlarmSet) Encode() []byte {
	SecondsBytes := []byte{}
	binary.BigEndian.PutUint16(SecondsBytes, c.Seconds)
	return []byte{
		0x13,
		c.node,
		byte(6),
		byte(SilenceAlarm),
		0x01,
		c.Mode,
		SecondsBytes[0],
		SecondsBytes[1],
		c.NumberofBitMasks,
		c.BitMask,
		0x25,
	}
}
