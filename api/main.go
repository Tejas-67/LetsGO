package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for Courses
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	Author      *Author `json:"author"`
	CoursePrice int     `json:"price"`
}

// fake database
var courses []Course

//middlewares

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

type Author struct {
	Name    string `json:"fullname"`
	Email   string `json:"email"`
	Website string `json:"website"`
}

func main() {
	fmt.Println("API ______________________________")
}

//controller

// serve-home route
func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Tejas</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//looping through the courses and find matching id and return response
	for index, c := range courses {
		if c.CourseId == params["id"] {
			json.NewEncoder(w).Encode(c)
			fmt.Println("found at: ", index)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id")
}

func createOneCourse(r *http.Request, w http.ResponseWriter) {
	fmt.Println("Add one course")
	w.Header().Set("Content-Type", "applicatoin/json")

	//what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Body is empty.")
	}
	//what if: {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Empty json")
		return
	}

	// generate unique id of type:string
	// append the new course into courses
	rand.Seed(time.Now().Unix())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}
