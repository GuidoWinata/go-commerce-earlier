    // Add Laptop
    product3, err3 := catalogB.CreateProduct("Laptop", 5_000_000, 20)
    if err3 != nil {
        fmt.Println(err3)
    } else {
        fmt.Println(product3)
    }

    // Get Kopi by Kopi ID
    kopi, err := catalogA.GetProductByID(1)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(kopi)
    }

    // Add Stock Kopi
    addStock, errAddStock := catalogA.AddStock(kopi.ID, 5)
    if errAddStock != nil {
        fmt.Println(errAddStock)
    } else {
        fmt.Println(addStock)
    }

    // Reduce Stock Kopi
    reduceStock, errReduceStock := catalogA.ReduceStock(kopi.ID, 3)
    if errReduceStock != nil {
        fmt.Println(errReduceStock)
    } else {
        fmt.Println(reduceStock)
    }

    // Invalid Reduce Stock Kopi
    invReduceStock, invErrReduceStock := catalogA.ReduceStock(kopi.ID, 100)
    if invErrReduceStock != nil {
        if errors.Is(invErrReduceStock, errInsufficientStock) {
            fmt.Println("Error: Tidak dapat mengurangi stock melebihi stock")
        } else {
            fmt.Println(invErrReduceStock)
        }
    } else {
        fmt.Println(invReduceStock)
    }

    // Get Kopi again
    getKopi, errGetKopi := catalogA.GetProductByID(1)
    if errGetKopi != nil {
        fmt.Println(errGetKopi)
    } else {
        fmt.Println(getKopi)
    }

    // Get Laptop by Laptop ID
    laptop, err := catalogB.GetProductByID(1)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(laptop)
    }

    // Invalid Add Stock Kopi
    invAddStock, invErrAddStock := catalogA.AddStock(kopi.ID, 0)
    if invErrAddStock != nil {
        fmt.Println(invErrAddStock)
    } else {
        fmt.Println(invAddStock)
    }

    // Get Invalid ID
    product, err := catalogA.GetProductByID(5)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(product)
    }