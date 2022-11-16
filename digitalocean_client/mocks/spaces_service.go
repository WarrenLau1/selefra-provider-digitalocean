package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "github.com/golang/mock/gomock"
)

type MockSpacesService struct {
	ctrl     *gomock.Controller
	recorder *MockSpacesServiceMockRecorder
}

type MockSpacesServiceMockRecorder struct {
	mock *MockSpacesService
}

func NewMockSpacesService(ctrl *gomock.Controller) *MockSpacesService {
	mock := &MockSpacesService{ctrl: ctrl}
	mock.recorder = &MockSpacesServiceMockRecorder{mock}
	return mock
}

func (m *MockSpacesService) EXPECT() *MockSpacesServiceMockRecorder {
	return m.recorder
}

func (m *MockSpacesService) GetBucketAcl(arg0 context.Context, arg1 *s3.GetBucketAclInput, arg2 ...func(*s3.Options)) (*s3.GetBucketAclOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketAcl, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketAclOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSpacesServiceMockRecorder) GetBucketAcl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketAcl, reflect.TypeOf((*MockSpacesService)(nil).GetBucketAcl), varargs...)
}

func (m *MockSpacesService) GetBucketCors(arg0 context.Context, arg1 *s3.GetBucketCorsInput, arg2 ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.GetBucketCors, varargs...)
	ret0, _ := ret[0].(*s3.GetBucketCorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSpacesServiceMockRecorder) GetBucketCors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.GetBucketCors, reflect.TypeOf((*MockSpacesService)(nil).GetBucketCors), varargs...)
}

func (m *MockSpacesService) ListBuckets(arg0 context.Context, arg1 *s3.ListBucketsInput, arg2 ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, constants.ListBuckets, varargs...)
	ret0, _ := ret[0].(*s3.ListBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSpacesServiceMockRecorder) ListBuckets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListBuckets, reflect.TypeOf((*MockSpacesService)(nil).ListBuckets), varargs...)
}
