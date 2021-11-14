package errs

import (
	"strings"
)

func GetNotificationErrs(err error) *ApiErros {
	errMsg := err.Error()

	if strings.Contains(errMsg, "Duplicate entry") && strings.Contains(errMsg, "cpf") {
		return &ApiErros{
			StatusCode: 400,
			Message:    "",
		}
	}

	return nil
}
