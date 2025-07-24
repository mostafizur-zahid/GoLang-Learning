package main

import "fmt"

func main() {
	s := []int{1, 3, 5}
	fmt.Println("Slice s are:", s)
	fmt.Println("Length of slice s:", len(s))
	fmt.Println("Total capacity:", cap(s))
}
