package main

import "fmt"

// import "test.go"
type human struct {
	name string
	age  int
}

func (h human) printInfo() {
	fmt.Println(h.name, h.age)
}

func main() {

	// var num int
	// num = 7

	var message string
	message = "Expecto Patronum!"

	// var num int
	// num := 7
	var num int = 7
	var platform float64 = 9.75
	// num := "Ice"
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(message)

	fmt.Println(num, platform)

	// อาร์เรย์ของข้อความจำนวนสามช่อง
	// var names [3]string

	// names[0] = "Somchai"
	// names[1] = "Somsree"
	// names[2] = "Somset"

	// fmt.Println(names)

	// var names []string

	// // เพิ่ม element เข้าไป
	// names = append(names, "Somchai")
	// names = append(names, "Somsree")
	// names = append(names, "Somset")

	// fmt.Println((names))

	names := []string{"Somchai", "Somsree", "Somset"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	// printFullName("111", "ssss")

	message = getMessage()
	fmt.Println(message)

	somchai := human{name: "Somchai", age: 23}
	// somsree := human{name: "Somsree", age: 32}

	fmt.Println((somchai.age))

}

func printFullName(firstName string, lastName string) {
	fmt.Println(firstName + " " + lastName)
}

func getMessage() string {
	return "IIII"

}
