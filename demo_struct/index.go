package main

import "fmt"

type product struct {
	name  string
	price int
	stock int
}

func (p product) clear() product {
	p.price = 0
	p.stock = 0
	return p
}

func main() {
	var p1 product
	p1.name = "Luffy"
	p1.price = 20
	p1.stock = 609

	show(p1)
	update(&p1)
	show(p1)

	// Demo()
	// p1 = p1.clear()
	p1 = p1.setDiscount(1)

	show(p1)
}

func show(p product) {
	fmt.Println(p)
}

func (p product) setDiscount(d int) product {
	p.price = p.price - d
	return p
}

func update(p *product) {
	fmt.Println(*p)
	p.price = p.price + 100
	p.stock = p.stock - 10
}
