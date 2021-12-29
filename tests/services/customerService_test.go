package services_test

import (
	"letscrud/src/domain/errs"
	"letscrud/src/services"
	"letscrud/tests"
	IMockInterfaces "letscrud/tests/mock/services/repositories"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewCustomer(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	expectedLastInsertedIdFromRepository := int(1)
	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().CreateNewCustomer(gomock.Any()).Return(
		expectedLastInsertedIdFromRepository,
		nil,
	)

	service := services.NewCustomerService(repo)
	customerRequest := tests.GetValidCustomerRequestForService()
	returnedLastInsertedId, returnedApiError := service.CreateNewCustomer(customerRequest)
	expectedLastInsertedId := int(1)

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
	expectedLastInsertedId := int(0)
	expectedApiError := errs.NewBadRequestError("O CPF informado é inválido!")

	assert.Equal(t, expectedLastInsertedId, returnedLastInsertedId)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestCreateNewCustomerErrorBirthDateBadFormatted(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	returnedApiErrorFromRepository := errs.NewBadRequestError("A data de nascimento informada está mal formadata! Formato correto: YYYY-MM-DD.")
	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().CreateNewCustomer(gomock.Any()).Return(int(0), returnedApiErrorFromRepository)

	service := services.NewCustomerService(repo)
	customerRequest := tests.GetCustomerRequestWithBirthDateBadFormattedForService()
	returnedLastInsertedId, returnedApiError := service.CreateNewCustomer(customerRequest)
	expectedLastInsertedId := int(0)
	expectedApiError := errs.NewBadRequestError("A data de nascimento informada está mal formadata! Formato correto: YYYY-MM-DD.")

	assert.Equal(t, expectedLastInsertedId, returnedLastInsertedId)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestCreateNewCustomerErrorCPFAlreadyRegistered(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	returnedApiErrorFromRepository := errs.NewBadRequestError("O CPF informado já foi cadastrado no banco de dados!")
	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().CreateNewCustomer(gomock.Any()).Return(int(0), returnedApiErrorFromRepository)

	service := services.NewCustomerService(repo)
	customerRequest := tests.GetValidCustomerRequestForService()
	returnedLastInsertedId, returnedApiError := service.CreateNewCustomer(customerRequest)
	expectedLastInsertedId := int(0)
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
		expectedLastInsertedId := int(0)
		expectedApiError := errs.NewBadRequestError("O nome deve possuir ao menos duas palavras!")

		assert.Equal(t, expectedLastInsertedId, returnedLastInsertedId)
		assert.Equal(t, expectedApiError, returnedApiError)
	})
	t.Run("TestNameWithOnlyOneWord", func(t *testing.T) {
		customerRequest := tests.GetCustomerRequestWithNameWithOnlyOneWord()
		returnedLastInsertedId, returnedApiError := service.CreateNewCustomer(customerRequest)
		expectedLastInsertedId := int(0)
		expectedApiError := errs.NewBadRequestError("O nome deve possuir ao menos duas palavras!")

		assert.Equal(t, expectedLastInsertedId, returnedLastInsertedId)
		assert.Equal(t, expectedApiError, returnedApiError)
	})
}

func TestCreateNewCustomerErrorInvalidName(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	service := services.NewCustomerService(repo)

	t.Run("TestNameWithSpecialCharacters", func(t *testing.T) {
		customerRequest := tests.GetCustomerRequestWithNameWithSpecialCharacters()
		returnedLastInsertedId, returnedError := service.CreateNewCustomer(customerRequest)
		expectedLastInsertedId := int(0)
		expectedError := errs.NewBadRequestError("O nome só pode conter letras e espaços!")

		assert.Equal(t, expectedLastInsertedId, returnedLastInsertedId)
		assert.Equal(t, expectedError, returnedError)
	})
	t.Run("TestNameWithNumber", func(t *testing.T) {
		customerRequest := tests.GetCustomerRequestWithNameWithNumber()
		returnedLastInsertedId, returnedError := service.CreateNewCustomer(customerRequest)
		expectedLastInsertedId := int(0)
		expectedError := errs.NewBadRequestError("O nome só pode conter letras e espaços!")

		assert.Equal(t, expectedLastInsertedId, returnedLastInsertedId)
		assert.Equal(t, expectedError, returnedError)
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

func TestReadCustomerById(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	returnedCustomerModel := tests.GetReturnedCustomerModel()
	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().ReadCustomerById(gomock.Any()).Return(returnedCustomerModel, nil)

	service := services.NewCustomerService(repo)
	returnedCustomerResponse, returnedApiError := service.ReadCustomerById(1)
	expectedCustomerResponse := tests.GetExpectedCustomerResponse()

	assert.Equal(t, expectedCustomerResponse, returnedCustomerResponse)
	assert.Nil(t, returnedApiError)
}

func TestReadCustomerByIdErrorCustomerNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	returnedApiErrorFromRepository := errs.NewBadRequestError("O cliente informado não foi encontrado!")
	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().ReadCustomerById(gomock.Any()).Return(nil, returnedApiErrorFromRepository)

	service := services.NewCustomerService(repo)
	returnedCustomerResponse, returnedApiError := service.ReadCustomerById(1)
	expectedApiError := errs.NewBadRequestError("O cliente informado não foi encontrado!")

	assert.Nil(t, returnedCustomerResponse)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestUpdateCustomerById(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().UpdateCustomerById(gomock.Any(), gomock.Any()).Return(true, nil)

	service := services.NewCustomerService(repo)
	customerRequest := tests.GetCustomerRequestUpdatedForService()
	returnedIsUpdated, returnedApiError := service.UpdateCustomerById(1, customerRequest)

	assert.Equal(t, true, returnedIsUpdated)
	assert.Nil(t, returnedApiError)
}

func TestUpdateCustomerByIdErrorInvalidCPF(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	service := services.NewCustomerService(repo)

	customerRequest := tests.GetCustomerRequestWithInvalidCPF()
	returnedIsUpdated, returnedApiError := service.UpdateCustomerById(1, customerRequest)
	expectedApiError := errs.NewBadRequestError("O CPF informado é inválido!")

	assert.Equal(t, false, returnedIsUpdated)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestUpdateCustomerByIdErrorCustomerNotFoundOrIdenticalData(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	returnedApiErrorFromRepository := errs.NewBadRequestError("O cliente informado não foi encontrado ou os dados informados são idênticos aos já cadastrados!")
	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().UpdateCustomerById(gomock.Any(), gomock.Any()).Return(false, returnedApiErrorFromRepository)

	service := services.NewCustomerService(repo)
	customerRequest := tests.GetValidCustomerRequestForService()
	returnedIsUpdated, returnedApiError := service.UpdateCustomerById(1, customerRequest)
	expectedApiError := errs.NewBadRequestError("O cliente informado não foi encontrado ou os dados informados são idênticos aos já cadastrados!")

	assert.Equal(t, false, returnedIsUpdated)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestUpdateCustomerByIdErrorNameTooShort(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	service := services.NewCustomerService(repo)

	t.Run("TestEmptyName", func(t *testing.T) {
		customerRequest := tests.GetCustomerRequestUpdatedWithEmptyName()
		returnedIsUpdated, returnedApiError := service.UpdateCustomerById(1, customerRequest)
		expectedIsUpdated := false
		expectedApiError := errs.NewBadRequestError("O nome deve possuir ao menos duas palavras!")

		assert.Equal(t, expectedIsUpdated, returnedIsUpdated)
		assert.Equal(t, expectedApiError, returnedApiError)
	})
	t.Run("TestNameWithOnlyOneWord", func(t *testing.T) {
		customerRequest := tests.GetCustomerRequestUpdatedWithNameWithOnlyOneWord()
		returnedIsUpdated, returnedApiError := service.UpdateCustomerById(1, customerRequest)
		expectedIsUpdated := false
		expectedApiError := errs.NewBadRequestError("O nome deve possuir ao menos duas palavras!")

		assert.Equal(t, expectedIsUpdated, returnedIsUpdated)
		assert.Equal(t, expectedApiError, returnedApiError)
	})
}

func TestDeleteCustomerById(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().DeleteCustomerById(gomock.Any()).Return(true, nil)

	service := services.NewCustomerService(repo)
	returnedIsDeleted, returnedApiError := service.DeleteCustomerById(1)

	assert.Equal(t, true, returnedIsDeleted)
	assert.Nil(t, returnedApiError)
}

func TestDeleteCustomerByIdCustomerNotFound(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repo := IMockInterfaces.NewMockICustomerRepository(controller)
	repo.EXPECT().DeleteCustomerById(gomock.Any()).Return(false, nil)

	service := services.NewCustomerService(repo)
	returnedIsDeleted, returnedApiError := service.DeleteCustomerById(1)

	assert.Equal(t, false, returnedIsDeleted)
	assert.Nil(t, returnedApiError)
}
