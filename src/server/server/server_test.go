package server

import (
	"context"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/fx"

	"go-api/.internal/mocks/mockservices"
	"go-api/src/common/config"
)

type fakeLifecycle struct{}

func (m *fakeLifecycle) Append(hook fx.Hook) {}

type ServerTestSuite struct {
	suite.Suite

	mockHealthcheckService *mockservices.MockHealthcheckService
	lifecycle              fx.Lifecycle
	config                 *config.Config
}

func (s *ServerTestSuite) SetupTest() {
	s.lifecycle = &fakeLifecycle{}
	s.mockHealthcheckService = mockservices.NewMockHealthcheckService(s.T())
	s.config = &config.Config{
		Port: "8080",
	}
}

func (s *ServerTestSuite) SetupSubTest() {
	s.SetupTest() // Clean up the mocks
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

func (s *ServerTestSuite) TestNew() {
	echoInstance := NewServer(Params{
		Lifecycle:          s.lifecycle,
		HealthcheckService: s.mockHealthcheckService,
		Config:             s.config,
	})
	assert.NotNil(s.T(), echoInstance)
}

func (s *ServerTestSuite) TestOnStartHook() {
	echoInstance := echo.New()

	hook := onStart(echoInstance, Params{
		Lifecycle:          s.lifecycle,
		HealthcheckService: s.mockHealthcheckService,
		Config:             s.config,
	})
	s.mockHealthcheckService.On("SetOnlineSince", mock.AnythingOfType("time.Time")).Return().Once()
	err := hook(context.Background())
	s.NoError(err)
}

func (s *ServerTestSuite) TestOnStopHook() {
	echoInstance := echo.New()

	hook := onStop(echoInstance, Params{
		Lifecycle:          s.lifecycle,
		HealthcheckService: s.mockHealthcheckService,
		Config:             s.config,
	})
	err := hook(context.Background())
	s.NoError(err)
}
