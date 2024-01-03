// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	fiber "github.com/gofiber/fiber/v2"
	mock "github.com/stretchr/testify/mock"
)

// BorrowersHandlerInterface is an autogenerated mock type for the BorrowersHandlerInterface type
type BorrowersHandlerInterface struct {
	mock.Mock
}

// GetAllBorrowers provides a mock function with given fields: c
func (_m *BorrowersHandlerInterface) GetAllBorrowers(c *fiber.Ctx) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBorrowersHandlerInterface creates a new instance of BorrowersHandlerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBorrowersHandlerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *BorrowersHandlerInterface {
	mock := &BorrowersHandlerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
