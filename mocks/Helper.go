// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	helper "sample/twirp/internal/helper"

	mock "github.com/stretchr/testify/mock"
)

// Helper is an autogenerated mock type for the Helper type
type Helper struct {
	mock.Mock
}

// GetAPIResponse provides a mock function with given fields:
func (_m *Helper) GetAPIResponse() helper.APIResponse {
	ret := _m.Called()

	var r0 helper.APIResponse
	if rf, ok := ret.Get(0).(func() helper.APIResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(helper.APIResponse)
	}

	return r0
}
