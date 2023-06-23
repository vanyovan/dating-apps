// Code generated by MockGen. DO NOT EDIT.
// Source: C:/Users/Yovan/Desktop/Deals/internal/usecase/otp.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/vanyovan/dating-apps/internal/entity"
)

// MockOtpServiceProvider is a mock of OtpServiceProvider interface.
type MockOtpServiceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockOtpServiceProviderMockRecorder
}

// MockOtpServiceProviderMockRecorder is the mock recorder for MockOtpServiceProvider.
type MockOtpServiceProviderMockRecorder struct {
	mock *MockOtpServiceProvider
}

// NewMockOtpServiceProvider creates a new mock instance.
func NewMockOtpServiceProvider(ctrl *gomock.Controller) *MockOtpServiceProvider {
	mock := &MockOtpServiceProvider{ctrl: ctrl}
	mock.recorder = &MockOtpServiceProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOtpServiceProvider) EXPECT() *MockOtpServiceProviderMockRecorder {
	return m.recorder
}

// SetOTP mocks base method.
func (m *MockOtpServiceProvider) SetOTP(ctx context.Context, param entity.OTPRequestParam) (entity.OTPResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOTP", ctx, param)
	ret0, _ := ret[0].(entity.OTPResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetOTP indicates an expected call of SetOTP.
func (mr *MockOtpServiceProviderMockRecorder) SetOTP(ctx, param interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOTP", reflect.TypeOf((*MockOtpServiceProvider)(nil).SetOTP), ctx, param)
}
