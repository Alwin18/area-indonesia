package utils

import "fmt"

func StringToInt(s string) int64 {
	var i int64
	fmt.Sscanf(s, "%d", &i)
	return i
}
