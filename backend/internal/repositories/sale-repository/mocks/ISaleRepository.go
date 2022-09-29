// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks_sale_repository

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

// Add provides a mock function with given fields: ctx, sale
func (_m *ISaleRepository) Add(ctx context.Context, sale []entities.Sale) *customerror.CustomError {
	ret := _m.Called(ctx, sale)

	var r0 *customerror.CustomError
	if rf, ok := ret.Get(0).(func(context.Context, []entities.Sale) *customerror.CustomError); ok {
		r0 = rf(ctx, sale)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customerror.CustomError)
		}
	}

	return r0
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