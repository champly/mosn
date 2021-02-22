// Code generated by MockGen. DO NOT EDIT.
// Source: ../../../api/network_filter.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "mosn.io/api"
	buffer "mosn.io/pkg/buffer"
)

// MockFilterManager is a mock of FilterManager interface
type MockFilterManager struct {
	ctrl     *gomock.Controller
	recorder *MockFilterManagerMockRecorder
}

// MockFilterManagerMockRecorder is the mock recorder for MockFilterManager
type MockFilterManagerMockRecorder struct {
	mock *MockFilterManager
}

// NewMockFilterManager creates a new mock instance
func NewMockFilterManager(ctrl *gomock.Controller) *MockFilterManager {
	mock := &MockFilterManager{ctrl: ctrl}
	mock.recorder = &MockFilterManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFilterManager) EXPECT() *MockFilterManagerMockRecorder {
	return m.recorder
}

// AddReadFilter mocks base method
func (m *MockFilterManager) AddReadFilter(rf api.ReadFilter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddReadFilter", rf)
}

// AddReadFilter indicates an expected call of AddReadFilter
func (mr *MockFilterManagerMockRecorder) AddReadFilter(rf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReadFilter", reflect.TypeOf((*MockFilterManager)(nil).AddReadFilter), rf)
}

// AddWriteFilter mocks base method
func (m *MockFilterManager) AddWriteFilter(wf api.WriteFilter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddWriteFilter", wf)
}

// AddWriteFilter indicates an expected call of AddWriteFilter
func (mr *MockFilterManagerMockRecorder) AddWriteFilter(wf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWriteFilter", reflect.TypeOf((*MockFilterManager)(nil).AddWriteFilter), wf)
}

// ListReadFilter mocks base method
func (m *MockFilterManager) ListReadFilter() []api.ReadFilter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReadFilter")
	ret0, _ := ret[0].([]api.ReadFilter)
	return ret0
}

// ListReadFilter indicates an expected call of ListReadFilter
func (mr *MockFilterManagerMockRecorder) ListReadFilter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReadFilter", reflect.TypeOf((*MockFilterManager)(nil).ListReadFilter))
}

// ListWriteFilters mocks base method
func (m *MockFilterManager) ListWriteFilters() []api.WriteFilter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWriteFilters")
	ret0, _ := ret[0].([]api.WriteFilter)
	return ret0
}

// ListWriteFilters indicates an expected call of ListWriteFilters
func (mr *MockFilterManagerMockRecorder) ListWriteFilters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWriteFilters", reflect.TypeOf((*MockFilterManager)(nil).ListWriteFilters))
}

// InitializeReadFilters mocks base method
func (m *MockFilterManager) InitializeReadFilters() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeReadFilters")
	ret0, _ := ret[0].(bool)
	return ret0
}

// InitializeReadFilters indicates an expected call of InitializeReadFilters
func (mr *MockFilterManagerMockRecorder) InitializeReadFilters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeReadFilters", reflect.TypeOf((*MockFilterManager)(nil).InitializeReadFilters))
}

// OnRead mocks base method
func (m *MockFilterManager) OnRead() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnRead")
}

// OnRead indicates an expected call of OnRead
func (mr *MockFilterManagerMockRecorder) OnRead() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnRead", reflect.TypeOf((*MockFilterManager)(nil).OnRead))
}

// OnWrite mocks base method
func (m *MockFilterManager) OnWrite(buffer []buffer.IoBuffer) api.FilterStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnWrite", buffer)
	ret0, _ := ret[0].(api.FilterStatus)
	return ret0
}

// OnWrite indicates an expected call of OnWrite
func (mr *MockFilterManagerMockRecorder) OnWrite(buffer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnWrite", reflect.TypeOf((*MockFilterManager)(nil).OnWrite), buffer)
}

// MockReadFilter is a mock of ReadFilter interface
type MockReadFilter struct {
	ctrl     *gomock.Controller
	recorder *MockReadFilterMockRecorder
}

// MockReadFilterMockRecorder is the mock recorder for MockReadFilter
type MockReadFilterMockRecorder struct {
	mock *MockReadFilter
}

// NewMockReadFilter creates a new mock instance
func NewMockReadFilter(ctrl *gomock.Controller) *MockReadFilter {
	mock := &MockReadFilter{ctrl: ctrl}
	mock.recorder = &MockReadFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadFilter) EXPECT() *MockReadFilterMockRecorder {
	return m.recorder
}

// OnData mocks base method
func (m *MockReadFilter) OnData(buffer buffer.IoBuffer) api.FilterStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnData", buffer)
	ret0, _ := ret[0].(api.FilterStatus)
	return ret0
}

// OnData indicates an expected call of OnData
func (mr *MockReadFilterMockRecorder) OnData(buffer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnData", reflect.TypeOf((*MockReadFilter)(nil).OnData), buffer)
}

// OnNewConnection mocks base method
func (m *MockReadFilter) OnNewConnection() api.FilterStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnNewConnection")
	ret0, _ := ret[0].(api.FilterStatus)
	return ret0
}

// OnNewConnection indicates an expected call of OnNewConnection
func (mr *MockReadFilterMockRecorder) OnNewConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNewConnection", reflect.TypeOf((*MockReadFilter)(nil).OnNewConnection))
}

// InitializeReadFilterCallbacks mocks base method
func (m *MockReadFilter) InitializeReadFilterCallbacks(cb api.ReadFilterCallbacks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitializeReadFilterCallbacks", cb)
}

// InitializeReadFilterCallbacks indicates an expected call of InitializeReadFilterCallbacks
func (mr *MockReadFilterMockRecorder) InitializeReadFilterCallbacks(cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeReadFilterCallbacks", reflect.TypeOf((*MockReadFilter)(nil).InitializeReadFilterCallbacks), cb)
}

// MockWriteFilter is a mock of WriteFilter interface
type MockWriteFilter struct {
	ctrl     *gomock.Controller
	recorder *MockWriteFilterMockRecorder
}

// MockWriteFilterMockRecorder is the mock recorder for MockWriteFilter
type MockWriteFilterMockRecorder struct {
	mock *MockWriteFilter
}

// NewMockWriteFilter creates a new mock instance
func NewMockWriteFilter(ctrl *gomock.Controller) *MockWriteFilter {
	mock := &MockWriteFilter{ctrl: ctrl}
	mock.recorder = &MockWriteFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriteFilter) EXPECT() *MockWriteFilterMockRecorder {
	return m.recorder
}

// OnWrite mocks base method
func (m *MockWriteFilter) OnWrite(buffer []buffer.IoBuffer) api.FilterStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnWrite", buffer)
	ret0, _ := ret[0].(api.FilterStatus)
	return ret0
}

// OnWrite indicates an expected call of OnWrite
func (mr *MockWriteFilterMockRecorder) OnWrite(buffer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnWrite", reflect.TypeOf((*MockWriteFilter)(nil).OnWrite), buffer)
}

// MockReadFilterCallbacks is a mock of ReadFilterCallbacks interface
type MockReadFilterCallbacks struct {
	ctrl     *gomock.Controller
	recorder *MockReadFilterCallbacksMockRecorder
}

// MockReadFilterCallbacksMockRecorder is the mock recorder for MockReadFilterCallbacks
type MockReadFilterCallbacksMockRecorder struct {
	mock *MockReadFilterCallbacks
}

// NewMockReadFilterCallbacks creates a new mock instance
func NewMockReadFilterCallbacks(ctrl *gomock.Controller) *MockReadFilterCallbacks {
	mock := &MockReadFilterCallbacks{ctrl: ctrl}
	mock.recorder = &MockReadFilterCallbacksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadFilterCallbacks) EXPECT() *MockReadFilterCallbacksMockRecorder {
	return m.recorder
}

// Connection mocks base method
func (m *MockReadFilterCallbacks) Connection() api.Connection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connection")
	ret0, _ := ret[0].(api.Connection)
	return ret0
}

// Connection indicates an expected call of Connection
func (mr *MockReadFilterCallbacksMockRecorder) Connection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connection", reflect.TypeOf((*MockReadFilterCallbacks)(nil).Connection))
}

// ContinueReading mocks base method
func (m *MockReadFilterCallbacks) ContinueReading() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ContinueReading")
}

// ContinueReading indicates an expected call of ContinueReading
func (mr *MockReadFilterCallbacksMockRecorder) ContinueReading() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContinueReading", reflect.TypeOf((*MockReadFilterCallbacks)(nil).ContinueReading))
}

// UpstreamHost mocks base method
func (m *MockReadFilterCallbacks) UpstreamHost() api.HostInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpstreamHost")
	ret0, _ := ret[0].(api.HostInfo)
	return ret0
}

// UpstreamHost indicates an expected call of UpstreamHost
func (mr *MockReadFilterCallbacksMockRecorder) UpstreamHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpstreamHost", reflect.TypeOf((*MockReadFilterCallbacks)(nil).UpstreamHost))
}

// SetUpstreamHost mocks base method
func (m *MockReadFilterCallbacks) SetUpstreamHost(upstreamHost api.HostInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUpstreamHost", upstreamHost)
}

// SetUpstreamHost indicates an expected call of SetUpstreamHost
func (mr *MockReadFilterCallbacksMockRecorder) SetUpstreamHost(upstreamHost interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUpstreamHost", reflect.TypeOf((*MockReadFilterCallbacks)(nil).SetUpstreamHost), upstreamHost)
}

// MockNetWorkFilterChainFactoryCallbacks is a mock of NetWorkFilterChainFactoryCallbacks interface
type MockNetWorkFilterChainFactoryCallbacks struct {
	ctrl     *gomock.Controller
	recorder *MockNetWorkFilterChainFactoryCallbacksMockRecorder
}

// MockNetWorkFilterChainFactoryCallbacksMockRecorder is the mock recorder for MockNetWorkFilterChainFactoryCallbacks
type MockNetWorkFilterChainFactoryCallbacksMockRecorder struct {
	mock *MockNetWorkFilterChainFactoryCallbacks
}

// NewMockNetWorkFilterChainFactoryCallbacks creates a new mock instance
func NewMockNetWorkFilterChainFactoryCallbacks(ctrl *gomock.Controller) *MockNetWorkFilterChainFactoryCallbacks {
	mock := &MockNetWorkFilterChainFactoryCallbacks{ctrl: ctrl}
	mock.recorder = &MockNetWorkFilterChainFactoryCallbacksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetWorkFilterChainFactoryCallbacks) EXPECT() *MockNetWorkFilterChainFactoryCallbacksMockRecorder {
	return m.recorder
}

// AddReadFilter mocks base method
func (m *MockNetWorkFilterChainFactoryCallbacks) AddReadFilter(rf api.ReadFilter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddReadFilter", rf)
}

// AddReadFilter indicates an expected call of AddReadFilter
func (mr *MockNetWorkFilterChainFactoryCallbacksMockRecorder) AddReadFilter(rf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReadFilter", reflect.TypeOf((*MockNetWorkFilterChainFactoryCallbacks)(nil).AddReadFilter), rf)
}

// AddWriteFilter mocks base method
func (m *MockNetWorkFilterChainFactoryCallbacks) AddWriteFilter(wf api.WriteFilter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddWriteFilter", wf)
}

// AddWriteFilter indicates an expected call of AddWriteFilter
func (mr *MockNetWorkFilterChainFactoryCallbacksMockRecorder) AddWriteFilter(wf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddWriteFilter", reflect.TypeOf((*MockNetWorkFilterChainFactoryCallbacks)(nil).AddWriteFilter), wf)
}

// MockNetworkFilterChainFactory is a mock of NetworkFilterChainFactory interface
type MockNetworkFilterChainFactory struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkFilterChainFactoryMockRecorder
}

// MockNetworkFilterChainFactoryMockRecorder is the mock recorder for MockNetworkFilterChainFactory
type MockNetworkFilterChainFactoryMockRecorder struct {
	mock *MockNetworkFilterChainFactory
}

// NewMockNetworkFilterChainFactory creates a new mock instance
func NewMockNetworkFilterChainFactory(ctrl *gomock.Controller) *MockNetworkFilterChainFactory {
	mock := &MockNetworkFilterChainFactory{ctrl: ctrl}
	mock.recorder = &MockNetworkFilterChainFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkFilterChainFactory) EXPECT() *MockNetworkFilterChainFactoryMockRecorder {
	return m.recorder
}

// CreateFilterChain mocks base method
func (m *MockNetworkFilterChainFactory) CreateFilterChain(context context.Context, callbacks api.NetWorkFilterChainFactoryCallbacks) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateFilterChain", context, callbacks)
}

// CreateFilterChain indicates an expected call of CreateFilterChain
func (mr *MockNetworkFilterChainFactoryMockRecorder) CreateFilterChain(context, callbacks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFilterChain", reflect.TypeOf((*MockNetworkFilterChainFactory)(nil).CreateFilterChain), context, callbacks)
}
