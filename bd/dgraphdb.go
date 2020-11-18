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

	op := &api.Operation{}
	op.Schema = `
		id: string @index(exact) .
		name: string .
		age: int .

		type Buyer {
			id
			name
			age
		}
	`

	ctx := context.Background()
	err := dgraphClient.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}

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


func SaveProducts(products []models.Product) {

	dgraphClient := newClient()
	op := &api.Operation{}
	op.Schema = `
		id: string @index(exact) .
		name: string .
		price: int .

		type Product {
			id
			name
			price
		}
	`
	ctx := context.Background()
	err := dgraphClient.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	jsonProducts, err := json.Marshal(products)
	if err != nil {
		 log.Fatal("failed to marshal ", err)
	}

	mu.SetJson = jsonProducts
	assigned, err := dgraphClient.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal("failed to mutate ", err)
	}
	 
	print("res: %v", assigned)

}


func SaveTransactions(transactions []models.Transaction) {

	dgraphClient := newClient()
	op := &api.Operation{}
	op.Schema = `
		id: string @index(exact) .
		buyer_id: string @index(exact) .
		ip: string @index(exact) .
		device: string .
		product_ids: [string] .

		type Transaction {
			id
			buyer_id
			ip
			device
			product_ids
		}
	`
	ctx := context.Background()
	err := dgraphClient.Alter(ctx, op)
	if err != nil {
		log.Fatal(err)
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	jsonTransactions, err := json.Marshal(transactions)
	if err != nil {
		 log.Fatal("failed to marshal ", err)
	}

	mu.SetJson = jsonTransactions
	assigned, err := dgraphClient.NewTxn().Mutate(ctx, mu)
	if err != nil {
		log.Fatal("failed to mutate ", err)
	}
	 
	print("res: %v", assigned)

}


func GetBuyers()([]models.Buyer, error) {

	type Data struct {
		Buyers []models.Buyer `json:"buyers"`
	}

	dgraphClient := newClient()

	ctx := context.Background()

	query := `
		{
    	buyers(func: has(age)){
    		id
      	name
				age
			}
		}
	`
		
	buyers, errQuery := dgraphClient.NewTxn().Query(ctx, query)
		
	var data Data

	err := json.Unmarshal(buyers.Json, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data.Buyers, errQuery
}

