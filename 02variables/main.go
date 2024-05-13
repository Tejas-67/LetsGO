package main

import "fmt"

const login = ""
const Login = ""

func main() {
	fmt.Println("Variables in GO")
	var username string = "tejas"
	fmt.Println(username)
	// var fraction float32 = 1.0

	var number uint = 255
	fmt.Printf("type: %T \n", number)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("this is: %T", anotherVariable)

	f := 3
	fmt.Println("fuck: ", f)

	fmt.Println(login)
	fmt.Println(Login)
}
