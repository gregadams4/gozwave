package commands
type LastWorkingRouteSet struct {
	node byte
	NodeID byte
	Repeater1 byte
	Repeater2 byte
	Repeater3 byte
	Repeater4 byte
	Speed byte
}

func NewLastWorkingRouteSet() LastWorkingRouteSet {
	return LastWorkingRouteSet{}
}

func (c *LastWorkingRouteSet) SetNode(node int) {
	c.node = byte(node)
}

func (c *LastWorkingRouteSet) Set(NodeID byte,Repeater1 byte,Repeater2 byte,Repeater3 byte,Repeater4 byte,Speed byte,) error {
	c.NodeID = NodeID
	c.Repeater1 = Repeater1
	c.Repeater2 = Repeater2
	c.Repeater3 = Repeater3
	c.Repeater4 = Repeater4
	c.Speed = Speed

	return nil
}

func (c *LastWorkingRouteSet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(8),
		byte(NetworkManagementInstallationMaintenance),
		0x01,
		c.NodeID,
		c.Repeater1,
		c.Repeater2,
		c.Repeater3,
		c.Repeater4,
		c.Speed,
		0x25,
	}
}
type LastWorkingRouteGet struct {
	node byte
	NodeID byte
}

func NewLastWorkingRouteGet() LastWorkingRouteGet {
	return LastWorkingRouteGet{}
}

func (c *LastWorkingRouteGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *LastWorkingRouteGet) Set(NodeID byte,) error {
	c.NodeID = NodeID

	return nil
}

func (c *LastWorkingRouteGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NetworkManagementInstallationMaintenance),
		0x02,
		c.NodeID,
		0x25,
	}
}
type LastWorkingRouteReport struct {
    *report
    NodeID byte
    Repeater1 byte
    Repeater2 byte
    Repeater3 byte
    Repeater4 byte
    Speed byte
    data []byte
}

func NewLastWorkingRouteReport(data []byte) *LastWorkingRouteReport {
    if len(data) < 6 {
				//may want to change this to just return nil
				for i := len(data); i < 6; i++ {
            data = append(data, 0x00)
        }
    }

    return &LastWorkingRouteReport{
				NodeID: data[0],
				Repeater1: data[1],
				Repeater2: data[2],
				Repeater3: data[3],
				Repeater4: data[4],
				Speed: data[5],
        data: data,
    }
}

type StatisticsGet struct {
	node byte
	NodeID byte
}

func NewStatisticsGet() StatisticsGet {
	return StatisticsGet{}
}

func (c *StatisticsGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *StatisticsGet) Set(NodeID byte,) error {
	c.NodeID = NodeID

	return nil
}

func (c *StatisticsGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(NetworkManagementInstallationMaintenance),
		0x04,
		c.NodeID,
		0x25,
	}
}
type StatisticsReport struct {
    *report
    NodeID byte
    data []byte
}

func NewStatisticsReport(data []byte) *StatisticsReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &StatisticsReport{
				NodeID: data[0],
        data: data,
    }
}

type StatisticsClear struct {
	node byte
}

func NewStatisticsClear() StatisticsClear {
	return StatisticsClear{}
}

func (c *StatisticsClear) SetNode(node int) {
	c.node = byte(node)
}

func (c *StatisticsClear) Set() error {

	return nil
}

func (c *StatisticsClear) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(NetworkManagementInstallationMaintenance),
		0x06,
		0x25,
	}
}
