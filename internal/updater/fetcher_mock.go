// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quay/claircore/internal/updater (interfaces: Fetcher)

// Package updater is a generated GoMock package.
package updater

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockFetcher is a mock of Fetcher interface
type MockFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockFetcherMockRecorder
}

// MockFetcherMockRecorder is the mock recorder for MockFetcher
type MockFetcherMockRecorder struct {
	mock *MockFetcher
}

// NewMockFetcher creates a new mock instance
func NewMockFetcher(ctrl *gomock.Controller) *MockFetcher {
	mock := &MockFetcher{ctrl: ctrl}
	mock.recorder = &MockFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFetcher) EXPECT() *MockFetcherMockRecorder {
	return m.recorder
}

// Fetch mocks base method
func (m *MockFetcher) Fetch(arg0 context.Context) (io.ReadCloser, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Fetch indicates an expected call of Fetch
func (mr *MockFetcherMockRecorder) Fetch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockFetcher)(nil).Fetch), arg0)
}
