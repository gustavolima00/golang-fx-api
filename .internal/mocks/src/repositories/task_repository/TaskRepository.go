// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	task "go-api/src/models/task"

	mock "github.com/stretchr/testify/mock"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

type TaskRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *TaskRepository) EXPECT() *TaskRepository_Expecter {
	return &TaskRepository_Expecter{mock: &_m.Mock}
}

// CreateTask provides a mock function with given fields: _a0
func (_m *TaskRepository) CreateTask(_a0 *task.Task) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*task.Task) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TaskRepository_CreateTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTask'
type TaskRepository_CreateTask_Call struct {
	*mock.Call
}

// CreateTask is a helper method to define mock.On call
//   - _a0 *task.Task
func (_e *TaskRepository_Expecter) CreateTask(_a0 interface{}) *TaskRepository_CreateTask_Call {
	return &TaskRepository_CreateTask_Call{Call: _e.mock.On("CreateTask", _a0)}
}

func (_c *TaskRepository_CreateTask_Call) Run(run func(_a0 *task.Task)) *TaskRepository_CreateTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*task.Task))
	})
	return _c
}

func (_c *TaskRepository_CreateTask_Call) Return(_a0 error) *TaskRepository_CreateTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TaskRepository_CreateTask_Call) RunAndReturn(run func(*task.Task) error) *TaskRepository_CreateTask_Call {
	_c.Call.Return(run)
	return _c
}

// GetTaskByID provides a mock function with given fields: id
func (_m *TaskRepository) GetTaskByID(id int) (*task.Task, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByID")
	}

	var r0 *task.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*task.Task, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *task.Task); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*task.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TaskRepository_GetTaskByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTaskByID'
type TaskRepository_GetTaskByID_Call struct {
	*mock.Call
}

// GetTaskByID is a helper method to define mock.On call
//   - id int
func (_e *TaskRepository_Expecter) GetTaskByID(id interface{}) *TaskRepository_GetTaskByID_Call {
	return &TaskRepository_GetTaskByID_Call{Call: _e.mock.On("GetTaskByID", id)}
}

func (_c *TaskRepository_GetTaskByID_Call) Run(run func(id int)) *TaskRepository_GetTaskByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *TaskRepository_GetTaskByID_Call) Return(_a0 *task.Task, _a1 error) *TaskRepository_GetTaskByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TaskRepository_GetTaskByID_Call) RunAndReturn(run func(int) (*task.Task, error)) *TaskRepository_GetTaskByID_Call {
	_c.Call.Return(run)
	return _c
}

// NewTaskRepository creates a new instance of TaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepository {
	mock := &TaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
