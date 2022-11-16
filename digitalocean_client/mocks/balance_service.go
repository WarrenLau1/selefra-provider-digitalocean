package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockBalanceService struct {
	ctrl     *gomock.Controller
	recorder *MockBalanceServiceMockRecorder
}

type MockBalanceServiceMockRecorder struct {
	mock *MockBalanceService
}

func NewMockBalanceService(ctrl *gomock.Controller) *MockBalanceService {
	mock := &MockBalanceService{ctrl: ctrl}
	mock.recorder = &MockBalanceServiceMockRecorder{mock}
	return mock
}

func (m *MockBalanceService) EXPECT() *MockBalanceServiceMockRecorder {
	return m.recorder
}

func (m *MockBalanceService) Get(arg0 context.Context) (*godo.Balance, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0)
	ret0, _ := ret[0].(*godo.Balance)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockBalanceServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockBalanceService)(nil).Get), arg0)
}
