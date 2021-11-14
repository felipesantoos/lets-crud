package handlers

import (
	"letscrud/domain/dtos"
	"letscrud/services"
	"letscrud/services/interfaces"
	"log"
	"net/http"
	"strconv"

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

	customer := new(dtos.CustomerDTO)
	if err := c.Bind(customer); err != nil {
		log.Println("Handler (CreateNewCustomer): " + err.Error())
	}

	lastInsertId := ch.service.CreateNewCustomer(*customer)

	return c.String(http.StatusOK, strconv.FormatInt(lastInsertId, 10))
}

func (ch CustomerHandler) ReadAllCustomers(c echo.Context) error {
	log.Println("Handler: ReadAllCustomers")

	ch.service.ReadAllCustomers()

	return c.String(http.StatusOK, "Read all customers")
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
