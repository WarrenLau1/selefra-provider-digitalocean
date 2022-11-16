package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockVpcsService struct {
	ctrl     *gomock.Controller
	recorder *MockVpcsServiceMockRecorder
}

type MockVpcsServiceMockRecorder struct {
	mock *MockVpcsService
}

func NewMockVpcsService(ctrl *gomock.Controller) *MockVpcsService {
	mock := &MockVpcsService{ctrl: ctrl}
	mock.recorder = &MockVpcsServiceMockRecorder{mock}
	return mock
}

func (m *MockVpcsService) EXPECT() *MockVpcsServiceMockRecorder {
	return m.recorder
}

func (m *MockVpcsService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]*godo.VPC, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]*godo.VPC)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockVpcsServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockVpcsService)(nil).List), arg0, arg1)
}

func (m *MockVpcsService) ListMembers(arg0 context.Context, arg1 string, arg2 *godo.VPCListMembersRequest, arg3 *godo.ListOptions) ([]*godo.VPCMember, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListMembers, arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*godo.VPCMember)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockVpcsServiceMockRecorder) ListMembers(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListMembers, reflect.TypeOf((*MockVpcsService)(nil).ListMembers), arg0, arg1, arg2, arg3)
}
