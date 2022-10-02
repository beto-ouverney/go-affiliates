// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	customerror "github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	entities "github.com/beto-ouverney/go-affiliates/backend/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// ISaleRepository is an autogenerated mock type for the ISaleRepository type
type ISaleRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, s
func (_m *ISaleRepository) Add(ctx context.Context, s []entities.Sale) *customerror.CustomError {
	ret := _m.Called(ctx, s)

	var r0 *customerror.CustomError
	if rf, ok := ret.Get(0).(func(context.Context, []entities.Sale) *customerror.CustomError); ok {
		r0 = rf(ctx, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customerror.CustomError)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *ISaleRepository) GetAll(ctx context.Context) (*[]entities.SaleResponse, *customerror.CustomError) {
	ret := _m.Called(ctx)

	var r0 *[]entities.SaleResponse
	if rf, ok := ret.Get(0).(func(context.Context) *[]entities.SaleResponse); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.SaleResponse)
		}
	}

	var r1 *customerror.CustomError
	if rf, ok := ret.Get(1).(func(context.Context) *customerror.CustomError); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*customerror.CustomError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewISaleRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewISaleRepository creates a new instance of ISaleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewISaleRepository(t mockConstructorTestingTNewISaleRepository) *ISaleRepository {
	mock := &ISaleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
