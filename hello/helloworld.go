package main

import "fmt"

var j int32 = 5

func main() {
	s := "hello"
	for i := range s {
		i++
		fmt.Println(i)
	}

}
