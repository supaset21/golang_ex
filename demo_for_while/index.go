package main

import "fmt"

func main() {
	// fnFor()
	// fnWhile()
	// fnWhileBreak()

	someValue := 10

	if someValue == 10 {
		fmt.Println("== 10")
	} else {
		fmt.Println("!= 10")
	}

	if someValue > 10 || someValue < 20 {
		fmt.Println("Hello")
	} else {
		fmt.Println("Hello2")
	}

	if result := doSomthing(); result == "ok" {
		fmt.Println("Ice")
	}

	fnSwitchCase()
}

func doSomthing() string {
	return "ok"
}

func fnFor() {
	for index := 0; index < 10; index++ {
		fmt.Printf("Index %d\n", index)
	}
}

func fnWhile() {
	index := 0
	for index < 5 {
		fmt.Printf("Index %d\n", index)
		index++
	}
}

func fnWhileBreak() {
	index := 0
	for true {
		if index > 5 {
			break
		}
		fmt.Printf("Index %d\n", index)
		index++
	}
}

func fnSwitchCase() {
	index := 1
	switch index {
	case 0:
		fmt.Println("JKKK")
	case 1:
		fmt.Println("Ice")
	default:
		fmt.Println("IIII")
	}
}
