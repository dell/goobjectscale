// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/dell/goobjectscale/pkg/client/model"
	mock "github.com/stretchr/testify/mock"
)

// ObjectUserInterface is an autogenerated mock type for the ObjectUserInterface type
type ObjectUserInterface struct {
	mock.Mock
}

// CreateSecret provides a mock function with given fields: ctx, uid, req, params
func (_m *ObjectUserInterface) CreateSecret(ctx context.Context, uid string, req model.ObjectUserSecretKeyCreateReq, params map[string]string) (*model.ObjectUserSecretKeyCreateRes, error) {
	ret := _m.Called(ctx, uid, req, params)

	var r0 *model.ObjectUserSecretKeyCreateRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.ObjectUserSecretKeyCreateReq, map[string]string) (*model.ObjectUserSecretKeyCreateRes, error)); ok {
		return rf(ctx, uid, req, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, model.ObjectUserSecretKeyCreateReq, map[string]string) *model.ObjectUserSecretKeyCreateRes); ok {
		r0 = rf(ctx, uid, req, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ObjectUserSecretKeyCreateRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, model.ObjectUserSecretKeyCreateReq, map[string]string) error); ok {
		r1 = rf(ctx, uid, req, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteSecret provides a mock function with given fields: ctx, uid, req, params
func (_m *ObjectUserInterface) DeleteSecret(ctx context.Context, uid string, req model.ObjectUserSecretKeyDeleteReq, params map[string]string) error {
	ret := _m.Called(ctx, uid, req, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.ObjectUserSecretKeyDeleteReq, map[string]string) error); ok {
		r0 = rf(ctx, uid, req, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInfo provides a mock function with given fields: ctx, uid, params
func (_m *ObjectUserInterface) GetInfo(ctx context.Context, uid string, params map[string]string) (*model.ObjectUserInfo, error) {
	ret := _m.Called(ctx, uid, params)

	var r0 *model.ObjectUserInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) (*model.ObjectUserInfo, error)); ok {
		return rf(ctx, uid, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) *model.ObjectUserInfo); ok {
		r0 = rf(ctx, uid, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ObjectUserInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, uid, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSecret provides a mock function with given fields: ctx, uid, params
func (_m *ObjectUserInterface) GetSecret(ctx context.Context, uid string, params map[string]string) (*model.ObjectUserSecret, error) {
	ret := _m.Called(ctx, uid, params)

	var r0 *model.ObjectUserSecret
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) (*model.ObjectUserSecret, error)); ok {
		return rf(ctx, uid, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) *model.ObjectUserSecret); ok {
		r0 = rf(ctx, uid, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ObjectUserSecret)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, uid, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, params
func (_m *ObjectUserInterface) List(ctx context.Context, params map[string]string) (*model.ObjectUserList, error) {
	ret := _m.Called(ctx, params)

	var r0 *model.ObjectUserList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string) (*model.ObjectUserList, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]string) *model.ObjectUserList); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ObjectUserList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]string) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewObjectUserInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewObjectUserInterface creates a new instance of ObjectUserInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewObjectUserInterface(t mockConstructorTestingTNewObjectUserInterface) *ObjectUserInterface {
	mock := &ObjectUserInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
