// Code generated by MockGen. DO NOT EDIT.
// Source: pgx_pool.go

// Package pgxpoolmock is a generated GoMock package.
package pgxpoolmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	pgconn "github.com/jackc/pgconn"
	v4 "github.com/jackc/pgx/v4"
)

// MockPgxPool is a mock of PgxPool interface.
type MockPgxPool struct {
	ctrl     *gomock.Controller
	recorder *MockPgxPoolMockRecorder
}

// MockPgxPoolMockRecorder is the mock recorder for MockPgxPool.
type MockPgxPoolMockRecorder struct {
	mock *MockPgxPool
}

// NewMockPgxPool creates a new mock instance.
func NewMockPgxPool(ctrl *gomock.Controller) *MockPgxPool {
	mock := &MockPgxPool{ctrl: ctrl}
	mock.recorder = &MockPgxPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPgxPool) EXPECT() *MockPgxPoolMockRecorder {
	return m.recorder
}

// Begin mocks base method.
func (m *MockPgxPool) Begin(ctx context.Context) (v4.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin", ctx)
	ret0, _ := ret[0].(v4.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Begin indicates an expected call of Begin.
func (mr *MockPgxPoolMockRecorder) Begin(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockPgxPool)(nil).Begin), ctx)
}

// BeginFunc mocks base method.
func (m *MockPgxPool) BeginFunc(ctx context.Context, f func(v4.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginFunc", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeginFunc indicates an expected call of BeginFunc.
func (mr *MockPgxPoolMockRecorder) BeginFunc(ctx, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginFunc", reflect.TypeOf((*MockPgxPool)(nil).BeginFunc), ctx, f)
}

// BeginTx mocks base method.
func (m *MockPgxPool) BeginTx(ctx context.Context, txOptions v4.TxOptions) (v4.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx", ctx, txOptions)
	ret0, _ := ret[0].(v4.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockPgxPoolMockRecorder) BeginTx(ctx, txOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockPgxPool)(nil).BeginTx), ctx, txOptions)
}

// BeginTxFunc mocks base method.
func (m *MockPgxPool) BeginTxFunc(ctx context.Context, txOptions v4.TxOptions, f func(v4.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTxFunc", ctx, txOptions, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeginTxFunc indicates an expected call of BeginTxFunc.
func (mr *MockPgxPoolMockRecorder) BeginTxFunc(ctx, txOptions, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTxFunc", reflect.TypeOf((*MockPgxPool)(nil).BeginTxFunc), ctx, txOptions, f)
}

// Close mocks base method.
func (m *MockPgxPool) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockPgxPoolMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPgxPool)(nil).Close))
}

// Exec mocks base method.
func (m *MockPgxPool) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range arguments {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockPgxPoolMockRecorder) Exec(ctx, sql interface{}, arguments ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, arguments...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockPgxPool)(nil).Exec), varargs...)
}

// Query mocks base method.
func (m *MockPgxPool) Query(ctx context.Context, sql string, args ...interface{}) (v4.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(v4.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockPgxPoolMockRecorder) Query(ctx, sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockPgxPool)(nil).Query), varargs...)
}

// QueryFunc mocks base method.
func (m *MockPgxPool) QueryFunc(ctx context.Context, sql string, args, scans []interface{}, f func(v4.QueryFuncRow) error) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFunc", ctx, sql, args, scans, f)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFunc indicates an expected call of QueryFunc.
func (mr *MockPgxPoolMockRecorder) QueryFunc(ctx, sql, args, scans, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFunc", reflect.TypeOf((*MockPgxPool)(nil).QueryFunc), ctx, sql, args, scans, f)
}

// QueryRow mocks base method.
func (m *MockPgxPool) QueryRow(ctx context.Context, sql string, args ...interface{}) v4.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(v4.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockPgxPoolMockRecorder) QueryRow(ctx, sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockPgxPool)(nil).QueryRow), varargs...)
}

// SendBatch mocks base method.
func (m *MockPgxPool) SendBatch(ctx context.Context, b *v4.Batch) v4.BatchResults {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendBatch", ctx, b)
	ret0, _ := ret[0].(v4.BatchResults)
	return ret0
}

// SendBatch indicates an expected call of SendBatch.
func (mr *MockPgxPoolMockRecorder) SendBatch(ctx, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendBatch", reflect.TypeOf((*MockPgxPool)(nil).SendBatch), ctx, b)
}
