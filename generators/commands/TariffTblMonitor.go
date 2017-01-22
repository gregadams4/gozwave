package commands
import "encoding/binary"
type TariffTblCostGet struct {
	node byte
	RateParameterSetID byte
	StartYear uint16
	StartMonth byte
	StartDay byte
	StartHourLocalTime byte
	StartMinuteLocalTime byte
	StopYear uint16
	StopMonth byte
	StopDay byte
	StopHourLocalTime byte
	StopMinuteLocalTime byte
}

func NewTariffTblCostGet() TariffTblCostGet {
	return TariffTblCostGet{}
}

func (c *TariffTblCostGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *TariffTblCostGet) Set(RateParameterSetID byte,StartYear uint16,StartMonth byte,StartDay byte,StartHourLocalTime byte,StartMinuteLocalTime byte,StopYear uint16,StopMonth byte,StopDay byte,StopHourLocalTime byte,StopMinuteLocalTime byte,) error {
	c.RateParameterSetID = RateParameterSetID
	c.StartYear = StartYear
	c.StartMonth = StartMonth
	c.StartDay = StartDay
	c.StartHourLocalTime = StartHourLocalTime
	c.StartMinuteLocalTime = StartMinuteLocalTime
	c.StopYear = StopYear
	c.StopMonth = StopMonth
	c.StopDay = StopDay
	c.StopHourLocalTime = StopHourLocalTime
	c.StopMinuteLocalTime = StopMinuteLocalTime

	return nil
}

func (c *TariffTblCostGet) Encode() []byte {
	StartYearBytes := []byte{}
	binary.BigEndian.PutUint16(StartYearBytes, c.StartYear)
	StopYearBytes := []byte{}
	binary.BigEndian.PutUint16(StopYearBytes, c.StopYear)
	return []byte{
		0x13,
		c.node,
		byte(13),
		byte(TariffTblMonitor),
		0x05,
		c.RateParameterSetID,
		StartYearBytes[0],
		StartYearBytes[1],
		c.StartMonth,
		c.StartDay,
		c.StartHourLocalTime,
		c.StartMinuteLocalTime,
		StopYearBytes[0],
		StopYearBytes[1],
		c.StopMonth,
		c.StopDay,
		c.StopHourLocalTime,
		c.StopMinuteLocalTime,
		0x25,
	}
}
type TariffTblCostReport struct {
    *report
    RateParameterSetID byte
    Properties1 byte
    StartYear uint16
    StartMonth byte
    StartDay byte
    StartHourLocalTime byte
    StartMinuteLocalTime byte
    StopYear uint16
    StopMonth byte
    StopDay byte
    StopHourLocalTime byte
    StopMinuteLocalTime byte
    Currency byte
    Properties2 byte
    CostValue uint64
    data []byte
}

func NewTariffTblCostReport(data []byte) *TariffTblCostReport {
    if len(data) < 20 {
				//may want to change this to just return nil
				for i := len(data); i < 20; i++ {
            data = append(data, 0x00)
        }
    }

    return &TariffTblCostReport{
				RateParameterSetID: data[0],
				Properties1: data[1],
				StartYear: binary.BigEndian.Uint16(data[2:4]),
				StartMonth: data[4],
				StartDay: data[5],
				StartHourLocalTime: data[6],
				StartMinuteLocalTime: data[7],
				StopYear: binary.BigEndian.Uint16(data[8:10]),
				StopMonth: data[10],
				StopDay: data[11],
				StopHourLocalTime: data[12],
				StopMinuteLocalTime: data[13],
				Currency: data[14],
				Properties2: data[15],
				CostValue: binary.BigEndian.Uint64(data[16:20]),
        data: data,
    }
}

type TariffTblGet struct {
	node byte
	RateParameterSetID byte
}

func NewTariffTblGet() TariffTblGet {
	return TariffTblGet{}
}

func (c *TariffTblGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *TariffTblGet) Set(RateParameterSetID byte,) error {
	c.RateParameterSetID = RateParameterSetID

	return nil
}

func (c *TariffTblGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(TariffTblMonitor),
		0x03,
		c.RateParameterSetID,
		0x25,
	}
}
type TariffTblReport struct {
    *report
    RateParameterSetID byte
    Properties1 byte
    TariffValue uint64
    data []byte
}

func NewTariffTblReport(data []byte) *TariffTblReport {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &TariffTblReport{
				RateParameterSetID: data[0],
				Properties1: data[1],
				TariffValue: binary.BigEndian.Uint64(data[2:6]),
        data: data,
    }
}

type TariffTblSupplierGet struct {
	node byte
}

func NewTariffTblSupplierGet() TariffTblSupplierGet {
	return TariffTblSupplierGet{}
}

func (c *TariffTblSupplierGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *TariffTblSupplierGet) Set() error {

	return nil
}

func (c *TariffTblSupplierGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(TariffTblMonitor),
		0x01,
		0x25,
	}
}
type TariffTblSupplierReport struct {
    *report
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
    data []byte
}

func NewTariffTblSupplierReport(data []byte) *TariffTblSupplierReport {
    if len(data) < 15 {
				//may want to change this to just return nil
				for i := len(data); i < 15; i++ {
            data = append(data, 0x00)
        }
    }

    return &TariffTblSupplierReport{
				Year: binary.BigEndian.Uint16(data[0:2]),
				Month: data[2],
				Day: data[3],
				HourLocalTime: data[4],
				MinuteLocalTime: data[5],
				SecondLocalTime: data[6],
				Currency: data[7],
				Properties1: data[8],
				StandingChargeValue: binary.BigEndian.Uint64(data[9:13]),
				Properties2: data[13],
				SupplierCharacter: data[14],
        data: data,
    }
}

