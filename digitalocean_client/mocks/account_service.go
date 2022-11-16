package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

func (m *MockAccountService) Get(arg0 context.Context) (*godo.Account, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0)
	ret0, _ := ret[0].(*godo.Account)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockAccountServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockAccountService)(nil).Get), arg0)
}
