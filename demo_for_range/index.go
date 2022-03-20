package main

import "fmt"

func main() {
	arr := []string{"Ice", "Ice2", "Ice3"}

	for index, item := range arr {
		fmt.Printf("%d. %s\n", index, item)
	}

	for _, item := range arr {
		fmt.Printf("%s.\n", item)
	}
}
