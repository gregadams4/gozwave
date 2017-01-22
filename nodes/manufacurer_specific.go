package nodes

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	commands "github.com/gregadams4/gozwave/commands"
)

type ManufacurerSpecific struct {
}

func (n *Node) RequestManufacturerSpecific() (*commands.ManufacturerSpecificReport, error) {
	cmd := commands.NewManufacturerSpecificGet()

	cmd.SetNode(n.Id)

	logrus.Infof("sending manufacturer")
	t, err := n.connection.WriteAndWaitForReport(&cmd, time.Second*10, 0x05) // Request node information
	if err != nil {
		return nil, err
	}

	report := <-t
	logrus.Infof("received manfacturer")
	switch cmd := report.(type) {
	case *commands.ManufacturerSpecificReport:
		logrus.Infof("%v", cmd)
		return cmd, nil
	default:
		return nil, fmt.Errorf("ManufacurerSpecific: Wrong report type: %v", report)
	}

	return nil, fmt.Errorf("Failed to get ManufacurerSpecific")
}
