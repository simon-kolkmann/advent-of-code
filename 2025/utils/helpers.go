package utils

import "strconv"

func UnsafeAtoi(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func Abs(i int) int {
	if i >= 0 {
		return i
	}

	return -i
}
