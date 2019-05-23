// Code generated by mockery v1.0.0. DO NOT EDIT.

package client

import agentpb "github.com/percona/pmm/api/agentpb"
import mock "github.com/stretchr/testify/mock"

// mockConnectionChecker is an autogenerated mock type for the connectionChecker type
type mockConnectionChecker struct {
	mock.Mock
}

// Check provides a mock function with given fields: msg
func (_m *mockConnectionChecker) Check(msg *agentpb.CheckConnectionRequest) error {
	ret := _m.Called(msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*agentpb.CheckConnectionRequest) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
