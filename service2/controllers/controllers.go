package controllers

import (
	"context"
	"net/http"
	"time"

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
		collection := client.Database("Cricket").Collection("players")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		_, err := collection.DeleteOne(ctx, bson.M{"_id": id})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "` + err.Error() + `"}`))
			return
		} else {
			// player not found
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "Player not found"}`))
			return
		}

	}
}
