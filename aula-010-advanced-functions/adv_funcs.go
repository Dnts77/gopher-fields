package main

import "fmt"

//	This stands out for a quite simple idea - functions as data.
//	As so, store a function as a value, pass a function as a parameter of other function, returns a function after another, etc

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

//	aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func main() {
	fmt.Println(aggregate(2, 3, 4, add))
	//	Prints 9
	fmt.Println(aggregate(2, 3, 4, mul))
	//	Prints 24
}

//	For most of the time, higher-order functions just add needless complexity into the code.
//	But, there are three use cases that these functions are pretty handy.
//	1. HTTP API handlers
//	2. Pub/Sub handlers
//	3. Onclick callbacks

//	DEFINITION - FIRST-CLASS
//	A first-class function is a function that can be treated like any other value.
//	Go supports first-class functions. A function's type is dependent on the types of its parameters and return values.
//	For example, these are different function types

//func() int
//func(string) int{}

//	DEFINITION - HIGHER-ORDER
//	A higher-order function is a function that takes a function as an argument or returns a function as a return value.
//	Go supports higher-order functions.
//	For example, this function takes a function as an argument

// func aggregate(a, b, c int, arithmetic func(int, int) int) int {
// 	return arithmetic(arithmetic(a, b), c)
// }

//	CURRYING
//	Function currying is the practice of writing a function that takes a function (or functions) as input, and returns a new function

func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

//	Currying a function is mostly used in middleware functions.

//	DEFER
//	The "defer" keyword is a fairly unique feature in Go. It allows functions to be executed automatically just before its enclosing function returns

//	ANONYMOUS FUNCTIONS
//	Anonymous functions are true to form in that they have no name.
//	Anonymous functions are used when defining a function that will only be used once or to create a quick closure
