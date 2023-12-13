// Code generated by mockery v2.37.1. DO NOT EDIT.

package mockpbconnectca

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	pbconnectca "github.com/hashicorp/consul/proto-public/pbconnectca"
)

// ConnectCAService_WatchRootsClient is an autogenerated mock type for the ConnectCAService_WatchRootsClient type
type ConnectCAService_WatchRootsClient struct {
	mock.Mock
}

type ConnectCAService_WatchRootsClient_Expecter struct {
	mock *mock.Mock
}

func (_m *ConnectCAService_WatchRootsClient) EXPECT() *ConnectCAService_WatchRootsClient_Expecter {
	return &ConnectCAService_WatchRootsClient_Expecter{mock: &_m.Mock}
}

// CloseSend provides a mock function with given fields:
func (_m *ConnectCAService_WatchRootsClient) CloseSend() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnectCAService_WatchRootsClient_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type ConnectCAService_WatchRootsClient_CloseSend_Call struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *ConnectCAService_WatchRootsClient_Expecter) CloseSend() *ConnectCAService_WatchRootsClient_CloseSend_Call {
	return &ConnectCAService_WatchRootsClient_CloseSend_Call{Call: _e.mock.On("CloseSend")}
}

func (_c *ConnectCAService_WatchRootsClient_CloseSend_Call) Run(run func()) *ConnectCAService_WatchRootsClient_CloseSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_CloseSend_Call) Return(_a0 error) *ConnectCAService_WatchRootsClient_CloseSend_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_CloseSend_Call) RunAndReturn(run func() error) *ConnectCAService_WatchRootsClient_CloseSend_Call {
	_c.Call.Return(run)
	return _c
}

// Context provides a mock function with given fields:
func (_m *ConnectCAService_WatchRootsClient) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// ConnectCAService_WatchRootsClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type ConnectCAService_WatchRootsClient_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *ConnectCAService_WatchRootsClient_Expecter) Context() *ConnectCAService_WatchRootsClient_Context_Call {
	return &ConnectCAService_WatchRootsClient_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *ConnectCAService_WatchRootsClient_Context_Call) Run(run func()) *ConnectCAService_WatchRootsClient_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_Context_Call) Return(_a0 context.Context) *ConnectCAService_WatchRootsClient_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_Context_Call) RunAndReturn(run func() context.Context) *ConnectCAService_WatchRootsClient_Context_Call {
	_c.Call.Return(run)
	return _c
}

// Header provides a mock function with given fields:
func (_m *ConnectCAService_WatchRootsClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	var r0 metadata.MD
	var r1 error
	if rf, ok := ret.Get(0).(func() (metadata.MD, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConnectCAService_WatchRootsClient_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type ConnectCAService_WatchRootsClient_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *ConnectCAService_WatchRootsClient_Expecter) Header() *ConnectCAService_WatchRootsClient_Header_Call {
	return &ConnectCAService_WatchRootsClient_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *ConnectCAService_WatchRootsClient_Header_Call) Run(run func()) *ConnectCAService_WatchRootsClient_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_Header_Call) Return(_a0 metadata.MD, _a1 error) *ConnectCAService_WatchRootsClient_Header_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_Header_Call) RunAndReturn(run func() (metadata.MD, error)) *ConnectCAService_WatchRootsClient_Header_Call {
	_c.Call.Return(run)
	return _c
}

// Recv provides a mock function with given fields:
func (_m *ConnectCAService_WatchRootsClient) Recv() (*pbconnectca.WatchRootsResponse, error) {
	ret := _m.Called()

	var r0 *pbconnectca.WatchRootsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func() (*pbconnectca.WatchRootsResponse, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *pbconnectca.WatchRootsResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pbconnectca.WatchRootsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConnectCAService_WatchRootsClient_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type ConnectCAService_WatchRootsClient_Recv_Call struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *ConnectCAService_WatchRootsClient_Expecter) Recv() *ConnectCAService_WatchRootsClient_Recv_Call {
	return &ConnectCAService_WatchRootsClient_Recv_Call{Call: _e.mock.On("Recv")}
}

func (_c *ConnectCAService_WatchRootsClient_Recv_Call) Run(run func()) *ConnectCAService_WatchRootsClient_Recv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_Recv_Call) Return(_a0 *pbconnectca.WatchRootsResponse, _a1 error) *ConnectCAService_WatchRootsClient_Recv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_Recv_Call) RunAndReturn(run func() (*pbconnectca.WatchRootsResponse, error)) *ConnectCAService_WatchRootsClient_Recv_Call {
	_c.Call.Return(run)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *ConnectCAService_WatchRootsClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnectCAService_WatchRootsClient_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type ConnectCAService_WatchRootsClient_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *ConnectCAService_WatchRootsClient_Expecter) RecvMsg(m interface{}) *ConnectCAService_WatchRootsClient_RecvMsg_Call {
	return &ConnectCAService_WatchRootsClient_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *ConnectCAService_WatchRootsClient_RecvMsg_Call) Run(run func(m interface{})) *ConnectCAService_WatchRootsClient_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_RecvMsg_Call) Return(_a0 error) *ConnectCAService_WatchRootsClient_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_RecvMsg_Call) RunAndReturn(run func(interface{}) error) *ConnectCAService_WatchRootsClient_RecvMsg_Call {
	_c.Call.Return(run)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *ConnectCAService_WatchRootsClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConnectCAService_WatchRootsClient_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type ConnectCAService_WatchRootsClient_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//   - m interface{}
func (_e *ConnectCAService_WatchRootsClient_Expecter) SendMsg(m interface{}) *ConnectCAService_WatchRootsClient_SendMsg_Call {
	return &ConnectCAService_WatchRootsClient_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *ConnectCAService_WatchRootsClient_SendMsg_Call) Run(run func(m interface{})) *ConnectCAService_WatchRootsClient_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_SendMsg_Call) Return(_a0 error) *ConnectCAService_WatchRootsClient_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_SendMsg_Call) RunAndReturn(run func(interface{}) error) *ConnectCAService_WatchRootsClient_SendMsg_Call {
	_c.Call.Return(run)
	return _c
}

// Trailer provides a mock function with given fields:
func (_m *ConnectCAService_WatchRootsClient) Trailer() metadata.MD {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// ConnectCAService_WatchRootsClient_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type ConnectCAService_WatchRootsClient_Trailer_Call struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *ConnectCAService_WatchRootsClient_Expecter) Trailer() *ConnectCAService_WatchRootsClient_Trailer_Call {
	return &ConnectCAService_WatchRootsClient_Trailer_Call{Call: _e.mock.On("Trailer")}
}

func (_c *ConnectCAService_WatchRootsClient_Trailer_Call) Run(run func()) *ConnectCAService_WatchRootsClient_Trailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_Trailer_Call) Return(_a0 metadata.MD) *ConnectCAService_WatchRootsClient_Trailer_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ConnectCAService_WatchRootsClient_Trailer_Call) RunAndReturn(run func() metadata.MD) *ConnectCAService_WatchRootsClient_Trailer_Call {
	_c.Call.Return(run)
	return _c
}

// NewConnectCAService_WatchRootsClient creates a new instance of ConnectCAService_WatchRootsClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConnectCAService_WatchRootsClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConnectCAService_WatchRootsClient {
	mock := &ConnectCAService_WatchRootsClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
