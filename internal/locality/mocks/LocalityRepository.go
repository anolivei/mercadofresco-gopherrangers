// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	locality "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/locality"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, id, localityName, provinceName, countryName
func (_m *Repository) Create(ctx context.Context, id string, localityName string, provinceName string, countryName string) (locality.Locality, error) {
	ret := _m.Called(ctx, id, localityName, provinceName, countryName)

	var r0 locality.Locality
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) locality.Locality); ok {
		r0 = rf(ctx, id, localityName, provinceName, countryName)
	} else {
		r0 = ret.Get(0).(locality.Locality)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, id, localityName, provinceName, countryName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]locality.Locality, error) {
	ret := _m.Called(ctx)

	var r0 []locality.Locality
	if rf, ok := ret.Get(0).(func(context.Context) []locality.Locality); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]locality.Locality)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *Repository) GetById(ctx context.Context, id string) (locality.Locality, error) {
	ret := _m.Called(ctx, id)

	var r0 locality.Locality
	if rf, ok := ret.Get(0).(func(context.Context, string) locality.Locality); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(locality.Locality)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReportSellers provides a mock function with given fields: ctx, id
func (_m *Repository) ReportSellers(ctx context.Context, id string) (locality.ReportSeller, error) {
	ret := _m.Called(ctx, id)

	var r0 locality.ReportSeller
	if rf, ok := ret.Get(0).(func(context.Context, string) locality.ReportSeller); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(locality.ReportSeller)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
