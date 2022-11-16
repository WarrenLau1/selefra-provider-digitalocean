package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockProjectsService struct {
	ctrl     *gomock.Controller
	recorder *MockProjectsServiceMockRecorder
}

type MockProjectsServiceMockRecorder struct {
	mock *MockProjectsService
}

func NewMockProjectsService(ctrl *gomock.Controller) *MockProjectsService {
	mock := &MockProjectsService{ctrl: ctrl}
	mock.recorder = &MockProjectsServiceMockRecorder{mock}
	return mock
}

func (m *MockProjectsService) EXPECT() *MockProjectsServiceMockRecorder {
	return m.recorder
}

func (m *MockProjectsService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Project, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.Project)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockProjectsServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockProjectsService)(nil).List), arg0, arg1)
}

func (m *MockProjectsService) ListResources(arg0 context.Context, arg1 string, arg2 *godo.ListOptions) ([]godo.ProjectResource, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListResources, arg0, arg1, arg2)
	ret0, _ := ret[0].([]godo.ProjectResource)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockProjectsServiceMockRecorder) ListResources(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListResources, reflect.TypeOf((*MockProjectsService)(nil).ListResources), arg0, arg1, arg2)
}
