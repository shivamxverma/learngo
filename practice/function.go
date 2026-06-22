package main

import "fmt"


func Solve(a int, b int) int {
	return a + b
}

func SolveSolve(a , b , c int) int {
	fmt.Println(a)
	return a + b + c;
}


func SecondFunction(x ,y int) int {
	return x + y
}

func MultiReturn(a ,b int) (int , int) {
	return 3, 7
}

func ReturnArray (a []int) []int {
	res := 0
	for i := 0 ; i < len(a) ; i++ {
		res += i
	}
	return a
}

func Sum(nums ...int) {
	fmt.Print(nums, " ")


}

// Closure

func inSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func CallFunc() {
	nextInt := inSeq()

	fmt.Println(nextInt())
}

// Recursion 

func Recursion () {
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}

