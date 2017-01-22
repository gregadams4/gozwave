package interfaces

import (
	"time"

	"github.com/gregadams4/gozwave/commands"
	"github.com/gregadams4/gozwave/serialapi"
)

type Encodable interface {
	Encode() []byte
}

type Writer interface {
	Write(Encodable) error
	WriteWithTimeout(Encodable, time.Duration) (<-chan *serialapi.Message, error)
	WriteAndWaitForReport(Encodable, time.Duration, byte) (<-chan commands.Report, error)
}

type LoadSaveable interface {
	LoadConfigurationFromFile() error
	SaveConfigurationToFile() error
}
