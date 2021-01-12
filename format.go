package main

import (
	"strconv"
	"strings"
)

// stringToInt. Little function that receives string values and return integer
func stringToInt(text string) int {
	integer, _ := strconv.Atoi(text)
	return integer
}

// makeFinalFormat. Returns the format that de configuration needs to have to the final file.
func makeFinalFormat(allConfig []string) string {
	var configText string

	// This for convert the array of values in one single string line
	for i := 0; i < len(allConfig); i++ {
		configText += allConfig[i]
	}

	// It returns the string line but adding the "," between the parentesis
	return strings.Replace(configText, ")(", "), (", -1)
}
