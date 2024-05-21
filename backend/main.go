package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Backend.....")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code: ", response.StatusCode)
	fmt.Println("content length: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("content: ", string(content))
	//OR
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("bytecount is: ", byteCount);
	//converting to string so that we can see the content
	responseToString := responseString.String()
	fmt.Println("reponse: ", responseToString)
}
