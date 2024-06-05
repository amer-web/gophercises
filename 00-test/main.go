package main

import "fmt"

func main() {
	var orders []orderEm
	emp := emp{name: "amer", id: 60}
	orders = append(orders, orderEm{emp, 10})
	fmt.Println(orders[0].getName())
}

type emp struct {
	name string
	id   int
}

type orderEm struct {
	emp
	index int
}

func (o *orderEm) getName() string {
	return o.name
}
