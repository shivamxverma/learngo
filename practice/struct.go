package main

import (
	"fmt"
)

// Defination of the struct
type person struct {
	name string
	age int
}

// How to Pass values to a function and return values from function
func newPerson(name string, age int) *person {
	p := person{name, age}
	return &p
}

type dog struct {
    name   string
    isGood bool
}

// How to enter data in struct type
func StructFunc() {
	fmt.Println(person{"shivam", 24})
	fmt.Println(person{"mandu", 6})

	p := person{"Moneky", 11}

	fmt.Println(p.name)
	fmt.Println(p.age)

	sp := &p // line 1
	sp1 := p // line 2

	sp.age = 12;
	fmt.Println(sp)
	fmt.Println(p)

	sp1.age = 12

	arr := []int{1,2,3}

	fmt.Println(arr);
}