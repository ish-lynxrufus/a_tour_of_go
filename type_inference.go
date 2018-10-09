package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)

	i := 42
	f := 3.142
	g := 0.867 + 0.5i
	s := "hoge"

	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
	fmt.Printf("s is of type %T\n", s)
}
