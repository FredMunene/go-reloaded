package reloaded

import "strings"

func capital(words []string, index int) []string {
	words[index-1] = capitalize(words[index-1])
	return append(words[:index], words[index+1:]...)
}

func toLow(words []string, index int) []string {
	words[index-1] = strings.ToLower(words[index-1])
	return append(words[:index], words[index+1:]...)
}

func toUp(words []string, index int) []string {
	words[index-1] = strings.ToUpper(words[index-1])
	return append(words[:index], words[index+1:]...)
}

func capitalize(word string) string {
	firstchar := word[0]
	if firstchar >= 'a' && firstchar <= 'z' {
		return string(firstchar-32) + word[1:]
	}
	return word
}
