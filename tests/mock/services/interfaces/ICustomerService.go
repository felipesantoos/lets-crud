// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/services/interfaces/ICustomerService.go

package mock_interfaces

import (
	gomock "github.com/golang/mock/gomock"
	request "letscrud/src/api/endpoints/dto/request"
	response "letscrud/src/api/endpoints/dto/response"
	errs "letscrud/src/domain/errs"
	reflect "reflect"
)

// MockICustomerService is a mock of ICustomerService interface
type MockICustomerService struct {
	ctrl     *gomock.Controller
	recorder *MockICustomerServiceMockRecorder
}

// MockICustomerServiceMockRecorder is the mock recorder for MockICustomerService
type MockICustomerServiceMockRecorder struct {
	mock *MockICustomerService
}

// NewMockICustomerService creates a new mock instance
func NewMockICustomerService(ctrl *gomock.Controller) *MockICustomerService {
	mock := &MockICustomerService{ctrl: ctrl}
	mock.recorder = &MockICustomerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockICustomerService) EXPECT() *MockICustomerServiceMockRecorder {
	return _m.recorder
}

// CreateNewCustomer mocks base method
func (_m *MockICustomerService) CreateNewCustomer(customerResquest request.CustomerRequest) (int, *errs.ApiError) {
	ret := _m.ctrl.Call(_m, "CreateNewCustomer", customerResquest)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*errs.ApiError)
	return ret0, ret1
}

// CreateNewCustomer indicates an expected call of CreateNewCustomer
func (_mr *MockICustomerServiceMockRecorder) CreateNewCustomer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateNewCustomer", reflect.TypeOf((*MockICustomerService)(nil).CreateNewCustomer), arg0)
}

// ReadAllCustomers mocks base method
func (_m *MockICustomerService) ReadAllCustomers() ([]response.CustomerResponse, *errs.ApiError) {
	ret := _m.ctrl.Call(_m, "ReadAllCustomers")
	ret0, _ := ret[0].([]response.CustomerResponse)
	ret1, _ := ret[1].(*errs.ApiError)
	return ret0, ret1
}

// ReadAllCustomers indicates an expected call of ReadAllCustomers
func (_mr *MockICustomerServiceMockRecorder) ReadAllCustomers() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadAllCustomers", reflect.TypeOf((*MockICustomerService)(nil).ReadAllCustomers))
}

// ReadCustomerById mocks base method
func (_m *MockICustomerService) ReadCustomerById(id int) (*response.CustomerResponse, *errs.ApiError) {
	ret := _m.ctrl.Call(_m, "ReadCustomerById", id)
	ret0, _ := ret[0].(*response.CustomerResponse)
	ret1, _ := ret[1].(*errs.ApiError)
	return ret0, ret1
}

// ReadCustomerById indicates an expected call of ReadCustomerById
func (_mr *MockICustomerServiceMockRecorder) ReadCustomerById(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadCustomerById", reflect.TypeOf((*MockICustomerService)(nil).ReadCustomerById), arg0)
}

// UpdateCustomerById mocks base method
func (_m *MockICustomerService) UpdateCustomerById(id int, customerRequest request.CustomerRequest) (bool, *errs.ApiError) {
	ret := _m.ctrl.Call(_m, "UpdateCustomerById", id, customerRequest)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*errs.ApiError)
	return ret0, ret1
}

// UpdateCustomerById indicates an expected call of UpdateCustomerById
func (_mr *MockICustomerServiceMockRecorder) UpdateCustomerById(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateCustomerById", reflect.TypeOf((*MockICustomerService)(nil).UpdateCustomerById), arg0, arg1)
}

// DeleteCustomerById mocks base method
func (_m *MockICustomerService) DeleteCustomerById(id int) (bool, *errs.ApiError) {
	ret := _m.ctrl.Call(_m, "DeleteCustomerById", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*errs.ApiError)
	return ret0, ret1
}

// DeleteCustomerById indicates an expected call of DeleteCustomerById
func (_mr *MockICustomerServiceMockRecorder) DeleteCustomerById(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteCustomerById", reflect.TypeOf((*MockICustomerService)(nil).DeleteCustomerById), arg0)
}
