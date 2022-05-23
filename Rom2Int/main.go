package main

import "fmt"

func switcher(target, case0, case1 byte, result0, result1 int) int {
	out := 0
	switch target {
	case case0:
		out = result0
	case case1:
		out = result1
	}
	return out
}

func romanToInt(s string) int {
	i := 0
	counter := 0
	temp := 0
	length := len(s)
	s += "0"
	for {
		temp = 0
		switch s[i] {
		case 'I':
			{
				temp = switcher(s[i+1], 'V', 'X', 4, 9)
				if temp != 0 {
					i++
				} else {
					temp = 1
				}
			}
		case 'V':
			temp = 5
		case 'X':
			{
				temp = switcher(s[i+1], 'L', 'C', 40, 90)
				if temp != 0 {
					i++
				} else {
					temp = 10
				}
			}
		case 'L':
			temp = 50
		case 'C':
			{
				temp = switcher(s[i+1], 'D', 'M', 400, 900)
				if temp != 0 {
					i++
				} else {
					temp = 100
				}
			}
		case 'D':
			temp = 500
		case 'M':
			temp = 1000
		}
		counter += temp
		i++
		if i > length {
			break
		}
	}
	return counter
}

func main() {
	var s_Rom string
	for {
		fmt.Println("请输入罗马字符串:")
		fmt.Scanln(&s_Rom)
		fmt.Println(romanToInt(s_Rom))
		fmt.Println("")
	}
}
