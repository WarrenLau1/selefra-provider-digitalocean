package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockDatabasesService struct {
	ctrl     *gomock.Controller
	recorder *MockDatabasesServiceMockRecorder
}

type MockDatabasesServiceMockRecorder struct {
	mock *MockDatabasesService
}

func NewMockDatabasesService(ctrl *gomock.Controller) *MockDatabasesService {
	mock := &MockDatabasesService{ctrl: ctrl}
	mock.recorder = &MockDatabasesServiceMockRecorder{mock}
	return mock
}

func (m *MockDatabasesService) EXPECT() *MockDatabasesServiceMockRecorder {
	return m.recorder
}

func (m *MockDatabasesService) GetFirewallRules(arg0 context.Context, arg1 string) ([]godo.DatabaseFirewallRule, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.GetFirewallRules, arg0, arg1)
	ret0, _ := ret[0].([]godo.DatabaseFirewallRule)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockDatabasesServiceMockRecorder) GetFirewallRules(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetFirewallRules, reflect.TypeOf((*MockDatabasesService)(nil).GetFirewallRules), arg0, arg1)
}

func (m *MockDatabasesService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Database, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.Database)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockDatabasesServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockDatabasesService)(nil).List), arg0, arg1)
}

func (m *MockDatabasesService) ListBackups(arg0 context.Context, arg1 string, arg2 *godo.ListOptions) ([]godo.DatabaseBackup, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListBackups, arg0, arg1, arg2)
	ret0, _ := ret[0].([]godo.DatabaseBackup)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockDatabasesServiceMockRecorder) ListBackups(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBackups, reflect.TypeOf((*MockDatabasesService)(nil).ListBackups), arg0, arg1, arg2)
}

func (m *MockDatabasesService) ListReplicas(arg0 context.Context, arg1 string, arg2 *godo.ListOptions) ([]godo.DatabaseReplica, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListReplicas, arg0, arg1, arg2)
	ret0, _ := ret[0].([]godo.DatabaseReplica)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockDatabasesServiceMockRecorder) ListReplicas(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListReplicas, reflect.TypeOf((*MockDatabasesService)(nil).ListReplicas), arg0, arg1, arg2)
}
