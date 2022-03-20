package main

import (
	"fmt"
	"time"
)

func routine1(c chan int, payload int) {
	c <- payload
	// fmt.Println("Step1")
	// c <- payload
	// fmt.Println("Step2")
	// c <- payload
	// fmt.Println("Step3")
}

func main() {
	ch := make(chan int, 5)
	go routine1(ch, 1)
	go routine1(ch, 2)
	go routine1(ch, 3)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}
