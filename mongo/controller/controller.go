package controller

import (
	"context"
	"fmt"
	"log"
	"mongo/model"

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

func updateOneCourse(courseId string, course model.Course) {
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
