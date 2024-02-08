package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/dgraph-io/dgo/v230/protos/api"
)

// ---------------------------------------------------------
type School struct {
	Name  string   `json:"name,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

type loc struct {
	Type   string    `json:"type,omitempty"`
	Coords []float64 `json:"coordinates,omitempty"`
}

// If omitempty is not set, then edges with empty values (0 for int/float, "" for string,
// false for bool) would be created for values not specified explicitly.

type Person struct {
	Uid      string     `json:"uid,omitempty"`
	Name     string     `json:"name,omitempty"`
	Age      int        `json:"age,omitempty"`
	Dob      *time.Time `json:"dob,omitempty"`
	Married  bool       `json:"married,omitempty"`
	Raw      []byte     `json:"raw_bytes,omitempty"`
	Friends  []Person   `json:"friend,omitempty"`
	Location loc        `json:"loc,omitempty"`
	School   []School   `json:"school,omitempty"`
	DType    []string   `json:"dgraph.type,omitempty"`
}

func Example_setObject() {
	/// a single example that shows CRUD operations.

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	// getting a Dgraph client and setting up a defer statement
	// to ensure the client is properly closed when the function ends.
	dg, cancel := getDgraphClient()
	defer cancel()

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	// Next, a Person struct is created with various properties.
	// The Uid field is set to "_:alice", which is a blank node identifier.
	//   This means that Dgraph will assign a new unique identifier (uid) to this node when it's created.
	// The Person struct also includes a Location of type loc,
	//   Friends which is a slice of Person,
	//   and School which is a slice of School.
	dob := time.Date(1980, 01, 01, 23, 0, 0, 0, time.UTC)
	// While setting an object if a struct has a Uid then its properties
	//  in the graph are updated else a new node is created.
	// In the example below new nodes for Alice, Bob and Charlie and
	//  school are created (since they don't have a Uid).
	p := Person{
		Uid:     "_:alice",
		Name:    "Alice",
		Age:     26,
		Married: true,
		DType:   []string{"Person"},
		Location: loc{
			Type:   "Point",
			Coords: []float64{1.1, 2},
		},
		Dob: &dob,
		Raw: []byte("raw_bytes"),
		Friends: []Person{{
			Name:  "Bob",
			Age:   24,
			DType: []string{"Person"},
		}, {
			Name:  "Charlie",
			Age:   29,
			DType: []string{"Person"},
		}},
		School: []School{{
			Name:  "Crown Public School",
			DType: []string{"Institution"},
		}},
	}

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	// The api.Operation struct is used to define the schema for the data.
	// This schema includes types for Person, Institution, and Loc,
	//   as well as various properties for these types.
	op := &api.Operation{}
	op.Schema = `
		name: string @index(exact) .
		age: int .
		married: bool .
		loc: geo .
		dob: datetime .
		Friend: [uid] .
		type: string .
		coords: float .

		type Person {
			name: string
			age: int
			married: bool
			Friend: [Person]
			loc: Loc
		}

		type Institution {
			name: string
		}

		type Loc {
			type: string
			coords: float
		}
	`

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	// The dg.Alter function is then called to apply this schema to the database.
	// If there's an error, the program will log the error and exit.
	ctx := context.Background()
	if err := dg.Alter(ctx, op); err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	// A mutation is then created to add the Person object to the database.
	// The Person struct is marshalled into JSON and set as the SetJson field of the mutation.
	// The mutation is then applied to the database using the Mutate function.
	// If there's an error, the program will log the error and exit.
	mu.SetJson = pb
	response, err := dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	// The response.Uids map contains the uids assigned to the nodes that were created.
	// In this case, the uid for the "alice" node is retrieved and used in a query to get the data for this node.

	// Assigned uids for nodes which were created would be returned in the response.Uids map.
	variables := map[string]string{"$id1": response.Uids["alice"]}
	q := `query Me($id1: string){
		me(func: uid($id1)) {
			name
			dob
			age
			loc
			raw_bytes
			married
			dgraph.type
			friend @filter(eq(name, "Bob")){
				name
				age
				dgraph.type
			}
			school {
				name
				dgraph.type
			}
		}
	}`

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	// The QueryWithVars function is used to execute the query,
	// with the uid for "alice" passed in as a variable.
	// The result is then unmarshalled into a Root struct,
	// which contains a slice of Person structs.
	// The result is then marshalled into JSON with indentation
	// for readability and printed to the console.
	resp, err := dg.NewTxn().QueryWithVars(ctx, q, variables)
	if err != nil {
		log.Fatal(err)
	}

	type Root struct {
		Me []Person `json:"me"`
	}

	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		log.Fatal(err)
	}

	// ---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---\/---
	// The output shows the properties of the "alice" node, including the
	// "Bob" friend node and the "Crown Public School" school node.
	out, _ := json.MarshalIndent(r, "", "\t")
	fmt.Printf("%s\n", out)
	// Output: {
	// 	"me": [
	// 		{
	// 			"name": "Alice",
	// 			"age": 26,
	// 			"dob": "1980-01-01T23:00:00Z",
	// 			"married": true,
	// 			"raw_bytes": "cmF3X2J5dGVz",
	// 			"friend": [
	// 				{
	// 					"name": "Bob",
	// 					"age": 24,
	// 					"loc": {},
	// 					"dgraph.type": [
	// 						"Person"
	// 					]
	// 				}
	// 			],
	// 			"loc": {
	// 				"type": "Point",
	// 				"coordinates": [
	// 					1.1,
	// 					2
	// 				]
	// 			},
	// 			"school": [
	// 				{
	// 					"name": "Crown Public School",
	// 					"dgraph.type": [
	// 						"Institution"
	// 					]
	// 				}
	// 			],
	// 			"dgraph.type": [
	// 				"Person"
	// 			]
	// 		}
	// 	]
	// }
}

// ---------------------------------------------------------

type CRUD_Person struct {
	Uid  string `json:"uid,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func example_Create_node(ctx context.Context, p CRUD_Person) {
	dg, cancel := getDgraphClient()
	defer cancel()

	pb, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	_, err = dg.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Created node: %s", p.Name)
}

func example_Read_node(ctx context.Context, p CRUD_Person) {
	dg, cancel := getDgraphClient()
	defer cancel()

	query := fmt.Sprintf(`{
        all(func: uid(%s)) {
            uid
            name
            age
        }
    }`, p.Uid)

	resp, err := dg.NewTxn().Query(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	type Root struct {
		All []Person `json:"all"`
	}

	var r Root
	err = json.Unmarshal(resp.Json, &r)
	if err != nil {
		log.Fatal(err)
	}

	for _, person := range r.All {
		fmt.Printf("Uid = %s , Name = %s , Age = %d\n", person.Uid, person.Name, person.Age)
	}
}

func example_Update_node(p CRUD_Person, ctx context.Context) {

}

func example_Delete_node(p CRUD_Person, ctx context.Context) {

}

// ---------------------------------------------------------

func example_Create_edge() {}

func example_Read_edge() {}

func example_Update_edge() {}

func example_Delete_edge() {}
