// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pefish/go-core-type/api-session (interfaces: IApiSession)

// Package mock_api_session is a generated GoMock package.
package mock_api_session

import (
	io "io"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/pefish/go-core-type/api"
	api_session "github.com/pefish/go-core-type/api-session"
	go_logger "github.com/pefish/go-logger"
)

// MockIApiSession is a mock of IApiSession interface.
type MockIApiSession struct {
	ctrl     *gomock.Controller
	recorder *MockIApiSessionMockRecorder
}

func (m *MockIApiSession) SetPathVars(vars map[string]string) {
	//TODO implement me
	panic("implement me")
}

func (m *MockIApiSession) PathVars() map[string]string {
	//TODO implement me
	panic("implement me")
}

func (m *MockIApiSession) Redirect(url string) {
	//TODO implement me
	panic("implement me")
}

func (m *MockIApiSession) Host() string {
	//TODO implement me
	panic("implement me")
}

// MockIApiSessionMockRecorder is the mock recorder for MockIApiSession.
type MockIApiSessionMockRecorder struct {
	mock *MockIApiSession
}

// NewMockIApiSession creates a new mock instance.
func NewMockIApiSession(ctrl *gomock.Controller) *MockIApiSession {
	mock := &MockIApiSession{ctrl: ctrl}
	mock.recorder = &MockIApiSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIApiSession) EXPECT() *MockIApiSessionMockRecorder {
	return m.recorder
}

// AddDefer mocks base method.
func (m *MockIApiSession) AddDefer(arg0 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddDefer", arg0)
}

// AddDefer indicates an expected call of AddDefer.
func (mr *MockIApiSessionMockRecorder) AddDefer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDefer", reflect.TypeOf((*MockIApiSession)(nil).AddDefer), arg0)
}

// Api mocks base method.
func (m *MockIApiSession) Api() api.IApi {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Api")
	ret0, _ := ret[0].(api.IApi)
	return ret0
}

// Api indicates an expected call of Api.
func (mr *MockIApiSessionMockRecorder) Api() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Api", reflect.TypeOf((*MockIApiSession)(nil).Api))
}

// Body mocks base method.
func (m *MockIApiSession) Body() io.ReadCloser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].(io.ReadCloser)
	return ret0
}

// Body indicates an expected call of Body.
func (mr *MockIApiSessionMockRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockIApiSession)(nil).Body))
}

// ClientType mocks base method.
func (m *MockIApiSession) ClientType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientType indicates an expected call of ClientType.
func (mr *MockIApiSessionMockRecorder) ClientType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientType", reflect.TypeOf((*MockIApiSession)(nil).ClientType))
}

// Data mocks base method.
func (m *MockIApiSession) Data(arg0 string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data", arg0)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Data indicates an expected call of Data.
func (mr *MockIApiSessionMockRecorder) Data(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockIApiSession)(nil).Data), arg0)
}

// Defers mocks base method.
func (m *MockIApiSession) Defers() []func() {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Defers")
	ret0, _ := ret[0].([]func())
	return ret0
}

// Defers indicates an expected call of Defers.
func (mr *MockIApiSessionMockRecorder) Defers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Defers", reflect.TypeOf((*MockIApiSession)(nil).Defers))
}

// FormValues mocks base method.
func (m *MockIApiSession) FormValues() (map[string][]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormValues")
	ret0, _ := ret[0].(map[string][]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormValues indicates an expected call of FormValues.
func (mr *MockIApiSessionMockRecorder) FormValues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormValues", reflect.TypeOf((*MockIApiSession)(nil).FormValues))
}

// Header mocks base method.
func (m *MockIApiSession) Header(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Header indicates an expected call of Header.
func (mr *MockIApiSessionMockRecorder) Header(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockIApiSession)(nil).Header), arg0)
}

// JwtBody mocks base method.
func (m *MockIApiSession) JwtBody() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JwtBody")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// JwtBody indicates an expected call of JwtBody.
func (mr *MockIApiSessionMockRecorder) JwtBody() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JwtBody", reflect.TypeOf((*MockIApiSession)(nil).JwtBody))
}

// JwtHeaderName mocks base method.
func (m *MockIApiSession) JwtHeaderName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JwtHeaderName")
	ret0, _ := ret[0].(string)
	return ret0
}

// JwtHeaderName indicates an expected call of JwtHeaderName.
func (mr *MockIApiSessionMockRecorder) JwtHeaderName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JwtHeaderName", reflect.TypeOf((*MockIApiSession)(nil).JwtHeaderName))
}

// Lang mocks base method.
func (m *MockIApiSession) Lang() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lang")
	ret0, _ := ret[0].(string)
	return ret0
}

// Lang indicates an expected call of Lang.
func (mr *MockIApiSessionMockRecorder) Lang() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lang", reflect.TypeOf((*MockIApiSession)(nil).Lang))
}

// Logger mocks base method.
func (m *MockIApiSession) Logger() go_logger.InterfaceLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(go_logger.InterfaceLogger)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockIApiSessionMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockIApiSession)(nil).Logger))
}

// Method mocks base method.
func (m *MockIApiSession) Method() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Method")
	ret0, _ := ret[0].(string)
	return ret0
}

// Method indicates an expected call of Method.
func (mr *MockIApiSessionMockRecorder) Method() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Method", reflect.TypeOf((*MockIApiSession)(nil).Method))
}

// OriginalParams mocks base method.
func (m *MockIApiSession) OriginalParams() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OriginalParams")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// OriginalParams indicates an expected call of OriginalParams.
func (mr *MockIApiSessionMockRecorder) OriginalParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OriginalParams", reflect.TypeOf((*MockIApiSession)(nil).OriginalParams))
}

// Params mocks base method.
func (m *MockIApiSession) Params() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Params")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Params indicates an expected call of Params.
func (mr *MockIApiSessionMockRecorder) Params() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Params", reflect.TypeOf((*MockIApiSession)(nil).Params))
}

// Path mocks base method.
func (m *MockIApiSession) Path() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(string)
	return ret0
}

// Path indicates an expected call of Path.
func (mr *MockIApiSessionMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockIApiSession)(nil).Path))
}

// ReadJSON mocks base method.
func (m *MockIApiSession) ReadJSON(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadJSON", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadJSON indicates an expected call of ReadJSON.
func (mr *MockIApiSessionMockRecorder) ReadJSON(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadJSON", reflect.TypeOf((*MockIApiSession)(nil).ReadJSON), arg0)
}

// RemoteAddress mocks base method.
func (m *MockIApiSession) RemoteAddress() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddress")
	ret0, _ := ret[0].(string)
	return ret0
}

// RemoteAddress indicates an expected call of RemoteAddress.
func (mr *MockIApiSessionMockRecorder) RemoteAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddress", reflect.TypeOf((*MockIApiSession)(nil).RemoteAddress))
}

// Request mocks base method.
func (m *MockIApiSession) Request() *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// Request indicates an expected call of Request.
func (mr *MockIApiSessionMockRecorder) Request() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockIApiSession)(nil).Request))
}

// ResponseWriter mocks base method.
func (m *MockIApiSession) ResponseWriter() http.ResponseWriter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponseWriter")
	ret0, _ := ret[0].(http.ResponseWriter)
	return ret0
}

// ResponseWriter indicates an expected call of ResponseWriter.
func (mr *MockIApiSessionMockRecorder) ResponseWriter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseWriter", reflect.TypeOf((*MockIApiSession)(nil).ResponseWriter))
}

// ScanParams mocks base method.
func (m *MockIApiSession) ScanParams(arg0 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ScanParams", arg0)
}

// ScanParams indicates an expected call of ScanParams.
func (mr *MockIApiSessionMockRecorder) ScanParams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanParams", reflect.TypeOf((*MockIApiSession)(nil).ScanParams), arg0)
}

// SetApi mocks base method.
func (m *MockIApiSession) SetApi(arg0 api.IApi) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetApi", arg0)
}

// SetApi indicates an expected call of SetApi.
func (mr *MockIApiSessionMockRecorder) SetApi(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetApi", reflect.TypeOf((*MockIApiSession)(nil).SetApi), arg0)
}

// SetClientType mocks base method.
func (m *MockIApiSession) SetClientType(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClientType", arg0)
}

// SetClientType indicates an expected call of SetClientType.
func (mr *MockIApiSessionMockRecorder) SetClientType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClientType", reflect.TypeOf((*MockIApiSession)(nil).SetClientType), arg0)
}

// SetData mocks base method.
func (m *MockIApiSession) SetData(arg0 string, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetData", arg0, arg1)
}

// SetData indicates an expected call of SetData.
func (mr *MockIApiSessionMockRecorder) SetData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetData", reflect.TypeOf((*MockIApiSession)(nil).SetData), arg0, arg1)
}

// SetHeader mocks base method.
func (m *MockIApiSession) SetHeader(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHeader", arg0, arg1)
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockIApiSessionMockRecorder) SetHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockIApiSession)(nil).SetHeader), arg0, arg1)
}

// SetJwtBody mocks base method.
func (m *MockIApiSession) SetJwtBody(arg0 map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetJwtBody", arg0)
}

// SetJwtBody indicates an expected call of SetJwtBody.
func (mr *MockIApiSessionMockRecorder) SetJwtBody(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJwtBody", reflect.TypeOf((*MockIApiSession)(nil).SetJwtBody), arg0)
}

// SetJwtHeaderName mocks base method.
func (m *MockIApiSession) SetJwtHeaderName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetJwtHeaderName", arg0)
}

// SetJwtHeaderName indicates an expected call of SetJwtHeaderName.
func (mr *MockIApiSessionMockRecorder) SetJwtHeaderName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJwtHeaderName", reflect.TypeOf((*MockIApiSession)(nil).SetJwtHeaderName), arg0)
}

// SetLang mocks base method.
func (m *MockIApiSession) SetLang(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLang", arg0)
}

// SetLang indicates an expected call of SetLang.
func (mr *MockIApiSessionMockRecorder) SetLang(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLang", reflect.TypeOf((*MockIApiSession)(nil).SetLang), arg0)
}

// SetOriginalParams mocks base method.
func (m *MockIApiSession) SetOriginalParams(arg0 map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOriginalParams", arg0)
}

// SetOriginalParams indicates an expected call of SetOriginalParams.
func (mr *MockIApiSessionMockRecorder) SetOriginalParams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOriginalParams", reflect.TypeOf((*MockIApiSession)(nil).SetOriginalParams), arg0)
}

// SetParams mocks base method.
func (m *MockIApiSession) SetParams(arg0 map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetParams", arg0)
}

// SetParams indicates an expected call of SetParams.
func (mr *MockIApiSessionMockRecorder) SetParams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParams", reflect.TypeOf((*MockIApiSession)(nil).SetParams), arg0)
}

// SetRequest mocks base method.
func (m *MockIApiSession) SetRequest(arg0 *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRequest", arg0)
}

// SetRequest indicates an expected call of SetRequest.
func (mr *MockIApiSessionMockRecorder) SetRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRequest", reflect.TypeOf((*MockIApiSession)(nil).SetRequest), arg0)
}

// SetResponseWriter mocks base method.
func (m *MockIApiSession) SetResponseWriter(arg0 http.ResponseWriter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetResponseWriter", arg0)
}

// SetResponseWriter indicates an expected call of SetResponseWriter.
func (mr *MockIApiSessionMockRecorder) SetResponseWriter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResponseWriter", reflect.TypeOf((*MockIApiSession)(nil).SetResponseWriter), arg0)
}

// SetStatusCode mocks base method.
func (m *MockIApiSession) SetStatusCode(arg0 api_session.StatusCode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStatusCode", arg0)
}

// SetStatusCode indicates an expected call of SetStatusCode.
func (mr *MockIApiSessionMockRecorder) SetStatusCode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatusCode", reflect.TypeOf((*MockIApiSession)(nil).SetStatusCode), arg0)
}

// SetUserId mocks base method.
func (m *MockIApiSession) SetUserId(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetUserId", arg0)
}

// SetUserId indicates an expected call of SetUserId.
func (mr *MockIApiSessionMockRecorder) SetUserId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUserId", reflect.TypeOf((*MockIApiSession)(nil).SetUserId), arg0)
}

// UrlParams mocks base method.
func (m *MockIApiSession) UrlParams() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UrlParams")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// UrlParams indicates an expected call of UrlParams.
func (mr *MockIApiSessionMockRecorder) UrlParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UrlParams", reflect.TypeOf((*MockIApiSession)(nil).UrlParams))
}

// UserId mocks base method.
func (m *MockIApiSession) UserId() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserId")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// UserId indicates an expected call of UserId.
func (mr *MockIApiSessionMockRecorder) UserId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserId", reflect.TypeOf((*MockIApiSession)(nil).UserId))
}

// WriteJson mocks base method.
func (m *MockIApiSession) WriteJson(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteJson", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteJson indicates an expected call of WriteJson.
func (mr *MockIApiSessionMockRecorder) WriteJson(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteJson", reflect.TypeOf((*MockIApiSession)(nil).WriteJson), arg0)
}

// WriteText mocks base method.
func (m *MockIApiSession) WriteText(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteText", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteText indicates an expected call of WriteText.
func (mr *MockIApiSessionMockRecorder) WriteText(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteText", reflect.TypeOf((*MockIApiSession)(nil).WriteText), arg0)
}
