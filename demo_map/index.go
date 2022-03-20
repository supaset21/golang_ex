package main

import "fmt"

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(numbers["three"])

	var colors = make(map[string]string)

	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"

	fmt.Println(colors["green"])

	var courses = make(map[string]map[string]int)

	// courses["IOS"] = map[string]int{"price": 200}
	courses["IOS"] = make(map[string]int)
	courses["IOS"]["price"] = 100
	courses["IOS"]["code"] = 1000

	fmt.Println(courses["IOS"]["price"])

}
