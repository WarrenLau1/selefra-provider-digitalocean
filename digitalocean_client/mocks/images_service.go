package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockImagesService struct {
	ctrl     *gomock.Controller
	recorder *MockImagesServiceMockRecorder
}

type MockImagesServiceMockRecorder struct {
	mock *MockImagesService
}

func NewMockImagesService(ctrl *gomock.Controller) *MockImagesService {
	mock := &MockImagesService{ctrl: ctrl}
	mock.recorder = &MockImagesServiceMockRecorder{mock}
	return mock
}

func (m *MockImagesService) EXPECT() *MockImagesServiceMockRecorder {
	return m.recorder
}

func (m *MockImagesService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Image, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.Image)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockImagesServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockImagesService)(nil).List), arg0, arg1)
}
