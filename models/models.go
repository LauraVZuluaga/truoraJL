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

type Transaction struct {
	ID string `json:"id"`
	BuyerID string `json:"buyer_id"`
	Ip string `json:"ip"`
	Device string `json:"device"`
	ProductIDs [] string `json:"product_ids"`
}