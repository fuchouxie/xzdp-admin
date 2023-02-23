package str

import (
	"strconv"
	"strings"
)

func IntStrToArray(str string, sep string) []int64 {
	var arr []int64
	strArr := strings.Split(str, sep)
	for _, v := range strArr {
		formatInt, _ := strconv.ParseInt(v, 10, 64)
		arr = append(arr, formatInt)
	}
	return arr
}
