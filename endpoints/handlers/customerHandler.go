package handlers

import (
	"letscrud/endpoints/request"
	"letscrud/services"
	"letscrud/services/interfaces"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	service interfaces.ICustomerService
}

func NewCustomerHandler() *CustomerHandler {
	service := services.NewCustomerService()

	return &CustomerHandler{service: service}
}

func (ch CustomerHandler) CreateNewCustomer(c echo.Context) error {
	log.Println("Handler: CreateNewCustomer")

	customer := new(request.CustomerRequest)
	if err := c.Bind(customer); err != nil {
		log.Println("Handler (CreateNewCustomer): " + err.Error())
	}

	lastInsertId, apiErr := ch.service.CreateNewCustomer(*customer)

	if apiErr != nil {
		response := map[string]interface{}{
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
	log.Println("Handler: ReadAllCustomers")

	customerList, apiErr := ch.service.ReadAllCustomers()
	if apiErr != nil {
		response := map[string]interface{}{
			"error": apiErr.Message,
		}

		return c.JSON(apiErr.StatusCode, response)
	}

	return c.JSON(http.StatusOK, customerList)
}

func (ch CustomerHandler) ReadCustomerById(c echo.Context) error {
	log.Println("Handler: ReadCustomerById")

	ch.service.ReadCustomerById()

	return c.String(http.StatusOK, "Read a specific customer")
}

func (ch CustomerHandler) UpdateCustomerById(c echo.Context) error {
	log.Println("Handler: UpdateCustomerById")

	ch.service.UpdateCustomerById()

	return c.String(http.StatusOK, "Update a specific customer")
}

func (ch CustomerHandler) DeleteCustomerById(c echo.Context) error {
	log.Println("Handler: DeleteCustomerById")

	ch.service.DeleteCustomerById()

	return c.String(http.StatusOK, "Delete a specific customer")
}
