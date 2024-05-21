package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// use any valid url to get the result else you receive an error.
const myurl string = "https://bunkbuddy-backend.onrender.com/allusers"

func main() {
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("type: %T", response)
		fmt.Println("response: ", response)
		defer response.Body.Close()
		//caller's responsibility to close the connection

		databytes, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		content := string(databytes)
		fmt.Println("content: ", content)
	}

	result, _ := url.Parse(myurl)
	fmt.Println("scheme", result.Scheme)
	fmt.Println("host", result.Host)
	fmt.Println("path", result.Path)
	fmt.Println("port", result.Port())
	fmt.Println("query", result.RawQuery)

	//get all the query parameters in this way and then do
	//qparams["name_of_query_parameter"] to get the value!!
	//good language!
	qparams := result.Query();
	fmt.Printf("type of query params: %T", qparams);

	//creating a Url
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "bunkbuddy-backend.onrender.com",
		Path: "/allusers",
	}
	parsedUrl := partsOfUrl.String();
	fmt.Println(parsedUrl)
}
