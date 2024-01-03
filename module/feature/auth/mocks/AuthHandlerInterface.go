// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	fiber "github.com/gofiber/fiber/v2"
	mock "github.com/stretchr/testify/mock"
)

// AuthHandlerInterface is an autogenerated mock type for the AuthHandlerInterface type
type AuthHandlerInterface struct {
	mock.Mock
}

// Login provides a mock function with given fields: c
func (_m *AuthHandlerInterface) Login(c *fiber.Ctx) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewAuthHandlerInterface creates a new instance of AuthHandlerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthHandlerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthHandlerInterface {
	mock := &AuthHandlerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
