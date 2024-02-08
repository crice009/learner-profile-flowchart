package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v230/protos/api"
)

func Example_getSchema() {
	/// This Go code is an example of how to interact with a Dgraph database,
	///    specifically how to get the schema of the database.

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	// Start by getting a Dgraph client and setting up a
	// defer statement to ensure the client is properly
	//   closed when the function ends.
	dg, cancel := getDgraphClient()
	defer cancel()

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	// Next, an api.Operation struct is created and its Schema field is
	// set to a string that defines the schema for the data.
	// This schema includes properties name, age, married, loc, and dob
	// with their respective types. The name property is also indexed for exact matches.
	op := &api.Operation{}
	op.Schema = `
		 name: string @index(exact) .
		 age: int .
		 married: bool .
		 loc: geo .
		 dob: datetime .
	 `

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	ctx := context.Background()
	// The dg.Alter function is then called to apply this schema to the database.
	// If there's an error, the program will log the error and exit.
	err := dg.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	// After the schema is set, a query is made to the database to get the types
	// of the name and age properties. The dg.NewTxn().Query function is used to
	// execute the query.
	// If there's an error, the program will log the error and exit.

	// Ask for the type of name and age.
	resp, err := dg.NewTxn().Query(ctx, `schema(pred: [name, age]) {type}`)
	if err != nil {
		log.Fatal(err)
	}

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	// The response from the query is stored in resp.
	// The Json field of resp contains the response to the schema query.
	// This is then printed to the console.
	// The output shows the types of the name and age properties,
	//   which are string and int respectively.

	// resp.Json contains the schema query response.
	fmt.Println(string(resp.Json))
	// Output: {"schema":[{"predicate":"age","type":"int"},{"predicate":"name","type":"string"}]}
}

// ---------------------------------------------------------
func example_Create_schema() {}

func example_Read_schema() {}

func example_Update_schema() {}

func example_Delete_schema() {}
