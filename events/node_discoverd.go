package events

import "github.com/gregadams4/gozwave/functions"

type NodeDiscoverd struct {
	Address int

	functions.FuncGetNodeProtocolInfo
}
