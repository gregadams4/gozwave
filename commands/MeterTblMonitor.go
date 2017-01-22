package commands
import "encoding/binary"
type MeterTblStatusReport struct {
    *report
    Reportstofollow byte
    CurrentOperatingStatus byte
    data []byte
}

func NewMeterTblStatusReport(data []byte) *MeterTblStatusReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblStatusReport{
				Reportstofollow: data[0],
				CurrentOperatingStatus: data[1],
        data: data,
    }
}

type MeterTblStatusDateGet struct {
	node byte
	MaximumReports byte
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

func NewMeterTblStatusDateGet() MeterTblStatusDateGet {
	return MeterTblStatusDateGet{}
}

func (c *MeterTblStatusDateGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblStatusDateGet) Set(MaximumReports byte,StartYear uint16,StartMonth byte,StartDay byte,StartHourLocalTime byte,StartMinuteLocalTime byte,StartSecondLocalTime byte,StopYear uint16,StopMonth byte,StopDay byte,StopHourLocalTime byte,StopMinuteLocalTime byte,StopSecondLocalTime byte,) error {
	c.MaximumReports = MaximumReports
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

func (c *MeterTblStatusDateGet) Encode() []byte {
	StartYearBytes := []byte{}
	binary.BigEndian.PutUint16(StartYearBytes, c.StartYear)
	StopYearBytes := []byte{}
	binary.BigEndian.PutUint16(StopYearBytes, c.StopYear)
	return []byte{
		0x13,
		c.node,
		byte(15),
		byte(MeterTblMonitor),
		0x0A,
		c.MaximumReports,
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
type MeterTblStatusDepthGet struct {
	node byte
	StatusEventLogDepth byte
}

func NewMeterTblStatusDepthGet() MeterTblStatusDepthGet {
	return MeterTblStatusDepthGet{}
}

func (c *MeterTblStatusDepthGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblStatusDepthGet) Set(StatusEventLogDepth byte,) error {
	c.StatusEventLogDepth = StatusEventLogDepth

	return nil
}

func (c *MeterTblStatusDepthGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MeterTblMonitor),
		0x09,
		c.StatusEventLogDepth,
		0x25,
	}
}
type MeterTblStatusSupportedGet struct {
	node byte
}

func NewMeterTblStatusSupportedGet() MeterTblStatusSupportedGet {
	return MeterTblStatusSupportedGet{}
}

func (c *MeterTblStatusSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblStatusSupportedGet) Set() error {

	return nil
}

func (c *MeterTblStatusSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterTblMonitor),
		0x07,
		0x25,
	}
}
type MeterTblStatusSupportedReport struct {
    *report
    SupportedOperatingStatus byte
    StatusEventLogDepth byte
    data []byte
}

func NewMeterTblStatusSupportedReport(data []byte) *MeterTblStatusSupportedReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblStatusSupportedReport{
				SupportedOperatingStatus: data[0],
				StatusEventLogDepth: data[1],
        data: data,
    }
}

type MeterTblCurrentDataGet struct {
	node byte
	DatasetRequested byte
}

func NewMeterTblCurrentDataGet() MeterTblCurrentDataGet {
	return MeterTblCurrentDataGet{}
}

func (c *MeterTblCurrentDataGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblCurrentDataGet) Set(DatasetRequested byte,) error {
	c.DatasetRequested = DatasetRequested

	return nil
}

func (c *MeterTblCurrentDataGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MeterTblMonitor),
		0x0C,
		c.DatasetRequested,
		0x25,
	}
}
type MeterTblCurrentDataReport struct {
    *report
    ReportstoFollow byte
    Properties1 byte
    Dataset byte
    Year uint16
    Month byte
    Day byte
    HourLocalTime byte
    MinuteLocalTime byte
    SecondLocalTime byte
    data []byte
}

func NewMeterTblCurrentDataReport(data []byte) *MeterTblCurrentDataReport {
    if len(data) < 10 {
				//may want to change this to just return nil
				for i := len(data); i < 10; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblCurrentDataReport{
				ReportstoFollow: data[0],
				Properties1: data[1],
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

type MeterTblHistoricalDataGet struct {
	node byte
	MaximumReports byte
	HistoricalDatasetRequested byte
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

func NewMeterTblHistoricalDataGet() MeterTblHistoricalDataGet {
	return MeterTblHistoricalDataGet{}
}

func (c *MeterTblHistoricalDataGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblHistoricalDataGet) Set(MaximumReports byte,HistoricalDatasetRequested byte,StartYear uint16,StartMonth byte,StartDay byte,StartHourLocalTime byte,StartMinuteLocalTime byte,StartSecondLocalTime byte,StopYear uint16,StopMonth byte,StopDay byte,StopHourLocalTime byte,StopMinuteLocalTime byte,StopSecondLocalTime byte,) error {
	c.MaximumReports = MaximumReports
	c.HistoricalDatasetRequested = HistoricalDatasetRequested
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

func (c *MeterTblHistoricalDataGet) Encode() []byte {
	StartYearBytes := []byte{}
	binary.BigEndian.PutUint16(StartYearBytes, c.StartYear)
	StopYearBytes := []byte{}
	binary.BigEndian.PutUint16(StopYearBytes, c.StopYear)
	return []byte{
		0x13,
		c.node,
		byte(16),
		byte(MeterTblMonitor),
		0x0E,
		c.MaximumReports,
		c.HistoricalDatasetRequested,
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
type MeterTblHistoricalDataReport struct {
    *report
    ReportstoFollow byte
    Properties1 byte
    Dataset byte
    Year uint16
    Month byte
    Day byte
    HourLocalTime byte
    MinuteLocalTime byte
    SecondLocalTime byte
    data []byte
}

func NewMeterTblHistoricalDataReport(data []byte) *MeterTblHistoricalDataReport {
    if len(data) < 10 {
				//may want to change this to just return nil
				for i := len(data); i < 10; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblHistoricalDataReport{
				ReportstoFollow: data[0],
				Properties1: data[1],
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

type MeterTblReport struct {
    *report
    Properties1 byte
    Properties2 byte
    DatasetSupported byte
    DatasetHistorySupported byte
    DataHistorySupported byte
    data []byte
}

func NewMeterTblReport(data []byte) *MeterTblReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblReport{
				Properties1: data[0],
				Properties2: data[1],
				DatasetSupported: data[2],
				DatasetHistorySupported: data[3],
				DataHistorySupported: data[4],
        data: data,
    }
}

type MeterTblTableCapabilityGet struct {
	node byte
}

func NewMeterTblTableCapabilityGet() MeterTblTableCapabilityGet {
	return MeterTblTableCapabilityGet{}
}

func (c *MeterTblTableCapabilityGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblTableCapabilityGet) Set() error {

	return nil
}

func (c *MeterTblTableCapabilityGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterTblMonitor),
		0x05,
		0x25,
	}
}
type MeterTblTableIdGet struct {
	node byte
}

func NewMeterTblTableIdGet() MeterTblTableIdGet {
	return MeterTblTableIdGet{}
}

func (c *MeterTblTableIdGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblTableIdGet) Set() error {

	return nil
}

func (c *MeterTblTableIdGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterTblMonitor),
		0x03,
		0x25,
	}
}
type MeterTblTableIdReport struct {
    *report
    Properties1 byte
    MeterIDCharacter byte
    data []byte
}

func NewMeterTblTableIdReport(data []byte) *MeterTblTableIdReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblTableIdReport{
				Properties1: data[0],
				MeterIDCharacter: data[1],
        data: data,
    }
}

type MeterTblTablePointAdmNoGet struct {
	node byte
}

func NewMeterTblTablePointAdmNoGet() MeterTblTablePointAdmNoGet {
	return MeterTblTablePointAdmNoGet{}
}

func (c *MeterTblTablePointAdmNoGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblTablePointAdmNoGet) Set() error {

	return nil
}

func (c *MeterTblTablePointAdmNoGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterTblMonitor),
		0x01,
		0x25,
	}
}
type MeterTblTablePointAdmNoReport struct {
    *report
    Properties1 byte
    MeterPointAdmNumberCharacter byte
    data []byte
}

func NewMeterTblTablePointAdmNoReport(data []byte) *MeterTblTablePointAdmNoReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblTablePointAdmNoReport{
				Properties1: data[0],
				MeterPointAdmNumberCharacter: data[1],
        data: data,
    }
}

