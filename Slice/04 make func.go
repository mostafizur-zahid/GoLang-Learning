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

	ss := make([]int, 3, 5) //[0, 0, 0] len = 3, cap = 5
	ss[0] = 10              //[10, 0, 0] len = 3, cap = 5
	ss[2] = 30              //[10, 0, 30] len = 3, cap = 5

	fmt.Println(ss)
	fmt.Println(len(ss))
	fmt.Println(cap(ss))
}
