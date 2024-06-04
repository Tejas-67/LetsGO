package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongo/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connnectionString = "mongodb+srv://tejasdev2023:h8Cz5KCFVUXpvcPT@project0.lphgs93.mongodb.net/?retryWrites=true&w=majority&appName=Project0"
const dbName = "cluster0"
const colName = "courses"

var collection *mongo.Collection

// connect with mongoDB

// specialised method -> runs at the very beginning once
func init() {
	clientOption := options.Client().ApplyURI(connnectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection successfull!")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("collection reference is ready!!!!")
}

func insertOneCourse(course model.Course) {
	inserted, err := collection.InsertOne(context.Background(), course)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a course: ", inserted.InsertedID)
}

func updateOneCourse(courseId string) {
	id, _ := primitive.ObjectIDFromHex(courseId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}};

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err!=nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount);
}

func deleteOneCourse(courseId string){
	id, _ := primitive.ObjectIDFromHex(courseId)
	filter := bson.M{"_id":id}
	count, err := collection.DeleteOne(context.Background(), filter)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("delete count: ", count.DeletedCount)
}

func deleteAllCourse(){
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("delete all count: ", deleteResult.DeletedCount)
}

func getAllCourses() []primitive.M{
	cursor, err := collection.Find(context.Background(), bson.M{})

	if err!=nil {
		log.Fatal(err)
	}

	var courses []primitive.M
	for cursor.Next(context.Background()){
		var course bson.M
		err := cursor.Decode(&course)

		if err!=nil {
			log.Fatal(err)
		}
		courses = append(courses, course)
	}
	
	defer cursor.Close(context.Background())
	return courses
}

// actual controller that will be imported
func GetAllCourses(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllCourses()
	json.NewEncoder(w).Encode(allMovies)
}
func MarkCourseWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Encode", "POST")
	
	params := mux.Vars(r)
	updateOneCourse(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func CreateCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var course model.Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	insertOneCourse(course)
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Encode", "POST")
	
	params := mux.Vars(r)
	deleteOneCourse(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Encode", "POST")
	
	deleteAllCourse()
	json.NewEncoder(w).Encode("Deleted all courses!!!")
}