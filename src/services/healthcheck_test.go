package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestSetOnlineSince will test the SetOnlineSince function
func TestSetOnlineSince(t *testing.T) {
	s := NewHealthcheckService()

	// Test if the onlineSince is nil
	_, err := s.OnlineSince()
	assert.Error(t, err)

	// Test if the onlineSince is set
	s.SetOnlineSince(time.Now())

	val, err := s.OnlineSince()
	assert.NoError(t, err)
	assert.NotEqual(t, time.Duration(0), val)
}
