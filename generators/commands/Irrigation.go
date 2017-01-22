package commands
import "encoding/binary"
type IrrigationSystemInfoGet struct {
	node byte
}

func NewIrrigationSystemInfoGet() IrrigationSystemInfoGet {
	return IrrigationSystemInfoGet{}
}

func (c *IrrigationSystemInfoGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IrrigationSystemInfoGet) Set() error {

	return nil
}

func (c *IrrigationSystemInfoGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Irrigation),
		0x01,
		0x25,
	}
}
type IrrigationSystemInfoReport struct {
    *report
    Properties1 byte
    TotalNumberofValves byte
    TotalNumberofValveTables byte
    Properties2 byte
    data []byte
}

func NewIrrigationSystemInfoReport(data []byte) *IrrigationSystemInfoReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &IrrigationSystemInfoReport{
				Properties1: data[0],
				TotalNumberofValves: data[1],
				TotalNumberofValveTables: data[2],
				Properties2: data[3],
        data: data,
    }
}

type IrrigationSystemStatusGet struct {
	node byte
}

func NewIrrigationSystemStatusGet() IrrigationSystemStatusGet {
	return IrrigationSystemStatusGet{}
}

func (c *IrrigationSystemStatusGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IrrigationSystemStatusGet) Set() error {

	return nil
}

func (c *IrrigationSystemStatusGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Irrigation),
		0x03,
		0x25,
	}
}
type IrrigationSystemStatusReport struct {
    *report
    SystemVoltage byte
    SensorStatus byte
    Properties1 byte
    FlowValue byte
    Properties2 byte
    PressureValue byte
    ShutoffDuration byte
    SystemErrorStatus byte
    Properties3 byte
    ValveID byte
    data []byte
}

func NewIrrigationSystemStatusReport(data []byte) *IrrigationSystemStatusReport {
    if len(data) < 10 {
				//may want to change this to just return nil
				for i := len(data); i < 10; i++ {
            data = append(data, 0x00)
        }
    }

    return &IrrigationSystemStatusReport{
				SystemVoltage: data[0],
				SensorStatus: data[1],
				Properties1: data[2],
				FlowValue: data[3],
				Properties2: data[4],
				PressureValue: data[5],
				ShutoffDuration: data[6],
				SystemErrorStatus: data[7],
				Properties3: data[8],
				ValveID: data[9],
        data: data,
    }
}

type IrrigationSystemConfigSet struct {
	node byte
	MasterValveDelay byte
	Properties1 byte
	HighPressureThresholdValue byte
	Properties2 byte
	LowPressureThresholdValue byte
	SensorPolarity byte
}

func NewIrrigationSystemConfigSet() IrrigationSystemConfigSet {
	return IrrigationSystemConfigSet{}
}

func (c *IrrigationSystemConfigSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IrrigationSystemConfigSet) Set(MasterValveDelay byte,Properties1 byte,HighPressureThresholdValue byte,Properties2 byte,LowPressureThresholdValue byte,SensorPolarity byte,) error {
	c.MasterValveDelay = MasterValveDelay
	c.Properties1 = Properties1
	c.HighPressureThresholdValue = HighPressureThresholdValue
	c.Properties2 = Properties2
	c.LowPressureThresholdValue = LowPressureThresholdValue
	c.SensorPolarity = SensorPolarity

	return nil
}

func (c *IrrigationSystemConfigSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(8),
		byte(Irrigation),
		0x05,
		c.MasterValveDelay,
		c.Properties1,
		c.HighPressureThresholdValue,
		c.Properties2,
		c.LowPressureThresholdValue,
		c.SensorPolarity,
		0x25,
	}
}
type IrrigationSystemConfigGet struct {
	node byte
}

func NewIrrigationSystemConfigGet() IrrigationSystemConfigGet {
	return IrrigationSystemConfigGet{}
}

func (c *IrrigationSystemConfigGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IrrigationSystemConfigGet) Set() error {

	return nil
}

func (c *IrrigationSystemConfigGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Irrigation),
		0x06,
		0x25,
	}
}
type IrrigationSystemConfigReport struct {
    *report
    MasterValveDelay byte
    Properties1 byte
    HighPressureThresholdValue byte
    Properties2 byte
    LowPressureThresholdValue byte
    SensorPolarity byte
    data []byte
}

func NewIrrigationSystemConfigReport(data []byte) *IrrigationSystemConfigReport {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &IrrigationSystemConfigReport{
				MasterValveDelay: data[0],
				Properties1: data[1],
				HighPressureThresholdValue: data[2],
				Properties2: data[3],
				LowPressureThresholdValue: data[4],
				SensorPolarity: data[5],
        data: data,
    }
}

type IrrigationValveInfoGet struct {
	node byte
	Properties1 byte
	ValveID byte
}

func NewIrrigationValveInfoGet() IrrigationValveInfoGet {
	return IrrigationValveInfoGet{}
}

func (c *IrrigationValveInfoGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IrrigationValveInfoGet) Set(Properties1 byte,ValveID byte,) error {
	c.Properties1 = Properties1
	c.ValveID = ValveID

	return nil
}

func (c *IrrigationValveInfoGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(Irrigation),
		0x08,
		c.Properties1,
		c.ValveID,
		0x25,
	}
}
type IrrigationValveInfoReport struct {
    *report
    Properties1 byte
    ValveID byte
    NominalCurrent byte
    ValveErrorStatus byte
    data []byte
}

func NewIrrigationValveInfoReport(data []byte) *IrrigationValveInfoReport {
    if len(data) < 4 {
				//may want to change this to just return nil
				for i := len(data); i < 4; i++ {
            data = append(data, 0x00)
        }
    }

    return &IrrigationValveInfoReport{
				Properties1: data[0],
				ValveID: data[1],
				NominalCurrent: data[2],
				ValveErrorStatus: data[3],
        data: data,
    }
}

type IrrigationValveConfigSet struct {
	node byte
	Properties1 byte
	ValveID byte
	NominalCurrentHighThreshold byte
	NominalCurrentLowThreshold byte
	Properties2 byte
	MaximumFlowValue byte
	Properties3 byte
	FlowHighThresholdValue byte
	Properties4 byte
	FlowLowThresholdValue byte
	SensorUsage byte
}

func NewIrrigationValveConfigSet() IrrigationValveConfigSet {
	return IrrigationValveConfigSet{}
}

func (c *IrrigationValveConfigSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IrrigationValveConfigSet) Set(Properties1 byte,ValveID byte,NominalCurrentHighThreshold byte,NominalCurrentLowThreshold byte,Properties2 byte,MaximumFlowValue byte,Properties3 byte,FlowHighThresholdValue byte,Properties4 byte,FlowLowThresholdValue byte,SensorUsage byte,) error {
	c.Properties1 = Properties1
	c.ValveID = ValveID
	c.NominalCurrentHighThreshold = NominalCurrentHighThreshold
	c.NominalCurrentLowThreshold = NominalCurrentLowThreshold
	c.Properties2 = Properties2
	c.MaximumFlowValue = MaximumFlowValue
	c.Properties3 = Properties3
	c.FlowHighThresholdValue = FlowHighThresholdValue
	c.Properties4 = Properties4
	c.FlowLowThresholdValue = FlowLowThresholdValue
	c.SensorUsage = SensorUsage

	return nil
}

func (c *IrrigationValveConfigSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(13),
		byte(Irrigation),
		0x0A,
		c.Properties1,
		c.ValveID,
		c.NominalCurrentHighThreshold,
		c.NominalCurrentLowThreshold,
		c.Properties2,
		c.MaximumFlowValue,
		c.Properties3,
		c.FlowHighThresholdValue,
		c.Properties4,
		c.FlowLowThresholdValue,
		c.SensorUsage,
		0x25,
	}
}
type IrrigationValveConfigGet struct {
	node byte
	Properties1 byte
	ValveID byte
}

func NewIrrigationValveConfigGet() IrrigationValveConfigGet {
	return IrrigationValveConfigGet{}
}

func (c *IrrigationValveConfigGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IrrigationValveConfigGet) Set(Properties1 byte,ValveID byte,) error {
	c.Properties1 = Properties1
	c.ValveID = ValveID

	return nil
}

func (c *IrrigationValveConfigGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(4),
		byte(Irrigation),
		0x0B,
		c.Properties1,
		c.ValveID,
		0x25,
	}
}
type IrrigationValveConfigReport struct {
    *report
    Properties1 byte
    ValveID byte
    NominalCurrentHighThreshold byte
    NominalCurrentLowThreshold byte
    Properties2 byte
    MaximumFlowValue byte
    Properties3 byte
    FlowHighThresholdValue byte
    Properties4 byte
    FlowLowThresholdValue byte
    SensorUsage byte
    data []byte
}

func NewIrrigationValveConfigReport(data []byte) *IrrigationValveConfigReport {
    if len(data) < 11 {
				//may want to change this to just return nil
				for i := len(data); i < 11; i++ {
            data = append(data, 0x00)
        }
    }

    return &IrrigationValveConfigReport{
				Properties1: data[0],
				ValveID: data[1],
				NominalCurrentHighThreshold: data[2],
				NominalCurrentLowThreshold: data[3],
				Properties2: data[4],
				MaximumFlowValue: data[5],
				Properties3: data[6],
				FlowHighThresholdValue: data[7],
				Properties4: data[8],
				FlowLowThresholdValue: data[9],
				SensorUsage: data[10],
        data: data,
    }
}

type IrrigationValveRun struct {
	node byte
	Properties1 byte
	ValveID byte
	Duration uint16
}

func NewIrrigationValveRun() IrrigationValveRun {
	return IrrigationValveRun{}
}

func (c *IrrigationValveRun) SetNode(node int) {
	c.node = byte(node)
}

func (c *IrrigationValveRun) Set(Properties1 byte,ValveID byte,Duration uint16,) error {
	c.Properties1 = Properties1
	c.ValveID = ValveID
	c.Duration = Duration

	return nil
}

func (c *IrrigationValveRun) Encode() []byte {
	DurationBytes := []byte{}
	binary.BigEndian.PutUint16(DurationBytes, c.Duration)
	return []byte{
		0x13,
		c.node,
		byte(5),
		byte(Irrigation),
		0x0D,
		c.Properties1,
		c.ValveID,
		DurationBytes[0],
		DurationBytes[1],
		0x25,
	}
}
type IrrigationValveTableSet struct {
	node byte
	ValveTableID byte
	Vg1 []IrrigationValveTableSetvg1
}
type IrrigationValveTableSetvg1 struct {
	ValveID byte
	Duration uint16
}

func NewIrrigationValveTableSet() IrrigationValveTableSet {
	return IrrigationValveTableSet{}
}

func (c *IrrigationValveTableSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IrrigationValveTableSet) Set(ValveTableID byte,Vg1 []IrrigationValveTableSetvg1,) error {
	c.ValveTableID = ValveTableID
	c.Vg1 = Vg1

	return nil
}

func (c *IrrigationValveTableSet) Encode() []byte {
	var data = []byte{
		0x13,
		c.node,
		byte(5),
		byte(Irrigation),
		0x0E,
		c.ValveTableID,
	}
	for _, v := range c.Vg1 {
		data = append(data, v.ValveID)
		DurationBytes := []byte{}
		binary.BigEndian.PutUint16(DurationBytes, v.Duration)
		data = append(data, DurationBytes...)
	}
	data = append(data, 0x25)
	return data
	
}
type IrrigationValveTableGet struct {
	node byte
	ValveTableID byte
}

func NewIrrigationValveTableGet() IrrigationValveTableGet {
	return IrrigationValveTableGet{}
}

func (c *IrrigationValveTableGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *IrrigationValveTableGet) Set(ValveTableID byte,) error {
	c.ValveTableID = ValveTableID

	return nil
}

func (c *IrrigationValveTableGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Irrigation),
		0x0F,
		c.ValveTableID,
		0x25,
	}
}
type IrrigationValveTableReport struct {
    *report
    ValveTableID byte
    data []byte
}

func NewIrrigationValveTableReport(data []byte) *IrrigationValveTableReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &IrrigationValveTableReport{
				ValveTableID: data[0],
        data: data,
    }
}

type IrrigationValveTableRun struct {
	node byte
	ValveTableID byte
}

func NewIrrigationValveTableRun() IrrigationValveTableRun {
	return IrrigationValveTableRun{}
}

func (c *IrrigationValveTableRun) SetNode(node int) {
	c.node = byte(node)
}

func (c *IrrigationValveTableRun) Set(ValveTableID byte,) error {
	c.ValveTableID = ValveTableID

	return nil
}

func (c *IrrigationValveTableRun) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Irrigation),
		0x11,
		c.ValveTableID,
		0x25,
	}
}
type IrrigationSystemShutoff struct {
	node byte
	Duration byte
}

func NewIrrigationSystemShutoff() IrrigationSystemShutoff {
	return IrrigationSystemShutoff{}
}

func (c *IrrigationSystemShutoff) SetNode(node int) {
	c.node = byte(node)
}

func (c *IrrigationSystemShutoff) Set(Duration byte,) error {
	c.Duration = Duration

	return nil
}

func (c *IrrigationSystemShutoff) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Irrigation),
		0x12,
		c.Duration,
		0x25,
	}
}
