package main

import "fmt"

func main() {
	fmt.Println("This is pointers....")


	//default value of any pointer is <nil>
	var one *int
	fmt.Println("pointer: ", one) //p[rints <nil>

	num:=26
	var pointerToNum = &num
	//& is used to provide a reference to the memory location of num
	fmt.Println("Pointer to num is: ", pointerToNum)
	//* will check the memory address and get its value
	fmt.Println("Value at memory location: ", *pointerToNum)

	*pointerToNum = *pointerToNum*2
	fmt.Println("double: ", num)
}

