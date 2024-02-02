package main

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	graph_db_username := os.Getenv("GRAPH_DB_USERNAME")
	graph_db_password := os.Getenv("GRAPH_DB_PASSWORD")

	db, err := NewGraphDB("bolt://localhost:7687", graph_db_username, graph_db_password)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Connected to Neo4j database, without error.")
		log.Printf("Database: %v", db)
	}

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
