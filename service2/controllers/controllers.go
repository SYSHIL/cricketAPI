package controllers

import (
	"context"
	"fmt"
	"net/http"

	models "service2/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// get player id in params and delete by id
func Delete_player(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		PlayerID := params["PlayerID"]
		id, _ := primitive.ObjectIDFromHex(PlayerID)
		collection := client.Database("cricket").Collection("players")
		fmt.Println("Deleting player with id: ", id)

		// get player by id
		var player models.Player
		filter := bson.M{"_id": id}
		err := collection.FindOne(context.TODO(), filter).Decode(&player)
		if err != nil {
			// player not found
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "Player not found"}`))
			return
		}
		// // print player
		fmt.Println("Player: ", player)

		// delete player by id
		filter = bson.M{"_id": id}
		_, err = collection.DeleteOne(context.TODO(), filter)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "` + err.Error() + `"}`))
			return
		}

		// Deletion successful
		w.WriteHeader(http.StatusNoContent)

	}
}
