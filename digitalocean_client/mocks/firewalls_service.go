package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockFirewallsService struct {
	ctrl     *gomock.Controller
	recorder *MockFirewallsServiceMockRecorder
}

type MockFirewallsServiceMockRecorder struct {
	mock *MockFirewallsService
}

func NewMockFirewallsService(ctrl *gomock.Controller) *MockFirewallsService {
	mock := &MockFirewallsService{ctrl: ctrl}
	mock.recorder = &MockFirewallsServiceMockRecorder{mock}
	return mock
}

func (m *MockFirewallsService) EXPECT() *MockFirewallsServiceMockRecorder {
	return m.recorder
}

func (m *MockFirewallsService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Firewall, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.Firewall)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockFirewallsServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockFirewallsService)(nil).List), arg0, arg1)
}
