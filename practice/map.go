package main

import (
	"fmt"
)

func Map() {
	m := make(map[string]int)
	mi := make(map[int]int)

	m["k1"] = 7
	m["k2"] = 13
	mi[1] = 15


	delete(m, "k2")
	clear(m)

	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:" , m2);
}