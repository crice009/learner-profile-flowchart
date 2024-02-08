package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/dgraph-io/dgo/v230"
	"github.com/dgraph-io/dgo/v230/protos/api"
	"google.golang.org/grpc"
)

const (
	dgraphAddress = "localhost:9080"
)

type CancelFunc func()

func getDgraphClient() (*dgo.Dgraph, CancelFunc) {
	/// Dial a gRPC connection and return for further use.
	conn, err := grpc.Dial(dgraphAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	// Perform login call. ~ not needed by default self-hosted Dgraph
	// ctx := context.Background()
	// clientAuth(ctx, dg)

	return dg, func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection:%v", err)
		}
	}
}

func clientAuth(ctx context.Context, dg *dgo.Dgraph) {
	/// login to the cluster

	var err error

	// Perform login call. If the Dgraph cluster does not have ACL and
	// enterprise features enabled, this call should be skipped.

	for {
		// Keep retrying until we succeed or receive a non-retriable error.
		err = dg.Login(ctx, "groot", "password")
		if err == nil || !strings.Contains(err.Error(), "Please retry") {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		log.Fatalf("While trying to login %v", err.Error())
	}

}
