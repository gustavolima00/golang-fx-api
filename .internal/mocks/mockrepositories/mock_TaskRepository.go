// Code generated by mockery. DO NOT EDIT.

package mockrepositories

import (
	models "go-api/src/models"

	mock "github.com/stretchr/testify/mock"
)

// MockTaskRepository is an autogenerated mock type for the TaskRepository type
type MockTaskRepository struct {
	mock.Mock
}

type MockTaskRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTaskRepository) EXPECT() *MockTaskRepository_Expecter {
	return &MockTaskRepository_Expecter{mock: &_m.Mock}
}

// CreateTask provides a mock function with given fields: task
func (_m *MockTaskRepository) CreateTask(task models.Task) error {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTaskRepository_CreateTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTask'
type MockTaskRepository_CreateTask_Call struct {
	*mock.Call
}

// CreateTask is a helper method to define mock.On call
//   - task models.Task
func (_e *MockTaskRepository_Expecter) CreateTask(task interface{}) *MockTaskRepository_CreateTask_Call {
	return &MockTaskRepository_CreateTask_Call{Call: _e.mock.On("CreateTask", task)}
}

func (_c *MockTaskRepository_CreateTask_Call) Run(run func(task models.Task)) *MockTaskRepository_CreateTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Task))
	})
	return _c
}

func (_c *MockTaskRepository_CreateTask_Call) Return(_a0 error) *MockTaskRepository_CreateTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTaskRepository_CreateTask_Call) RunAndReturn(run func(models.Task) error) *MockTaskRepository_CreateTask_Call {
	_c.Call.Return(run)
	return _c
}

// GetTaskByID provides a mock function with given fields: id
func (_m *MockTaskRepository) GetTaskByID(id int) (*models.Task, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByID")
	}

	var r0 *models.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*models.Task, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *models.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTaskRepository_GetTaskByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTaskByID'
type MockTaskRepository_GetTaskByID_Call struct {
	*mock.Call
}

// GetTaskByID is a helper method to define mock.On call
//   - id int
func (_e *MockTaskRepository_Expecter) GetTaskByID(id interface{}) *MockTaskRepository_GetTaskByID_Call {
	return &MockTaskRepository_GetTaskByID_Call{Call: _e.mock.On("GetTaskByID", id)}
}

func (_c *MockTaskRepository_GetTaskByID_Call) Run(run func(id int)) *MockTaskRepository_GetTaskByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *MockTaskRepository_GetTaskByID_Call) Return(_a0 *models.Task, _a1 error) *MockTaskRepository_GetTaskByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTaskRepository_GetTaskByID_Call) RunAndReturn(run func(int) (*models.Task, error)) *MockTaskRepository_GetTaskByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTaskRepository creates a new instance of MockTaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTaskRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTaskRepository {
	mock := &MockTaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
