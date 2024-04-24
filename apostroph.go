package reloaded

func apostrophe(str []string, index int) []string {
	if index < len(str)-1 {
		str[index+1] = str[index] + str[index+1]
		return append(str[:index], str[index+1:]...)
	}
	if index == len(str)-1 {
		str[index-1] = str[index-1] + str[index]
		return str[:index]
	}
	return str
}
