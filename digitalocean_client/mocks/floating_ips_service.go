package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockFloatingIpsService struct {
	ctrl     *gomock.Controller
	recorder *MockFloatingIpsServiceMockRecorder
}

type MockFloatingIpsServiceMockRecorder struct {
	mock *MockFloatingIpsService
}

func NewMockFloatingIpsService(ctrl *gomock.Controller) *MockFloatingIpsService {
	mock := &MockFloatingIpsService{ctrl: ctrl}
	mock.recorder = &MockFloatingIpsServiceMockRecorder{mock}
	return mock
}

func (m *MockFloatingIpsService) EXPECT() *MockFloatingIpsServiceMockRecorder {
	return m.recorder
}

func (m *MockFloatingIpsService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.FloatingIP, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.FloatingIP)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockFloatingIpsServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockFloatingIpsService)(nil).List), arg0, arg1)
}
