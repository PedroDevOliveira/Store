package models

import (
	"Store/db"
	"log"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
	CreatedAt   string
}

func GetProductList() []Product {
	dbConnection := db.ConnectToPostgres()

	selectAllProducts, err := dbConnection.Query("select * from products")
	if err != nil {
		log.Fatal(err)
	}

	product := Product{}
	productList := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description, create_at string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity, &create_at)
		if err != nil {
			log.Fatal(err)
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.CreatedAt = create_at

		productList = append(productList, product)
	}
	defer dbConnection.Close()
	return productList
}
