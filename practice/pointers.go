/*
we have three things 
i , &i , *i
i -> normal value
&i -> value ka address
*i -> point to a value means it can store address of a value

i 
p = &i
p -> address of i
*p --> value at address i 

*/

package main

func zeroval(val int) {
	val = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func pointers(){
	i := 1
	zeroval(i) // value of i will be 0 in the function but outside the function value will be same
}