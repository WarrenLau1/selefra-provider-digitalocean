package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockBillingHistoryService struct {
	ctrl     *gomock.Controller
	recorder *MockBillingHistoryServiceMockRecorder
}

type MockBillingHistoryServiceMockRecorder struct {
	mock *MockBillingHistoryService
}

func NewMockBillingHistoryService(ctrl *gomock.Controller) *MockBillingHistoryService {
	mock := &MockBillingHistoryService{ctrl: ctrl}
	mock.recorder = &MockBillingHistoryServiceMockRecorder{mock}
	return mock
}

func (m *MockBillingHistoryService) EXPECT() *MockBillingHistoryServiceMockRecorder {
	return m.recorder
}

func (m *MockBillingHistoryService) List(arg0 context.Context, arg1 *godo.ListOptions) (*godo.BillingHistory, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].(*godo.BillingHistory)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockBillingHistoryServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockBillingHistoryService)(nil).List), arg0, arg1)
}
