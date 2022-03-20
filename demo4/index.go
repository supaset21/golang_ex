package main

import "fmt"

func main() {
	fn1()
	fn2("Hello")
	fn3("Ice", 3)
	const a, b = 2, 3
	fmt.Printf("%d+%d=%d\n", a, b, sum(a, b))
	x, y := swap(a, b)

	fmt.Printf("%d,%d\n", x, y)

	_x, _y, title := swapV2(a, b)

	fmt.Printf("%d,%d = %s", _x, _y, title)

}

func fn1() {
	fmt.Println("Display")
}

func fn2(msg string) {
	fmt.Println(msg)
}

func fn3(title string, version int) {
	fmt.Println(title)
	fmt.Println(version)
}

func sum(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func swapV2(a int, b int) (x, y int, title string) {
	y = a
	x = b
	title = "Result"
	return
}
