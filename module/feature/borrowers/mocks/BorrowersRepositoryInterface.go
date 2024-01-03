// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	entities "debtomate/module/entities"

	mock "github.com/stretchr/testify/mock"
)

// BorrowersRepositoryInterface is an autogenerated mock type for the BorrowersRepositoryInterface type
type BorrowersRepositoryInterface struct {
	mock.Mock
}

// GetPaginatedBorrowers provides a mock function with given fields: page, pageSize
func (_m *BorrowersRepositoryInterface) GetPaginatedBorrowers(page int, pageSize int) ([]*entities.BorrowersModels, error) {
	ret := _m.Called(page, pageSize)

	var r0 []*entities.BorrowersModels
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*entities.BorrowersModels, error)); ok {
		return rf(page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*entities.BorrowersModels); ok {
		r0 = rf(page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.BorrowersModels)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalItems provides a mock function with given fields:
func (_m *BorrowersRepositoryInterface) GetTotalItems() (int64, error) {
	ret := _m.Called()

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func() (int64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBorrowersRepositoryInterface creates a new instance of BorrowersRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBorrowersRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *BorrowersRepositoryInterface {
	mock := &BorrowersRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}