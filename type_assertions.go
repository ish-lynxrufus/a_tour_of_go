package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	str, ok := i.(string)
	fmt.Println(str, ok)

	flo, ok := i.(float64)
	fmt.Println(flo, ok)

	f := i.(float64)
	fmt.Println(f)
}
