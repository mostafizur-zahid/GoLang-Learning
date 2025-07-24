package main

import "fmt"

func main() { //make is a built-in function
	s := make([]int, 3) //[0, 0, 0] len = 3, cap = 3

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s[0] = 10 //[10, 0, 0] len = 3, cap = 3

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
