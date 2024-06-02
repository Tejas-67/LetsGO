package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Author struct {
	Name    string `json:"fullname"`
	Email   string `json:"email"`
	Website string `json:"website"`
}
