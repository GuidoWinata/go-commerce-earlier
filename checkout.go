package main

import (
	"math/rand/v2"
)

var (
	CompleteMsg  = "Completed"
	PendingMsg   = "Pending"
	CancelledMsg = "Cancelled"
)

type Order struct {
	OrderID    uint64
	CustomerID int
	Total      int64
	Status     string
	Items      []OrderItem
}

type OrderItem struct {
	Qty         int
	UnitPrice   int64
	ProductName string
	Subtotal    int64
}

func Checkout(customerID int, c *Cart, catalog *CatalogLists) (*Order, error) {
	var items []OrderItem
	orderID := rand.Uint64()
	var total int64

	for _, item := range c.Items {
		product, err := catalog.GetProductByID(item.ProductID)
		if err != nil {
			return nil, err
		}

		if !product.Active {
			return nil, ErrInactiveProduct
		}

		if product.Stock < item.Qty {
			return nil, ErrSoldStock
		}

		subtotal := int64(item.Qty) * product.Price

		items = append(items, OrderItem{
			Qty:         item.Qty,
			UnitPrice:   product.Price,
			ProductName: product.Name,
			Subtotal:    subtotal,
		})

		product.Stock -= item.Qty
		total += subtotal
	}

	return &Order{
		OrderID:    orderID,
		CustomerID: customerID,
		Total:      total,
		Status:     PendingMsg,
		Items:      items,
	}, nil

}
