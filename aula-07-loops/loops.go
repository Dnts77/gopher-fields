package main

import "fmt"

func main() {
	loopI()
	// omittedLoop()
	whileLoop()
}

//	Sintatically, loops in Go are very similar to JavaScript, the only difference is that the parenthesis are not used

func loopI() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, "\n")
	}
}

//	It's totally fine to omit conditions from a for loop in Go
func omittedLoop() {
	for i := 0; ; i++ { //	The condition was omitted, so this loop will run FOREVER
		fmt.Print(i)
	}
}

//	There is no WHILE loop in Go, instead, a for loop is used with only its condition

func whileLoop() {
	i := 0
	for i < 10 {
		fmt.Print(i, "\n")
		i++
	}
	fmt.Printf("O número atingiu o valor 10,\nValor atual de i: %v", i) //	Execute this line only the for loop is over
}

//	Modulo operator
//	The modulo operator is a percent sign (%)
//	It calculates remainders
//	For example, in a 4/3 division, the result would be 1.333, but, it is not possible to divide integers and have a floating point number result.
// 	So Go would just chop off the .333 and the result would be just 1
// 	But with 4%3, the remainder would also be 1. Whereas 5%3, the remainder would be 2. In 6%3, the remainder would be 0. And so on

//	Logical Operators
//	Logical AND operator -> && 		true && false
//	Logical OR operator -> || 		true || false

//	The "continue" keyword stops the current iteration of a loop and continues to the next iteration
//	The "break" keyword stops the current iteration of a loop and exits the loop
