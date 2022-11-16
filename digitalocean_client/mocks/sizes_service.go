package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockSizesService struct {
	ctrl     *gomock.Controller
	recorder *MockSizesServiceMockRecorder
}

type MockSizesServiceMockRecorder struct {
	mock *MockSizesService
}

func NewMockSizesService(ctrl *gomock.Controller) *MockSizesService {
	mock := &MockSizesService{ctrl: ctrl}
	mock.recorder = &MockSizesServiceMockRecorder{mock}
	return mock
}

func (m *MockSizesService) EXPECT() *MockSizesServiceMockRecorder {
	return m.recorder
}

func (m *MockSizesService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Size, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.Size)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockSizesServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockSizesService)(nil).List), arg0, arg1)
}
