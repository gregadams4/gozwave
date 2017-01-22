package commands
import "encoding/binary"
type BasicTariffInfoGet struct {
	node byte
}

func NewBasicTariffInfoGet() BasicTariffInfoGet {
	return BasicTariffInfoGet{}
}

func (c *BasicTariffInfoGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *BasicTariffInfoGet) Set() error {

	return nil
}

func (c *BasicTariffInfoGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(BasicTariffInfo),
		0x01,
		0x25,
	}
}
type BasicTariffInfoReport struct {
    *report
    Properties1 byte
    Properties2 byte
    E1RateConsumptionRegister uint64
    E1TimeforNextRateHours byte
    E1TimeforNextRateMinutes byte
    E1TimeforNextRateSeconds byte
    Properties3 byte
    E2RateConsumptionRegister uint64
    data []byte
}

func NewBasicTariffInfoReport(data []byte) *BasicTariffInfoReport {
    if len(data) < 14 {
				//may want to change this to just return nil
				for i := len(data); i < 14; i++ {
            data = append(data, 0x00)
        }
    }

    return &BasicTariffInfoReport{
				Properties1: data[0],
				Properties2: data[1],
				E1RateConsumptionRegister: binary.BigEndian.Uint64(data[2:6]),
				E1TimeforNextRateHours: data[6],
				E1TimeforNextRateMinutes: data[7],
				E1TimeforNextRateSeconds: data[8],
				Properties3: data[9],
				E2RateConsumptionRegister: binary.BigEndian.Uint64(data[10:14]),
        data: data,
    }
}

