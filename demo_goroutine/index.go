package main

import (
	"fmt"
	"time"
)

func run1() {

	for i := 0; i < 100; i++ {
		// fmt.Println("Run1...")
		fmt.Printf("Run1 %d \n", i)
	}

}
func run2() {
	for i := 0; i < 100; i++ {
		fmt.Printf("Run2 %d \n", i)
		// fmt.Println("Run2...")
	}

}

func main() {
	go run1()
	go run2()

	time.Sleep(5 * time.Second)
}
