package reloaded

func charIsPresent(char string, s string) bool {
	for index, value := range s {
		if value == rune(char[0]) && index <= 1 {
			return true
		}
	}
	return false
}
