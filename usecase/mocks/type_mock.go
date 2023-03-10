// Code generated by mockery v2.15.0. DO NOT EDIT.

package usecasemock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	entity "github.com/winartodev/go-pokedex/entity"
)

// TypeUsecaseItf is an autogenerated mock type for the TypeUsecaseItf type
type TypeUsecaseItf struct {
	mock.Mock
}

// CreateType provides a mock function with given fields: ctx, data
func (_m *TypeUsecaseItf) CreateType(ctx context.Context, data entity.Type) (int64, error) {
	ret := _m.Called(ctx, data)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, entity.Type) int64); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.Type) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GeTypeByID provides a mock function with given fields: ctx, id
func (_m *TypeUsecaseItf) GeTypeByID(ctx context.Context, id int64) (entity.Type, error) {
	ret := _m.Called(ctx, id)

	var r0 entity.Type
	if rf, ok := ret.Get(0).(func(context.Context, int64) entity.Type); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Type)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllType provides a mock function with given fields: ctx
func (_m *TypeUsecaseItf) GetAllType(ctx context.Context) ([]entity.Type, error) {
	ret := _m.Called(ctx)

	var r0 []entity.Type
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Type); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Type)
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

// UpdateType provides a mock function with given fields: ctx, id, data
func (_m *TypeUsecaseItf) UpdateType(ctx context.Context, id int64, data entity.Type) error {
	ret := _m.Called(ctx, id, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, entity.Type) error); ok {
		r0 = rf(ctx, id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTypeUsecaseItf interface {
	mock.TestingT
	Cleanup(func())
}

// NewTypeUsecaseItf creates a new instance of TypeUsecaseItf. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTypeUsecaseItf(t mockConstructorTestingTNewTypeUsecaseItf) *TypeUsecaseItf {
	mock := &TypeUsecaseItf{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
