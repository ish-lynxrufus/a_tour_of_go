package main

import (
	"fmt"
	"reflect"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var ti I
	fmt.Println(ti, reflect.TypeOf(ti))

	var i I = T{"hello"}
	fmt.Println(i, reflect.TypeOf(i))
	i.M()
}
