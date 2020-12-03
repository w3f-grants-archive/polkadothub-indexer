// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/figment-networks/polkadothub-indexer/indexer (interfaces: ConfigParser,FetcherClient,RewardsCalculator)

// Package mock_indexer is a generated GoMock package.
package mock_indexer

import (
	pipeline "github.com/figment-networks/indexing-engine/pipeline"
	heightpb "github.com/figment-networks/polkadothub-proxy/grpc/height/heightpb"
	gomock "github.com/golang/mock/gomock"
	big "math/big"
	reflect "reflect"
)

// MockConfigParser is a mock of ConfigParser interface
type MockConfigParser struct {
	ctrl     *gomock.Controller
	recorder *MockConfigParserMockRecorder
}

// MockConfigParserMockRecorder is the mock recorder for MockConfigParser
type MockConfigParserMockRecorder struct {
	mock *MockConfigParser
}

// NewMockConfigParser creates a new mock instance
func NewMockConfigParser(ctrl *gomock.Controller) *MockConfigParser {
	mock := &MockConfigParser{ctrl: ctrl}
	mock.recorder = &MockConfigParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigParser) EXPECT() *MockConfigParserMockRecorder {
	return m.recorder
}

// GetAllAvailableTasks mocks base method
func (m *MockConfigParser) GetAllAvailableTasks() []pipeline.TaskName {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAvailableTasks")
	ret0, _ := ret[0].([]pipeline.TaskName)
	return ret0
}

// GetAllAvailableTasks indicates an expected call of GetAllAvailableTasks
func (mr *MockConfigParserMockRecorder) GetAllAvailableTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAvailableTasks", reflect.TypeOf((*MockConfigParser)(nil).GetAllAvailableTasks))
}

// GetAllVersionedTasks mocks base method
func (m *MockConfigParser) GetAllVersionedTasks() ([]pipeline.TaskName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVersionedTasks")
	ret0, _ := ret[0].([]pipeline.TaskName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllVersionedTasks indicates an expected call of GetAllVersionedTasks
func (mr *MockConfigParserMockRecorder) GetAllVersionedTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVersionedTasks", reflect.TypeOf((*MockConfigParser)(nil).GetAllVersionedTasks))
}

// GetAllVersionedVersionIds mocks base method
func (m *MockConfigParser) GetAllVersionedVersionIds() []int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVersionedVersionIds")
	ret0, _ := ret[0].([]int64)
	return ret0
}

// GetAllVersionedVersionIds indicates an expected call of GetAllVersionedVersionIds
func (mr *MockConfigParserMockRecorder) GetAllVersionedVersionIds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVersionedVersionIds", reflect.TypeOf((*MockConfigParser)(nil).GetAllVersionedVersionIds))
}

// GetCurrentVersionId mocks base method
func (m *MockConfigParser) GetCurrentVersionId() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentVersionId")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetCurrentVersionId indicates an expected call of GetCurrentVersionId
func (mr *MockConfigParserMockRecorder) GetCurrentVersionId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentVersionId", reflect.TypeOf((*MockConfigParser)(nil).GetCurrentVersionId))
}

// GetTasksByTargetIds mocks base method
func (m *MockConfigParser) GetTasksByTargetIds(arg0 []int64) ([]pipeline.TaskName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasksByTargetIds", arg0)
	ret0, _ := ret[0].([]pipeline.TaskName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksByTargetIds indicates an expected call of GetTasksByTargetIds
func (mr *MockConfigParserMockRecorder) GetTasksByTargetIds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksByTargetIds", reflect.TypeOf((*MockConfigParser)(nil).GetTasksByTargetIds), arg0)
}

// GetTasksByVersionIds mocks base method
func (m *MockConfigParser) GetTasksByVersionIds(arg0 []int64) ([]pipeline.TaskName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasksByVersionIds", arg0)
	ret0, _ := ret[0].([]pipeline.TaskName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksByVersionIds indicates an expected call of GetTasksByVersionIds
func (mr *MockConfigParserMockRecorder) GetTasksByVersionIds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksByVersionIds", reflect.TypeOf((*MockConfigParser)(nil).GetTasksByVersionIds), arg0)
}

// IsAnyVersionSequential mocks base method
func (m *MockConfigParser) IsAnyVersionSequential(arg0 []int64) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAnyVersionSequential", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAnyVersionSequential indicates an expected call of IsAnyVersionSequential
func (mr *MockConfigParserMockRecorder) IsAnyVersionSequential(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAnyVersionSequential", reflect.TypeOf((*MockConfigParser)(nil).IsAnyVersionSequential), arg0)
}

// IsLastInEra mocks base method
func (m *MockConfigParser) IsLastInEra() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLastInEra")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLastInEra indicates an expected call of IsLastInEra
func (mr *MockConfigParserMockRecorder) IsLastInEra() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLastInEra", reflect.TypeOf((*MockConfigParser)(nil).IsLastInEra))
}

// IsLastInSession mocks base method
func (m *MockConfigParser) IsLastInSession() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLastInSession")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLastInSession indicates an expected call of IsLastInSession
func (mr *MockConfigParserMockRecorder) IsLastInSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLastInSession", reflect.TypeOf((*MockConfigParser)(nil).IsLastInSession))
}

// MockFetcherClient is a mock of FetcherClient interface
type MockFetcherClient struct {
	ctrl     *gomock.Controller
	recorder *MockFetcherClientMockRecorder
}

// MockFetcherClientMockRecorder is the mock recorder for MockFetcherClient
type MockFetcherClientMockRecorder struct {
	mock *MockFetcherClient
}

// NewMockFetcherClient creates a new mock instance
func NewMockFetcherClient(ctrl *gomock.Controller) *MockFetcherClient {
	mock := &MockFetcherClient{ctrl: ctrl}
	mock.recorder = &MockFetcherClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFetcherClient) EXPECT() *MockFetcherClientMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockFetcherClient) GetAll(arg0 int64) (*heightpb.GetAllResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].(*heightpb.GetAllResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockFetcherClientMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockFetcherClient)(nil).GetAll), arg0)
}

// MockRewardsCalculator is a mock of RewardsCalculator interface
type MockRewardsCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockRewardsCalculatorMockRecorder
}

// MockRewardsCalculatorMockRecorder is the mock recorder for MockRewardsCalculator
type MockRewardsCalculatorMockRecorder struct {
	mock *MockRewardsCalculator
}

// NewMockRewardsCalculator creates a new mock instance
func NewMockRewardsCalculator(ctrl *gomock.Controller) *MockRewardsCalculator {
	mock := &MockRewardsCalculator{ctrl: ctrl}
	mock.recorder = &MockRewardsCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRewardsCalculator) EXPECT() *MockRewardsCalculatorMockRecorder {
	return m.recorder
}

// commissionPayout mocks base method
func (m *MockRewardsCalculator) commissionPayout(arg0, arg1 int64) (big.Int, big.Int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "commissionPayout", arg0, arg1)
	ret0, _ := ret[0].(big.Int)
	ret1, _ := ret[1].(big.Int)
	return ret0, ret1
}

// commissionPayout indicates an expected call of commissionPayout
func (mr *MockRewardsCalculatorMockRecorder) commissionPayout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "commissionPayout", reflect.TypeOf((*MockRewardsCalculator)(nil).commissionPayout), arg0, arg1)
}

// nominatorPayout mocks base method
func (m *MockRewardsCalculator) nominatorPayout(arg0, arg1, arg2 big.Int) big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "nominatorPayout", arg0, arg1, arg2)
	ret0, _ := ret[0].(big.Int)
	return ret0
}

// nominatorPayout indicates an expected call of nominatorPayout
func (mr *MockRewardsCalculatorMockRecorder) nominatorPayout(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "nominatorPayout", reflect.TypeOf((*MockRewardsCalculator)(nil).nominatorPayout), arg0, arg1, arg2)
}
