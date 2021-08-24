// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	evidence "github.com/tendermint/tendermint/pkg/evidence"
	light "github.com/tendermint/tendermint/pkg/light"

	mock "github.com/stretchr/testify/mock"
)

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// LightBlock provides a mock function with given fields: ctx, height
func (_m *Provider) LightBlock(ctx context.Context, height int64) (*light.LightBlock, error) {
	ret := _m.Called(ctx, height)

	var r0 *light.LightBlock
	if rf, ok := ret.Get(0).(func(context.Context, int64) *light.LightBlock); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*light.LightBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReportEvidence provides a mock function with given fields: _a0, _a1
func (_m *Provider) ReportEvidence(_a0 context.Context, _a1 evidence.Evidence) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, evidence.Evidence) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
