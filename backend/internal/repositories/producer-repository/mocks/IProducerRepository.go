// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks_product_repository

import (
	context "context"

	customerror "github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	entities "github.com/beto-ouverney/go-affiliates/backend/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// IProducerRepository is an autogenerated mock type for the IProducerRepository type
type IProducerRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, p
func (_m *IProducerRepository) Add(ctx context.Context, p []entities.Producer) *customerror.CustomError {
	ret := _m.Called(ctx, p)

	var r0 *customerror.CustomError
	if rf, ok := ret.Get(0).(func(context.Context, []entities.Producer) *customerror.CustomError); ok {
		r0 = rf(ctx, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customerror.CustomError)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *IProducerRepository) GetAll(ctx context.Context) (*[]entities.Producer, *customerror.CustomError) {
	ret := _m.Called(ctx)

	var r0 *[]entities.Producer
	if rf, ok := ret.Get(0).(func(context.Context) *[]entities.Producer); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.Producer)
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

type mockConstructorTestingTNewIProducerRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIProducerRepository creates a new instance of IProducerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIProducerRepository(t mockConstructorTestingTNewIProducerRepository) *IProducerRepository {
	mock := &IProducerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
