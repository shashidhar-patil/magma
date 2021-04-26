// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "github.com/facebookincubator/prometheus-edge-hub/grpc"
	mock "github.com/stretchr/testify/mock"
)

// EdgeHubServer is an autogenerated mock type for the EdgeHubServer type
type EdgeHubServer struct {
	mock.Mock
}

// Collect provides a mock function with given fields: ctx, families
func (_m *EdgeHubServer) Collect(ctx context.Context, families *grpc.MetricFamilies) (*grpc.Void, error) {
	ret := _m.Called(ctx, families)

	var r0 *grpc.Void
	if rf, ok := ret.Get(0).(func(context.Context, *grpc.MetricFamilies) *grpc.Void); ok {
		r0 = rf(ctx, families)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grpc.Void)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *grpc.MetricFamilies) error); ok {
		r1 = rf(ctx, families)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForTestsOnly provides a mock function with given fields:
func (_m *EdgeHubServer) ForTestsOnly() {
	_m.Called()
}