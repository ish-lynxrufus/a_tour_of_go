package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*y + y*y))
	var z uint = uint(f)
	// var f float64 = x*y + y*y
	// var z uint = f
	fmt.Println(x, y, z)
}
