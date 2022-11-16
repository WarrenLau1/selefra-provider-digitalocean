package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockDomainsService struct {
	ctrl     *gomock.Controller
	recorder *MockDomainsServiceMockRecorder
}

type MockDomainsServiceMockRecorder struct {
	mock *MockDomainsService
}

func NewMockDomainsService(ctrl *gomock.Controller) *MockDomainsService {
	mock := &MockDomainsService{ctrl: ctrl}
	mock.recorder = &MockDomainsServiceMockRecorder{mock}
	return mock
}

func (m *MockDomainsService) EXPECT() *MockDomainsServiceMockRecorder {
	return m.recorder
}

func (m *MockDomainsService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Domain, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.Domain)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockDomainsServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockDomainsService)(nil).List), arg0, arg1)
}

func (m *MockDomainsService) Records(arg0 context.Context, arg1 string, arg2 *godo.ListOptions) ([]godo.DomainRecord, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.Records, arg0, arg1, arg2)
	ret0, _ := ret[0].([]godo.DomainRecord)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockDomainsServiceMockRecorder) Records(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.Records, reflect.TypeOf((*MockDomainsService)(nil).Records), arg0, arg1, arg2)
}
