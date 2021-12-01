package tests

import (
	"letscrud/endpoints/dto/request"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func RunGenericPostRequestTest(
	t *testing.T,
	target string,
	contextParams *ContextParams,
	bodyJSON string,
	handlerFunc func(ctx echo.Context) error,
	expectedStatusCode int,
	expectedJSON string,
) {
	ast := assert.New(t)

	request := httptest.NewRequest(
		http.MethodPost,
		target,
		strings.NewReader(bodyJSON),
	)

	e := echo.New()
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := e.NewContext(request, recorder)

	if contextParams != nil {
		c.SetParamNames(contextParams.Keys...)
		c.SetParamValues(contextParams.Values...)
	}

	if ast.NoError(handlerFunc(c)) {
		ast.Equal(expectedStatusCode, recorder.Code)
		if expectedJSON != "" {
			ast.JSONEq(expectedJSON, recorder.Body.String())
		}
	}
}

func GetValidCustomerRequestForHandler() request.CustomerRequest {
	return request.CustomerRequest{
		CPF:       "098.629.560-46",
		Name:      "Lyara Caparica Onofre",
		BirthDate: "2001-01-01",
	}
}

func GetCreateNewCustomerResponse() map[string]interface{} {
	return map[string]interface{}{
		"id": 1,
	}
}
