package main

import (
	"fmt"
)

func Sqrt(x float64) (z float64) {
	z = 1.
	temp := 0.
	for i := 0; i < 10; i++ {
		temp = (z*z - x) / (2 * z)
		if temp > 1e-8 || temp < -1e-8 {
			z -= temp
		} else {
			return
		}
	}
	return
}

func main() {
	fmt.Println(Sqrt(2))
}
