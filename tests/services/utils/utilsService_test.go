package service_test

import (
	"letscrud/services/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatCPFRemovingPunctuation(t *testing.T) {
	returnedCPF_1 := utils.FormatCPFRemovingPunctuation("878.723.850-07")
	expectedCPF_1 := "87872385007"
	returnedCPF_2 := utils.FormatCPFRemovingPunctuation("324.957.640-96")
	expectedCPF_2 := "32495764096"
	returnedCPF_3 := utils.FormatCPFRemovingPunctuation("373.922.270-09")
	expectedCPF_3 := "37392227009"
	returnedCPF_4 := utils.FormatCPFRemovingPunctuation("785.829.780-57")
	expectedCPF_4 := "78582978057"
	returnedCPF_5 := utils.FormatCPFRemovingPunctuation("922.391.710-73")
	expectedCPF_5 := "92239171073"
	returnedCPF_6 := utils.FormatCPFRemovingPunctuation("702.567.430-37")
	expectedCPF_6 := "70256743037"
	returnedCPF_7 := utils.FormatCPFRemovingPunctuation("873.821.960-38")
	expectedCPF_7 := "87382196038"
	returnedCPF_8 := utils.FormatCPFRemovingPunctuation("249.889.370-97")
	expectedCPF_8 := "24988937097"
	returnedCPF_9 := utils.FormatCPFRemovingPunctuation("020.135.550-75")
	expectedCPF_9 := "02013555075"
	returnedCPF_10 := utils.FormatCPFRemovingPunctuation("849.601.790-74")
	expectedCPF_10 := "84960179074"

	assert.Equal(t, expectedCPF_1, returnedCPF_1)
	assert.Equal(t, expectedCPF_2, returnedCPF_2)
	assert.Equal(t, expectedCPF_3, returnedCPF_3)
	assert.Equal(t, expectedCPF_4, returnedCPF_4)
	assert.Equal(t, expectedCPF_5, returnedCPF_5)
	assert.Equal(t, expectedCPF_6, returnedCPF_6)
	assert.Equal(t, expectedCPF_7, returnedCPF_7)
	assert.Equal(t, expectedCPF_8, returnedCPF_8)
	assert.Equal(t, expectedCPF_9, returnedCPF_9)
	assert.Equal(t, expectedCPF_10, returnedCPF_10)
}

func TestFormatCPFAddingPunctuation(t *testing.T) {
	returnedCPF_1 := utils.FormatCPFAddingPunctuation("87872385007")
	expectedCPF_1 := "878.723.850-07"
	returnedCPF_2 := utils.FormatCPFAddingPunctuation("32495764096")
	expectedCPF_2 := "324.957.640-96"
	returnedCPF_3 := utils.FormatCPFAddingPunctuation("37392227009")
	expectedCPF_3 := "373.922.270-09"
	returnedCPF_4 := utils.FormatCPFAddingPunctuation("78582978057")
	expectedCPF_4 := "785.829.780-57"
	returnedCPF_5 := utils.FormatCPFAddingPunctuation("92239171073")
	expectedCPF_5 := "922.391.710-73"
	returnedCPF_6 := utils.FormatCPFAddingPunctuation("70256743037")
	expectedCPF_6 := "702.567.430-37"
	returnedCPF_7 := utils.FormatCPFAddingPunctuation("87382196038")
	expectedCPF_7 := "873.821.960-38"
	returnedCPF_8 := utils.FormatCPFAddingPunctuation("24988937097")
	expectedCPF_8 := "249.889.370-97"
	returnedCPF_9 := utils.FormatCPFAddingPunctuation("02013555075")
	expectedCPF_9 := "020.135.550-75"
	returnedCPF_10 := utils.FormatCPFAddingPunctuation("84960179074")
	expectedCPF_10 := "849.601.790-74"

	assert.Equal(t, expectedCPF_1, returnedCPF_1)
	assert.Equal(t, expectedCPF_2, returnedCPF_2)
	assert.Equal(t, expectedCPF_3, returnedCPF_3)
	assert.Equal(t, expectedCPF_4, returnedCPF_4)
	assert.Equal(t, expectedCPF_5, returnedCPF_5)
	assert.Equal(t, expectedCPF_6, returnedCPF_6)
	assert.Equal(t, expectedCPF_7, returnedCPF_7)
	assert.Equal(t, expectedCPF_8, returnedCPF_8)
	assert.Equal(t, expectedCPF_9, returnedCPF_9)
	assert.Equal(t, expectedCPF_10, returnedCPF_10)
}
