package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"service1/models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create_player(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I am inside create functionx")
		var player_var models.Player
		err := json.NewDecoder(r.Body).Decode(&player_var)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		collection := client.Database("cricket").Collection("players")
		player_var.ID = primitive.NewObjectID()

		result, err := collection.InsertOne(context.TODO(), player_var)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(result)

	}
}

// read all players from mongodb
// Read_all_players reads all players from the MongoDB database
func Read_all_players(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get a handle to the "Cricket" database and the "players" collection
		db := client.Database("cricket")
		collection := db.Collection("players")

		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Define filter to find all players (empty filter matches all)
		filter := bson.M{}

		// Find all players in the collection
		cursor, err := collection.Find(ctx, filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)
		players := []models.Player{}
		err = cursor.All(context.TODO(), &players)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "internal sever error"})
			return
		}
		json.NewEncoder(w).Encode(players)
	}
}

// func Update_player(client *mongo.Client, userID string, updateFields bson.M) http.HandlerFunc {
// 	// Get a handle to the "Cricket" database and the "users" collection
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		db := client.Database("Cricket")
// 		collection := db.Collection("players")

// 		// Create a context with timeout
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()

// 	}
// }

// func Update_player(client *mongo.Client) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		params := mux.Vars(r)
// 		Name := params["Name"]

// 		objectID, err := primitive.ObjectIDFromHex(Name)
// 		if err != nil {
// 			http.Error(w, "Invalid user ID", http.StatusBadRequest)
// 			return
// 		}

// 		var user models.Player
// 		err = json.NewDecoder(r.Body).Decode(&user)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}

// 		collection := client.Database("Cricket").Collection("players")
// 		filter := bson.M{"_id": objectID}
// 		update := bson.M{"$set": user}

// 		_, err = collection.UpdateOne(context.TODO(), filter, update)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

//			json.NewEncoder(w).Encode("User updated successfully")
//		}
//	}

func Update_player(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		PlayerID := params["PlayerID"]

		objectID, err := primitive.ObjectIDFromHex(PlayerID)
		if err != nil {
			http.Error(w, "Invalid player ID", http.StatusBadRequest)
			return
		}
		if r.Body == nil {
			http.Error(w, "Please send a request body", http.StatusBadRequest)
			return
		}

		// Fetch the existing player from the database
		collection := client.Database("cricket").Collection("players")
		filter := bson.M{"_id": objectID}
		var existingPlayer models.Player
		err = collection.FindOne(context.TODO(), filter).Decode(&existingPlayer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var player models.Player
		err = json.NewDecoder(r.Body).Decode(&player)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update the player with the new values
		if player.Name != "" {
			existingPlayer.Name = player.Name
		}
		if player.Jersey != "" {
			existingPlayer.Jersey = player.Jersey
		}
		if player.Age != 0 {
			existingPlayer.Age = player.Age
		}
		if player.PrimaryRole != "" {
			existingPlayer.PrimaryRole = player.PrimaryRole
		}
		if player.SecondaryRole != nil {
			existingPlayer.SecondaryRole = player.SecondaryRole
		}
		if player.Matches != 0 {
			existingPlayer.Matches = player.Matches
		}
		if player.Runs != 0 {
			existingPlayer.Runs = player.Runs
		}
		if player.Wickets != 0 {
			existingPlayer.Wickets = player.Wickets
		}

		// Create an update filter based on the PlayerID
		updateFilter := bson.M{"_id": objectID}

		// Create an update document that sets all fields
		update := bson.M{"$set": existingPlayer}

		_, err = collection.UpdateOne(context.TODO(), updateFilter, update)
		if err != nil {
			log.Printf("Error updating player: %v\n", err)
			http.Error(w, "Player update failed", http.StatusInternalServerError)
			return
		}

		// Fetch the updated player to ensure it was successfully updated
		collection.FindOne(context.Background(), updateFilter).Decode(&existingPlayer)

		fmt.Println("Player updated successfully")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(existingPlayer)
	}
}
