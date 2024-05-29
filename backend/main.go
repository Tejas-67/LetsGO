package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Backend.....")
	// PerformGetRequest()
	// performPostJsonRequest()
	performPostFormRequest();
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

	content, _ := io.ReadAll(response.Body)
	fmt.Println("content: ", string(content))
	//OR
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("bytecount is: ", byteCount)
	//converting to string so that we can see the content
	responseToString := responseString.String()
	fmt.Println("reponse: ", responseToString)
}

func performPostJsonRequest() {
	const myrul = "http://localhost:8000/postJson"

	//fake payload as json
	payload := strings.NewReader(`
	{
		"email": "tejasjha3210@gmail",
		"name": "pass"
	}
	`)
	response, err := http.Post(myrul, "application/json", payload)
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}
	content, _ := io.ReadAll(response.Body)
	fmt.Println("response: ", string(content))
}

func performPostFormRequest() {
	const myurl = "http://localhost:8000/postForm"

	data := url.Values{}
	data.Add("name", "tejas");
	data.Add("email", "emiaa");

	response, err := http.PostForm(myurl, data)
	// defer response.Body.Close()
	if err!=nil{
		panic(err)
	}
	content, _:= io.ReadAll(response.Body);
	fmt.Println("response: ", string(content));
}
