package main

import (
	"fmt"
)

/* We are calling handlePanic() before the occurence of panic
-> This is because the program will be terminated if it encounters panic and we want
		to execute handlePanic before termination */

/* We are using defer() to call handlePanic() 
-> Its becuase we want to handle the panic only after its occurs, so we are deferring its execution  */

func handlePanic() {
	a := recover()

	if a != nil {
		fmt.Println("RECOVER ", a)
	}
}

func division(num1, num2 int) int {
	defer handlePanic()	

	if num2 == 0 {
		panic("mathematically incorrect, cannot divide by zero")
	}

	return num1 / num2
}

func main() {
	fmt.Println(division(4, 0))
	fmt.Println(division(4, 2))
}
