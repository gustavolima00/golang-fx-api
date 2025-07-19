package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"go-api/.internal/mocks/mockhandlers"
)

type RoutesTestSuite struct {
	suite.Suite

	mockEcho               *echo.Echo
	mockHealthcheckHandler *mockhandlers.MockHealthcheckHandler
}

func (s *RoutesTestSuite) SetupTest() {
	s.mockEcho = echo.New()
	s.mockHealthcheckHandler = mockhandlers.NewMockHealthcheckHandler(s.T())
	Register(Params{
		Echo:               s.mockEcho,
		HealthcheckHandler: s.mockHealthcheckHandler,
	})
}

func TestRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(RoutesTestSuite))
}

func (s *RoutesTestSuite) TestGetAPIStatus() {
	// Mock the GetAPIStatus method of HealthcheckHandler
	s.mockHealthcheckHandler.On("GetAPIStatus", mock.AnythingOfType("*echo.context")).Return(nil).Once()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	s.mockEcho.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func (s *RoutesTestSuite) TestSwaggerRoute() {
	// Test Swagger route (less critical to mock, just ensure it's registered)
	req := httptest.NewRequest(http.MethodGet, "/swagger/index.html", nil)
	rec := httptest.NewRecorder()
	s.mockEcho.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}
