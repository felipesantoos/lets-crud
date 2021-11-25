// Code generated by MockGen. DO NOT EDIT.
// Source: data/interfaces/interfaces.go

package mock_interfaces

import (
	gomock "github.com/golang/mock/gomock"
	errs "letscrud/domain/errs"
	models "letscrud/domain/models"
	request "letscrud/endpoints/dto/request"
	reflect "reflect"
)

// MockICustomerRepository is a mock of ICustomerRepository interface
type MockICustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICustomerRepositoryMockRecorder
}

// MockICustomerRepositoryMockRecorder is the mock recorder for MockICustomerRepository
type MockICustomerRepositoryMockRecorder struct {
	mock *MockICustomerRepository
}

// NewMockICustomerRepository creates a new mock instance
func NewMockICustomerRepository(ctrl *gomock.Controller) *MockICustomerRepository {
	mock := &MockICustomerRepository{ctrl: ctrl}
	mock.recorder = &MockICustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockICustomerRepository) EXPECT() *MockICustomerRepositoryMockRecorder {
	return _m.recorder
}

// CreateNewCustomer mocks base method
func (_m *MockICustomerRepository) CreateNewCustomer(customerRequest request.CustomerRequest) (int64, *errs.ApiError) {
	ret := _m.ctrl.Call(_m, "CreateNewCustomer", customerRequest)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(*errs.ApiError)
	return ret0, ret1
}

// CreateNewCustomer indicates an expected call of CreateNewCustomer
func (_mr *MockICustomerRepositoryMockRecorder) CreateNewCustomer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateNewCustomer", reflect.TypeOf((*MockICustomerRepository)(nil).CreateNewCustomer), arg0)
}

// ReadAllCustomers mocks base method
func (_m *MockICustomerRepository) ReadAllCustomers() ([]models.Customer, *errs.ApiError) {
	ret := _m.ctrl.Call(_m, "ReadAllCustomers")
	ret0, _ := ret[0].([]models.Customer)
	ret1, _ := ret[1].(*errs.ApiError)
	return ret0, ret1
}

// ReadAllCustomers indicates an expected call of ReadAllCustomers
func (_mr *MockICustomerRepositoryMockRecorder) ReadAllCustomers() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadAllCustomers", reflect.TypeOf((*MockICustomerRepository)(nil).ReadAllCustomers))
}

// ReadCustomerById mocks base method
func (_m *MockICustomerRepository) ReadCustomerById(id int64) (*models.Customer, *errs.ApiError) {
	ret := _m.ctrl.Call(_m, "ReadCustomerById", id)
	ret0, _ := ret[0].(*models.Customer)
	ret1, _ := ret[1].(*errs.ApiError)
	return ret0, ret1
}

// ReadCustomerById indicates an expected call of ReadCustomerById
func (_mr *MockICustomerRepositoryMockRecorder) ReadCustomerById(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ReadCustomerById", reflect.TypeOf((*MockICustomerRepository)(nil).ReadCustomerById), arg0)
}

// UpdateCustomerById mocks base method
func (_m *MockICustomerRepository) UpdateCustomerById(id int64, customerRequest request.CustomerRequest) (bool, *errs.ApiError) {
	ret := _m.ctrl.Call(_m, "UpdateCustomerById", id, customerRequest)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*errs.ApiError)
	return ret0, ret1
}

// UpdateCustomerById indicates an expected call of UpdateCustomerById
func (_mr *MockICustomerRepositoryMockRecorder) UpdateCustomerById(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "UpdateCustomerById", reflect.TypeOf((*MockICustomerRepository)(nil).UpdateCustomerById), arg0, arg1)
}

// DeleteCustomerById mocks base method
func (_m *MockICustomerRepository) DeleteCustomerById(id int64) (bool, *errs.ApiError) {
	ret := _m.ctrl.Call(_m, "DeleteCustomerById", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*errs.ApiError)
	return ret0, ret1
}

// DeleteCustomerById indicates an expected call of DeleteCustomerById
func (_mr *MockICustomerRepositoryMockRecorder) DeleteCustomerById(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteCustomerById", reflect.TypeOf((*MockICustomerRepository)(nil).DeleteCustomerById), arg0)
}