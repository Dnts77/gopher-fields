package main

import "fmt"

func main() {
	x := 5
	y := x
	z := &x //	& -> Pointer syntax. Because of the "&", the stored value in "z" is the address of "x". So here, "z" is a pointer type
	*z = 6
	//	So, what this means is:
	//	If the value that "z" points to gets updated, under the hood "x" will be updated
	//	It's possible to pass pointers into functions -> change the underlying value and the value outside of that function is also changed.

	//	So, if we want to update the value of the variable "x" without having access to the original variable "x", what we can do is:
	//	Use the dereference operator " *z ", where "z" is the variable being dereferenced.
	//	So, above, where " z := &x " and then " *z = 6 ", essentially means: "z" stores the address in memory of "x" as its value. Then "z" is dereferenced and stores 6, which updates "x" to 6 aswell.
	//	"x" got its value updated without even using the name "x".

	//	Pointers are useful in two primary situations:
	//		1. You want to be able to pass a value into a function and change the valuea and have those changes "persisting" outside of the function
	//			because normally when you pass a value into a function, a copy is passed in. So if you want to pass the "original" value to be changed
	//			updated, you might use a pointer

	//		2.	If you're very concerned about the performance of your program. Every time you create a copy of a variable in memory, you have to copy that
	//			 variable in memory, which takes some time and consume memory.

	fmt.Println(x)  //	Prints 6, because "x" got its value changed by "*z"
	fmt.Println(y)  //	Prints 5, because it stores a copy of "x" and not the original "x"
	fmt.Println(z)  //	Prints the memory address of "x", in hexadecimal values
	fmt.Println(*z) //	Prints 6 again, because it stores 6 and updates "x"

	//	Pointers can be very dangerous.
	//	If a pointer points to nothing, then dereferencing it will cause a runtime error that crashes the program.
	//	Generally speaking, whenever dealing with pointers, it's necessary to check if it's nil before trying to dereference it.

	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "blue"
}

// 	x := 5
//	Under the hood, what happens is -> somewhere in the computer's memory, that variable's value get stored
//	If another variable is created and set to the previous variable -> y := x, y gets a new copy of the value.
//	So, if "x" lives in address 169 in the computer's memory, y could live in address 170 being a copy of x (5)

//	Pointer receivers
type car struct {
	color string
}

func (c *car) setColor(color string) { //Methods usually take a pointer receiver because the method will be making changes to the instance of the type itself
	c.color = color
}
