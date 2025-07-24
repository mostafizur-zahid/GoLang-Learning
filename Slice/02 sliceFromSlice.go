package main

import "fmt"

func main() {
	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}
	fmt.Println(arr)

	s := arr[1:4]  //slice from a aray (arr)
	fmt.Println(s) //slice hold 3 type(pointer, length, capacity)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	fmt.Println("Print S1 from s:")

	s1 := s[1:2] //slice (s1) from slice (s)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

}
