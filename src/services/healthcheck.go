package services

import (
	"errors"
	"time"
)

// Service interface define functions
// that returns the database connection status
// last time the sync was done and the system status
type HealthcheckService interface {
	// SetOnlineSince sets the time the system was online
	SetOnlineSince(time.Time)

	// OnlineSince returns the time since the system was online
	OnlineSince() (time.Duration, error)
}

type service struct {
	onlineSince *time.Time
}

// New returns an implementation of Healthcheck interface
func NewHealthcheckService() HealthcheckService {
	return &service{}
}

func (s *service) SetOnlineSince(t time.Time) {
	s.onlineSince = &t
}

func (s *service) OnlineSince() (time.Duration, error) {
	if s.onlineSince == nil {
		return time.Duration(0), errors.New("online since is not set")
	}
	return time.Since(*s.onlineSince), nil
}
