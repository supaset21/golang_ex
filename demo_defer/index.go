package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		defer finish(i)
	}

	defer fmt.Println("Finish")

	for i := 0; i < 10; i++ {
		fmt.Println("", i)
	}
}

func finish(i int) {
	fmt.Println("Finish ", i)

}
