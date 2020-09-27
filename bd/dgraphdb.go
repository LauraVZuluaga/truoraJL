package bd

import (
	"context"
	"log"
	"encoding/json"
	"../models"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

func newClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	connect, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(connect),
	)
}

func SaveBuyers(buyers []models.Buyer) {

	dgraphClient := newClient()

	ctx := context.Background()

	mu := &api.Mutation{
		CommitNow: true,
	}
	jsonBuyers, err := json.Marshal(buyers)
	if err != nil {
		 log.Fatal("failed to marshal ", err)
	}

	mu.SetJson = jsonBuyers
	assigned, err := dgraphClient.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal("failed to mutate ", err)
	}
	 
	print("res: %v", assigned)

}

