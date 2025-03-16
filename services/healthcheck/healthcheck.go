package healthcheck

import (
	"time"

	"go.uber.org/fx"
)

// Service interface define functions
// that returns the database connection status
// last time the sync was done and the system status
type Service interface {
	// SetOnlineSince sets the time the system was online
	SetOnlineSince(time.Time)

	// OnlineSince returns the time since the system was online
	OnlineSince() time.Duration
}

// Params defines the dependencies that the healthcheck module needs
type Params struct {
	fx.In
}

type service struct {
	onlineSince time.Time
}

// New returns an implementation of Healthcheck interface
func New(p Params) Service {
	return &service{}
}

func (s *service) SetOnlineSince(t time.Time) {
	s.onlineSince = t
}

func (s *service) OnlineSince() time.Duration {
	return time.Since(s.onlineSince)
}
