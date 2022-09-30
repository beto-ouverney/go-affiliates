// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocksaffiliaterepository

import (
	context "context"

	customerror "github.com/beto-ouverney/go-affiliates/backend/internal/customerror"
	entities "github.com/beto-ouverney/go-affiliates/backend/internal/entities"

	mock "github.com/stretchr/testify/mock"
)

// IAffiliateRepository is an autogenerated mock type for the IAffiliateRepository type
type IAffiliateRepository struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, af
func (_m *IAffiliateRepository) Add(ctx context.Context, af []entities.Affiliate) *customerror.CustomError {
	ret := _m.Called(ctx, af)

	var r0 *customerror.CustomError
	if rf, ok := ret.Get(0).(func(context.Context, []entities.Affiliate) *customerror.CustomError); ok {
		r0 = rf(ctx, af)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customerror.CustomError)
		}
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *IAffiliateRepository) GetAll(ctx context.Context) (*[]entities.Affiliate, *customerror.CustomError) {
	ret := _m.Called(ctx)

	var r0 *[]entities.Affiliate
	if rf, ok := ret.Get(0).(func(context.Context) *[]entities.Affiliate); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.Affiliate)
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

type mockConstructorTestingTNewIAffiliateRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIAffiliateRepository creates a new instance of IAffiliateRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIAffiliateRepository(t mockConstructorTestingTNewIAffiliateRepository) *IAffiliateRepository {
	mock := &IAffiliateRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
