package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Patient is ...
type Patient struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Gender  string             `json:"gender" bson:"gender,omitempty"`
	Address *Address           `json:"address" bson:"address,omitempty"`
}

// Address is ...
type Address struct {
	State   string `json:"state,omitempty" bson:"state,omitempty"`
	Country string `json:"country,omitempty" bson:"country,omitempty"`
}
