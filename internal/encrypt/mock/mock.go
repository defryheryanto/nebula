// Code generated by MockGen. DO NOT EDIT.
// Source: encryptor.go
//
// Generated by this command:
//
//	mockgen -source encryptor.go -package encryptormock -destination mock/mock.go
//
// Package encryptormock is a generated GoMock package.
package encryptormock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockEncryptor is a mock of Encryptor interface.
type MockEncryptor struct {
	ctrl     *gomock.Controller
	recorder *MockEncryptorMockRecorder
}

// MockEncryptorMockRecorder is the mock recorder for MockEncryptor.
type MockEncryptorMockRecorder struct {
	mock *MockEncryptor
}

// NewMockEncryptor creates a new mock instance.
func NewMockEncryptor(ctrl *gomock.Controller) *MockEncryptor {
	mock := &MockEncryptor{ctrl: ctrl}
	mock.recorder = &MockEncryptorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncryptor) EXPECT() *MockEncryptorMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockEncryptor) Check(ctx context.Context, realString, encryptedString string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", ctx, realString, encryptedString)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockEncryptorMockRecorder) Check(ctx, realString, encryptedString any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockEncryptor)(nil).Check), ctx, realString, encryptedString)
}

// Encrypt mocks base method.
func (m *MockEncryptor) Encrypt(ctx context.Context, raw string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encrypt", ctx, raw)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockEncryptorMockRecorder) Encrypt(ctx, raw any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockEncryptor)(nil).Encrypt), ctx, raw)
}
