package handlers_test

import (
	"encoding/json"
	"letscrud/endpoints/handlers"
	"letscrud/tests"
	IMockInterfaces "letscrud/tests/mock/services"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreateNewCustomer(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	service := IMockInterfaces.NewMockICustomerService(controller)
	service.EXPECT().CreateNewCustomer(gomock.Any()).Return(int64(1), nil)
	handler := handlers.NewCustomerHandlerWithCustomService(service)

	requestedJSON, _ := json.Marshal(tests.GetValidCustomerRequestForHandler())
	expectedJSON, _ := json.Marshal(tests.GetCreateNewCustomerResponse())

	tests.RunGenericPostRequestTest(
		t,
		"/customer",
		tests.NewContextParams(),
		string(requestedJSON),
		handler.CreateNewCustomer,
		http.StatusCreated,
		string(expectedJSON),
	)
}

/*
	- TestCreateNewCustomerErrorInvalidCPF
	- TestCreateNewCustomerErrorBirthDateBadFormatted
	- TestCreateNewCustomerErrorCPFAlreadyRegistered
	- TestCreateNewCustomerErrorNameTooShort
		- TestEmptyName
		- TestNameWithOnlyOneWord
	- TestCreateNewCustomerErrorInvalidName
		- TestNameWithSpecialCharacters
		- TestNameWithNumber
	- TestReadAllCustomers
	- TestReadAllCustomersErrorZeroRecordsReturned
	- TestReadCustomerById
	- TestReadCustomerByIdErrorCustomerNotFound
	- TestUpdateCustomerById
	- TestUpdateCustomerByIdErrorInvalidCPF
	- TestUpdateCustomerByIdErrorCustomerNotFoundOrIdenticalData
	- TestUpdateCustomerByIdErrorNameTooShort
		- TestEmptyName
		- TestNameWithOnlyOneWord
	- TestDeleteCustomerById
	- TestDeleteCustomerByIdCustomerNotFound
	- Testes específicos da camada de handler.
*/
