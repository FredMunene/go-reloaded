package reloaded

import (
	"strings"
)

func punctuate(words []string) {
	chars := ".,!?:;"
	for index, word := range words {
		for _, punct := range chars {
			if len(word) > 0 && strings.HasPrefix(word, string(punct)) && index < len(words)-1 {
				words[index-1] += words[index][:1]
				words[index] = words[index][1:]
				// fmt.Println(words[index])
			}
			if len(word) > 0 && strings.HasPrefix(word, string(punct)) && strings.HasSuffix(word, string(punct)) && index == len(words)-1 {
				// words[index-1] += words[index][:1]
				words[index-1] += words[index]
				// fmt.Println(words[index])
				words = words[:index]
			}

			if len(word) > 0 && strings.HasPrefix(word, string(punct)) && index < len(words)-1 && strings.HasSuffix(word, string(punct)) {
				words[index-1] += words[index]
				// fmt.Println(words[ilsndex])
				words = append(words[:index], words[index+1:]...)
			}
		}
	}
	for index, word := range words {
		for _, punct := range chars {
			if len(word) > 0 && strings.HasPrefix(word, string(punct)) && index == len(words)-1 {
				words[index-1] += words[index]
				words = words[:index-5]
			}
		}
	}
}
