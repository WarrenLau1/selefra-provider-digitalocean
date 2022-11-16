package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockCdnService struct {
	ctrl     *gomock.Controller
	recorder *MockCdnServiceMockRecorder
}

type MockCdnServiceMockRecorder struct {
	mock *MockCdnService
}

func NewMockCdnService(ctrl *gomock.Controller) *MockCdnService {
	mock := &MockCdnService{ctrl: ctrl}
	mock.recorder = &MockCdnServiceMockRecorder{mock}
	return mock
}

func (m *MockCdnService) EXPECT() *MockCdnServiceMockRecorder {
	return m.recorder
}

func (m *MockCdnService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.CDN, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.CDN)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockCdnServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockCdnService)(nil).List), arg0, arg1)
}
