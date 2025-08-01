// Code generated by mockery. DO NOT EDIT.

package mockclients

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// MockPostgresClient is an autogenerated mock type for the PostgresClient type
type MockPostgresClient struct {
	mock.Mock
}

type MockPostgresClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPostgresClient) EXPECT() *MockPostgresClient_Expecter {
	return &MockPostgresClient_Expecter{mock: &_m.Mock}
}

// GetConnection provides a mock function with no fields
func (_m *MockPostgresClient) GetConnection() *gorm.DB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetConnection")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// MockPostgresClient_GetConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnection'
type MockPostgresClient_GetConnection_Call struct {
	*mock.Call
}

// GetConnection is a helper method to define mock.On call
func (_e *MockPostgresClient_Expecter) GetConnection() *MockPostgresClient_GetConnection_Call {
	return &MockPostgresClient_GetConnection_Call{Call: _e.mock.On("GetConnection")}
}

func (_c *MockPostgresClient_GetConnection_Call) Run(run func()) *MockPostgresClient_GetConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPostgresClient_GetConnection_Call) Return(_a0 *gorm.DB) *MockPostgresClient_GetConnection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPostgresClient_GetConnection_Call) RunAndReturn(run func() *gorm.DB) *MockPostgresClient_GetConnection_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPostgresClient creates a new instance of MockPostgresClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPostgresClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPostgresClient {
	mock := &MockPostgresClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
