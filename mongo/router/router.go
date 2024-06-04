package router

import (
	"mongo/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/courses", controller.GetAllCourses).Methods("GET")
	router.HandleFunc("/course", controller.CreateCourse).Methods("POST")
	router.HandleFunc("/course/{id}", controller.MarkCourseWatched).Methods("PUT")
	router.HandleFunc("/course/delete/{id}", controller.DeleteOneCourse).Methods("DELETE")
	router.HandleFunc("/course/delete", controller.DeleteAllCourse).Methods("DELETE")

	return router
}