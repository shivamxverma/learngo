package main

import (
	"fmt"
)

func SliceExample() {
	var x []string

	if x == nil || len(x) == 0 {

	}

	x = make([]string, 3)

	arr := make([][]int, 3)

	for i := range 3 {
		for j := range 3 {
			arr[i][j] = i+j;
		}
	}


	fmt.Println("Hello i am here")
}