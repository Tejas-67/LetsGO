package main

import "fmt"

func main() {
	fmt.Println("first time writing go")
	for index := 0; index < 10; index++ {
		fmt.Println(index % 2)
	}
	toString()
}

func toString() string {
	return ""
}
