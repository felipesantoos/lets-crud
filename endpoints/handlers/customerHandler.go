package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateNewCustomer(c echo.Context) error {
	return c.String(http.StatusOK, "Create a new customer")
}

func ReadAllCustomers(c echo.Context) error {
	return c.String(http.StatusOK, "Read all customers")
}

func ReadCustomerById(c echo.Context) error {
	return c.String(http.StatusOK, "Read a specific customer")
}

func UpdateCustomerById(c echo.Context) error {
	return c.String(http.StatusOK, "Update a specific customer")
}

func DeleteCustomerById(c echo.Context) error {
	return c.String(http.StatusOK, "Delete a specific customer")
}
