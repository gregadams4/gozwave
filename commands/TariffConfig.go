package commands
import "encoding/binary"
type TariffTblRemove struct {
	node byte
	Properties1 byte
	RateParameterSetID byte
}

func NewTariffTblRemove() TariffTblRemove {
	return TariffTblRemove{}
}

func (c *TariffTblRemove) SetNode(node int) {
	c.node = byte(node)
}

func (c *TariffTblRemove) Set(Properties1 byte,RateParameterSetID byte,) error {
	c.Properties1 = Properties1
	c.RateParameterSetID = RateParameterSetID

	return nil
}

func (c *TariffTblRemove) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(TariffConfig),
		0x03,
		c.Properties1,
		c.RateParameterSetID,
		0x25,
	}
}
type TariffTblSet struct {
	node byte
	RateParameterSetID byte
	Properties1 byte
	TariffValue uint64
}

func NewTariffTblSet() TariffTblSet {
	return TariffTblSet{}
}

func (c *TariffTblSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *TariffTblSet) Set(RateParameterSetID byte,Properties1 byte,TariffValue uint64,) error {
	c.RateParameterSetID = RateParameterSetID
	c.Properties1 = Properties1
	c.TariffValue = TariffValue

	return nil
}

func (c *TariffTblSet) Encode() []byte {
	TariffValueBytes := []byte{}
	binary.BigEndian.PutUint64(TariffValueBytes, c.TariffValue)
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(TariffConfig),
		0x02,
		c.RateParameterSetID,
		c.Properties1,
		TariffValueBytes[0],
		TariffValueBytes[1],
		TariffValueBytes[2],
		TariffValueBytes[3],
		0x25,
	}
}
type TariffTblSupplierSet struct {
	node byte
	Year uint16
	Month byte
	Day byte
	HourLocalTime byte
	MinuteLocalTime byte
	SecondLocalTime byte
	Currency byte
	Properties1 byte
	StandingChargeValue uint64
	Properties2 byte
	SupplierCharacter byte
}

func NewTariffTblSupplierSet() TariffTblSupplierSet {
	return TariffTblSupplierSet{}
}

func (c *TariffTblSupplierSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *TariffTblSupplierSet) Set(Year uint16,Month byte,Day byte,HourLocalTime byte,MinuteLocalTime byte,SecondLocalTime byte,Currency byte,Properties1 byte,StandingChargeValue uint64,Properties2 byte,SupplierCharacter byte,) error {
	c.Year = Year
	c.Month = Month
	c.Day = Day
	c.HourLocalTime = HourLocalTime
	c.MinuteLocalTime = MinuteLocalTime
	c.SecondLocalTime = SecondLocalTime
	c.Currency = Currency
	c.Properties1 = Properties1
	c.StandingChargeValue = StandingChargeValue
	c.Properties2 = Properties2
	c.SupplierCharacter = SupplierCharacter

	return nil
}

func (c *TariffTblSupplierSet) Encode() []byte {
	YearBytes := []byte{}
	binary.BigEndian.PutUint16(YearBytes, c.Year)
	StandingChargeValueBytes := []byte{}
	binary.BigEndian.PutUint64(StandingChargeValueBytes, c.StandingChargeValue)
	return []byte{
		0x13,
		c.node,
		byte(13),
		byte(TariffConfig),
		0x01,
		YearBytes[0],
		YearBytes[1],
		c.Month,
		c.Day,
		c.HourLocalTime,
		c.MinuteLocalTime,
		c.SecondLocalTime,
		c.Currency,
		c.Properties1,
		StandingChargeValueBytes[0],
		StandingChargeValueBytes[1],
		StandingChargeValueBytes[2],
		StandingChargeValueBytes[3],
		c.Properties2,
		c.SupplierCharacter,
		0x25,
	}
}
