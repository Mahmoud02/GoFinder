package mocks

import (
	"GoFinder/linkgraph/graph"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"reflect"
	"time"
)

// MockGraphAPI is a mock of GraphAPI interface
type MockGraphAPI struct {
	ctrl     *gomock.Controller
	recorder *MockGraphAPIMockRecorder
}

// MockGraphAPIMockRecorder is the mock recorder for MockGraphAPI
type MockGraphAPIMockRecorder struct {
	mock *MockGraphAPI
}

// NewMockGraphAPI creates a new mock instance
func NewMockGraphAPI(ctrl *gomock.Controller) *MockGraphAPI {
	mock := &MockGraphAPI{ctrl: ctrl}
	mock.recorder = &MockGraphAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGraphAPI) EXPECT() *MockGraphAPIMockRecorder {
	return m.recorder
}

// Edges mocks base method
func (m *MockGraphAPI) Edges(arg0, arg1 uuid.UUID, arg2 time.Time) (graph.EdgeIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edges", arg0, arg1, arg2)
	ret0, _ := ret[0].(graph.EdgeIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Edges indicates an expected call of Edges
func (mr *MockGraphAPIMockRecorder) Edges(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edges", reflect.TypeOf((*MockGraphAPI)(nil).Edges), arg0, arg1, arg2)
}

// Links mocks base method
func (m *MockGraphAPI) Links(arg0, arg1 uuid.UUID, arg2 time.Time) (graph.LinkIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Links", arg0, arg1, arg2)
	ret0, _ := ret[0].(graph.LinkIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Links indicates an expected call of Links
func (mr *MockGraphAPIMockRecorder) Links(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Links", reflect.TypeOf((*MockGraphAPI)(nil).Links), arg0, arg1, arg2)
}

// MockIndexAPI is a mock of IndexAPI interface
type MockIndexAPI struct {
	ctrl     *gomock.Controller
	recorder *MockIndexAPIMockRecorder
}

// MockIndexAPIMockRecorder is the mock recorder for MockIndexAPI
type MockIndexAPIMockRecorder struct {
	mock *MockIndexAPI
}

// NewMockIndexAPI creates a new mock instance
func NewMockIndexAPI(ctrl *gomock.Controller) *MockIndexAPI {
	mock := &MockIndexAPI{ctrl: ctrl}
	mock.recorder = &MockIndexAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIndexAPI) EXPECT() *MockIndexAPIMockRecorder {
	return m.recorder
}

// UpdateScore mocks base method
func (m *MockIndexAPI) UpdateScore(arg0 uuid.UUID, arg1 float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScore", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateScore indicates an expected call of UpdateScore
func (mr *MockIndexAPIMockRecorder) UpdateScore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScore", reflect.TypeOf((*MockIndexAPI)(nil).UpdateScore), arg0, arg1)
}
