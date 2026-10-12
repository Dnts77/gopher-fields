package main

import "fmt"

func main() {
	fmt.Println("a")
	sub(10, 2)
	sum(10, 10)
	calculateSalary(3.000, 2.000)
	concatenateStrings("Hello ", "World!")
}

//Declaring a function
func sub(x int8, y int8) int8 { // The data type between the parenthesis and the curly braces means that this function has an int8 return type
	result := x - y
	fmt.Printf("O resultado é: %v \n", result)
	return result
}

func sum(x, y int16) int16 { //Here, the type declaration only happens after the last paremeter as Go considers both of them are the same type
	result := x + y
	fmt.Printf("O resultado é: %v \n", result)
	return result
}

func calculateSalary(baseSalary, discount float32) float32 {
	liquidSalary := baseSalary - discount

	if liquidSalary >= 10.000 {
		fmt.Printf("O seu salário é %f, logo, você está bem de vida\n", liquidSalary)
	} else if liquidSalary >= 7.000 {
		fmt.Printf("Seu salário é %f, então, você ainda vive bem\n", liquidSalary)
	} else if liquidSalary <= 2.000 {
		fmt.Printf("Seu salário é %f, então, você precisa gerenciar suas despesas melhor\n", liquidSalary)
	}
	return liquidSalary
}

func concatenateStrings(text1, text2 string) {
	concatenatedText := text1 + text2
	fmt.Println(concatenatedText)
}

func getCoords() (x, y int) { //Two returns
	return
	//It's the same as returning x and y
}
