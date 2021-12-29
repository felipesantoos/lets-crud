package service_test

import (
	"letscrud/src/services/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatCPFRemovingPunctuation(t *testing.T) {
	returnedCPF1 := utils.FormatCPFRemovingPunctuation("878.723.850-07")
	expectedCPF1 := "87872385007"
	returnedCPF2 := utils.FormatCPFRemovingPunctuation("324.957.640-96")
	expectedCPF2 := "32495764096"
	returnedCPF3 := utils.FormatCPFRemovingPunctuation("373.922.270-09")
	expectedCPF3 := "37392227009"
	returnedCPF4 := utils.FormatCPFRemovingPunctuation("785.829.780-57")
	expectedCPF4 := "78582978057"
	returnedCPF5 := utils.FormatCPFRemovingPunctuation("922.391.710-73")
	expectedCPF5 := "92239171073"
	returnedCPF6 := utils.FormatCPFRemovingPunctuation("702.567.430-37")
	expectedCPF6 := "70256743037"
	returnedCPF7 := utils.FormatCPFRemovingPunctuation("873.821.960-38")
	expectedCPF7 := "87382196038"
	returnedCPF8 := utils.FormatCPFRemovingPunctuation("249.889.370-97")
	expectedCPF8 := "24988937097"
	returnedCPF9 := utils.FormatCPFRemovingPunctuation("020.135.550-75")
	expectedCPF9 := "02013555075"
	returnedCPF10 := utils.FormatCPFRemovingPunctuation("849.601.790-74")
	expectedCPF10 := "84960179074"

	assert.Equal(t, expectedCPF1, returnedCPF1)
	assert.Equal(t, expectedCPF2, returnedCPF2)
	assert.Equal(t, expectedCPF3, returnedCPF3)
	assert.Equal(t, expectedCPF4, returnedCPF4)
	assert.Equal(t, expectedCPF5, returnedCPF5)
	assert.Equal(t, expectedCPF6, returnedCPF6)
	assert.Equal(t, expectedCPF7, returnedCPF7)
	assert.Equal(t, expectedCPF8, returnedCPF8)
	assert.Equal(t, expectedCPF9, returnedCPF9)
	assert.Equal(t, expectedCPF10, returnedCPF10)
}

func TestFormatCPFAddingPunctuation(t *testing.T) {
	returnedCPF1 := utils.FormatCPFAddingPunctuation("87872385007")
	expectedCPF1 := "878.723.850-07"
	returnedCPF2 := utils.FormatCPFAddingPunctuation("32495764096")
	expectedCPF2 := "324.957.640-96"
	returnedCPF3 := utils.FormatCPFAddingPunctuation("37392227009")
	expectedCPF3 := "373.922.270-09"
	returnedCPF4 := utils.FormatCPFAddingPunctuation("78582978057")
	expectedCPF4 := "785.829.780-57"
	returnedCPF5 := utils.FormatCPFAddingPunctuation("92239171073")
	expectedCPF5 := "922.391.710-73"
	returnedCPF6 := utils.FormatCPFAddingPunctuation("70256743037")
	expectedCPF6 := "702.567.430-37"
	returnedCPF7 := utils.FormatCPFAddingPunctuation("87382196038")
	expectedCPF7 := "873.821.960-38"
	returnedCPF8 := utils.FormatCPFAddingPunctuation("24988937097")
	expectedCPF8 := "249.889.370-97"
	returnedCPF9 := utils.FormatCPFAddingPunctuation("02013555075")
	expectedCPF9 := "020.135.550-75"
	returnedCPF10 := utils.FormatCPFAddingPunctuation("84960179074")
	expectedCPF10 := "849.601.790-74"

	assert.Equal(t, expectedCPF1, returnedCPF1)
	assert.Equal(t, expectedCPF2, returnedCPF2)
	assert.Equal(t, expectedCPF3, returnedCPF3)
	assert.Equal(t, expectedCPF4, returnedCPF4)
	assert.Equal(t, expectedCPF5, returnedCPF5)
	assert.Equal(t, expectedCPF6, returnedCPF6)
	assert.Equal(t, expectedCPF7, returnedCPF7)
	assert.Equal(t, expectedCPF8, returnedCPF8)
	assert.Equal(t, expectedCPF9, returnedCPF9)
	assert.Equal(t, expectedCPF10, returnedCPF10)
}

func TestIsValidCPF(t *testing.T) {
	t.Run("TestValidCPF", func(t *testing.T) {
		isValidCPF1 := utils.IsValidCPF("22104245001")
		isValidCPF2 := utils.IsValidCPF("07177327037")
		isValidCPF3 := utils.IsValidCPF("28790175050")
		isValidCPF4 := utils.IsValidCPF("37797110018")
		isValidCPF5 := utils.IsValidCPF("85186381004")
		isValidCPF6 := utils.IsValidCPF("03339451079")
		isValidCPF7 := utils.IsValidCPF("15131687039")
		isValidCPF8 := utils.IsValidCPF("22233708024")
		isValidCPF9 := utils.IsValidCPF("71395182035")
		isValidCPF10 := utils.IsValidCPF("51835402097")

		assert.Equal(t, true, isValidCPF1)
		assert.Equal(t, true, isValidCPF2)
		assert.Equal(t, true, isValidCPF3)
		assert.Equal(t, true, isValidCPF4)
		assert.Equal(t, true, isValidCPF5)
		assert.Equal(t, true, isValidCPF6)
		assert.Equal(t, true, isValidCPF7)
		assert.Equal(t, true, isValidCPF8)
		assert.Equal(t, true, isValidCPF9)
		assert.Equal(t, true, isValidCPF10)
	})
	t.Run("TestInvalidCPF", func(t *testing.T) {
		isValidCPF1 := utils.IsValidCPF("15905630002")
		isValidCPF2 := utils.IsValidCPF("22854396013")
		isValidCPF3 := utils.IsValidCPF("85760369082")
		isValidCPF4 := utils.IsValidCPF("66311603089")
		isValidCPF5 := utils.IsValidCPF("13340235065")
		isValidCPF6 := utils.IsValidCPF("08293484097")
		isValidCPF7 := utils.IsValidCPF("65328398091")
		isValidCPF8 := utils.IsValidCPF("68853302053")
		isValidCPF9 := utils.IsValidCPF("16489494018")
		isValidCPF10 := utils.IsValidCPF("66749821056")

		assert.Equal(t, false, isValidCPF1)
		assert.Equal(t, false, isValidCPF2)
		assert.Equal(t, false, isValidCPF3)
		assert.Equal(t, false, isValidCPF4)
		assert.Equal(t, false, isValidCPF5)
		assert.Equal(t, false, isValidCPF6)
		assert.Equal(t, false, isValidCPF7)
		assert.Equal(t, false, isValidCPF8)
		assert.Equal(t, false, isValidCPF9)
		assert.Equal(t, false, isValidCPF10)
	})
	t.Run("TestCPFTooShort", func(t *testing.T) {
		isValidCPF1 := utils.IsValidCPF("4918237207")
		isValidCPF2 := utils.IsValidCPF("164832580")
		isValidCPF3 := utils.IsValidCPF("44668216")
		isValidCPF4 := utils.IsValidCPF("8754113")
		isValidCPF5 := utils.IsValidCPF("927116")
		isValidCPF6 := utils.IsValidCPF("60063")
		isValidCPF7 := utils.IsValidCPF("0644")
		isValidCPF8 := utils.IsValidCPF("801")
		isValidCPF9 := utils.IsValidCPF("96")
		isValidCPF10 := utils.IsValidCPF("8")
		isValidCPF11 := utils.IsValidCPF("")

		assert.Equal(t, false, isValidCPF1)
		assert.Equal(t, false, isValidCPF2)
		assert.Equal(t, false, isValidCPF3)
		assert.Equal(t, false, isValidCPF4)
		assert.Equal(t, false, isValidCPF5)
		assert.Equal(t, false, isValidCPF6)
		assert.Equal(t, false, isValidCPF7)
		assert.Equal(t, false, isValidCPF8)
		assert.Equal(t, false, isValidCPF9)
		assert.Equal(t, false, isValidCPF10)
		assert.Equal(t, false, isValidCPF11)
	})
}

func TestIsValidName(t *testing.T) {
	t.Run("TestValidName", func(t *testing.T) {
		isValidName1 := utils.IsValidName("Evan Vilar Marinho")
		isValidName2 := utils.IsValidName("Belchior Castanho Guerra")
		isValidName3 := utils.IsValidName("Dante Carmona Coimbra")
		isValidName4 := utils.IsValidName("Anastacia Ulhoa Filgueiras")
		isValidName5 := utils.IsValidName("Rosarinho Xavier Vilela")
		isValidName6 := utils.IsValidName("Isac Leite Bulhosa")
		isValidName7 := utils.IsValidName("Melanie Quina Murtinho")
		isValidName8 := utils.IsValidName("Mário Andrade Laureano")
		isValidName9 := utils.IsValidName("Luca Cabeça de Vaca Paz")
		isValidName10 := utils.IsValidName("Inês Brásio Cascais")

		assert.Equal(t, true, isValidName1)
		assert.Equal(t, true, isValidName2)
		assert.Equal(t, true, isValidName3)
		assert.Equal(t, true, isValidName4)
		assert.Equal(t, true, isValidName5)
		assert.Equal(t, true, isValidName6)
		assert.Equal(t, true, isValidName7)
		assert.Equal(t, true, isValidName8)
		assert.Equal(t, true, isValidName9)
		assert.Equal(t, true, isValidName10)
	})
	t.Run("TestNameWithSpecialCharacters", func(t *testing.T) {
		isValidName1 := utils.IsValidName("!Raquel Morão da Costa")
		isValidName2 := utils.IsValidName("@Máximo Orriça Vilariça")
		isValidName3 := utils.IsValidName("#José Sardo Gomide")
		isValidName4 := utils.IsValidName("$Ionara Veleda Santos")
		isValidName5 := utils.IsValidName("%Guransh Aguiar Jorge")
		isValidName6 := utils.IsValidName("¨Rebecca Júdice Pires")
		isValidName7 := utils.IsValidName("&Soraya Bettencourt Arruda")
		isValidName8 := utils.IsValidName("*Lisandro Carvalho Arruda")
		isValidName9 := utils.IsValidName("(Nara Borges Bastos")
		isValidName10 := utils.IsValidName(")Osvaldo Dutra Bicalho")
		isValidName11 := utils.IsValidName("-Lucas Casalinho Freitas")
		isValidName12 := utils.IsValidName("_Ohana Regueira Espírito Santo")
		isValidName13 := utils.IsValidName("=André Castilho Gadelha")
		isValidName14 := utils.IsValidName("+Aaron Ornelas Carmona")
		isValidName15 := utils.IsValidName("\\Lisandra Alcoforado CanedoFred Seixas Barreno")
		isValidName16 := utils.IsValidName("|Nadia Pontes Álvares")
		isValidName17 := utils.IsValidName(",Zara Carvalho Patrício")
		isValidName18 := utils.IsValidName("<Rayane Assis Milhães")
		isValidName19 := utils.IsValidName(">Yassna Maior Estrela")
		isValidName20 := utils.IsValidName(":Katie Mata Noite")
		isValidName21 := utils.IsValidName("?Virgílio Vides Garcia")

		assert.Equal(t, false, isValidName1)
		assert.Equal(t, false, isValidName2)
		assert.Equal(t, false, isValidName3)
		assert.Equal(t, false, isValidName4)
		assert.Equal(t, false, isValidName5)
		assert.Equal(t, false, isValidName6)
		assert.Equal(t, false, isValidName7)
		assert.Equal(t, false, isValidName8)
		assert.Equal(t, false, isValidName9)
		assert.Equal(t, false, isValidName10)
		assert.Equal(t, false, isValidName11)
		assert.Equal(t, false, isValidName12)
		assert.Equal(t, false, isValidName13)
		assert.Equal(t, false, isValidName14)
		assert.Equal(t, false, isValidName15)
		assert.Equal(t, false, isValidName16)
		assert.Equal(t, false, isValidName17)
		assert.Equal(t, false, isValidName18)
		assert.Equal(t, false, isValidName19)
		assert.Equal(t, false, isValidName20)
		assert.Equal(t, false, isValidName21)
	})
	t.Run("NameWithNumber", func(t *testing.T) {
		isValidName0 := utils.IsValidName("0Mélanie Gouveia Álvares")
		isValidName1 := utils.IsValidName("1Allana Faro Barreto")
		isValidName2 := utils.IsValidName("2Jil Rosa Caparica")
		isValidName3 := utils.IsValidName("3Yiyi Martinho Carneiro")
		isValidName4 := utils.IsValidName("4Oséias Amorim Carrasco")
		isValidName5 := utils.IsValidName("5Sérgio Amarante Cabreira")
		isValidName6 := utils.IsValidName("6Emília Rodovalho Figueiroa")
		isValidName7 := utils.IsValidName("7Djeyson Freire Dias")
		isValidName8 := utils.IsValidName("8Aliyha Salgado Canedo")
		isValidName9 := utils.IsValidName("9Ulisses Capucho Espírito Santo")

		assert.Equal(t, false, isValidName0)
		assert.Equal(t, false, isValidName1)
		assert.Equal(t, false, isValidName2)
		assert.Equal(t, false, isValidName3)
		assert.Equal(t, false, isValidName4)
		assert.Equal(t, false, isValidName5)
		assert.Equal(t, false, isValidName6)
		assert.Equal(t, false, isValidName7)
		assert.Equal(t, false, isValidName8)
		assert.Equal(t, false, isValidName9)
	})
}
