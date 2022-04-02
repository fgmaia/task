// Code generated by mockery v2.10.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	output "github.com/fgmaia/task/internal/usecases/ports/output"
)

// FindTaskUseCase is an autogenerated mock type for the FindTaskUseCase type
type FindTaskUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, userID, taskID
func (_m *FindTaskUseCase) Execute(ctx context.Context, userID string, taskID string) (*output.FindTaskOutput, error) {
	ret := _m.Called(ctx, userID, taskID)

	var r0 *output.FindTaskOutput
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *output.FindTaskOutput); ok {
		r0 = rf(ctx, userID, taskID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.FindTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, taskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
