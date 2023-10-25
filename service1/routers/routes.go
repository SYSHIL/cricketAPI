package routers

import (
	"service1/controllers"

	"service1/db"

	"github.com/gorilla/mux"
)

// func myconnection() *mongo.Client {

// 	// const dbname = "GoProject"
// 	// const colname = "users"

// 	// var collection *mongo.Collection
// 	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
// 	opts := options.Client().ApplyURI("mongodb+srv://shivaniniranjan30:Shivani30!@cluster0.0q2rc6i.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

// 	// Create a new client and connect to the server
// 	client, err := mongo.Connect(context.TODO(), opts)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// defer func() {
// 	//   if err = client.Disconnect(context.TODO()); err != nil {
// 	// 	panic(err)
// 	//   }
// 	// }()

// 	// const collection := client.Database(dbname).Collection(colname)

// 	// Send a ping to confirm a successful connection
// 	if err := client.Database("Cricket").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
// 	return client
// }

func Route_func(router *mux.Router) {
	client := db.GetGlobalMongoClient()
	router.HandleFunc("/create_player", controllers.Create_player(client)).Methods("POST")
	router.HandleFunc("/read_all_players", controllers.Read_all_players(client)).Methods("GET")
	router.HandleFunc("/update_player/{PlayerID}", controllers.Update_player(client)).Methods("POST")
}
