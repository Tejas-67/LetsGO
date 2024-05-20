package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//use any valid url to get the result else you receive an error.
const url = "https://lco.dev"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("type: %T", response)
		fmt.Println("response: ", response)
		defer response.Body.Close() 
		//caller's responsibility to close the connection

		databytes, err := ioutil.ReadAll(response.Body)
		if err!=nil {
			panic(err)
		}
		content := string(databytes)
		fmt.Println("content: ", content)
	}

}
