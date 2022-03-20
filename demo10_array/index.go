package main

import "fmt"

func main() {
	var array1 []int = []int{1, 2, 3, 4}
	var array2 = []int{1, 2, 3, 4}
	var array3 [3]string

	fmt.Println(array1)
	fmt.Println(array2)

	for _, item := range array1 {
		fmt.Print(item, ",")
	}

	array3[0] = "Ice"

	for _, item := range array3 {
		fmt.Print(item, ",")
	}

}
