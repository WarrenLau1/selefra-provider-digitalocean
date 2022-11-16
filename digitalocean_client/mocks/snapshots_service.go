package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockSnapshotsService struct {
	ctrl     *gomock.Controller
	recorder *MockSnapshotsServiceMockRecorder
}

type MockSnapshotsServiceMockRecorder struct {
	mock *MockSnapshotsService
}

func NewMockSnapshotsService(ctrl *gomock.Controller) *MockSnapshotsService {
	mock := &MockSnapshotsService{ctrl: ctrl}
	mock.recorder = &MockSnapshotsServiceMockRecorder{mock}
	return mock
}

func (m *MockSnapshotsService) EXPECT() *MockSnapshotsServiceMockRecorder {
	return m.recorder
}

func (m *MockSnapshotsService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.Snapshot, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.Snapshot)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockSnapshotsServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockSnapshotsService)(nil).List), arg0, arg1)
}
