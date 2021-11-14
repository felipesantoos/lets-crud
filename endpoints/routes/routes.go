package routes

import (
	"letscrud/endpoints/handlers"

	"github.com/labstack/echo/v4"
)

func GetEcho() *echo.Echo {
	e := echo.New()

	e.POST("/customer", handlers.CreateNewCustomer)
	e.GET("/customers", handlers.ReadAllCustomers)
	e.GET("/customer/:id", handlers.ReadCustomerById)
	e.PUT("/customer/:id", handlers.UpdateCustomerById)
	e.DELETE("/customer/:id", handlers.DeleteCustomerById)

	return e
}
