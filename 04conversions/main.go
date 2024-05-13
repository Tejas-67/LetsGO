package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to hell")
	fmt.Println("enter a number")
	reader := bufio.NewReader(os.Stdin)

	//taking input as a string
	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for entering: ", input)

	//converting input into a number, and also checking if there is any error.
	//by default the input has a endline suffix, so to remove that we use strings.TrimSpace()
	//also if input is not a number, then we handle the error
	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 10, 32) 

	//checks if there is an error.
	//if error present, then using panic() program ends and error is logged into the terminal
	//if no error then printing the value
	if err != nil {
		panic(err)
	} else {
		fmt.Println("numRating: ", numRating)
	}
}
