package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Jersey        string             `json:"jersey,omitempty" bson:"jersey,omitempty"`
	Age           int                `json:"age,omitempty" bson:"age,omitempty"`
	PrimaryRole   string             `json:"primary_role,omitempty" bson:"primary_role,omitempty"`
	SecondaryRole []string           `json:"secondary_role,omitempty" bson:"secondary_role,omitempty"`
	Matches       int                `json:"matches,omitempty" bson:"matches,omitempty"`
	Runs          int                `json:"runs,omitempty" bson:"runs,omitempty"`
	Wickets       int                `json:"wickets,omitempty" bson:"wickets,omitempty"`
}
