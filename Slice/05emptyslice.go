package main

import "fmt"

func main() {
	var s []int //empty silce or nill slice [] len = 0, cap = 0
	fmt.Println(s)

	s = append(s, 1) //[1] len = 1, cap = 1
	fmt.Println(s)
}
