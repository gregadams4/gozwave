package commands
type ChimneyFanAlarmLogGet struct {
	node byte
}

func NewChimneyFanAlarmLogGet() ChimneyFanAlarmLogGet {
	return ChimneyFanAlarmLogGet{}
}

func (c *ChimneyFanAlarmLogGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanAlarmLogGet) Set() error {

	return nil
}

func (c *ChimneyFanAlarmLogGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x20,
		0x25,
	}
}
type ChimneyFanAlarmLogReport struct {
    *report
    AlarmEvent1 byte
    AlarmEvent2 byte
    AlarmEvent3 byte
    AlarmEvent4 byte
    AlarmEvent5 byte
    data []byte
}

func NewChimneyFanAlarmLogReport(data []byte) *ChimneyFanAlarmLogReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanAlarmLogReport{
				AlarmEvent1: data[0],
				AlarmEvent2: data[1],
				AlarmEvent3: data[2],
				AlarmEvent4: data[3],
				AlarmEvent5: data[4],
        data: data,
    }
}

type ChimneyFanAlarmLogSet struct {
	node byte
	Message byte
}

func NewChimneyFanAlarmLogSet() ChimneyFanAlarmLogSet {
	return ChimneyFanAlarmLogSet{}
}

func (c *ChimneyFanAlarmLogSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanAlarmLogSet) Set(Message byte,) error {
	c.Message = Message

	return nil
}

func (c *ChimneyFanAlarmLogSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ChimneyFan),
		0x1F,
		c.Message,
		0x25,
	}
}
type ChimneyFanAlarmStatusGet struct {
	node byte
}

func NewChimneyFanAlarmStatusGet() ChimneyFanAlarmStatusGet {
	return ChimneyFanAlarmStatusGet{}
}

func (c *ChimneyFanAlarmStatusGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanAlarmStatusGet) Set() error {

	return nil
}

func (c *ChimneyFanAlarmStatusGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x23,
		0x25,
	}
}
type ChimneyFanAlarmStatusReport struct {
    *report
    AlarmStatus byte
    data []byte
}

func NewChimneyFanAlarmStatusReport(data []byte) *ChimneyFanAlarmStatusReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanAlarmStatusReport{
				AlarmStatus: data[0],
        data: data,
    }
}

type ChimneyFanAlarmStatusSet struct {
	node byte
	Message byte
}

func NewChimneyFanAlarmStatusSet() ChimneyFanAlarmStatusSet {
	return ChimneyFanAlarmStatusSet{}
}

func (c *ChimneyFanAlarmStatusSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanAlarmStatusSet) Set(Message byte,) error {
	c.Message = Message

	return nil
}

func (c *ChimneyFanAlarmStatusSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ChimneyFan),
		0x22,
		c.Message,
		0x25,
	}
}
type ChimneyFanAlarmTempGet struct {
	node byte
}

func NewChimneyFanAlarmTempGet() ChimneyFanAlarmTempGet {
	return ChimneyFanAlarmTempGet{}
}

func (c *ChimneyFanAlarmTempGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanAlarmTempGet) Set() error {

	return nil
}

func (c *ChimneyFanAlarmTempGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x0E,
		0x25,
	}
}
type ChimneyFanAlarmTempReport struct {
    *report
    Properties1 byte
    Value byte
    data []byte
}

func NewChimneyFanAlarmTempReport(data []byte) *ChimneyFanAlarmTempReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanAlarmTempReport{
				Properties1: data[0],
				Value: data[1],
        data: data,
    }
}

type ChimneyFanAlarmTempSet struct {
	node byte
	Properties1 byte
	Value byte
}

func NewChimneyFanAlarmTempSet() ChimneyFanAlarmTempSet {
	return ChimneyFanAlarmTempSet{}
}

func (c *ChimneyFanAlarmTempSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanAlarmTempSet) Set(Properties1 byte,Value byte,) error {
	c.Properties1 = Properties1
	c.Value = Value

	return nil
}

func (c *ChimneyFanAlarmTempSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ChimneyFan),
		0x0D,
		c.Properties1,
		c.Value,
		0x25,
	}
}
type ChimneyFanBoostTimeGet struct {
	node byte
}

func NewChimneyFanBoostTimeGet() ChimneyFanBoostTimeGet {
	return ChimneyFanBoostTimeGet{}
}

func (c *ChimneyFanBoostTimeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanBoostTimeGet) Set() error {

	return nil
}

func (c *ChimneyFanBoostTimeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x11,
		0x25,
	}
}
type ChimneyFanBoostTimeReport struct {
    *report
    Time byte
    data []byte
}

func NewChimneyFanBoostTimeReport(data []byte) *ChimneyFanBoostTimeReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanBoostTimeReport{
				Time: data[0],
        data: data,
    }
}

type ChimneyFanBoostTimeSet struct {
	node byte
	Time byte
}

func NewChimneyFanBoostTimeSet() ChimneyFanBoostTimeSet {
	return ChimneyFanBoostTimeSet{}
}

func (c *ChimneyFanBoostTimeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanBoostTimeSet) Set(Time byte,) error {
	c.Time = Time

	return nil
}

func (c *ChimneyFanBoostTimeSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ChimneyFan),
		0x10,
		c.Time,
		0x25,
	}
}
type ChimneyFanDefaultSet struct {
	node byte
}

func NewChimneyFanDefaultSet() ChimneyFanDefaultSet {
	return ChimneyFanDefaultSet{}
}

func (c *ChimneyFanDefaultSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanDefaultSet) Set() error {

	return nil
}

func (c *ChimneyFanDefaultSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x28,
		0x25,
	}
}
type ChimneyFanMinSpeedGet struct {
	node byte
}

func NewChimneyFanMinSpeedGet() ChimneyFanMinSpeedGet {
	return ChimneyFanMinSpeedGet{}
}

func (c *ChimneyFanMinSpeedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanMinSpeedGet) Set() error {

	return nil
}

func (c *ChimneyFanMinSpeedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x26,
		0x25,
	}
}
type ChimneyFanMinSpeedReport struct {
    *report
    MinSpeed byte
    data []byte
}

func NewChimneyFanMinSpeedReport(data []byte) *ChimneyFanMinSpeedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanMinSpeedReport{
				MinSpeed: data[0],
        data: data,
    }
}

type ChimneyFanMinSpeedSet struct {
	node byte
	MinSpeed byte
}

func NewChimneyFanMinSpeedSet() ChimneyFanMinSpeedSet {
	return ChimneyFanMinSpeedSet{}
}

func (c *ChimneyFanMinSpeedSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanMinSpeedSet) Set(MinSpeed byte,) error {
	c.MinSpeed = MinSpeed

	return nil
}

func (c *ChimneyFanMinSpeedSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ChimneyFan),
		0x25,
		c.MinSpeed,
		0x25,
	}
}
type ChimneyFanModeGet struct {
	node byte
}

func NewChimneyFanModeGet() ChimneyFanModeGet {
	return ChimneyFanModeGet{}
}

func (c *ChimneyFanModeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanModeGet) Set() error {

	return nil
}

func (c *ChimneyFanModeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x17,
		0x25,
	}
}
type ChimneyFanModeReport struct {
    *report
    Mode byte
    data []byte
}

func NewChimneyFanModeReport(data []byte) *ChimneyFanModeReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanModeReport{
				Mode: data[0],
        data: data,
    }
}

type ChimneyFanModeSet struct {
	node byte
	Mode byte
}

func NewChimneyFanModeSet() ChimneyFanModeSet {
	return ChimneyFanModeSet{}
}

func (c *ChimneyFanModeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanModeSet) Set(Mode byte,) error {
	c.Mode = Mode

	return nil
}

func (c *ChimneyFanModeSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ChimneyFan),
		0x16,
		c.Mode,
		0x25,
	}
}
type ChimneyFanSetupGet struct {
	node byte
}

func NewChimneyFanSetupGet() ChimneyFanSetupGet {
	return ChimneyFanSetupGet{}
}

func (c *ChimneyFanSetupGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanSetupGet) Set() error {

	return nil
}

func (c *ChimneyFanSetupGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x1A,
		0x25,
	}
}
type ChimneyFanSetupReport struct {
    *report
    Mode byte
    BoostTime byte
    StopTime byte
    MinSpeed byte
    Properties1 byte
    StartTemperature byte
    Properties2 byte
    StopTemperature byte
    Properties3 byte
    AlarmTemperatureValue byte
    data []byte
}

func NewChimneyFanSetupReport(data []byte) *ChimneyFanSetupReport {
    if len(data) < 10 {
				//may want to change this to just return nil
				for i := len(data); i < 10; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanSetupReport{
				Mode: data[0],
				BoostTime: data[1],
				StopTime: data[2],
				MinSpeed: data[3],
				Properties1: data[4],
				StartTemperature: data[5],
				Properties2: data[6],
				StopTemperature: data[7],
				Properties3: data[8],
				AlarmTemperatureValue: data[9],
        data: data,
    }
}

type ChimneyFanSetupSet struct {
	node byte
	Mode byte
	BoostTime byte
	StopTime byte
	MinSpeed byte
	Properties1 byte
	StartTemperature byte
	Properties2 byte
	StopTemperature byte
	Properties3 byte
	AlarmTemperatureValue byte
}

func NewChimneyFanSetupSet() ChimneyFanSetupSet {
	return ChimneyFanSetupSet{}
}

func (c *ChimneyFanSetupSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanSetupSet) Set(Mode byte,BoostTime byte,StopTime byte,MinSpeed byte,Properties1 byte,StartTemperature byte,Properties2 byte,StopTemperature byte,Properties3 byte,AlarmTemperatureValue byte,) error {
	c.Mode = Mode
	c.BoostTime = BoostTime
	c.StopTime = StopTime
	c.MinSpeed = MinSpeed
	c.Properties1 = Properties1
	c.StartTemperature = StartTemperature
	c.Properties2 = Properties2
	c.StopTemperature = StopTemperature
	c.Properties3 = Properties3
	c.AlarmTemperatureValue = AlarmTemperatureValue

	return nil
}

func (c *ChimneyFanSetupSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(12),
		byte(ChimneyFan),
		0x19,
		c.Mode,
		c.BoostTime,
		c.StopTime,
		c.MinSpeed,
		c.Properties1,
		c.StartTemperature,
		c.Properties2,
		c.StopTemperature,
		c.Properties3,
		c.AlarmTemperatureValue,
		0x25,
	}
}
type ChimneyFanSpeedGet struct {
	node byte
}

func NewChimneyFanSpeedGet() ChimneyFanSpeedGet {
	return ChimneyFanSpeedGet{}
}

func (c *ChimneyFanSpeedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanSpeedGet) Set() error {

	return nil
}

func (c *ChimneyFanSpeedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x05,
		0x25,
	}
}
type ChimneyFanSpeedReport struct {
    *report
    Speed byte
    data []byte
}

func NewChimneyFanSpeedReport(data []byte) *ChimneyFanSpeedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanSpeedReport{
				Speed: data[0],
        data: data,
    }
}

type ChimneyFanSpeedSet struct {
	node byte
	Speed byte
}

func NewChimneyFanSpeedSet() ChimneyFanSpeedSet {
	return ChimneyFanSpeedSet{}
}

func (c *ChimneyFanSpeedSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanSpeedSet) Set(Speed byte,) error {
	c.Speed = Speed

	return nil
}

func (c *ChimneyFanSpeedSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ChimneyFan),
		0x04,
		c.Speed,
		0x25,
	}
}
type ChimneyFanStartTempGet struct {
	node byte
}

func NewChimneyFanStartTempGet() ChimneyFanStartTempGet {
	return ChimneyFanStartTempGet{}
}

func (c *ChimneyFanStartTempGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanStartTempGet) Set() error {

	return nil
}

func (c *ChimneyFanStartTempGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x08,
		0x25,
	}
}
type ChimneyFanStartTempReport struct {
    *report
    Properties1 byte
    Value byte
    data []byte
}

func NewChimneyFanStartTempReport(data []byte) *ChimneyFanStartTempReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanStartTempReport{
				Properties1: data[0],
				Value: data[1],
        data: data,
    }
}

type ChimneyFanStartTempSet struct {
	node byte
	Properties1 byte
	Value byte
}

func NewChimneyFanStartTempSet() ChimneyFanStartTempSet {
	return ChimneyFanStartTempSet{}
}

func (c *ChimneyFanStartTempSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanStartTempSet) Set(Properties1 byte,Value byte,) error {
	c.Properties1 = Properties1
	c.Value = Value

	return nil
}

func (c *ChimneyFanStartTempSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ChimneyFan),
		0x07,
		c.Properties1,
		c.Value,
		0x25,
	}
}
type ChimneyFanStateGet struct {
	node byte
}

func NewChimneyFanStateGet() ChimneyFanStateGet {
	return ChimneyFanStateGet{}
}

func (c *ChimneyFanStateGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanStateGet) Set() error {

	return nil
}

func (c *ChimneyFanStateGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x02,
		0x25,
	}
}
type ChimneyFanStateReport struct {
    *report
    State byte
    data []byte
}

func NewChimneyFanStateReport(data []byte) *ChimneyFanStateReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanStateReport{
				State: data[0],
        data: data,
    }
}

type ChimneyFanStateSet struct {
	node byte
	State byte
}

func NewChimneyFanStateSet() ChimneyFanStateSet {
	return ChimneyFanStateSet{}
}

func (c *ChimneyFanStateSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanStateSet) Set(State byte,) error {
	c.State = State

	return nil
}

func (c *ChimneyFanStateSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ChimneyFan),
		0x01,
		c.State,
		0x25,
	}
}
type ChimneyFanStatusGet struct {
	node byte
}

func NewChimneyFanStatusGet() ChimneyFanStatusGet {
	return ChimneyFanStatusGet{}
}

func (c *ChimneyFanStatusGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanStatusGet) Set() error {

	return nil
}

func (c *ChimneyFanStatusGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x1D,
		0x25,
	}
}
type ChimneyFanStatusReport struct {
    *report
    State byte
    Speed byte
    AlarmStatus byte
    Properties1 byte
    Value byte
    data []byte
}

func NewChimneyFanStatusReport(data []byte) *ChimneyFanStatusReport {
    if len(data) < 5 {
				//may want to change this to just return nil
				for i := len(data); i < 5; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanStatusReport{
				State: data[0],
				Speed: data[1],
				AlarmStatus: data[2],
				Properties1: data[3],
				Value: data[4],
        data: data,
    }
}

type ChimneyFanStopTempGet struct {
	node byte
}

func NewChimneyFanStopTempGet() ChimneyFanStopTempGet {
	return ChimneyFanStopTempGet{}
}

func (c *ChimneyFanStopTempGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanStopTempGet) Set() error {

	return nil
}

func (c *ChimneyFanStopTempGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x0B,
		0x25,
	}
}
type ChimneyFanStopTempReport struct {
    *report
    Properties1 byte
    Value byte
    data []byte
}

func NewChimneyFanStopTempReport(data []byte) *ChimneyFanStopTempReport {
    if len(data) < 2 {
				//may want to change this to just return nil
				for i := len(data); i < 2; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanStopTempReport{
				Properties1: data[0],
				Value: data[1],
        data: data,
    }
}

type ChimneyFanStopTempSet struct {
	node byte
	Properties1 byte
	Value byte
}

func NewChimneyFanStopTempSet() ChimneyFanStopTempSet {
	return ChimneyFanStopTempSet{}
}

func (c *ChimneyFanStopTempSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanStopTempSet) Set(Properties1 byte,Value byte,) error {
	c.Properties1 = Properties1
	c.Value = Value

	return nil
}

func (c *ChimneyFanStopTempSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(ChimneyFan),
		0x0A,
		c.Properties1,
		c.Value,
		0x25,
	}
}
type ChimneyFanStopTimeGet struct {
	node byte
}

func NewChimneyFanStopTimeGet() ChimneyFanStopTimeGet {
	return ChimneyFanStopTimeGet{}
}

func (c *ChimneyFanStopTimeGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanStopTimeGet) Set() error {

	return nil
}

func (c *ChimneyFanStopTimeGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(ChimneyFan),
		0x14,
		0x25,
	}
}
type ChimneyFanStopTimeReport struct {
    *report
    Time byte
    data []byte
}

func NewChimneyFanStopTimeReport(data []byte) *ChimneyFanStopTimeReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &ChimneyFanStopTimeReport{
				Time: data[0],
        data: data,
    }
}

type ChimneyFanStopTimeSet struct {
	node byte
	Time byte
}

func NewChimneyFanStopTimeSet() ChimneyFanStopTimeSet {
	return ChimneyFanStopTimeSet{}
}

func (c *ChimneyFanStopTimeSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *ChimneyFanStopTimeSet) Set(Time byte,) error {
	c.Time = Time

	return nil
}

func (c *ChimneyFanStopTimeSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(ChimneyFan),
		0x13,
		c.Time,
		0x25,
	}
}
