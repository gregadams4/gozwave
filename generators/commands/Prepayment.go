package commands
import "encoding/binary"
type PrepaymentBalanceGet struct {
	node byte
	Properties1 byte
}

func NewPrepaymentBalanceGet() PrepaymentBalanceGet {
	return PrepaymentBalanceGet{}
}

func (c *PrepaymentBalanceGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *PrepaymentBalanceGet) Set(Properties1 byte,) error {
	c.Properties1 = Properties1

	return nil
}

func (c *PrepaymentBalanceGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(3),
		byte(Prepayment),
		0x01,
		c.Properties1,
		0x25,
	}
}
type PrepaymentBalanceReport struct {
    *report
    Properties1 byte
    Properties2 byte
    BalanceValue uint64
    Properties3 byte
    Debt uint64
    Properties4 byte
    EmerCredit uint64
    Currency byte
    DebtRecoveryPercentage byte
    data []byte
}

func NewPrepaymentBalanceReport(data []byte) *PrepaymentBalanceReport {
    if len(data) < 18 {
				//may want to change this to just return nil
				for i := len(data); i < 18; i++ {
            data = append(data, 0x00)
        }
    }

    return &PrepaymentBalanceReport{
				Properties1: data[0],
				Properties2: data[1],
				BalanceValue: binary.BigEndian.Uint64(data[2:6]),
				Properties3: data[6],
				Debt: binary.BigEndian.Uint64(data[7:11]),
				Properties4: data[11],
				EmerCredit: binary.BigEndian.Uint64(data[12:16]),
				Currency: data[16],
				DebtRecoveryPercentage: data[17],
        data: data,
    }
}

type PrepaymentSupportedGet struct {
	node byte
}

func NewPrepaymentSupportedGet() PrepaymentSupportedGet {
	return PrepaymentSupportedGet{}
}

func (c *PrepaymentSupportedGet) SetNode(node int) {
	c.node = byte(node)
}

func (c *PrepaymentSupportedGet) Set() error {

	return nil
}

func (c *PrepaymentSupportedGet) Encode() []byte {
	return []byte{
		0x13,
		c.node,
		byte(2),
		byte(Prepayment),
		0x03,
		0x25,
	}
}
type PrepaymentSupportedReport struct {
    *report
    Properties1 byte
    data []byte
}

func NewPrepaymentSupportedReport(data []byte) *PrepaymentSupportedReport {
    if len(data) < 1 {
				//may want to change this to just return nil
				for i := len(data); i < 1; i++ {
            data = append(data, 0x00)
        }
    }

    return &PrepaymentSupportedReport{
				Properties1: data[0],
        data: data,
    }
}

