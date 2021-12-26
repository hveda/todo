package main

import (
	"log"
	"net/http"

	"github.com/hveda/todo/src/routes"
	"github.com/hveda/todo/src/types"
	"github.com/hveda/todo/src/utils"
	"github.com/gorilla/mux"
)

func main() {
	// Attempt connection with DB
	conn, err := utils.DbConnect()
	if err != nil {
		log.Fatal("Could not connect to database.\n", err)
		return
	}

	// Migrate db when success to connect
	conn.AutoMigrate(&types.Activity{}, &types.ToDo{})

	// Create router instance
	router := mux.NewRouter()
	// Create app base with DB connection
	app := routes.New(conn)
	app.CreateRoutes(router)

	const port = "3030"
	log.Printf("Server started on port %s\n", port)
	if err := http.ListenAndServe(":" + port, router); err != nil {
		log.Fatalf("Server already started on port %s\n\n", port)
		log.Fatal(err)
	}
}
