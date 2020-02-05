package main

import "fmt"

func main() {
	fmt.Println(myAtoi("    0000000000000   "))
}

func myAtoi(str string) int {
	const (
		MININT32 = -2147483648
		MAXINT32 = 2147483647
	)
	var num int
	isPositive := true
	// length check
	if len(str) == 0 {
		return 0
	}
	// trim prefix space
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			str = str[i:]
			break
		}
	}
	// 符号和数字检查
	switch {
	case str[0] == '+':
		str = str[1:]
	case str[0] == '-':
		isPositive = false
		str = str[1:]
	case str[0] > '9' || str[0] < '0': // 无有效数字
		return 0
	}
	// number string catch
	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			str = str[:i]
		}
	}
	// trim prefix zero
	for len(str) > 0 && str[0] == '0' {
		str = str[1:]
	}
	// parse
	numLen := len(str)
	// int64 boundary
	if numLen > 10 {
		if isPositive {
			return MAXINT32
		}
		return MININT32
	}
	for i := 0; i < numLen; i++ {
		power := 1
		for j := 0; j < numLen-i-1; j++ {
			power *= 10
		}
		num += power * (int(str[i]) - 48)
	}
	if !isPositive {
		num = -num
	}
	// int32 boundary
	switch {
	case num > MAXINT32:
		return MAXINT32
	case num < MININT32:
		return MININT32
	}
	return num
}
