package reloaded

import "strings"

func ProcessText(content string) string {
	words := strings.Fields(content)
	for index, word := range words {
		switch word {
		case "(hex)":
			words = hexToDec(words, index)
		case "(bin)":
			words = binToDec(words, index)
		case "(cap)":
			words = capital(words, index)
		case "(low)":
			words = toLow(words, index)
		case "(up)":
			words = toUp(words, index)
		case "'":
			words = apostrophe(words, index)
		case "(cap,":
			words = capitalSpecial(words, index)
		case "(low,":
			words = toLowSpecial(words, index)
		case "(up,":
			words = toUpSpecial(words, index)
		case "a":
			words = aToAn(words, index)
		case "A":
			words = capAToAn(words, index)
		}
	}
	punctuate(words)

	return strings.Join(words, " ")
}
