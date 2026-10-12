package main //main package - whenever created an individual file that's gonna be executed - use package main

import "fmt" //This package allows output things to the console

func main() { //Main function
	fmt.Println("Hello, World!")
	sayHello()
}

func sayHello() {
	fmt.Print("Hello again!")
}
