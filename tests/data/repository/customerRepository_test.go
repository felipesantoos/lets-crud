package repository_test

import (
	repository "letscrud/data/repository"
	"letscrud/domain/errs"
	"letscrud/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewCustomer(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	customerRequest := tests.GetValidCustomerRequestForRepository()
	returnedId, returnedApiError := repo.CreateNewCustomer(customerRequest)
	expectedId := int64(1)

	assert.Equal(t, expectedId, returnedId)
	assert.Nil(t, returnedApiError)
}

func TestCreateNewCustomerErrorCPFTooLong(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	customerRequest := tests.GetCustomerRequestWithCPFTooLongForRepository()
	returnedId, returnedApiError := repo.CreateNewCustomer(customerRequest)
	expectedId := int64(0)
	expectedApiError := errs.NewBadRequestError("O CPF informado é muito longo!")

	assert.Equal(t, expectedId, returnedId)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestCreateNewCustomerErrorBirthDateBadFormatted(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	customerRequest := tests.GetCustomerRequestWithBirthDateBadFormatted()
	returnedId, returnedApiError := repo.CreateNewCustomer(customerRequest)
	expectedId := int64(0)
	expectedApiError := errs.NewBadRequestError("A data de nascimento informada está mal formadata! Formato correto: YYYY-MM-DD.")

	assert.Equal(t, expectedId, returnedId)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestCreateNewCustomerErrorCPFAlreadyRegistered(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
		tests.INSERT_CUSTOMER_ID_1,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	customerRequest := tests.GetValidCustomerRequestForRepository()
	returnedId, returnedApiError := repo.CreateNewCustomer(customerRequest)
	expectedId := int64(0)
	expectedApiError := errs.NewBadRequestError("O CPF informado já foi cadastrado no banco de dados!")

	assert.Equal(t, expectedId, returnedId)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestReadAllCustomers(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
		tests.INSERT_CUSTOMER_ID_1,
		tests.INSERT_CUSTOMER_ID_2,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	returnedCustomerModelList, returnedApiError := repo.ReadAllCustomers()
	expectedCustomerModelList := tests.GetExpectedCustomerModelList()

	assert.Equal(t, expectedCustomerModelList, returnedCustomerModelList)
	assert.Nil(t, returnedApiError)
}

func TestReadAllCustomersErrorZeroRecordsReturned(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	returnedCustomerModelList, returnedApiError := repo.ReadAllCustomers()
	expectedApiError := errs.NewBadRequestError("Nenhum registro foi retornado!")

	assert.Nil(t, returnedCustomerModelList)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestReadCustomerById(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
		tests.INSERT_CUSTOMER_ID_1,
		tests.INSERT_CUSTOMER_ID_2,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()

	t.Run("TestReadCustomerById1", func(t *testing.T) {
		returnedCustomerModel, returnedApiError := repo.ReadCustomerById(1)
		expectedCustomerModel := tests.GetExpectedCustomerModelWithId1()

		assert.Equal(t, expectedCustomerModel, returnedCustomerModel)
		assert.Nil(t, returnedApiError)
	})
	t.Run("TestReadCustomerById2", func(t *testing.T) {
		returnedCustomerModel, returnedApiError := repo.ReadCustomerById(2)
		expectedCustomerModel := tests.GetExpectedCustomerModelWithId2()

		assert.Equal(t, expectedCustomerModel, returnedCustomerModel)
		assert.Nil(t, returnedApiError)
	})
}

func TestReadCustomerByIdErrorCustomerNotFound(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	returnedCustomerModel, returnedApiError := repo.ReadCustomerById(1)
	expectedApiError := errs.NewBadRequestError("O cliente informado não foi encontrado!")

	assert.Nil(t, returnedCustomerModel)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestUpdateCustomerById(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
		tests.INSERT_CUSTOMER_ID_1,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	customerRequest := tests.GetCustomerRequestUpdated()
	returnedIsUpdated, returnedApiError := repo.UpdateCustomerById(1, customerRequest)
	expectedIsUpdated := true

	assert.Equal(t, expectedIsUpdated, returnedIsUpdated)
	assert.Nil(t, returnedApiError)
}

func TestUpdateCustomerByIdErrorCPFTooLong(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
		tests.INSERT_CUSTOMER_ID_1,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	customerRequest := tests.GetCustomerRequestUpdatedWithCPFTooLong()
	returnedIsUpdated, returnedApiError := repo.UpdateCustomerById(1, customerRequest)
	expectedIsUpdated := false
	expectedApiError := errs.NewBadRequestError("O CPF informado é muito longo!")

	assert.Equal(t, expectedIsUpdated, returnedIsUpdated)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestUpdateCustomerByIdErrorBirthDateBadFormatted(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
		tests.INSERT_CUSTOMER_ID_1,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	customerRequest := tests.GetCustomerRequestUpdatedWithBirthDateBadFormatted()
	returnedIsUpdated, returnedApiError := repo.UpdateCustomerById(1, customerRequest)
	expectedIsUpdated := false
	expectedApiError := errs.NewBadRequestError("A data de nascimento informada está mal formadata! Formato correto: YYYY-MM-DD.")

	assert.Equal(t, expectedIsUpdated, returnedIsUpdated)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestUpdateCustomerByIdErrorCPFAlreadyRegistered(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
		tests.INSERT_CUSTOMER_ID_1,
		tests.INSERT_CUSTOMER_ID_2,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	customerRequest := tests.GetCustomerRequestUpdatedWithCPFAlreadyRegistered()
	returnedIsUpdated, returnedApiError := repo.UpdateCustomerById(1, customerRequest)
	expectedIsUpdated := false
	expectedApiError := errs.NewBadRequestError("O CPF informado já foi cadastrado no banco de dados!")

	assert.Equal(t, expectedIsUpdated, returnedIsUpdated)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestUpdateCustomerByIdErrorCustomerNotFound(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
		tests.INSERT_CUSTOMER_ID_1,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	customerRequest := tests.GetCustomerRequestUpdated()
	returnedIsUpdated, returnedApiError := repo.UpdateCustomerById(2, customerRequest)
	expectedIsUpdated := false
	expectedApiError := errs.NewBadRequestError("O cliente informado não foi encontrado ou os dados informados são idênticos aos já cadastrados!")

	assert.Equal(t, expectedIsUpdated, returnedIsUpdated)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestUpdateCustomerByIdErrorIdenticalData(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
		tests.INSERT_CUSTOMER_ID_1,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	customerRequest := tests.GetCustomerRequestUpdatedWithIdenticalData()
	returnedIsUpdated, returnedApiError := repo.UpdateCustomerById(1, customerRequest)
	expectedIsUpdated := false
	expectedApiError := errs.NewBadRequestError("O cliente informado não foi encontrado ou os dados informados são idênticos aos já cadastrados!")

	assert.Equal(t, expectedIsUpdated, returnedIsUpdated)
	assert.Equal(t, expectedApiError, returnedApiError)
}

func TestDeleteCustomerById(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
		tests.INSERT_CUSTOMER_ID_1,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	returnedIsDeleted, returnedApiError := repo.DeleteCustomerById(1)
	expectedIsDeleted := true

	assert.Equal(t, expectedIsDeleted, returnedIsDeleted)
	assert.Nil(t, returnedApiError)
}

func TestDeleteCustomerByIdErrorCustomerNotFound(t *testing.T) {
	queries := []string{
		tests.CLEAR_CUSTOMER_TABLE,
		tests.RESET_AUTO_INCREMET,
		tests.INSERT_CUSTOMER_ID_1,
	}
	tests.SetUp(queries)

	repo := repository.NewCustomerRepository()
	returnedIsDeleted, returnedApiError := repo.DeleteCustomerById(2)
	expectedIsDeleted := false

	assert.Equal(t, expectedIsDeleted, returnedIsDeleted)
	assert.Nil(t, returnedApiError)
}
