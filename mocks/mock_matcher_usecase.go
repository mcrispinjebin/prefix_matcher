// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/matcher.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	models "prefix_matcher/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWordMatcher is a mock of WordMatcher interface.
type MockWordMatcher struct {
	ctrl     *gomock.Controller
	recorder *MockWordMatcherMockRecorder
}

// MockWordMatcherMockRecorder is the mock recorder for MockWordMatcher.
type MockWordMatcherMockRecorder struct {
	mock *MockWordMatcher
}

// NewMockWordMatcher creates a new mock instance.
func NewMockWordMatcher(ctrl *gomock.Controller) *MockWordMatcher {
	mock := &MockWordMatcher{ctrl: ctrl}
	mock.recorder = &MockWordMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWordMatcher) EXPECT() *MockWordMatcherMockRecorder {
	return m.recorder
}

// AppropriateByBinarySearch mocks base method.
func (m *MockWordMatcher) AppropriateByBinarySearch(wordsList []string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppropriateByBinarySearch", wordsList)
	ret0, _ := ret[0].(int)
	return ret0
}

// AppropriateByBinarySearch indicates an expected call of AppropriateByBinarySearch.
func (mr *MockWordMatcherMockRecorder) AppropriateByBinarySearch(wordsList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppropriateByBinarySearch", reflect.TypeOf((*MockWordMatcher)(nil).AppropriateByBinarySearch), wordsList)
}

// FindExactLongWord mocks base method.
func (m *MockWordMatcher) FindExactLongWord(wordsList []string, startValue, incrementalValue int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindExactLongWord", wordsList, startValue, incrementalValue)
}

// FindExactLongWord indicates an expected call of FindExactLongWord.
func (mr *MockWordMatcherMockRecorder) FindExactLongWord(wordsList, startValue, incrementalValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindExactLongWord", reflect.TypeOf((*MockWordMatcher)(nil).FindExactLongWord), wordsList, startValue, incrementalValue)
}

// FindWordByProbing mocks base method.
func (m *MockWordMatcher) FindWordByProbing(wordsList []string, i int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FindWordByProbing", wordsList, i)
}

// FindWordByProbing indicates an expected call of FindWordByProbing.
func (mr *MockWordMatcherMockRecorder) FindWordByProbing(wordsList, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindWordByProbing", reflect.TypeOf((*MockWordMatcher)(nil).FindWordByProbing), wordsList, i)
}

// Process mocks base method.
func (m *MockWordMatcher) Process(c chan<- models.MatcherParam) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Process", c)
}

// Process indicates an expected call of Process.
func (mr *MockWordMatcherMockRecorder) Process(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockWordMatcher)(nil).Process), c)
}

// ReadAppropriateSubFile mocks base method.
func (m *MockWordMatcher) ReadAppropriateSubFile() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAppropriateSubFile")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ReadAppropriateSubFile indicates an expected call of ReadAppropriateSubFile.
func (mr *MockWordMatcherMockRecorder) ReadAppropriateSubFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAppropriateSubFile", reflect.TypeOf((*MockWordMatcher)(nil).ReadAppropriateSubFile))
}
