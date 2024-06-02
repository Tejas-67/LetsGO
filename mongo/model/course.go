package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	CourseId primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CourseName  string  `json:"coursename"`
	Author      *Author `json:"author"`
	CoursePrice int     `json:"price"`
	Watched bool `json:"watched,omitempty"`
}

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

