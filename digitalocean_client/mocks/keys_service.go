package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockKeysService struct {
	ctrl     *gomock.Controller
	recorder *MockKeysServiceMockRecorder
}

type MockKeysServiceMockRecorder struct {
	mock *MockKeysService
}

func NewMockKeysService(ctrl *gomock.Controller) *MockKeysService {
	mock := &MockKeysService{ctrl: ctrl}
	mock.recorder = &MockKeysServiceMockRecorder{mock}
	return mock
}

func (m *MockKeysService) EXPECT() *MockKeysServiceMockRecorder {
	return m.recorder
}

func (m *MockKeysService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Key, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.Key)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockKeysServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockKeysService)(nil).List), arg0, arg1)
}
