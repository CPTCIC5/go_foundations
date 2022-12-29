package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := "Welcome to the user input"
	//var name string = "Welcome to user input"
	fmt.Println(name)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza")

	//comma ok syntax || err ok
	input, _ := reader.ReadString('\n') //_ is used as a catch/except
	fmt.Println("thanks for rating", input)
	fmt.Printf("Data type of this input is %T", input)
}
