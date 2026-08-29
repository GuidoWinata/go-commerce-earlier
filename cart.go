package main

type Cart struct {
	ID         int
	CustomerID int
	Items      []CartItem
}

type CartItem struct {
	ProductID   int
	ProductName string
	Qty         int
	Price       int64
}

type InvoiceItem struct {
	ProductID   int
	ProductName string
	Qty         int
	SubTotal    int64
	UnitPrice   int64
}

type Invoice struct {
	Items []InvoiceItem
	Total int64
}

func (c *Cart) AddItem(productID int, qty int, catalog *CatalogLists) error {
	if qty <= 0 {
		return ErrInvalidQuantity
	}

	prod, err := catalog.GetProductByID(productID)
	if err != nil {
		return err
	}

	if !prod.Active {
		return ErrInactiveProduct
	}

	for i, item := range c.Items {
		if item.ProductID == productID {
			if item.Qty+qty > prod.Stock {
				return ErrInsufficientStock
			}
			c.Items[i].Qty += qty
			return nil
		}
	}

	if prod.Stock < qty {
		return ErrInsufficientStock
	}

	newItem := CartItem{
		ProductID:   productID,
		ProductName: prod.Name,
		Qty:         qty,
		Price:       prod.Price,
	}

	c.Items = append(c.Items, newItem)

	return nil
}

func (c *Cart) GetItems() []CartItem {
	return c.Items
}

func (c *Cart) Invoice() *Invoice {
	var items []InvoiceItem
	var total int64

	for _, item := range c.Items {
		subtotal := int64(item.Qty) * item.Price

		items = append(items, InvoiceItem{
			ProductID:   item.ProductID,
			ProductName: item.ProductName,
			Qty:         item.Qty,
			UnitPrice:   item.Price,
			SubTotal:    subtotal,
		})

		total += subtotal
	}

	return &Invoice{
		Items: items,
		Total: total,
	}
}
