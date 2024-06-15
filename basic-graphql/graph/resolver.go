package graph

import "graphql-template/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver functions (in resolver.go)

type Resolver struct {
	FoodStore     map[string]model.Food
	CustomerStore map[string]model.Customer
	InvoiceStore  map[string]model.Invoice
	OrderStore    map[string]model.Order
}
