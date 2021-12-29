package errs

import (
	"strings"
)

func GetCustomerError(err error) *ApiError {
	errMsg := err.Error()

	if strings.Contains(errMsg, "Duplicate entry") && strings.Contains(errMsg, "cpf") {
		return NewBadRequestError("O CPF informado já foi cadastrado no banco de dados!")
	} else if strings.Contains(errMsg, "Data too long for column") && strings.Contains(errMsg, "cpf") {
		return NewBadRequestError("O CPF informado é muito longo!")
	} else if strings.Contains(errMsg, "Incorrect date value") && strings.Contains(errMsg, "birthDate") {
		return NewBadRequestError("A data de nascimento informada está mal formadata! Formato correto: YYYY-MM-DD.")
	}

	return NewUnexpectedError("Ocorreu um erro inesperado: " + errMsg)
}
