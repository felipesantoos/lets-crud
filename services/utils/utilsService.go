package utils

import "strings"

func FormatCPF_RemovePunctuation(cpf string) string {
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)

	return cpf
}

func FormatCPF_AddPunctuation(cpf string) string {
	cpf = cpf[:3] + "." + cpf[3:]
	cpf = cpf[:7] + "." + cpf[7:]
	cpf = cpf[:11] + "-" + cpf[11:]

	return cpf
}
