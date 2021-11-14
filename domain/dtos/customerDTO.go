package dtos

type CustomerDTO struct {
	CPF       string `json:"cpf"`
	Name      string `json:"name"`
	BirthDate string `json:"birthDate"`
}
