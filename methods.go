package main

import (
	"fmt"
	"math"
)

type Vertext struct {
	X, Y float64
}

func Abs(v Vertext) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertext{3, 4}
	fmt.Println(Abs(v))
}
