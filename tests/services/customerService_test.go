package services_test

import (
	"letscrud/domain/errs"
	"letscrud/services"
	"letscrud/tests"
	IMockInterfaces "letscrud/tests/mock/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewCustomer(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	expectedLastInsertedIdFromRepository := int64(1)
	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().CreateNewCustomer(gomock.Any()).Return(
		expectedLastInsertedIdFromRepository,
		nil,
	)

	service := services.NewCustomerService(repo)
	customerRequest := tests.GetValidCustomerRequestForService()
	returnedLastInsertedId, returnedApiError := service.CreateNewCustomer(customerRequest)
	expectedLastInsertedId := int64(1)

	assert.Equal(t, expectedLastInsertedId, returnedLastInsertedId)
	assert.Nil(t, returnedApiError)
}
func TestCreateNewCustomerErrorInvalidCPF(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	service := services.NewCustomerService(repo)

	customerRequest := tests.GetCustomerRequestWithInvalidCPF()
	returnedLastInsertedId, returnedApiError := service.CreateNewCustomer(customerRequest)
	expectedLastInsertedId := int64(0)
	expectedApiError := errs.NewBadRequestError("O CPF informado é inválido!")

	assert.Equal(t, expectedLastInsertedId, returnedLastInsertedId)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestCreateNewCustomerErrorBirthDateBadFormatted(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	returnedApiErrorFromRepository := errs.NewBadRequestError("A data de nascimento informada está mal formadata! Formato correto: YYYY-MM-DD.")
	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().CreateNewCustomer(gomock.Any()).Return(int64(0), returnedApiErrorFromRepository)

	service := services.NewCustomerService(repo)
	customerRequest := tests.GetCustomerRequestWithBirthDateBadFormattedForService()
	returnedLastInsertedId, returnedApiError := service.CreateNewCustomer(customerRequest)
	expectedLastInsertedId := int64(0)
	expectedApiError := errs.NewBadRequestError("A data de nascimento informada está mal formadata! Formato correto: YYYY-MM-DD.")

	assert.Equal(t, expectedLastInsertedId, returnedLastInsertedId)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestCreateNewCustomerErrorCPFAlreadyRegistered(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	returnedApiErrorFromRepository := errs.NewBadRequestError("O CPF informado já foi cadastrado no banco de dados!")
	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().CreateNewCustomer(gomock.Any()).Return(int64(0), returnedApiErrorFromRepository)

	service := services.NewCustomerService(repo)
	customerRequest := tests.GetValidCustomerRequestForService()
	returnedLastInsertedId, returnedApiError := service.CreateNewCustomer(customerRequest)
	expectedLastInsertedId := int64(0)
	expectedApiError := errs.NewBadRequestError("O CPF informado já foi cadastrado no banco de dados!")

	assert.Equal(t, expectedLastInsertedId, returnedLastInsertedId)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestCreateNewCustomerErrorNameTooShort(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	service := services.NewCustomerService(repo)

	t.Run("TestEmptyName", func(t *testing.T) {
		customerRequest := tests.GetCustomerRequestWithEmptyName()
		returnedLastInsertedId, returnedApiError := service.CreateNewCustomer(customerRequest)
		expectedLastInsertedId := int64(0)
		expectedApiError := errs.NewBadRequestError("O nome deve possuir ao menos duas palavras!")

		assert.Equal(t, expectedLastInsertedId, returnedLastInsertedId)
		assert.Equal(t, expectedApiError, returnedApiError)
	})
	t.Run("TestNameWithOnlyOneWord", func(t *testing.T) {
		customerRequest := tests.GetCustomerRequestWithNameWithOnlyOneWord()
		returnedLastInsertedId, returnedApiError := service.CreateNewCustomer(customerRequest)
		expectedLastInsertedId := int64(0)
		expectedApiError := errs.NewBadRequestError("O nome deve possuir ao menos duas palavras!")

		assert.Equal(t, expectedLastInsertedId, returnedLastInsertedId)
		assert.Equal(t, expectedApiError, returnedApiError)
	})
}

func TestReadAllCustomers(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().ReadAllCustomers().Return(tests.GetReturnedCustomerModelList(), nil)

	service := services.NewCustomerService(repo)
	returnedCustomerResponseList, returnedApiError := service.ReadAllCustomers()
	expectedCustomerResponseList := tests.GetExpectedCustomerResponseList()

	assert.Equal(t, expectedCustomerResponseList, returnedCustomerResponseList)
	assert.Nil(t, returnedApiError)
}

func TestReadAllCustomersErrorZeroRecordsReturned(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	returnedApiErrorFromRepository := errs.NewBadRequestError("Nenhum registro foi retornado!")
	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().ReadAllCustomers().Return(nil, returnedApiErrorFromRepository)

	service := services.NewCustomerService(repo)
	returnedCustomerModelList, returnedApiError := service.ReadAllCustomers()
	expectedApiError := errs.NewBadRequestError("Nenhum registro foi retornado!")

	assert.Nil(t, returnedCustomerModelList)
	assert.Equal(t, expectedApiError, returnedApiError)
}

/*
	- TestReadCustomerById
	- TestUpdateCustomerById
	- TestDeleteCustomerById
	- Testar removedor de pontuação.
	- Testar adicionador de pontuação.
	- Testar validade do CPF.
	- Testar invalidade do CPF.
	- Testar validade do nome.
	- Testar nome vazio.
	- Testar nome com uma única palavra.
	- Testar nome com caracteres especiais e números.
*/
