package main

import (
	"fmt"
	"log"
	"mongo/router"
	"net/http"
)

func main() {
	fmt.Println("MON-GODDDD")

	r := router.Router()
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
	
}
