package main

import "fmt"

func Solve1() {
	a := []int{2,3,4}
	sum := 0

	for _,val := range a {
		sum += val
	} 

	kvs := map[string]string{"a": "apple", "b": "banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k , v)
	}

	for k := range kvs {
		fmt.Println("Key:", k)
	}

	fmt.Println(sum);
}