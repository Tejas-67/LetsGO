package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to hell"
	fmt.Println(welcome)
	//taking input from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	//comma ok || err ok
	rating, err := reader.ReadString('\n')
	fmt.Println("rating is: ", rating)
	fmt.Printf("type of rating is: %T \n", rating)
	fmt.Println("error: ", err)
}
