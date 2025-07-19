package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"go-api/.internal/mocks/mockservices"
)

// HandlerTestSuite ...
type HealthcheckHandlerTestSuite struct {
	suite.Suite

	mockService *mockservices.MockHealthcheckService
	handler     HealthcheckHandler
}

// SetupTest ...
func (s *HealthcheckHandlerTestSuite) SetupTest() {
	t := s.T()
	s.mockService = mockservices.NewMockHealthcheckService(t)
	s.handler = NewHealthcheckHandler(Params{
		HealthcheckService: s.mockService,
	})
}

// SetupSubTest ...
func (s *HealthcheckHandlerTestSuite) SetupSubTest() {
	s.SetupTest() // Clean up the mocks
}

// TestHandlerTestSuite ...
func TestHealthcheckHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HealthcheckHandlerTestSuite))
}

// TestGetAPIStatus ...
func (s *HealthcheckHandlerTestSuite) TestGetAPIStatus() {
	tests := map[string]struct {
		MockSetup      func()
		ExpectedStatus int
		ExpectedBody   string
	}{
		"success": {
			MockSetup: func() {
				s.mockService.On("OnlineSince").Return(time.Duration(10*time.Second), nil)
			},
			ExpectedStatus: http.StatusOK,
			ExpectedBody:   "{\"online_time\":\"10s\"}\n",
		},
		"fail to get online time": {
			MockSetup: func() {
				s.mockService.On("OnlineSince").Return(time.Duration(0), assert.AnError)
			},
			ExpectedStatus: http.StatusInternalServerError,
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			if tc.MockSetup != nil {
				tc.MockSetup()
			}

			resp, err := runHandler(s.handler.GetAPIStatus)

			s.NoError(err)
			s.Equal(tc.ExpectedStatus, resp.Code)
			if tc.ExpectedBody != "" {
				s.Equal(tc.ExpectedBody, resp.Body.String())
			}
		})
	}
}

func runHandler(f func(e echo.Context) error) (*httptest.ResponseRecorder, error) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := f(c)
	return rec, err
}
