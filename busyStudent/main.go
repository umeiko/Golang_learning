package main

import "fmt"

func busyStudent(startTime []int, endTime []int, queryTime int) (result int) {
	result = 0
	for i := range startTime {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			result++
		}
	}
	return
}

func main() {
	fmt.Print(busyStudent([]int{1, 2, 3}, []int{3, 2, 7}, 4))
}
