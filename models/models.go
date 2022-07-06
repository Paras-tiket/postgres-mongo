package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SeatNo string             `json:"SeatNo" bson:"SeatNo,omitempty"`
	TimeT string             `json:"TimeT" bson:"TimeT,omitempty"`
	User   *User              `json:"User" bson:"User,omitempty"`
}

type User struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
type Train struct {
	TrainID   string `json:"TrainID"`
	TrainName string `json:"TrainName"`
	CoachNo   string `json:"CoachNo"`
	SeatNo    string `json:"SeatNo"`
	TimeT     string `json:"TimeT"`
	ClassT    string `json:"ClassT"`
}

type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Train `json:"data"`
	Message string  `json:"message"`
}
