// Code generated by MockGen. DO NOT EDIT.
// Source: nmc_reconciler.go
//
// Generated by this command:
//
//	mockgen -source=nmc_reconciler.go -package=controllers -destination=mock_nmc_reconciler.go pullSecretHelper
//
// Package controllers is a generated GoMock package.
package controllers

import (
	context "context"
	reflect "reflect"

	v1beta1 "github.com/kubernetes-sigs/kernel-module-management/api/v1beta1"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/core/v1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockworkerHelper is a mock of workerHelper interface.
type MockworkerHelper struct {
	ctrl     *gomock.Controller
	recorder *MockworkerHelperMockRecorder
}

// MockworkerHelperMockRecorder is the mock recorder for MockworkerHelper.
type MockworkerHelperMockRecorder struct {
	mock *MockworkerHelper
}

// NewMockworkerHelper creates a new mock instance.
func NewMockworkerHelper(ctrl *gomock.Controller) *MockworkerHelper {
	mock := &MockworkerHelper{ctrl: ctrl}
	mock.recorder = &MockworkerHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockworkerHelper) EXPECT() *MockworkerHelperMockRecorder {
	return m.recorder
}

// ProcessModuleSpec mocks base method.
func (m *MockworkerHelper) ProcessModuleSpec(ctx context.Context, nmc *v1beta1.NodeModulesConfig, spec *v1beta1.NodeModuleSpec, status *v1beta1.NodeModuleStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessModuleSpec", ctx, nmc, spec, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessModuleSpec indicates an expected call of ProcessModuleSpec.
func (mr *MockworkerHelperMockRecorder) ProcessModuleSpec(ctx, nmc, spec, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessModuleSpec", reflect.TypeOf((*MockworkerHelper)(nil).ProcessModuleSpec), ctx, nmc, spec, status)
}

// ProcessOrphanModuleStatus mocks base method.
func (m *MockworkerHelper) ProcessOrphanModuleStatus(ctx context.Context, nmc *v1beta1.NodeModulesConfig, status *v1beta1.NodeModuleStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessOrphanModuleStatus", ctx, nmc, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessOrphanModuleStatus indicates an expected call of ProcessOrphanModuleStatus.
func (mr *MockworkerHelperMockRecorder) ProcessOrphanModuleStatus(ctx, nmc, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessOrphanModuleStatus", reflect.TypeOf((*MockworkerHelper)(nil).ProcessOrphanModuleStatus), ctx, nmc, status)
}

// RemoveOrphanFinalizers mocks base method.
func (m *MockworkerHelper) RemoveOrphanFinalizers(ctx context.Context, nodeName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveOrphanFinalizers", ctx, nodeName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveOrphanFinalizers indicates an expected call of RemoveOrphanFinalizers.
func (mr *MockworkerHelperMockRecorder) RemoveOrphanFinalizers(ctx, nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveOrphanFinalizers", reflect.TypeOf((*MockworkerHelper)(nil).RemoveOrphanFinalizers), ctx, nodeName)
}

// SyncStatus mocks base method.
func (m *MockworkerHelper) SyncStatus(ctx context.Context, nmc *v1beta1.NodeModulesConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatus", ctx, nmc)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStatus indicates an expected call of SyncStatus.
func (mr *MockworkerHelperMockRecorder) SyncStatus(ctx, nmc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatus", reflect.TypeOf((*MockworkerHelper)(nil).SyncStatus), ctx, nmc)
}

// MockpodManager is a mock of podManager interface.
type MockpodManager struct {
	ctrl     *gomock.Controller
	recorder *MockpodManagerMockRecorder
}

// MockpodManagerMockRecorder is the mock recorder for MockpodManager.
type MockpodManagerMockRecorder struct {
	mock *MockpodManager
}

// NewMockpodManager creates a new mock instance.
func NewMockpodManager(ctrl *gomock.Controller) *MockpodManager {
	mock := &MockpodManager{ctrl: ctrl}
	mock.recorder = &MockpodManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpodManager) EXPECT() *MockpodManagerMockRecorder {
	return m.recorder
}

// CreateLoaderPod mocks base method.
func (m *MockpodManager) CreateLoaderPod(ctx context.Context, nmc client.Object, nms *v1beta1.NodeModuleSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoaderPod", ctx, nmc, nms)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLoaderPod indicates an expected call of CreateLoaderPod.
func (mr *MockpodManagerMockRecorder) CreateLoaderPod(ctx, nmc, nms any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoaderPod", reflect.TypeOf((*MockpodManager)(nil).CreateLoaderPod), ctx, nmc, nms)
}

// CreateUnloaderPod mocks base method.
func (m *MockpodManager) CreateUnloaderPod(ctx context.Context, nmc client.Object, nms *v1beta1.NodeModuleStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUnloaderPod", ctx, nmc, nms)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUnloaderPod indicates an expected call of CreateUnloaderPod.
func (mr *MockpodManagerMockRecorder) CreateUnloaderPod(ctx, nmc, nms any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUnloaderPod", reflect.TypeOf((*MockpodManager)(nil).CreateUnloaderPod), ctx, nmc, nms)
}

// DeletePod mocks base method.
func (m *MockpodManager) DeletePod(ctx context.Context, pod *v1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePod", ctx, pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePod indicates an expected call of DeletePod.
func (mr *MockpodManagerMockRecorder) DeletePod(ctx, pod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePod", reflect.TypeOf((*MockpodManager)(nil).DeletePod), ctx, pod)
}

// ListWorkerPodsOnNode mocks base method.
func (m *MockpodManager) ListWorkerPodsOnNode(ctx context.Context, nodeName string) ([]v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkerPodsOnNode", ctx, nodeName)
	ret0, _ := ret[0].([]v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkerPodsOnNode indicates an expected call of ListWorkerPodsOnNode.
func (mr *MockpodManagerMockRecorder) ListWorkerPodsOnNode(ctx, nodeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkerPodsOnNode", reflect.TypeOf((*MockpodManager)(nil).ListWorkerPodsOnNode), ctx, nodeName)
}

// MockpullSecretHelper is a mock of pullSecretHelper interface.
type MockpullSecretHelper struct {
	ctrl     *gomock.Controller
	recorder *MockpullSecretHelperMockRecorder
}

// MockpullSecretHelperMockRecorder is the mock recorder for MockpullSecretHelper.
type MockpullSecretHelperMockRecorder struct {
	mock *MockpullSecretHelper
}

// NewMockpullSecretHelper creates a new mock instance.
func NewMockpullSecretHelper(ctrl *gomock.Controller) *MockpullSecretHelper {
	mock := &MockpullSecretHelper{ctrl: ctrl}
	mock.recorder = &MockpullSecretHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockpullSecretHelper) EXPECT() *MockpullSecretHelperMockRecorder {
	return m.recorder
}

// VolumesAndVolumeMounts mocks base method.
func (m *MockpullSecretHelper) VolumesAndVolumeMounts(ctx context.Context, nms *v1beta1.ModuleItem) ([]v1.Volume, []v1.VolumeMount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumesAndVolumeMounts", ctx, nms)
	ret0, _ := ret[0].([]v1.Volume)
	ret1, _ := ret[1].([]v1.VolumeMount)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VolumesAndVolumeMounts indicates an expected call of VolumesAndVolumeMounts.
func (mr *MockpullSecretHelperMockRecorder) VolumesAndVolumeMounts(ctx, nms any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumesAndVolumeMounts", reflect.TypeOf((*MockpullSecretHelper)(nil).VolumesAndVolumeMounts), ctx, nms)
}
