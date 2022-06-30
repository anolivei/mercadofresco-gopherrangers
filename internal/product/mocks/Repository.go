// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	products "github.com/Gopher-Rangers/mercadofresco-gopherrangers/internal/product"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *Repository) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() ([]products.Product, error) {
	ret := _m.Called()

	var r0 []products.Product
	if rf, ok := ret.Get(0).(func() []products.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]products.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *Repository) GetById(id int) (products.Product, error) {
	ret := _m.Called(id)

	var r0 products.Product
	if rf, ok := ret.Get(0).(func(int) products.Product); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(products.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LastID provides a mock function with given fields:
func (_m *Repository) LastID() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: prod, id
func (_m *Repository) Store(prod products.Product, id int) (products.Product, error) {
	ret := _m.Called(prod, id)

	var r0 products.Product
	if rf, ok := ret.Get(0).(func(products.Product, int) products.Product); ok {
		r0 = rf(prod, id)
	} else {
		r0 = ret.Get(0).(products.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(products.Product, int) error); ok {
		r1 = rf(prod, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: prod, id
func (_m *Repository) Update(prod products.Product, id int) (products.Product, error) {
	ret := _m.Called(prod, id)

	var r0 products.Product
	if rf, ok := ret.Get(0).(func(products.Product, int) products.Product); ok {
		r0 = rf(prod, id)
	} else {
		r0 = ret.Get(0).(products.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(products.Product, int) error); ok {
		r1 = rf(prod, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t NewRepositoryT) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
