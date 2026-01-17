package main

import (
	adapter "coin-keeper/internal/adapter/engine"
	"coin-keeper/internal/database/sql/engine"
	handlers "coin-keeper/internal/server/handlers/engine"
	"coin-keeper/internal/server/router"
	"net/http"
)

const connStr = "host=localhost port=5432 user=postgres password=1234 dbname=mydb sslmode=disable"

func main() {
	dbEngine, err := engine.NewDatabaseEngine(connStr)
	if err != nil {
		panic(err)
	}
	adapter := adapter.NewAdapter(dbEngine)
	handlersKeeper := handlers.NewHandlersKeeper(adapter)
	router := router.NewRouter(handlersKeeper)
	router.SetupRoutes()

	http.ListenAndServe(":8080", nil)
}
