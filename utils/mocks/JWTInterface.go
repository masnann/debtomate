// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	jwt "github.com/golang-jwt/jwt"
	mock "github.com/stretchr/testify/mock"
)

// JWTInterface is an autogenerated mock type for the JWTInterface type
type JWTInterface struct {
	mock.Mock
}

// GenerateJWT provides a mock function with given fields: userID, email
func (_m *JWTInterface) GenerateJWT(userID uint64, email string) (string, error) {
	ret := _m.Called(userID, email)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, string) (string, error)); ok {
		return rf(userID, email)
	}
	if rf, ok := ret.Get(0).(func(uint64, string) string); ok {
		r0 = rf(userID, email)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(uint64, string) error); ok {
		r1 = rf(userID, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateToken provides a mock function with given fields: tokenString
func (_m *JWTInterface) ValidateToken(tokenString string) (*jwt.Token, error) {
	ret := _m.Called(tokenString)

	var r0 *jwt.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*jwt.Token, error)); ok {
		return rf(tokenString)
	}
	if rf, ok := ret.Get(0).(func(string) *jwt.Token); ok {
		r0 = rf(tokenString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewJWTInterface creates a new instance of JWTInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWTInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWTInterface {
	mock := &JWTInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
