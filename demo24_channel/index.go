package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println("Step1")
	fmt.Println(<-ch)

	ch <- 2
	fmt.Println("Step2")
	fmt.Println(<-ch)

	time.Sleep(1 * time.Second)
	// s := []int{7, 2, 8, -9, 4, 0}

	// c := make(chan int)
	// go sum(s[:len(s)/2], c)
	// go sum(s[len(s)/2:], c)
	// x, y := <-c, <-c // receive from c

	// fmt.Println(x, y, x+y)
}
