package commands
import "encoding/binary"
type RateTblRemove struct {
	node byte
	Properties1 byte
	RateParameterSetID byte
}

func NewRateTblRemove() RateTblRemove {
	return RateTblRemove{}
}

func (c *RateTblRemove) SetNode(node int) {
	c.node = byte(node)
}

func (c *RateTblRemove) Set(Properties1 byte,RateParameterSetID byte,) error {
	c.Properties1 = Properties1
	c.RateParameterSetID = RateParameterSetID

	return nil
}

func (c *RateTblRemove) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(RateTblConfig),
		0x02,
		c.Properties1,
		c.RateParameterSetID,
		0x25,
	}
}
type RateTblSet struct {
	node byte
	RateParameterSetID byte
	Properties1 byte
	RateCharacter byte
	StartHourLocalTime byte
	StartMinuteLocalTime byte
	DurationMinute uint16
	Properties2 byte
	MinConsumptionValue uint64
	MaxConsumptionValue uint64
	Properties3 byte
	MaxDemandValue uint64
	DCPRateID byte
}

func NewRateTblSet() RateTblSet {
	return RateTblSet{}
}

func (c *RateTblSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *RateTblSet) Set(RateParameterSetID byte,Properties1 byte,RateCharacter byte,StartHourLocalTime byte,StartMinuteLocalTime byte,DurationMinute uint16,Properties2 byte,MinConsumptionValue uint64,MaxConsumptionValue uint64,Properties3 byte,MaxDemandValue uint64,DCPRateID byte,) error {
	c.RateParameterSetID = RateParameterSetID
	c.Properties1 = Properties1
	c.RateCharacter = RateCharacter
	c.StartHourLocalTime = StartHourLocalTime
	c.StartMinuteLocalTime = StartMinuteLocalTime
	c.DurationMinute = DurationMinute
	c.Properties2 = Properties2
	c.MinConsumptionValue = MinConsumptionValue
	c.MaxConsumptionValue = MaxConsumptionValue
	c.Properties3 = Properties3
	c.MaxDemandValue = MaxDemandValue
	c.DCPRateID = DCPRateID

	return nil
}

func (c *RateTblSet) Encode() []byte {
	DurationMinuteBytes := []byte{}
	binary.BigEndian.PutUint16(DurationMinuteBytes, c.DurationMinute)
	MinConsumptionValueBytes := []byte{}
	binary.BigEndian.PutUint64(MinConsumptionValueBytes, c.MinConsumptionValue)
	MaxConsumptionValueBytes := []byte{}
	binary.BigEndian.PutUint64(MaxConsumptionValueBytes, c.MaxConsumptionValue)
	MaxDemandValueBytes := []byte{}
	binary.BigEndian.PutUint64(MaxDemandValueBytes, c.MaxDemandValue)
	return []byte{
		0x13,
		c.node,
		byte(14),
		byte(RateTblConfig),
		0x01,
		c.RateParameterSetID,
		c.Properties1,
		c.RateCharacter,
		c.StartHourLocalTime,
		c.StartMinuteLocalTime,
		DurationMinuteBytes[0],
		DurationMinuteBytes[1],
		c.Properties2,
		MinConsumptionValueBytes[0],
		MinConsumptionValueBytes[1],
		MinConsumptionValueBytes[2],
		MinConsumptionValueBytes[3],
		MaxConsumptionValueBytes[0],
		MaxConsumptionValueBytes[1],
		MaxConsumptionValueBytes[2],
		MaxConsumptionValueBytes[3],
		c.Properties3,
		MaxDemandValueBytes[0],
		MaxDemandValueBytes[1],
		MaxDemandValueBytes[2],
		MaxDemandValueBytes[3],
		c.DCPRateID,
		0x25,
	}
}
