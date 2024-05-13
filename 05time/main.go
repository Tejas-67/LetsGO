package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time")

	presentTime := time.Now()
	fmt.Println("presentTime: ", presentTime)

	//for date in mm-dd-yyyy, use string layout = "01-02-2006"
	//for time use layout = 15:04:05
	//for day use layout = "Monday"
	fmt.Println("presentTime also: ", presentTime.Format("15:04:05"))

	createdDate := time.Date(2020, time.January, 11, 0, 0, 0, 0, time.Local)
	fmt.Println("createdDate: ", createdDate.Format("Monday"))
}
