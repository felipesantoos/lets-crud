package response

type CustomerResponse struct {
	Id        int64  `json:"id"`
	CPF       string `json:"cpf"`
	Name      string `json:"name"`
	Birthdate string `json:"birthDate,omitempty"`
}
