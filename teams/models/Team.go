package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	ID            primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string               `json:"name,omitempty" bson:"name,omitempty"`
	BattingPower  int                  `json:"battingPower,omitempty" bson:"battingPower,omitempty"`
	BowlingPower  int                  `json:"bowlingPower,omitempty" bson:"bowlingPower,omitempty"`
	FieldingPower int                  `json:"fieldingPower,omitempty" bson:"fieldingPower,omitempty"`
	CaptainID     primitive.ObjectID   `json:"captainId,omitempty" bson:"captainId,omitempty"`
	Players       []primitive.ObjectID `json:"players,omitempty" bson:"players,omitempty"`
}
