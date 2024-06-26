package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"graphql-template/graph/model"
	"time"
)

// Initialize and populate some sample data (replace with your actual logic)
func NewResolver() *Resolver {
	resolver := &Resolver{
		FoodStore:     make(map[string]model.Food),
		CustomerStore: make(map[string]model.Customer),
		OrderStore:    make(map[string]model.Order),
		InvoiceStore:  make(map[string]model.Invoice),
	}

	// Add some sample data
	resolver.FoodStore["1"] = model.Food{ID: "1", Name: "Pizza", Price: 9.99, Description: "Delicious pizza"}
	resolver.FoodStore["2"] = model.Food{ID: "2", Name: "Burger", Price: 7.99, Description: "Juicy burger"}

	resolver.CustomerStore["1"] = model.Customer{ID: "1", Name: "John Doe", Email: "john.doe@example.com"}
	resolver.CustomerStore["2"] = model.Customer{ID: "2", Name: "Jane Smith", Email: "jane.smith@example.com"}

	// Assume orders and invoices are related to customers and foods, populate them accordingly
	// ...

	return resolver
}

func (r *Resolver) Foods(ctx context.Context) ([]*model.Food, error) {
	var foods []*model.Food
	for _, food := range r.FoodStore {
		foods = append(foods, &food)
	}
	return foods, nil
}

func (r *Resolver) Food(ctx context.Context, id string) (*model.Food, error) {
	food, ok := r.FoodStore[id]
	if !ok {
		return nil, fmt.Errorf("food with ID %s not found", id)
	}
	return &food, nil
}

func (r *Resolver) Customers(ctx context.Context) ([]*model.Customer, error) {
	var customers []*model.Customer
	for _, customer := range r.CustomerStore {
		customers = append(customers, &customer)
	}
	return customers, nil
}

func (r *Resolver) Customer(ctx context.Context, id string) (*model.Customer, error) {
	customer, ok := r.CustomerStore[id]
	if !ok {
		return nil, fmt.Errorf("customer with ID %s not found", id)
	}
	return &customer, nil
}

func (r *Resolver) Orders(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order
	for _, order := range r.OrderStore {
		orders = append(orders, &order)
	}
	return orders, nil
}

func (r *Resolver) Order(ctx context.Context, id string) (*model.Order, error) {
	order, ok := r.OrderStore[id]
	if !ok {
		return nil, fmt.Errorf("order with ID %s not found", id)
	}
	return &order, nil
}

func (r *Resolver) Invoices(ctx context.Context) ([]*model.Invoice, error) {
	var invoices []*model.Invoice
	for _, invoice := range r.InvoiceStore {
		invoices = append(invoices, &invoice)
	}
	return invoices, nil
}

func (r *Resolver) Invoice(ctx context.Context, id string) (*model.Invoice, error) {
	invoice, ok := r.InvoiceStore[id]
	if !ok {
		return nil, fmt.Errorf("invoice with ID %s not found", id)
	}
	return &invoice, nil
}

func (r *Resolver) AddFood(ctx context.Context, name string, price float64, description string) (*model.Food, error) {
	// Generate a new ID (simplified for example, replace with proper ID generation logic)
	id := fmt.Sprintf("%d", time.Now().UnixNano())

	food := model.Food{
		ID:          id,
		Name:        name,
		Price:       price,
		Description: description,
	}

	r.FoodStore[id] = food
	return &food, nil
}

func (r *Resolver) UpdateFood(ctx context.Context, id string, name *string, price *float64, description string) (*model.Food, error) {
	food, ok := r.FoodStore[id]
	if !ok {
		return nil, fmt.Errorf("food with ID %s not found", id)
	}

	if name != nil {
		food.Name = *name
	}
	if price != nil {
		food.Price = *price
	}
	if description != "" {
		food.Description = description
	}

	r.FoodStore[id] = food
	return &food, nil
}

func (r *Resolver) DeleteFood(ctx context.Context, id string) (*bool, error) {
	_, ok := r.FoodStore[id]
	if !ok {
		return nil, fmt.Errorf("food with ID %s not found", id)
	}

	delete(r.FoodStore, id)
	return BoolPtr(true), nil
}

func (r *Resolver) AddCustomer(ctx context.Context, name string, email string, phone string) (*model.Customer, error) {
	// Generate a new ID (simplified for example, replace with proper ID generation logic)
	id := fmt.Sprintf("%d", time.Now().UnixNano())

	customer := model.Customer{
		ID:    id,
		Name:  name,
		Email: email,
		Phone: &phone,
	}

	r.CustomerStore[id] = customer
	return &customer, nil
}

func (r *Resolver) UpdateCustomer(ctx context.Context, id string, name *string, email *string, phone string) (*model.Customer, error) {
	customer, ok := r.CustomerStore[id]
	if !ok {
		return nil, fmt.Errorf("customer with ID %s not found", id)
	}

	if name != nil {
		customer.Name = *name
	}
	if email != nil {
		customer.Email = *email
	}
	if phone != "" {
		customer.Phone = &phone
	}

	r.CustomerStore[id] = customer
	return &customer, nil
}

func (r *Resolver) DeleteCustomer(ctx context.Context, id string) (*bool, error) {
	_, ok := r.CustomerStore[id]
	if !ok {
		return nil, fmt.Errorf("customer with ID %s not found", id)
	}

	delete(r.CustomerStore, id)
	return BoolPtr(true), nil
}

func (r *Resolver) AddOrder(ctx context.Context, customerID string, foodItemIds []*string) (*model.Order, error) {
	// Generate a new ID (simplified for example, replace with proper ID generation logic)
	id := fmt.Sprintf("%d", time.Now().UnixNano())

	customer, ok := r.CustomerStore[customerID]
	if !ok {
		return nil, fmt.Errorf("customer with ID %s not found", customerID)
	}

	var foodItems []*model.Food
	for _, foodID := range foodItemIds {
		food, ok := r.FoodStore[*foodID]
		if !ok {
			return nil, fmt.Errorf("food with ID %s not found", *foodID)
		}
		foodItems = append(foodItems, &food)
	}

	order := model.Order{
		ID:        id,
		Customer:  &customer,
		FoodItems: foodItems,
		Total:     calculateTotal(foodItems),
		Status:    "Pending", // Example status
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	r.OrderStore[id] = order
	return &order, nil
}

func (r *Resolver) UpdateOrder(ctx context.Context, id string, status string) (*model.Order, error) {
	order, ok := r.OrderStore[id]
	if !ok {
		return nil, fmt.Errorf("order with ID %s not found", id)
	}

	order.Status = status
	order.UpdatedAt = time.Now().Format(time.RFC3339)

	r.OrderStore[id] = order
	return &order, nil
}

func (r *Resolver) DeleteOrder(ctx context.Context, id string) (*bool, error) {
	_, ok := r.OrderStore[id]
	if !ok {
		return nil, fmt.Errorf("order with ID %s not found", id)
	}

	delete(r.OrderStore, id)
	return BoolPtr(true), nil
}

func (r *Resolver) AddInvoice(ctx context.Context, orderID string, totalAmount float64, paymentStatus string) (*model.Invoice, error) {
	// Generate a new ID (simplified for example, replace with proper ID generation logic)
	id := fmt.Sprintf("%d", time.Now().UnixNano())

	order, ok := r.OrderStore[orderID]
	if !ok {
		return nil, fmt.Errorf("order with ID %s not found", orderID)
	}

	invoice := model.Invoice{
		ID:            id,
		Order:         &order,
		TotalAmount:   totalAmount,
		PaymentStatus: paymentStatus,
		IssuedAt:      time.Now().Format(time.RFC3339),
	}

	r.InvoiceStore[id] = invoice
	return &invoice, nil
}

func (r *Resolver) UpdateInvoice(ctx context.Context, id string, paymentStatus string) (*model.Invoice, error) {
	invoice, ok := r.InvoiceStore[id]
	if !ok {
		return nil, fmt.Errorf("invoice with ID %s not found", id)
	}

	invoice.PaymentStatus = paymentStatus

	r.InvoiceStore[id] = invoice
	return &invoice, nil
}

func (r *Resolver) DeleteInvoice(ctx context.Context, id string) (*bool, error) {
	_, ok := r.InvoiceStore[id]
	if !ok {
		return nil, fmt.Errorf("invoice with ID %s not found", id)
	}

	delete(r.InvoiceStore, id)
	return BoolPtr(true), nil
}

// Helper function to convert bool to *bool
func BoolPtr(b bool) *bool {
	return &b
}

// Helper function to calculate total price of food items in an order
func calculateTotal(foodItems []*model.Food) float64 {
	var total float64
	for _, food := range foodItems {
		total += food.Price
	}
	return total
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
