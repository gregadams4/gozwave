package commands
import "encoding/binary"
type RateTblActiveRateGet struct {
	node byte
}

func NewRateTblActiveRateGet() RateTblActiveRateGet {
	return RateTblActiveRateGet{}
}

func (c *RateTblActiveRateGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *RateTblActiveRateGet) Set() error {

	return nil
}

func (c *RateTblActiveRateGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(RateTblMonitor),
		0x05,
		0x25,
	}
}
type RateTblActiveRateReport struct {
    *report
    RateParameterSetID byte
    data []byte
}

func NewRateTblActiveRateReport(data []byte) *RateTblActiveRateReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &RateTblActiveRateReport{
				RateParameterSetID: data[0],
        data: data,
    }
}

type RateTblCurrentDataGet struct {
	node byte
	RateParameterSetID byte
	DatasetRequested byte
}

func NewRateTblCurrentDataGet() RateTblCurrentDataGet {
	return RateTblCurrentDataGet{}
}

func (c *RateTblCurrentDataGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *RateTblCurrentDataGet) Set(RateParameterSetID byte,DatasetRequested byte,) error {
	c.RateParameterSetID = RateParameterSetID
	c.DatasetRequested = DatasetRequested

	return nil
}

func (c *RateTblCurrentDataGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(RateTblMonitor),
		0x07,
		c.RateParameterSetID,
		c.DatasetRequested,
		0x25,
	}
}
type RateTblCurrentDataReport struct {
    *report
    ReportstoFollow byte
    RateParameterSetID byte
    Dataset byte
    Year uint16
    Month byte
    Day byte
    HourLocalTime byte
    MinuteLocalTime byte
    SecondLocalTime byte
    data []byte
}

func NewRateTblCurrentDataReport(data []byte) *RateTblCurrentDataReport {
    if len(data) < 10 {
				//may want to change this to just return nil
				for i := len(data); i < 10; i++ {
            data = append(data, 0x00)
        }
    }

    return &RateTblCurrentDataReport{
				ReportstoFollow: data[0],
				RateParameterSetID: data[1],
				Dataset: data[2],
				Year: binary.BigEndian.Uint16(data[3:5]),
				Month: data[5],
				Day: data[6],
				HourLocalTime: data[7],
				MinuteLocalTime: data[8],
				SecondLocalTime: data[9],
        data: data,
    }
}

type RateTblGet struct {
	node byte
	RateParameterSetID byte
}

func NewRateTblGet() RateTblGet {
	return RateTblGet{}
}

func (c *RateTblGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *RateTblGet) Set(RateParameterSetID byte,) error {
	c.RateParameterSetID = RateParameterSetID

	return nil
}

func (c *RateTblGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(RateTblMonitor),
		0x03,
		c.RateParameterSetID,
		0x25,
	}
}
type RateTblHistoricalDataGet struct {
	node byte
	MaximumReports byte
	RateParameterSetID byte
	DatasetRequested byte
	StartYear uint16
	StartMonth byte
	StartDay byte
	StartHourLocalTime byte
	StartMinuteLocalTime byte
	StartSecondLocalTime byte
	StopYear uint16
	StopMonth byte
	StopDay byte
	StopHourLocalTime byte
	StopMinuteLocalTime byte
	StopSecondLocalTime byte
}

func NewRateTblHistoricalDataGet() RateTblHistoricalDataGet {
	return RateTblHistoricalDataGet{}
}

func (c *RateTblHistoricalDataGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *RateTblHistoricalDataGet) Set(MaximumReports byte,RateParameterSetID byte,DatasetRequested byte,StartYear uint16,StartMonth byte,StartDay byte,StartHourLocalTime byte,StartMinuteLocalTime byte,StartSecondLocalTime byte,StopYear uint16,StopMonth byte,StopDay byte,StopHourLocalTime byte,StopMinuteLocalTime byte,StopSecondLocalTime byte,) error {
	c.MaximumReports = MaximumReports
	c.RateParameterSetID = RateParameterSetID
	c.DatasetRequested = DatasetRequested
	c.StartYear = StartYear
	c.StartMonth = StartMonth
	c.StartDay = StartDay
	c.StartHourLocalTime = StartHourLocalTime
	c.StartMinuteLocalTime = StartMinuteLocalTime
	c.StartSecondLocalTime = StartSecondLocalTime
	c.StopYear = StopYear
	c.StopMonth = StopMonth
	c.StopDay = StopDay
	c.StopHourLocalTime = StopHourLocalTime
	c.StopMinuteLocalTime = StopMinuteLocalTime
	c.StopSecondLocalTime = StopSecondLocalTime

	return nil
}

func (c *RateTblHistoricalDataGet) Encode() []byte {
	StartYearBytes := []byte{}
	binary.BigEndian.PutUint16(StartYearBytes, c.StartYear)
	StopYearBytes := []byte{}
	binary.BigEndian.PutUint16(StopYearBytes, c.StopYear)
	return []byte{
		0x13,
		c.node,
		byte(17),
		byte(RateTblMonitor),
		0x09,
		c.MaximumReports,
		c.RateParameterSetID,
		c.DatasetRequested,
		StartYearBytes[0],
		StartYearBytes[1],
		c.StartMonth,
		c.StartDay,
		c.StartHourLocalTime,
		c.StartMinuteLocalTime,
		c.StartSecondLocalTime,
		StopYearBytes[0],
		StopYearBytes[1],
		c.StopMonth,
		c.StopDay,
		c.StopHourLocalTime,
		c.StopMinuteLocalTime,
		c.StopSecondLocalTime,
		0x25,
	}
}
type RateTblHistoricalDataReport struct {
    *report
    ReportstoFollow byte
    RateParameterSetID byte
    Dataset byte
    Year uint16
    Month byte
    Day byte
    HourLocalTime byte
    MinuteLocalTime byte
    SecondLocalTime byte
    data []byte
}

func NewRateTblHistoricalDataReport(data []byte) *RateTblHistoricalDataReport {
    if len(data) < 10 {
				//may want to change this to just return nil
				for i := len(data); i < 10; i++ {
            data = append(data, 0x00)
        }
    }

    return &RateTblHistoricalDataReport{
				ReportstoFollow: data[0],
				RateParameterSetID: data[1],
				Dataset: data[2],
				Year: binary.BigEndian.Uint16(data[3:5]),
				Month: data[5],
				Day: data[6],
				HourLocalTime: data[7],
				MinuteLocalTime: data[8],
				SecondLocalTime: data[9],
        data: data,
    }
}

type RateTblReport struct {
    *report
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
    data []byte
}

func NewRateTblReport(data []byte) *RateTblReport {
    if len(data) < 22 {
				//may want to change this to just return nil
				for i := len(data); i < 22; i++ {
            data = append(data, 0x00)
        }
    }

    return &RateTblReport{
				RateParameterSetID: data[0],
				Properties1: data[1],
				RateCharacter: data[2],
				StartHourLocalTime: data[3],
				StartMinuteLocalTime: data[4],
				DurationMinute: binary.BigEndian.Uint16(data[5:7]),
				Properties2: data[7],
				MinConsumptionValue: binary.BigEndian.Uint64(data[8:12]),
				MaxConsumptionValue: binary.BigEndian.Uint64(data[12:16]),
				Properties3: data[16],
				MaxDemandValue: binary.BigEndian.Uint64(data[17:21]),
				DCPRateID: data[21],
        data: data,
    }
}

type RateTblSupportedGet struct {
	node byte
}

func NewRateTblSupportedGet() RateTblSupportedGet {
	return RateTblSupportedGet{}
}

func (c *RateTblSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *RateTblSupportedGet) Set() error {

	return nil
}

func (c *RateTblSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(RateTblMonitor),
		0x01,
		0x25,
	}
}
type RateTblSupportedReport struct {
    *report
    RatesSupported byte
    ParameterSetSupportedBitMask uint16
    data []byte
}

func NewRateTblSupportedReport(data []byte) *RateTblSupportedReport {
    if len(data) < 3 {
				//may want to change this to just return nil
				for i := len(data); i < 3; i++ {
            data = append(data, 0x00)
        }
    }

    return &RateTblSupportedReport{
				RatesSupported: data[0],
				ParameterSetSupportedBitMask: binary.BigEndian.Uint16(data[1:3]),
        data: data,
    }
}

