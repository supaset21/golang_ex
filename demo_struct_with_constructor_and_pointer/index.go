package main

import (
	"hello/employee"
)

func main() {
	e := employee.New("Ice", "Ice", 30, 20)

	e.LeavesRemaining()

}
