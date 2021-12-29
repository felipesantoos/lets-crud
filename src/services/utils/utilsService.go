package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var cpfFirstDigitTable = []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
var cpfSecondDigitTable = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

func FormatCPFRemovingPunctuation(cpf string) string {
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)

	return cpf
}

func FormatCPFAddingPunctuation(cpf string) string {
	cpf = cpf[:3] + "." + cpf[3:]
	cpf = cpf[:7] + "." + cpf[7:]
	cpf = cpf[:11] + "-" + cpf[11:]

	return cpf
}

func IsValidCPF(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	firstPart := cpf[0:9]
	sum := sumDigit(firstPart, cpfFirstDigitTable)

	r1 := sum % 11
	d1 := 0

	if r1 >= 2 {
		d1 = 11 - r1
	}

	secondPart := firstPart + strconv.Itoa(d1)

	dsum := sumDigit(secondPart, cpfSecondDigitTable)

	r2 := dsum % 11
	d2 := 0

	if r2 >= 2 {
		d2 = 11 - r2
	}

	finalPart := fmt.Sprintf("%s%d%d", firstPart, d1, d2)

	return finalPart == cpf
}

func sumDigit(s string, table []int) int {
	if len(s) != len(table) {
		return 0
	}

	sum := 0

	for i, v := range table {
		c := string(s[i])
		d, err := strconv.Atoi(c)
		if err == nil {
			sum += v * d
		}
	}

	return sum
}

func IsValidName(name string) bool {
	for _, l := range name {
		if !unicode.IsLetter(l) && l != ' ' {
			return false
		}
	}

	return true
}
