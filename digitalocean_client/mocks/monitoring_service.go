package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockMonitoringService struct {
	ctrl     *gomock.Controller
	recorder *MockMonitoringServiceMockRecorder
}

type MockMonitoringServiceMockRecorder struct {
	mock *MockMonitoringService
}

func NewMockMonitoringService(ctrl *gomock.Controller) *MockMonitoringService {
	mock := &MockMonitoringService{ctrl: ctrl}
	mock.recorder = &MockMonitoringServiceMockRecorder{mock}
	return mock
}

func (m *MockMonitoringService) EXPECT() *MockMonitoringServiceMockRecorder {
	return m.recorder
}

func (m *MockMonitoringService) ListAlertPolicies(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.AlertPolicy, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListAlertPolicies, arg0, arg1)
	ret0, _ := ret[0].([]godo.AlertPolicy)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockMonitoringServiceMockRecorder) ListAlertPolicies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListAlertPolicies, reflect.TypeOf((*MockMonitoringService)(nil).ListAlertPolicies), arg0, arg1)
}
