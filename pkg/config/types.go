package config

import (
	eh "github.com/looplab/eventhorizon"
	"github.com/sirupsen/logrus"
)

// Config represents the configuration for the service
type Config struct {
	Logger *logrus.Entry

	EventStore eh.EventStore
	EventBus   eh.EventBus
	Repo       eh.ReadWriteRepo
}
