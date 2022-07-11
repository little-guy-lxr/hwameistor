// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package storage is a generated GoMock package.
package storage

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/local-storage/v1alpha1"
)

// MockLocalPoolManager is a mock of LocalPoolManager interface.
type MockLocalPoolManager struct {
	ctrl     *gomock.Controller
	recorder *MockLocalPoolManagerMockRecorder
}

// MockLocalPoolManagerMockRecorder is the mock recorder for MockLocalPoolManager.
type MockLocalPoolManagerMockRecorder struct {
	mock *MockLocalPoolManager
}

// NewMockLocalPoolManager creates a new mock instance.
func NewMockLocalPoolManager(ctrl *gomock.Controller) *MockLocalPoolManager {
	mock := &MockLocalPoolManager{ctrl: ctrl}
	mock.recorder = &MockLocalPoolManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalPoolManager) EXPECT() *MockLocalPoolManagerMockRecorder {
	return m.recorder
}

// ExtendPools mocks base method.
func (m *MockLocalPoolManager) ExtendPools(localDisks []*v1alpha1.LocalDisk) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendPools", localDisks)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExtendPools indicates an expected call of ExtendPools.
func (mr *MockLocalPoolManagerMockRecorder) ExtendPools(localDisks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendPools", reflect.TypeOf((*MockLocalPoolManager)(nil).ExtendPools), localDisks)
}

// ExtendPoolsInfo mocks base method.
func (m *MockLocalPoolManager) ExtendPoolsInfo(localDisks map[string]*v1alpha1.LocalDisk) (map[string]*v1alpha1.LocalPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendPoolsInfo", localDisks)
	ret0, _ := ret[0].(map[string]*v1alpha1.LocalPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtendPoolsInfo indicates an expected call of ExtendPoolsInfo.
func (mr *MockLocalPoolManagerMockRecorder) ExtendPoolsInfo(localDisks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendPoolsInfo", reflect.TypeOf((*MockLocalPoolManager)(nil).ExtendPoolsInfo), localDisks)
}

// GetReplicas mocks base method.
func (m *MockLocalPoolManager) GetReplicas() (map[string]*v1alpha1.LocalVolumeReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicas")
	ret0, _ := ret[0].(map[string]*v1alpha1.LocalVolumeReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplicas indicates an expected call of GetReplicas.
func (mr *MockLocalPoolManagerMockRecorder) GetReplicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicas", reflect.TypeOf((*MockLocalPoolManager)(nil).GetReplicas))
}

// MockLocalVolumeReplicaManager is a mock of LocalVolumeReplicaManager interface.
type MockLocalVolumeReplicaManager struct {
	ctrl     *gomock.Controller
	recorder *MockLocalVolumeReplicaManagerMockRecorder
}

// MockLocalVolumeReplicaManagerMockRecorder is the mock recorder for MockLocalVolumeReplicaManager.
type MockLocalVolumeReplicaManagerMockRecorder struct {
	mock *MockLocalVolumeReplicaManager
}

// NewMockLocalVolumeReplicaManager creates a new mock instance.
func NewMockLocalVolumeReplicaManager(ctrl *gomock.Controller) *MockLocalVolumeReplicaManager {
	mock := &MockLocalVolumeReplicaManager{ctrl: ctrl}
	mock.recorder = &MockLocalVolumeReplicaManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalVolumeReplicaManager) EXPECT() *MockLocalVolumeReplicaManagerMockRecorder {
	return m.recorder
}

// ConsistencyCheck mocks base method.
func (m *MockLocalVolumeReplicaManager) ConsistencyCheck() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ConsistencyCheck")
}

// ConsistencyCheck indicates an expected call of ConsistencyCheck.
func (mr *MockLocalVolumeReplicaManagerMockRecorder) ConsistencyCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsistencyCheck", reflect.TypeOf((*MockLocalVolumeReplicaManager)(nil).ConsistencyCheck))
}

// CreateVolumeReplica mocks base method.
func (m *MockLocalVolumeReplicaManager) CreateVolumeReplica(replica *v1alpha1.LocalVolumeReplica) (*v1alpha1.LocalVolumeReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolumeReplica", replica)
	ret0, _ := ret[0].(*v1alpha1.LocalVolumeReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolumeReplica indicates an expected call of CreateVolumeReplica.
func (mr *MockLocalVolumeReplicaManagerMockRecorder) CreateVolumeReplica(replica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolumeReplica", reflect.TypeOf((*MockLocalVolumeReplicaManager)(nil).CreateVolumeReplica), replica)
}

// DeleteVolumeReplica mocks base method.
func (m *MockLocalVolumeReplicaManager) DeleteVolumeReplica(replica *v1alpha1.LocalVolumeReplica) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolumeReplica", replica)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolumeReplica indicates an expected call of DeleteVolumeReplica.
func (mr *MockLocalVolumeReplicaManagerMockRecorder) DeleteVolumeReplica(replica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolumeReplica", reflect.TypeOf((*MockLocalVolumeReplicaManager)(nil).DeleteVolumeReplica), replica)
}

// ExpandVolumeReplica mocks base method.
func (m *MockLocalVolumeReplicaManager) ExpandVolumeReplica(replica *v1alpha1.LocalVolumeReplica, newCapacityBytes int64) (*v1alpha1.LocalVolumeReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpandVolumeReplica", replica, newCapacityBytes)
	ret0, _ := ret[0].(*v1alpha1.LocalVolumeReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExpandVolumeReplica indicates an expected call of ExpandVolumeReplica.
func (mr *MockLocalVolumeReplicaManagerMockRecorder) ExpandVolumeReplica(replica, newCapacityBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpandVolumeReplica", reflect.TypeOf((*MockLocalVolumeReplicaManager)(nil).ExpandVolumeReplica), replica, newCapacityBytes)
}

// GetVolumeReplica mocks base method.
func (m *MockLocalVolumeReplicaManager) GetVolumeReplica(replica *v1alpha1.LocalVolumeReplica) (*v1alpha1.LocalVolumeReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeReplica", replica)
	ret0, _ := ret[0].(*v1alpha1.LocalVolumeReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeReplica indicates an expected call of GetVolumeReplica.
func (mr *MockLocalVolumeReplicaManagerMockRecorder) GetVolumeReplica(replica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeReplica", reflect.TypeOf((*MockLocalVolumeReplicaManager)(nil).GetVolumeReplica), replica)
}

// TestVolumeReplica mocks base method.
func (m *MockLocalVolumeReplicaManager) TestVolumeReplica(replica *v1alpha1.LocalVolumeReplica) (*v1alpha1.LocalVolumeReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestVolumeReplica", replica)
	ret0, _ := ret[0].(*v1alpha1.LocalVolumeReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TestVolumeReplica indicates an expected call of TestVolumeReplica.
func (mr *MockLocalVolumeReplicaManagerMockRecorder) TestVolumeReplica(replica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestVolumeReplica", reflect.TypeOf((*MockLocalVolumeReplicaManager)(nil).TestVolumeReplica), replica)
}

// MockLocalRegistry is a mock of LocalRegistry interface.
type MockLocalRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockLocalRegistryMockRecorder
}

// MockLocalRegistryMockRecorder is the mock recorder for MockLocalRegistry.
type MockLocalRegistryMockRecorder struct {
	mock *MockLocalRegistry
}

// NewMockLocalRegistry creates a new mock instance.
func NewMockLocalRegistry(ctrl *gomock.Controller) *MockLocalRegistry {
	mock := &MockLocalRegistry{ctrl: ctrl}
	mock.recorder = &MockLocalRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalRegistry) EXPECT() *MockLocalRegistryMockRecorder {
	return m.recorder
}

// Disks mocks base method.
func (m *MockLocalRegistry) Disks() map[string]*v1alpha1.LocalDisk {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disks")
	ret0, _ := ret[0].(map[string]*v1alpha1.LocalDisk)
	return ret0
}

// Disks indicates an expected call of Disks.
func (mr *MockLocalRegistryMockRecorder) Disks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disks", reflect.TypeOf((*MockLocalRegistry)(nil).Disks))
}

// HasVolumeReplica mocks base method.
func (m *MockLocalRegistry) HasVolumeReplica(replica *v1alpha1.LocalVolumeReplica) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasVolumeReplica", replica)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasVolumeReplica indicates an expected call of HasVolumeReplica.
func (mr *MockLocalRegistryMockRecorder) HasVolumeReplica(replica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasVolumeReplica", reflect.TypeOf((*MockLocalRegistry)(nil).HasVolumeReplica), replica)
}

// Init mocks base method.
func (m *MockLocalRegistry) Init() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init")
}

// Init indicates an expected call of Init.
func (mr *MockLocalRegistryMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockLocalRegistry)(nil).Init))
}

// Pools mocks base method.
func (m *MockLocalRegistry) Pools() map[string]*v1alpha1.LocalPool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pools")
	ret0, _ := ret[0].(map[string]*v1alpha1.LocalPool)
	return ret0
}

// Pools indicates an expected call of Pools.
func (mr *MockLocalRegistryMockRecorder) Pools() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pools", reflect.TypeOf((*MockLocalRegistry)(nil).Pools))
}

// SyncResourcesToNodeCRD mocks base method.
func (m *MockLocalRegistry) SyncResourcesToNodeCRD(localDisks map[string]*v1alpha1.LocalDisk) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncResourcesToNodeCRD", localDisks)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncResourcesToNodeCRD indicates an expected call of SyncResourcesToNodeCRD.
func (mr *MockLocalRegistryMockRecorder) SyncResourcesToNodeCRD(localDisks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncResourcesToNodeCRD", reflect.TypeOf((*MockLocalRegistry)(nil).SyncResourcesToNodeCRD), localDisks)
}

// UpdateNodeForVolumeReplica mocks base method.
func (m *MockLocalRegistry) UpdateNodeForVolumeReplica(replica *v1alpha1.LocalVolumeReplica) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateNodeForVolumeReplica", replica)
}

// UpdateNodeForVolumeReplica indicates an expected call of UpdateNodeForVolumeReplica.
func (mr *MockLocalRegistryMockRecorder) UpdateNodeForVolumeReplica(replica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNodeForVolumeReplica", reflect.TypeOf((*MockLocalRegistry)(nil).UpdateNodeForVolumeReplica), replica)
}

// VolumeReplicas mocks base method.
func (m *MockLocalRegistry) VolumeReplicas() map[string]*v1alpha1.LocalVolumeReplica {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeReplicas")
	ret0, _ := ret[0].(map[string]*v1alpha1.LocalVolumeReplica)
	return ret0
}

// VolumeReplicas indicates an expected call of VolumeReplicas.
func (mr *MockLocalRegistryMockRecorder) VolumeReplicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeReplicas", reflect.TypeOf((*MockLocalRegistry)(nil).VolumeReplicas))
}

// MockLocalVolumeExecutor is a mock of LocalVolumeExecutor interface.
type MockLocalVolumeExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockLocalVolumeExecutorMockRecorder
}

// MockLocalVolumeExecutorMockRecorder is the mock recorder for MockLocalVolumeExecutor.
type MockLocalVolumeExecutorMockRecorder struct {
	mock *MockLocalVolumeExecutor
}

// NewMockLocalVolumeExecutor creates a new mock instance.
func NewMockLocalVolumeExecutor(ctrl *gomock.Controller) *MockLocalVolumeExecutor {
	mock := &MockLocalVolumeExecutor{ctrl: ctrl}
	mock.recorder = &MockLocalVolumeExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalVolumeExecutor) EXPECT() *MockLocalVolumeExecutorMockRecorder {
	return m.recorder
}

// ConsistencyCheck mocks base method.
func (m *MockLocalVolumeExecutor) ConsistencyCheck(crdReplicas map[string]*v1alpha1.LocalVolumeReplica) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ConsistencyCheck", crdReplicas)
}

// ConsistencyCheck indicates an expected call of ConsistencyCheck.
func (mr *MockLocalVolumeExecutorMockRecorder) ConsistencyCheck(crdReplicas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsistencyCheck", reflect.TypeOf((*MockLocalVolumeExecutor)(nil).ConsistencyCheck), crdReplicas)
}

// CreateVolumeReplica mocks base method.
func (m *MockLocalVolumeExecutor) CreateVolumeReplica(replica *v1alpha1.LocalVolumeReplica) (*v1alpha1.LocalVolumeReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolumeReplica", replica)
	ret0, _ := ret[0].(*v1alpha1.LocalVolumeReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVolumeReplica indicates an expected call of CreateVolumeReplica.
func (mr *MockLocalVolumeExecutorMockRecorder) CreateVolumeReplica(replica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolumeReplica", reflect.TypeOf((*MockLocalVolumeExecutor)(nil).CreateVolumeReplica), replica)
}

// DeleteVolumeReplica mocks base method.
func (m *MockLocalVolumeExecutor) DeleteVolumeReplica(replica *v1alpha1.LocalVolumeReplica) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolumeReplica", replica)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolumeReplica indicates an expected call of DeleteVolumeReplica.
func (mr *MockLocalVolumeExecutorMockRecorder) DeleteVolumeReplica(replica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolumeReplica", reflect.TypeOf((*MockLocalVolumeExecutor)(nil).DeleteVolumeReplica), replica)
}

// ExpandVolumeReplica mocks base method.
func (m *MockLocalVolumeExecutor) ExpandVolumeReplica(replica *v1alpha1.LocalVolumeReplica, newCapacityBytes int64) (*v1alpha1.LocalVolumeReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExpandVolumeReplica", replica, newCapacityBytes)
	ret0, _ := ret[0].(*v1alpha1.LocalVolumeReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExpandVolumeReplica indicates an expected call of ExpandVolumeReplica.
func (mr *MockLocalVolumeExecutorMockRecorder) ExpandVolumeReplica(replica, newCapacityBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpandVolumeReplica", reflect.TypeOf((*MockLocalVolumeExecutor)(nil).ExpandVolumeReplica), replica, newCapacityBytes)
}

// GetReplicas mocks base method.
func (m *MockLocalVolumeExecutor) GetReplicas() (map[string]*v1alpha1.LocalVolumeReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicas")
	ret0, _ := ret[0].(map[string]*v1alpha1.LocalVolumeReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplicas indicates an expected call of GetReplicas.
func (mr *MockLocalVolumeExecutorMockRecorder) GetReplicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicas", reflect.TypeOf((*MockLocalVolumeExecutor)(nil).GetReplicas))
}

// TestVolumeReplica mocks base method.
func (m *MockLocalVolumeExecutor) TestVolumeReplica(replica *v1alpha1.LocalVolumeReplica) (*v1alpha1.LocalVolumeReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestVolumeReplica", replica)
	ret0, _ := ret[0].(*v1alpha1.LocalVolumeReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TestVolumeReplica indicates an expected call of TestVolumeReplica.
func (mr *MockLocalVolumeExecutorMockRecorder) TestVolumeReplica(replica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestVolumeReplica", reflect.TypeOf((*MockLocalVolumeExecutor)(nil).TestVolumeReplica), replica)
}

// MockLocalPoolExecutor is a mock of LocalPoolExecutor interface.
type MockLocalPoolExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockLocalPoolExecutorMockRecorder
}

// MockLocalPoolExecutorMockRecorder is the mock recorder for MockLocalPoolExecutor.
type MockLocalPoolExecutorMockRecorder struct {
	mock *MockLocalPoolExecutor
}

// NewMockLocalPoolExecutor creates a new mock instance.
func NewMockLocalPoolExecutor(ctrl *gomock.Controller) *MockLocalPoolExecutor {
	mock := &MockLocalPoolExecutor{ctrl: ctrl}
	mock.recorder = &MockLocalPoolExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalPoolExecutor) EXPECT() *MockLocalPoolExecutorMockRecorder {
	return m.recorder
}

// ExtendPools mocks base method.
func (m *MockLocalPoolExecutor) ExtendPools(localDisks []*v1alpha1.LocalDisk) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendPools", localDisks)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExtendPools indicates an expected call of ExtendPools.
func (mr *MockLocalPoolExecutorMockRecorder) ExtendPools(localDisks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendPools", reflect.TypeOf((*MockLocalPoolExecutor)(nil).ExtendPools), localDisks)
}

// ExtendPoolsInfo mocks base method.
func (m *MockLocalPoolExecutor) ExtendPoolsInfo(localDisks map[string]*v1alpha1.LocalDisk) (map[string]*v1alpha1.LocalPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendPoolsInfo", localDisks)
	ret0, _ := ret[0].(map[string]*v1alpha1.LocalPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExtendPoolsInfo indicates an expected call of ExtendPoolsInfo.
func (mr *MockLocalPoolExecutorMockRecorder) ExtendPoolsInfo(localDisks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendPoolsInfo", reflect.TypeOf((*MockLocalPoolExecutor)(nil).ExtendPoolsInfo), localDisks)
}

// GetReplicas mocks base method.
func (m *MockLocalPoolExecutor) GetReplicas() (map[string]*v1alpha1.LocalVolumeReplica, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicas")
	ret0, _ := ret[0].(map[string]*v1alpha1.LocalVolumeReplica)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplicas indicates an expected call of GetReplicas.
func (mr *MockLocalPoolExecutorMockRecorder) GetReplicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicas", reflect.TypeOf((*MockLocalPoolExecutor)(nil).GetReplicas))
}
