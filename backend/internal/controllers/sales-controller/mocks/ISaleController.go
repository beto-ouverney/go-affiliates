// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	customerror "github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	mock "github.com/stretchr/testify/mock"
)

// ISaleController is an autogenerated mock type for the ISaleController type
type ISaleController struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, nameFile
func (_m *ISaleController) Add(ctx context.Context, nameFile string) ([]byte, *customerror.CustomError) {
	ret := _m.Called(ctx, nameFile)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, string) []byte); ok {
		r0 = rf(ctx, nameFile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 *customerror.CustomError
	if rf, ok := ret.Get(1).(func(context.Context, string) *customerror.CustomError); ok {
		r1 = rf(ctx, nameFile)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*customerror.CustomError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewISaleController interface {
	mock.TestingT
	Cleanup(func())
}

// NewISaleController creates a new instance of ISaleController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewISaleController(t mockConstructorTestingTNewISaleController) *ISaleController {
	mock := &ISaleController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
