package routes

import (
	"letscrud/src/api/endpoints/handlers"

	"github.com/labstack/echo/v4"
)

func GetEcho() *echo.Echo {
	e := echo.New()

	ch := handlers.NewCustomerHandler()

	e.POST("/customer", ch.CreateNewCustomer)
	e.GET("/customers", ch.ReadAllCustomers)
	e.GET("/customer/:id", ch.ReadCustomerById)
	e.PUT("/customer/:id", ch.UpdateCustomerById)
	e.DELETE("/customer/:id", ch.DeleteCustomerById)

	return e
}
