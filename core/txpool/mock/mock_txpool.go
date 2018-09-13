// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/iost-official/Go-IOS-Protocol/core/txpool (interfaces: TxPool)

// Package txpool_mock is a generated GoMock package.
package txpool_mock

import (
	gomock "github.com/golang/mock/gomock"
	block "github.com/iost-official/Go-IOS-Protocol/core/block"
	blockcache "github.com/iost-official/Go-IOS-Protocol/core/blockcache"
	tx "github.com/iost-official/Go-IOS-Protocol/core/tx"
	txpool "github.com/iost-official/Go-IOS-Protocol/core/txpool"
	reflect "reflect"
)

// MockTxPool is a mock of TxPool interface
type MockTxPool struct {
	ctrl     *gomock.Controller
	recorder *MockTxPoolMockRecorder
}

// MockTxPoolMockRecorder is the mock recorder for MockTxPool
type MockTxPoolMockRecorder struct {
	mock *MockTxPool
}

// NewMockTxPool creates a new mock instance
func NewMockTxPool(ctrl *gomock.Controller) *MockTxPool {
	mock := &MockTxPool{ctrl: ctrl}
	mock.recorder = &MockTxPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTxPool) EXPECT() *MockTxPoolMockRecorder {
	return m.recorder
}

// AddLinkedNode mocks base method
func (m *MockTxPool) AddLinkedNode(arg0, arg1 *blockcache.BlockCacheNode) error {
	ret := m.ctrl.Call(m, "AddLinkedNode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLinkedNode indicates an expected call of AddLinkedNode
func (mr *MockTxPoolMockRecorder) AddLinkedNode(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLinkedNode", reflect.TypeOf((*MockTxPool)(nil).AddLinkedNode), arg0, arg1)
}

// AddTx mocks base method
func (m *MockTxPool) AddTx(arg0 *tx.Tx) txpool.TAddTx {
	ret := m.ctrl.Call(m, "AddTx", arg0)
	ret0, _ := ret[0].(txpool.TAddTx)
	return ret0
}

// AddTx indicates an expected call of AddTx
func (mr *MockTxPoolMockRecorder) AddTx(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTx", reflect.TypeOf((*MockTxPool)(nil).AddTx), arg0)
}

// CheckTxs mocks base method
func (m *MockTxPool) CheckTxs(arg0 []*tx.Tx, arg1 *block.Block) (*tx.Tx, error) {
	ret := m.ctrl.Call(m, "CheckTxs", arg0, arg1)
	ret0, _ := ret[0].(*tx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckTxs indicates an expected call of CheckTxs
func (mr *MockTxPoolMockRecorder) CheckTxs(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTxs", reflect.TypeOf((*MockTxPool)(nil).CheckTxs), arg0, arg1)
}

// DelTx mocks base method
func (m *MockTxPool) DelTx(arg0 []byte) error {
	ret := m.ctrl.Call(m, "DelTx", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelTx indicates an expected call of DelTx
func (mr *MockTxPoolMockRecorder) DelTx(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelTx", reflect.TypeOf((*MockTxPool)(nil).DelTx), arg0)
}

// ExistTxs mocks base method
func (m *MockTxPool) ExistTxs(arg0 []byte, arg1 *block.Block) (txpool.FRet, error) {
	ret := m.ctrl.Call(m, "ExistTxs", arg0, arg1)
	ret0, _ := ret[0].(txpool.FRet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistTxs indicates an expected call of ExistTxs
func (mr *MockTxPoolMockRecorder) ExistTxs(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistTxs", reflect.TypeOf((*MockTxPool)(nil).ExistTxs), arg0, arg1)
}

// Lease mocks base method
func (m *MockTxPool) Lease() {
	m.ctrl.Call(m, "Lease")
}

// Lease indicates an expected call of Lease
func (mr *MockTxPoolMockRecorder) Lease() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lease", reflect.TypeOf((*MockTxPool)(nil).Lease))
}

// Lock mocks base method
func (m *MockTxPool) Lock() {
	m.ctrl.Call(m, "Lock")
}

// Lock indicates an expected call of Lock
func (mr *MockTxPoolMockRecorder) Lock() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockTxPool)(nil).Lock))
}

// PendingTxs mocks base method
func (m *MockTxPool) PendingTxs(arg0 int) (txpool.TxsList, *blockcache.BlockCacheNode, error) {
	ret := m.ctrl.Call(m, "PendingTxs", arg0)
	ret0, _ := ret[0].(txpool.TxsList)
	ret1, _ := ret[1].(*blockcache.BlockCacheNode)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PendingTxs indicates an expected call of PendingTxs
func (mr *MockTxPoolMockRecorder) PendingTxs(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PendingTxs", reflect.TypeOf((*MockTxPool)(nil).PendingTxs), arg0)
}

// Start mocks base method
func (m *MockTxPool) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockTxPoolMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockTxPool)(nil).Start))
}

// Stop mocks base method
func (m *MockTxPool) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockTxPoolMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTxPool)(nil).Stop))
}
