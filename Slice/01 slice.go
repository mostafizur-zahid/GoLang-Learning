package main

import "fmt"

func main() {
	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}
	fmt.Println(arr)

	s := arr[1:4]  // slice from a array
	fmt.Println(s) //slice hold 3 type(pointer, length, capacity)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
