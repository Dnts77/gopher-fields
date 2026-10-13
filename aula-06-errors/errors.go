package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*  user, err := getUser()  //So what happens here is: the function "getUser()" returns an user and an error.
	 if err != nil {         //If the error is different from nil (null, so there's an error), it is treated and
	 	fmt.Println(err)  //only after that, the happy path happens
	 	return
	 }

	 profile, err := getUserProfile(user.id)
	 if err != nil {
	 	fmt.Println(err)
	 	return
	 }

	// Dealing with "user" after treating the error
	 fmt.Println("Tudo certo com o usuário> %v", user) */

	convertStrToInt()

}

/* func getUser() (User, error) { //Differently from other programming languages, Go makes it clear to what the function returns right at its signature
		do some get user logic here
	return user
} */

//	Go programs express errors with "error" values. An Error is any type that implements the error interface

func convertStrToInt() {
	i, err := strconv.Atoi("46")
	if err != nil {
		fmt.Println("Não foi possível converter:", err)
		return
	}
	//If the code makes it through here, it means that "i" was succefully converted
	fmt.Printf("i: %v\n", i)

}

// Custom Error interfaces
type userError struct {
	name string
}

func (e *userError) Error() string { // The function name "Error()" implements the original Error interface
	return fmt.Sprintf("%v possui um problema com a conta", e.name)
}
