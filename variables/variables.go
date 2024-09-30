package main

import "fmt"

// only var and const can be used

func main() {

	var name string = "john"
	var age int = 28
	var text string

	fmt.Println(name)
	fmt.Println(age)

	age = 30 // re-assigning

	text = "hello"
	fmt.Println(age)
	fmt.Println(text)

	// constants

	const pi = 3.14159
	const url = "https://www.google.com"

	fmt.Println(url)
	fmt.Println(pi)

}
