package main

import "fmt"

func main() {
	msg := "some"
	var msgPointer *string = &msg
	fmt.Println(msg)
	fmt.Println(*msgPointer)

	changeMsg(&msg, "new m")
	fmt.Println(msg)

	changeMsg(msgPointer, "new m2")
	fmt.Println(msg)
}

func changeMsg(aPointer *string, newMsg string) {
	*aPointer = newMsg
}
