package reloaded

import "strings"

func aToAn(words []string, index int) []string {
	nextword := words[index+1]
	if vowelIsPresent(nextword) {
		words[index] = "an"
		words = append(words[:index+1], words[index+1:]...)
	}
	return words
}

func capAToAn(words []string, index int) []string {
	nextword := words[index+1]
	if vowelIsPresent(nextword) {
		words[index] = "An"
		words = append(words[:index+1], words[index+1:]...)
	}
	return words
}

func vowelIsPresent(word string) bool {
	chars := "aeiouhAEIOUH"
	if len(word) > 0 && strings.ContainsRune(chars, rune(word[0])) {
		return true
	}
	return false
}
