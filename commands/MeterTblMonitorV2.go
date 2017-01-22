package commands
import "encoding/binary"
type MeterTblStatusReportV2 struct {
    *report
    Reportstofollow byte
    CurrentOperatingStatus byte
    data []byte
}

func NewMeterTblStatusReportV2(data []byte) *MeterTblStatusReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblStatusReportV2{
				Reportstofollow: data[0],
				CurrentOperatingStatus: data[1],
        data: data,
    }
}

type MeterTblStatusDateGetV2 struct {
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

func NewMeterTblStatusDateGetV2() MeterTblStatusDateGetV2 {
	return MeterTblStatusDateGetV2{}
}

func (c *MeterTblStatusDateGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblStatusDateGetV2) Set(MaximumReports byte,StartYear uint16,StartMonth byte,StartDay byte,StartHourLocalTime byte,StartMinuteLocalTime byte,StartSecondLocalTime byte,StopYear uint16,StopMonth byte,StopDay byte,StopHourLocalTime byte,StopMinuteLocalTime byte,StopSecondLocalTime byte,) error {
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

func (c *MeterTblStatusDateGetV2) Encode() []byte {
	StartYearBytes := []byte{}
	binary.BigEndian.PutUint16(StartYearBytes, c.StartYear)
	StopYearBytes := []byte{}
	binary.BigEndian.PutUint16(StopYearBytes, c.StopYear)
	return []byte{
		0x13,
		c.node,
		byte(15),
		byte(MeterTblMonitorV2),
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
type MeterTblStatusDepthGetV2 struct {
	node byte
	StatusEventLogDepth byte
}

func NewMeterTblStatusDepthGetV2() MeterTblStatusDepthGetV2 {
	return MeterTblStatusDepthGetV2{}
}

func (c *MeterTblStatusDepthGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblStatusDepthGetV2) Set(StatusEventLogDepth byte,) error {
	c.StatusEventLogDepth = StatusEventLogDepth

	return nil
}

func (c *MeterTblStatusDepthGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MeterTblMonitorV2),
		0x09,
		c.StatusEventLogDepth,
		0x25,
	}
}
type MeterTblStatusSupportedGetV2 struct {
	node byte
}

func NewMeterTblStatusSupportedGetV2() MeterTblStatusSupportedGetV2 {
	return MeterTblStatusSupportedGetV2{}
}

func (c *MeterTblStatusSupportedGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblStatusSupportedGetV2) Set() error {

	return nil
}

func (c *MeterTblStatusSupportedGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterTblMonitorV2),
		0x07,
		0x25,
	}
}
type MeterTblStatusSupportedReportV2 struct {
    *report
    SupportedOperatingStatus byte
    StatusEventLogDepth byte
    data []byte
}

func NewMeterTblStatusSupportedReportV2(data []byte) *MeterTblStatusSupportedReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblStatusSupportedReportV2{
				SupportedOperatingStatus: data[0],
				StatusEventLogDepth: data[1],
        data: data,
    }
}

type MeterTblCurrentDataGetV2 struct {
	node byte
	DatasetRequested byte
}

func NewMeterTblCurrentDataGetV2() MeterTblCurrentDataGetV2 {
	return MeterTblCurrentDataGetV2{}
}

func (c *MeterTblCurrentDataGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblCurrentDataGetV2) Set(DatasetRequested byte,) error {
	c.DatasetRequested = DatasetRequested

	return nil
}

func (c *MeterTblCurrentDataGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(MeterTblMonitorV2),
		0x0C,
		c.DatasetRequested,
		0x25,
	}
}
type MeterTblCurrentDataReportV2 struct {
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

func NewMeterTblCurrentDataReportV2(data []byte) *MeterTblCurrentDataReportV2 {
    if len(data) < 10 {
				//may want to change this to just return nil
				for i := len(data); i < 10; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblCurrentDataReportV2{
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

type MeterTblHistoricalDataGetV2 struct {
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

func NewMeterTblHistoricalDataGetV2() MeterTblHistoricalDataGetV2 {
	return MeterTblHistoricalDataGetV2{}
}

func (c *MeterTblHistoricalDataGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblHistoricalDataGetV2) Set(MaximumReports byte,HistoricalDatasetRequested byte,StartYear uint16,StartMonth byte,StartDay byte,StartHourLocalTime byte,StartMinuteLocalTime byte,StartSecondLocalTime byte,StopYear uint16,StopMonth byte,StopDay byte,StopHourLocalTime byte,StopMinuteLocalTime byte,StopSecondLocalTime byte,) error {
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

func (c *MeterTblHistoricalDataGetV2) Encode() []byte {
	StartYearBytes := []byte{}
	binary.BigEndian.PutUint16(StartYearBytes, c.StartYear)
	StopYearBytes := []byte{}
	binary.BigEndian.PutUint16(StopYearBytes, c.StopYear)
	return []byte{
		0x13,
		c.node,
		byte(16),
		byte(MeterTblMonitorV2),
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
type MeterTblHistoricalDataReportV2 struct {
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

func NewMeterTblHistoricalDataReportV2(data []byte) *MeterTblHistoricalDataReportV2 {
    if len(data) < 10 {
				//may want to change this to just return nil
				for i := len(data); i < 10; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblHistoricalDataReportV2{
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

type MeterTblReportV2 struct {
    *report
    Properties1 byte
    Properties2 byte
    DatasetSupported byte
    DatasetHistorySupported byte
    DataHistorySupported byte
    data []byte
}

func NewMeterTblReportV2(data []byte) *MeterTblReportV2 {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblReportV2{
				Properties1: data[0],
				Properties2: data[1],
				DatasetSupported: data[2],
				DatasetHistorySupported: data[3],
				DataHistorySupported: data[4],
        data: data,
    }
}

type MeterTblTableCapabilityGetV2 struct {
	node byte
}

func NewMeterTblTableCapabilityGetV2() MeterTblTableCapabilityGetV2 {
	return MeterTblTableCapabilityGetV2{}
}

func (c *MeterTblTableCapabilityGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblTableCapabilityGetV2) Set() error {

	return nil
}

func (c *MeterTblTableCapabilityGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterTblMonitorV2),
		0x05,
		0x25,
	}
}
type MeterTblTableIdGetV2 struct {
	node byte
}

func NewMeterTblTableIdGetV2() MeterTblTableIdGetV2 {
	return MeterTblTableIdGetV2{}
}

func (c *MeterTblTableIdGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblTableIdGetV2) Set() error {

	return nil
}

func (c *MeterTblTableIdGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterTblMonitorV2),
		0x03,
		0x25,
	}
}
type MeterTblTableIdReportV2 struct {
    *report
    Properties1 byte
    MeterIDCharacter byte
    data []byte
}

func NewMeterTblTableIdReportV2(data []byte) *MeterTblTableIdReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblTableIdReportV2{
				Properties1: data[0],
				MeterIDCharacter: data[1],
        data: data,
    }
}

type MeterTblTablePointAdmNoGetV2 struct {
	node byte
}

func NewMeterTblTablePointAdmNoGetV2() MeterTblTablePointAdmNoGetV2 {
	return MeterTblTablePointAdmNoGetV2{}
}

func (c *MeterTblTablePointAdmNoGetV2) SetNode(node int) {
	c.node = byte(node)
}

func (c *MeterTblTablePointAdmNoGetV2) Set() error {

	return nil
}

func (c *MeterTblTablePointAdmNoGetV2) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(MeterTblMonitorV2),
		0x01,
		0x25,
	}
}
type MeterTblTablePointAdmNoReportV2 struct {
    *report
    Properties1 byte
    MeterPointAdmNumberCharacter byte
    data []byte
}

func NewMeterTblTablePointAdmNoReportV2(data []byte) *MeterTblTablePointAdmNoReportV2 {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &MeterTblTablePointAdmNoReportV2{
				Properties1: data[0],
				MeterPointAdmNumberCharacter: data[1],
        data: data,
    }
}

