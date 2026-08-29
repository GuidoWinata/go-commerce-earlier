package main

import (
	"errors"
	"strings"
)

var (
	ErrProductNotFound   = errors.New("product not found")
	ErrInactiveProduct   = errors.New("product was not activated anymore")
	ErrEmptyProductName  = errors.New("product name cannot be empty")
	ErrInvalidPrice      = errors.New("price must be greater than 0")
	ErrInvalidStock      = errors.New("stock must be greater or equal than 0")
	ErrInvalidID         = errors.New("product ID must be greater than 0")
	ErrInvalidQuantity   = errors.New("quantity must be greater than 0")
	ErrInsufficientStock = errors.New("insufficient stock: cannot reduce below zero")
	ErrSoldStock         = errors.New("Stock was sold")
)

type Product struct {
	ID     int
	Name   string
	Price  int64
	Stock  int
	Active bool
}

type CatalogLists struct {
	nextID   int
	products map[int]Product
}

func NewProductCatalog() *CatalogLists {
	return &CatalogLists{
		products: make(map[int]Product),
		nextID:   1,
	}
}

func (p *CatalogLists) CreateProduct(name string, price int64, stock int) (Product, error) {

	cleanName := strings.TrimSpace(name)

	if cleanName == "" {
		return Product{}, ErrEmptyProductName
	}

	if price <= 0 {
		return Product{}, ErrInvalidPrice
	}

	if stock < 0 {
		return Product{}, ErrInvalidStock
	}

	newProduct := Product{
		ID:     p.nextID,
		Name:   cleanName,
		Price:  price,
		Stock:  stock,
		Active: true,
	}

	p.products[newProduct.ID] = newProduct

	p.nextID++

	return newProduct, nil
}

func (p *CatalogLists) ReduceStock(productID int, qty int) (Product, error) {
	if productID <= 0 {
		return Product{}, ErrInvalidID
	}

	if qty <= 0 {
		return Product{}, ErrInvalidQuantity
	}

	product, exist := p.products[productID]

	if !exist {
		return Product{}, ErrProductNotFound
	}

	if product.Stock < qty {
		return Product{}, ErrInsufficientStock
	}

	product.Stock -= qty

	p.products[productID] = product

	return product, nil
}

func (p *CatalogLists) AddStock(productID int, qty int) (Product, error) {

	if productID <= 0 {
		return Product{}, ErrInvalidID
	}

	if qty <= 0 {
		return Product{}, ErrInvalidQuantity
	}

	product, exist := p.products[productID]

	if !exist {
		return Product{}, ErrProductNotFound
	}

	product.Stock += qty

	p.products[productID] = product

	return product, nil
}

func (p CatalogLists) GetProductByID(productID int) (Product, error) {

	if productID <= 0 {
		return Product{}, ErrInvalidID
	}

	product, exist := p.products[productID]
	if !exist {
		return Product{}, ErrProductNotFound
	}

	return product, nil
}

func (p CatalogLists) GetAllProducts() []Product {
	var allProducts []Product
	for _, product := range p.products {
		allProducts = append(allProducts, product)
	}

	return allProducts
}
