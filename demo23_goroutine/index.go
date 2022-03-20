package main

import (
	"fmt"
	"time"
)

const count = 100

func run1() {
	for i := 0; i < count; i++ {
		fmt.Printf("Run1 %d \n", i)
	}

}
func run2() {
	for i := 0; i < count; i++ {
		fmt.Printf("Run2 %d \n", i)
	}

}

func main() {
	go run1()
	go run2()

	time.Sleep(5 * time.Second)
}
