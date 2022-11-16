package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockRegionsService struct {
	ctrl     *gomock.Controller
	recorder *MockRegionsServiceMockRecorder
}

type MockRegionsServiceMockRecorder struct {
	mock *MockRegionsService
}

func NewMockRegionsService(ctrl *gomock.Controller) *MockRegionsService {
	mock := &MockRegionsService{ctrl: ctrl}
	mock.recorder = &MockRegionsServiceMockRecorder{mock}
	return mock
}

func (m *MockRegionsService) EXPECT() *MockRegionsServiceMockRecorder {
	return m.recorder
}

func (m *MockRegionsService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Region, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.Region)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockRegionsServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockRegionsService)(nil).List), arg0, arg1)
}
