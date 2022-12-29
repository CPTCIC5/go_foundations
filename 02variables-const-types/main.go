package main

import "fmt"

//var jwtToken int = 100
const LoginToken string = "efhe" //capital == Public vc

func main() {
	var username string = "hitest"
	fmt.Println("Hey", username)

	var isLoggedin bool = true
	fmt.Println("Logged in", isLoggedin)

	var age uint8 = 0 //uint8 limits from 0 to 250
	age = 250
	fmt.Println("Logged in", age)

	var marks float32 = 250.455 //0 to 255.45 smthg
	marks = 21.69
	fmt.Println("Logged in", marks)

	var anotherVar int
	fmt.Println(anotherVar)

	var username2 = "hitest" //we can add the datatype or not it's upto us lexical will handle it
	fmt.Println(username2)
	fmt.Println(LoginToken)
}
