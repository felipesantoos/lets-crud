package errs

import (
	"strings"
)

func GetCustomerError(err error) *ApiError {
	errMsg := err.Error()

	if strings.Contains(errMsg, "Duplicate entry") && strings.Contains(errMsg, "cpf") {
		return NewBadRequestError("O CPF informado já foi cadastrado no banco de dados!")
	}

	return NewUnexpectedError("Ocorreu um erro inesperado: " + errMsg)
}
