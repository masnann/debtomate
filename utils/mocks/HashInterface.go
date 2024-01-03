// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// HashInterface is an autogenerated mock type for the HashInterface type
type HashInterface struct {
	mock.Mock
}

// ComparePassword provides a mock function with given fields: _a0, password
func (_m *HashInterface) ComparePassword(_a0 string, password string) (bool, error) {
	ret := _m.Called(_a0, password)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(_a0, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(_a0, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateHash provides a mock function with given fields: password
func (_m *HashInterface) GenerateHash(password string) (string, error) {
	ret := _m.Called(password)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(password)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewHashInterface creates a new instance of HashInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHashInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *HashInterface {
	mock := &HashInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}