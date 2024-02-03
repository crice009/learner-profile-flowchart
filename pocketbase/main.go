package main

import (
	"context"
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	// create a new PocketBase instance
	app := pocketbase.New()

	// build a new GraphDB driver, to connect to neo4j
	gdb, err := NewGraphDB("bolt://neo4j:7687", os.Getenv("NEO4J_USERNAME"), os.Getenv("NEO4J_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}

	// the 'helloWorld' shows we can connect to neo4j
	ctx := context.TODO()
	log.Printf(helloWorld(ctx, gdb))

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/static/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
