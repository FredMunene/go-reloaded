package reloaded

import (
	"fmt"
	"strconv"
)

func hexToDec(str []string, index int) []string {
	num, _ := strconv.ParseInt(str[index-1], 16, 64)
	str[index-1] = fmt.Sprint(num)
	return append(str[:index], str[index+1:]...)
}
