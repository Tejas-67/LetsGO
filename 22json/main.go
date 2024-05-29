package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"username"`
	Price    int `json:"amount"`
	Age      float32
	Password string `json:"-"` // i dont want this feild in json
	Email    string

}

func main() {
	fmt.Println("JSON.......")
	// EncodeJson()
	DecodeJson()

}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
	{
		"username": "Sohan",
		"amount": 30,
		"Age": 20,
		"Email": "email__"
}
	`)
	
	var courseObject course
	isValid := json.Valid(jsonDataFromWeb)
	if isValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &courseObject)
		fmt.Printf("%#v\n", courseObject);
	}else {
		fmt.Println("JSON was not valid")
	}

	// other method - storing as key-value pair
	var myMap map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myMap)
	fmt.Println(myMap);

	for k, v := range myMap {
		fmt.Printf("Key is %v and value is %v and type: %T \n", k, v, v);
	}
}

func EncodeJson(){
	object :=  []course{
		{"sdjgb", 10, 10.0, "dljnaa", "email"},
		{"adoub", 20, 20.0, "adsjb", "email_"},
		{"ansifi", 30, 20.0, "djbasm", "email__"},
	}
	//package this data as json

	finalJson, err := json.MarshalIndent(object, "", "\t")
	//last parameter is \t which will add required number of tab-spaces

	if err!=nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson);
}