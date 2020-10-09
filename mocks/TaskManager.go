// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	tasks "tasks/gen/tasks"

	mock "github.com/stretchr/testify/mock"
)

// TaskManager is an autogenerated mock type for the TaskManager type
type TaskManager struct {
	mock.Mock
}

// Add provides a mock function with given fields: p
func (_m *TaskManager) Add(p *tasks.Task) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(*tasks.Task) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields:
func (_m *TaskManager) List() (tasks.StoredTaskCollection, error) {
	ret := _m.Called()

	var r0 tasks.StoredTaskCollection
	if rf, ok := ret.Get(0).(func() tasks.StoredTaskCollection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tasks.StoredTaskCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Remove provides a mock function with given fields: id
func (_m *TaskManager) Remove(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Show provides a mock function with given fields: id
func (_m *TaskManager) Show(id int) (*tasks.StoredTask, error) {
	ret := _m.Called(id)

	var r0 *tasks.StoredTask
	if rf, ok := ret.Get(0).(func(int) *tasks.StoredTask); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tasks.StoredTask)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields: id, status
func (_m *TaskManager) Status(id int, status string) error {
	ret := _m.Called(id, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: p
func (_m *TaskManager) Update(p *tasks.Task) error {
	ret := _m.Called(p)

	var r0 error
	if rf, ok := ret.Get(0).(func(*tasks.Task) error); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}