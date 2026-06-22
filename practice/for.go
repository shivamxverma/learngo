package main

import "fmt"

func forLoop() {
	i := 1

	for i <= 3 {
		i++
		fmt.Println(i)
	}

	for j := 0; j < 9 ; j ++ {
		fmt.Println("Perfect Loop in go", j)
	}

	if 8%2 == 0 {
		fmt.Println("Hello chalo bhai")
	}

	for i := range 3 {
		fmt.Println(i);
	}

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	}

	
}