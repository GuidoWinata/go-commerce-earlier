package main

import (
	"fmt"
)

func handleErr(err error, successMsg Product) {
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(successMsg)
	}
}

func main() {
	catalogA := NewProductCatalog()
	cart := &Cart{ID: 1, CustomerID: 1}

	// Add Kopi
	product1, err1 := catalogA.CreateProduct("Kopi", 5000, 10)
	handleErr(err1, product1)

	// Add Gula
	product2, err2 := catalogA.CreateProduct("Gula", 15000, 20)
	handleErr(err2, product2)
	// Add Item
	err := cart.AddItem(product1.ID, 2, catalogA)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Cart after adding item: %+v\n", cart)
	}

	err3 := cart.AddItem(product1.ID, 3, catalogA)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Printf("Cart after adding item: %+v\n", cart)
	}

	err4 := cart.AddItem(product1.ID, 6, catalogA)
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Printf("Cart after adding itssem: %+v\n", cart)
	}

	err5 := cart.AddItem(product2.ID, 3, catalogA)
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Printf("Cart after adding item: %+v\n", cart)
	}

	fmt.Println(cart.Invoice())
}
