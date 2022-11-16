package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockRegistryService struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryServiceMockRecorder
}

type MockRegistryServiceMockRecorder struct {
	mock *MockRegistryService
}

func NewMockRegistryService(ctrl *gomock.Controller) *MockRegistryService {
	mock := &MockRegistryService{ctrl: ctrl}
	mock.recorder = &MockRegistryServiceMockRecorder{mock}
	return mock
}

func (m *MockRegistryService) EXPECT() *MockRegistryServiceMockRecorder {
	return m.recorder
}

func (m *MockRegistryService) Get(arg0 context.Context) (*godo.Registry, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Get, arg0)
	ret0, _ := ret[0].(*godo.Registry)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockRegistryServiceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Get, reflect.TypeOf((*MockRegistryService)(nil).Get), arg0)
}

func (m *MockRegistryService) ListRepositories(arg0 context.Context, arg1 string, arg2 *godo.ListOptions) ([]*godo.Repository, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListRepositories, arg0, arg1, arg2)
	ret0, _ := ret[0].([]*godo.Repository)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockRegistryServiceMockRecorder) ListRepositories(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListRepositories, reflect.TypeOf((*MockRegistryService)(nil).ListRepositories), arg0, arg1, arg2)
}
