package reloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func capitalSpecial(words []string, index int) []string {
	num := words[index+1][0]
	start, err := strconv.Atoi(string(num))
	if err != nil {
		fmt.Println("Error :", err)
	}
	for i := index - start; i < index; i++ {
		words[i] = capitalize(words[i])
	}
	return append(words[:index], words[index+2:]...)
}

func toLowSpecial(words []string, index int) []string {
	num := words[index+1][0]
	start, err := strconv.Atoi(string(num))
	if err != nil {
		fmt.Println("Error :", err)
	}
	for i := index - start; i < index; i++ {
		// prev := index - i
		words[i] = strings.ToLower(words[i])
	}
	return append(words[:index], words[index+2:]...)
}

func toUpSpecial(words []string, index int) []string {
	num := words[index+1][0]
	start, err := strconv.Atoi(string(num))
	if err != nil {
		fmt.Println("Error :", err)
	}
	for i := index - start; i < index; i++ {
		// prev := index - i
		words[i] = strings.ToUpper(words[i])
	}
	return append(words[:index], words[index+2:]...)
}
