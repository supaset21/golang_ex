package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func myInt() func() string {
	y := 0
	return func() string {
		y++
		return fmt.Sprintf("Y = %d", y)
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	newInts2 := myInt()
	fmt.Println(newInts2())
	fmt.Println(newInts2())
	fmt.Println(newInts2())

}
