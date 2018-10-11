package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // 7, 2, 8 = 17
	go sum(s[len(s)/2:], c) // -9, 4, 0 = -5
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
