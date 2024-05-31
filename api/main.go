package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	router := mux.NewRouter()
	
	//seeding 
	courses = append(courses, Course{
		CourseId: "1",
		CourseName: "course1",
		CoursePrice: 100,
		Author: &Author{
			Name: "tejas",
			Email: "tejas&mm",
			Website: "tejasdotcom",
		},
	})
	courses = append(courses, Course{
		CourseId: "2",
		CourseName: "course2",
		CoursePrice: 200,
		Author: &Author{
			Name: "sohan",
			Email: "sohan&mm",
			Website: "sohandotcom",
		},
	})
	courses = append(courses, Course{
		CourseId: "3",
		CourseName: "course3",
		CoursePrice: 300,
		Author: &Author{
			Name: "aryan",
			Email: "aryan&mm",
			Website: "aryandotcom",
		},
	})
 
	router.HandleFunc("/", serverHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", router))
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

func createOneCourse(w http.ResponseWriter, r *http.Request) {
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

func updateCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("updating a course......................")
	w.Header().Set("Content-Type", "application/json")

	//first -> grab id from req
	params := mux.Vars(r)
	
	//loop, id, remove, add with my ID
	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			var c Course
			_ = json.NewDecoder(r.Body).Decode(&c)
			courses = append(courses, c)
			c.CourseId = params["id"]
			courses = append(courses, c)		
			json.NewEncoder(w).Encode(c)
			return	
		}
	}
	// course not found -> send a response
	json.NewEncoder(w).Encode("Course not found!!")
}


func deleteCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Deleting")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Successfully deleted course")
			return
		}
	}
	json.NewEncoder(w).Encode("Couldn't find course with provided ID")
}