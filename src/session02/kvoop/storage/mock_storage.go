// Automatically generated by MockGen. DO NOT EDIT!
// Source: src/session02/kvoop/storage/storage.go

package storage

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockStorageRecorder
}

// Recorder for MockStorage (not exported)
type _MockStorageRecorder struct {
	mock *MockStorage
}

func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &_MockStorageRecorder{mock}
	return mock
}

func (_m *MockStorage) EXPECT() *_MockStorageRecorder {
	return _m.recorder
}

func (_m *MockStorage) Get(key string) (interface{}, bool) {
	ret := _m.ctrl.Call(_m, "Get", key)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

func (_mr *_MockStorageRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockStorage) Set(key string, value string) {
	_m.ctrl.Call(_m, "Set", key, value)
}

func (_mr *_MockStorageRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Set", arg0, arg1)
}

func (_m *MockStorage) All() KeyValueMap {
	ret := _m.ctrl.Call(_m, "All")
	ret0, _ := ret[0].(KeyValueMap)
	return ret0
}

func (_mr *_MockStorageRecorder) All() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "All")
}