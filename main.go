package main

import (
	"fmt"
	"hallo-corona-be/database"
	"hallo-corona-be/pkg/mysql"
	"hallo-corona-be/routes"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// env
    errEnv := godotenv.Load()
    if errEnv != nil {
      panic("Failed to load env file")
    } else {
		fmt.Println("Env success loaded.")
	}

	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()
	
	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}