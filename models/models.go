package models

type Buyer struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}