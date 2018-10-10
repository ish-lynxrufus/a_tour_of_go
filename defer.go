package main

import "fmt"

func main() {
	defer fmt.Print("World!\n")

	fmt.Print("Hello ")
}
