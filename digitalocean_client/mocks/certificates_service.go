package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockCertificatesService struct {
	ctrl     *gomock.Controller
	recorder *MockCertificatesServiceMockRecorder
}

type MockCertificatesServiceMockRecorder struct {
	mock *MockCertificatesService
}

func NewMockCertificatesService(ctrl *gomock.Controller) *MockCertificatesService {
	mock := &MockCertificatesService{ctrl: ctrl}
	mock.recorder = &MockCertificatesServiceMockRecorder{mock}
	return mock
}

func (m *MockCertificatesService) EXPECT() *MockCertificatesServiceMockRecorder {
	return m.recorder
}

func (m *MockCertificatesService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Certificate, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.Certificate)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockCertificatesServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockCertificatesService)(nil).List), arg0, arg1)
}
