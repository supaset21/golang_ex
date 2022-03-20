package main

import "fmt"

func main() {
	// var numbers []int
	// numbers = append(numbers, 1)
	// numbers = append(numbers, 2)
	// showSlice(numbers)

	// var numbers2 = make([]int, 3, 5)
	// numbers2 = append(numbers2, 1)
	// numbers2 = append(numbers2, 2)
	// numbers2 = append(numbers2, 3)

	// showSlice(numbers2)
	// main2()
}

func main2() {
	var numbers = []int{1, 2, 3, 4, 5, 6}
	showSlice(numbers)

	// leading remove
	// numbers = numbers[1:len(numbers)]
	// showSlice(numbers)

	// trailing remove
	// numbers = numbers[0 : len(numbers)-1]
	// showSlice(numbers)

	// numbers = removeIndex(numbers, 2)
	// showSlice(numbers)

}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// func removeIndex(s []int, index int) []int {
// 	return append(s[:index], s[index+1:]...)
// }
