package main

import "fmt"

func main() {
	var s []int
	printSlice(s) // 0

	s = append(s, 0) // 1
	printSlice(s)

	s = append(s, 1, 2) // 4
	printSlice(s)

	s = append(s, 2, 3, 4) // 8
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
