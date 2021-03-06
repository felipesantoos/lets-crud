package handlers

import (
	"letscrud/src/api/endpoints/dto/request"
	"letscrud/src/infra/mysql"
	"letscrud/src/services"
	"letscrud/src/services/interfaces"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	service interfaces.ICustomerService
}

func NewCustomerHandler() *CustomerHandler {
	repo := mysql.NewCustomerMySqlRepository()
	service := services.NewCustomerService(repo)

	return &CustomerHandler{service: service}
}

func NewCustomerHandlerWithCustomService(service interfaces.ICustomerService) *CustomerHandler {
	return &CustomerHandler{service}
}

func (ch CustomerHandler) CreateNewCustomer(c echo.Context) error {
	log.Println("H [CreateNewCustomer]")

	customer := new(request.CustomerRequest)
	if err := c.Bind(customer); err != nil {
		log.Println("H [CreateNewCustomer]: " + err.Error())
	}

	lastInsertId, apiErr := ch.service.CreateNewCustomer(*customer)

	if apiErr != nil {
		response := map[string]string{
			"error": apiErr.Message,
		}

		return c.JSON(apiErr.StatusCode, response)
	}

	response := map[string]interface{}{
		"id": lastInsertId,
	}

	return c.JSON(http.StatusCreated, response)
}

func (ch CustomerHandler) ReadAllCustomers(c echo.Context) error {
	log.Println("H [ReadAllCustomers]")

	customerList, apiErr := ch.service.ReadAllCustomers()
	if apiErr != nil {
		response := map[string]string{
			"error": apiErr.Message,
		}

		return c.JSON(apiErr.StatusCode, response)
	}

	return c.JSON(http.StatusOK, customerList)
}

func (ch CustomerHandler) ReadCustomerById(c echo.Context) error {
	log.Println("H [ReadCustomerById]")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("H [ReadCustomerById]: " + err.Error())
		response := map[string]string{
			"error": "Verifique o tipo do dado informado no parâmentro :id.",
		}

		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	customer, apiError := ch.service.ReadCustomerById(id)
	if apiError != nil {
		response := map[string]string{
			"error": apiError.Message,
		}

		return c.JSON(apiError.StatusCode, response)
	}

	return c.JSON(http.StatusOK, customer)
}

func (ch CustomerHandler) UpdateCustomerById(c echo.Context) error {
	log.Println("H [UpdateCustomerById]")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("H [UpdateCustomerById]: " + err.Error())

		response := map[string]string{
			"error": "Verifique o tipo do dado informado no parâmetro :id.",
		}

		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	customer := new(request.CustomerRequest)
	if err := c.Bind(customer); err != nil {
		log.Println("H [UpdateCustomerById]: " + err.Error())

		response := map[string]string{
			"error": "Os dados informados são inválidos!",
		}

		return c.JSON(http.StatusBadRequest, response)
	}

	isUpdated, apiError := ch.service.UpdateCustomerById(id, *customer)
	if apiError != nil {
		response := map[string]string{
			"error": apiError.Message,
		}

		return c.JSON(apiError.StatusCode, response)
	}

	response := map[string]bool{
		"isUpdated": isUpdated,
	}

	return c.JSON(http.StatusOK, response)
}

func (ch CustomerHandler) DeleteCustomerById(c echo.Context) error {
	log.Println("H [DeleteCustomerById]")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("H [DeleteCustomerById]: " + err.Error())
	}

	isDeleted, apiError := ch.service.DeleteCustomerById(id)
	if apiError != nil {
		response := map[string]string{
			"error": apiError.Message,
		}

		return c.JSON(apiError.StatusCode, response)
	}

	response := map[string]bool{
		"isDeleted": isDeleted,
	}

	return c.JSON(http.StatusOK, response)
}
