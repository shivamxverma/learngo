package main

import "fmt"

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b,c int = 1, 2
	
	var res int = b*c

	var d = true

	fmt.Println(d)

	var e int

	fmt.Println(e)

	f := "apple"

	const res1 string = "constant"

	fmt.Println(res1)

	fmt.Println(f)

	fmt.Println(res)
}