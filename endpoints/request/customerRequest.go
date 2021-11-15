package request

type CustomerRequest struct {
	CPF       string `json:"cpf"`
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
}
