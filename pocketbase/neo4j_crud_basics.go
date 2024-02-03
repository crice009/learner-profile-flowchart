// neo4j_crud_basics.go
package main

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type GraphDB struct {
	driver neo4j.DriverWithContext
}

// NewGraphDB creates a new GraphDB instance
func NewGraphDB(uri, username, password string) (*GraphDB, error) {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	return &GraphDB{driver: driver}, nil
}

func (db *GraphDB) CreateNode(label string, properties map[string]interface{}) (*neo4j.Node, error) {
	// Implement node creation here
	return nil, nil
}

func (db *GraphDB) ReadNode(id int64) (*neo4j.Node, error) {
	// Implement node reading here
	return nil, nil
}

func (db *GraphDB) UpdateNode(id int64, properties map[string]interface{}) (*neo4j.Node, error) {
	// Implement node updating here
	return nil, nil
}

func (db *GraphDB) DeleteNode(id int64) error {
	// Implement node deletion here
	return nil
}

func (db *GraphDB) CreateRelationship(startNodeID, endNodeID int64, relationshipType string, properties map[string]interface{}) (*neo4j.Relationship, error) {
	// Implement relationship creation here
	return nil, nil
}

func (db *GraphDB) ReadRelationship(id int64) (*neo4j.Relationship, error) {
	// Implement relationship reading here
	return nil, nil
}

func (db *GraphDB) UpdateRelationship(id int64, properties map[string]interface{}) (*neo4j.Relationship, error) {
	// Implement relationship updating here
	return nil, nil
}

func (db *GraphDB) DeleteRelationship(id int64) error {
	// Implement relationship deletion here
	return nil
}

// ---------------------------------------------------------
func helloWorld(ctx context.Context, db *GraphDB) (string, error) {
	session := db.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	greeting, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx,
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]any{"message": "hello, world"})
		if err != nil {
			return nil, err
		}

		if result.Next(ctx) {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}

	return greeting.(string), nil
}
