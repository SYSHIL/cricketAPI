package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"teams/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create_team(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("I am inside create functionx")
		var player_var models.Team
		err := json.NewDecoder(r.Body).Decode(&player_var)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		collection := client.Database("cricket").Collection("teams")
		player_var.ID = primitive.NewObjectID()

		result, err := collection.InsertOne(context.TODO(), player_var)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(result)

	}
}

// read all teams from mongodb
// Read_all_teams reads all teams from the MongoDB database
func Read_teams(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get a handle to the "cricket" database and the "teams" collection
		db := client.Database("cricket")
		collection := db.Collection("teams")

		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Define filter to find all teams (empty filter matches all)
		filter := bson.M{}
		// Find all teams in the collection
		cursor, err := collection.Find(
			ctx,
			filter,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)
		teams := []models.Team{}
		err = cursor.All(context.TODO(), &teams)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "internal sever error"})
			return
		}
		json.NewEncoder(w).Encode(teams)
	}
}

// Simulate_league simulates a league by creating a fixture list and running each fixture
// each team plays the other team only once,
// maintain a league table that maps team id to points
// win means 2 points, loss means 0 points
// finally top 2 teams in the league table must play a final match
// return the winner of the final match

func Simulate_league(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get all teams from teams collection
		// create a fixture list
		// run each fixture
		// maintain a league table that maps team id to points
		// win means 2 points, loss means 0 points
		// finally top 2 teams in the league table must play a final match
		// return the winner of the final match
		// Get a handle to the "cricket" database and the "teams" collection
		db := client.Database("cricket")
		collection := db.Collection("teams")

		// Create a context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Define filter to find all teams (empty filter matches all)
		filter := bson.M{}

		// Find all teams in the collection
		cursor, err := collection.Find(
			ctx,
			filter,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)
		teams := []models.Team{}
		err = cursor.All(context.TODO(), &teams)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "internal sever error"})
			return
		}
		// run a loop to create a fixture list

		// create a map for storing the points of each team, key is team id, value is points
		leagueTable := map[primitive.ObjectID]int{}
		// prepopulate leaguetable with 0 points for each team
		for _, team := range teams {
			leagueTable[team.ID] = 0
		}

		// run a loop to simulate each fixture

		for i := 0; i < len(teams); i++ {
			for j := i + 1; j < len(teams); j++ {
				// team i vs team j
				// winner is decided by the following score
				// score = battingPower + bowlingPower + fieldingPower/2
				team1Score := teams[i].BattingPower + teams[i].BowlingPower + teams[i].FieldingPower/2
				team2Score := teams[j].BattingPower + teams[j].BowlingPower + teams[j].FieldingPower/2
				if team1Score > team2Score {
					leagueTable[teams[i].ID] += 2
				} else if team2Score > team1Score {
					leagueTable[teams[j].ID] += 2
				} else {
					leagueTable[teams[i].ID] += 1
					leagueTable[teams[j].ID] += 1
				}

			}
		}

		// get top 2 teams from league table
		// sort the league table by points
		// create a slice of teams
		// append each team to the slice
		// return the slice
		type teamPoints struct {
			TeamID primitive.ObjectID `json:"teamID"`
			Points int                `json:"points"`
		}
		teamPointsSlice := []teamPoints{}
		for teamID, points := range leagueTable {
			teamPointsSlice = append(teamPointsSlice, teamPoints{TeamID: teamID, Points: points})
		}

		// sort the slice
		for i := 0; i < len(teamPointsSlice); i++ {
			for j := i + 1; j < len(teamPointsSlice); j++ {
				if teamPointsSlice[i].Points < teamPointsSlice[j].Points {
					teamPointsSlice[i], teamPointsSlice[j] = teamPointsSlice[j], teamPointsSlice[i]
				}
			}
		}
		// fmt.Println(teamPointsSlice)
		// return the league table for testing

		// json.NewEncoder(w).Encode(teamPointsSlice)

		// get top 2 teams
		// top2Teams := []primitive.ObjectID{teamPointsSlice[0].teamID, teamPointsSlice[1].teamID}

		// // play a final game between top 2 teams
		// // winner is decided by the following score
		// // score = battingPower + bowlingPower + fieldingPower/2
		team1Score := teams[0].BattingPower + teams[0].BowlingPower + teams[0].FieldingPower/2
		team2Score := teams[1].BattingPower + teams[1].BowlingPower + teams[1].FieldingPower/2
		var winner primitive.ObjectID
		if team1Score > team2Score {
			winner = teams[0].ID
		}
		if team2Score > team1Score {
			winner = teams[1].ID
		}
		// // get team by winner id from the collection
		// // return the team
		// // Define filter to find all teams (empty filter matches all)
		filter = bson.M{"_id": winner}

		// // Find all teams in the collection
		cursor, err = collection.Find(
			ctx,
			filter,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)
		team := []models.Team{}
		err = cursor.All(context.TODO(), &team)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "internal sever error"})
			return
		}
		json.NewEncoder(w).Encode(team)

	}
}
