package fixtures

import "math/rand"

var username = "admin"
var password = "admin1234"

type Product struct {
	ID    int
	Title string
}

var Products = []Product{
	{
		ID:    1,
		Title: "Delivery fee",
	},
	{
		ID:    2,
		Title: "Commission",
	},
	{
		ID:    3,
		Title: "Product 3",
	},
	{
		ID:    4,
		Title: "Product 4",
	},
	{
		ID:    5,
		Title: "Product 5",
	},
	{
		ID:    6,
		Title: "Product 6",
	},
	{
		ID:    7,
		Title: "Product 7",
	},
	{
		ID:    8,
		Title: "Product 8",
	},
	{
		ID:    9,
		Title: "Product 9",
	},
	{
		ID:    10,
		Title: "Product 10",
	},
}

func RandomProduct() Product {
	return Products[rand.Intn(len(Products))]
}

var DepositSources = []string{
	"credit_card",
	"cash",
	"bank_transfer",
	"mobile_payment",
	"other",
	"bonus",
}

func RandomDepositSource() string {
	return DepositSources[rand.Intn(len(DepositSources))]
}
