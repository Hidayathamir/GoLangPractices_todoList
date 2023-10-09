// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/ArMo-Team/GoLangPractices_todoList/internal/app/domain"
	mock "github.com/stretchr/testify/mock"
)

// TodoRepository is an autogenerated mock type for the TodoRepository type
type TodoRepository struct {
	mock.Mock
}

type TodoRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *TodoRepository) EXPECT() *TodoRepository_Expecter {
	return &TodoRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0
func (_m *TodoRepository) Create(_a0 *domain.Todo) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Todo) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TodoRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type TodoRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 *domain.Todo
func (_e *TodoRepository_Expecter) Create(_a0 interface{}) *TodoRepository_Create_Call {
	return &TodoRepository_Create_Call{Call: _e.mock.On("Create", _a0)}
}

func (_c *TodoRepository_Create_Call) Run(run func(_a0 *domain.Todo)) *TodoRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Todo))
	})
	return _c
}

func (_c *TodoRepository_Create_Call) Return(_a0 error) *TodoRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TodoRepository_Create_Call) RunAndReturn(run func(*domain.Todo) error) *TodoRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *TodoRepository) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TodoRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type TodoRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id string
func (_e *TodoRepository_Expecter) Delete(id interface{}) *TodoRepository_Delete_Call {
	return &TodoRepository_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *TodoRepository_Delete_Call) Run(run func(id string)) *TodoRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TodoRepository_Delete_Call) Return(_a0 error) *TodoRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TodoRepository_Delete_Call) RunAndReturn(run func(string) error) *TodoRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *TodoRepository) GetAll() ([]*domain.Todo, error) {
	ret := _m.Called()

	var r0 []*domain.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*domain.Todo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*domain.Todo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TodoRepository_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type TodoRepository_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *TodoRepository_Expecter) GetAll() *TodoRepository_GetAll_Call {
	return &TodoRepository_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *TodoRepository_GetAll_Call) Run(run func()) *TodoRepository_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *TodoRepository_GetAll_Call) Return(_a0 []*domain.Todo, _a1 error) *TodoRepository_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TodoRepository_GetAll_Call) RunAndReturn(run func() ([]*domain.Todo, error)) *TodoRepository_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: id
func (_m *TodoRepository) GetByID(id string) (*domain.Todo, error) {
	ret := _m.Called(id)

	var r0 *domain.Todo
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Todo, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Todo); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Todo)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TodoRepository_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type TodoRepository_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - id string
func (_e *TodoRepository_Expecter) GetByID(id interface{}) *TodoRepository_GetByID_Call {
	return &TodoRepository_GetByID_Call{Call: _e.mock.On("GetByID", id)}
}

func (_c *TodoRepository_GetByID_Call) Run(run func(id string)) *TodoRepository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *TodoRepository_GetByID_Call) Return(_a0 *domain.Todo, _a1 error) *TodoRepository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TodoRepository_GetByID_Call) RunAndReturn(run func(string) (*domain.Todo, error)) *TodoRepository_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0
func (_m *TodoRepository) Update(_a0 *domain.Todo) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Todo) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TodoRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type TodoRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 *domain.Todo
func (_e *TodoRepository_Expecter) Update(_a0 interface{}) *TodoRepository_Update_Call {
	return &TodoRepository_Update_Call{Call: _e.mock.On("Update", _a0)}
}

func (_c *TodoRepository_Update_Call) Run(run func(_a0 *domain.Todo)) *TodoRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Todo))
	})
	return _c
}

func (_c *TodoRepository_Update_Call) Return(_a0 error) *TodoRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TodoRepository_Update_Call) RunAndReturn(run func(*domain.Todo) error) *TodoRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewTodoRepository creates a new instance of TodoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTodoRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TodoRepository {
	mock := &TodoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}