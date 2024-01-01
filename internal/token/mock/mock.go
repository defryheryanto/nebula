// Code generated by MockGen. DO NOT EDIT.
// Source: token.go
//
// Generated by this command:
//
//	mockgen -source token.go -package tokenmock -destination mock/mock.go
//
// Package tokenmock is a generated GoMock package.
package tokenmock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockTokener is a mock of Tokener interface.
type MockTokener[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockTokenerMockRecorder[T]
}

// MockTokenerMockRecorder is the mock recorder for MockTokener.
type MockTokenerMockRecorder[T any] struct {
	mock *MockTokener[T]
}

// NewMockTokener creates a new mock instance.
func NewMockTokener[T any](ctrl *gomock.Controller) *MockTokener[T] {
	mock := &MockTokener[T]{ctrl: ctrl}
	mock.recorder = &MockTokenerMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokener[T]) EXPECT() *MockTokenerMockRecorder[T] {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockTokener[T]) GenerateToken(ctx context.Context, payload T, expiryTime *time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", ctx, payload, expiryTime)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockTokenerMockRecorder[T]) GenerateToken(ctx, payload, expiryTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockTokener[T])(nil).GenerateToken), ctx, payload, expiryTime)
}

// Validate mocks base method.
func (m *MockTokener[T]) Validate(ctx context.Context, token string) (T, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ctx, token)
	ret0, _ := ret[0].(T)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validate indicates an expected call of Validate.
func (mr *MockTokenerMockRecorder[T]) Validate(ctx, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockTokener[T])(nil).Validate), ctx, token)
}
