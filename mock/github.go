// Code generated by MockGen. DO NOT EDIT.
// Source: github.go
//
// Generated by this command:
//
//	mockgen -source=github.go -destination=../../mock/github.go
//
// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	entity "github.com/cocoide/commitify/internal/entity"
	service "github.com/cocoide/commitify/internal/service"
	"github.com/golang/mock/gomock"
)

// MockGithubService is a mock of GithubService interface.
type MockGithubService struct {
	ctrl     *gomock.Controller
	recorder *MockGithubServiceMockRecorder
}

// MockGithubServiceMockRecorder is the mock recorder for MockGithubService.
type MockGithubServiceMockRecorder struct {
	mock *MockGithubService
}

// NewMockGithubService creates a new mock instance.
func NewMockGithubService(ctrl *gomock.Controller) *MockGithubService {
	mock := &MockGithubService{ctrl: ctrl}
	mock.recorder = &MockGithubServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGithubService) EXPECT() *MockGithubServiceMockRecorder {
	return m.recorder
}

// CreatePullRequest mocks base method.
func (m *MockGithubService) CreatePullRequest(pr *entity.PullRequest, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePullRequest", pr, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePullRequest indicates an expected call of CreatePullRequest.
func (mr *MockGithubServiceMockRecorder) CreatePullRequest(pr, token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePullRequest", reflect.TypeOf((*MockGithubService)(nil).CreatePullRequest), pr, token)
}

// GetCurrentBranch mocks base method.
func (m *MockGithubService) GetCurrentBranch() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentBranch")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentBranch indicates an expected call of GetCurrentBranch.
func (mr *MockGithubServiceMockRecorder) GetCurrentBranch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentBranch", reflect.TypeOf((*MockGithubService)(nil).GetCurrentBranch))
}

// GetCurrentRepoDetails mocks base method.
func (m *MockGithubService) GetCurrentRepoDetails() (*service.GetRepoDetailsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentRepoDetails")
	ret0, _ := ret[0].(*service.GetRepoDetailsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentRepoDetails indicates an expected call of GetCurrentRepoDetails.
func (mr *MockGithubServiceMockRecorder) GetCurrentRepoDetails() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentRepoDetails", reflect.TypeOf((*MockGithubService)(nil).GetCurrentRepoDetails))
}

// GetRecentUpdatedBranch mocks base method.
func (m *MockGithubService) GetRecentUpdatedBranch() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecentUpdatedBranch")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecentUpdatedBranch indicates an expected call of GetRecentUpdatedBranch.
func (mr *MockGithubServiceMockRecorder) GetRecentUpdatedBranch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecentUpdatedBranch", reflect.TypeOf((*MockGithubService)(nil).GetRecentUpdatedBranch))
}

// GetStagingCodeDiff mocks base method.
func (m *MockGithubService) GetStagingCodeDiff() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStagingCodeDiff")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStagingCodeDiff indicates an expected call of GetStagingCodeDiff.
func (mr *MockGithubServiceMockRecorder) GetStagingCodeDiff() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStagingCodeDiff", reflect.TypeOf((*MockGithubService)(nil).GetStagingCodeDiff))
}

// GetUnPushedCommits mocks base method.
func (m *MockGithubService) GetUnPushedCommits(base string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnPushedCommits", base)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnPushedCommits indicates an expected call of GetUnPushedCommits.
func (mr *MockGithubServiceMockRecorder) GetUnPushedCommits(base any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnPushedCommits", reflect.TypeOf((*MockGithubService)(nil).GetUnPushedCommits), base)
}

// PushCurrentBranch mocks base method.
func (m *MockGithubService) PushCurrentBranch() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushCurrentBranch")
	ret0, _ := ret[0].(error)
	return ret0
}

// PushCurrentBranch indicates an expected call of PushCurrentBranch.
func (mr *MockGithubServiceMockRecorder) PushCurrentBranch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushCurrentBranch", reflect.TypeOf((*MockGithubService)(nil).PushCurrentBranch))
}