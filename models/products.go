package models

import (
	"Store/db"
	"Store/shared"
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

	selectAllProducts, err := dbConnection.Query("SELECT id, name, description, price, quantity, created_at as createdAt FROM go_store.public.products;")
	if err != nil {
		log.Fatal(err)
	}

	product := Product{}
	var productList []Product

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description, createdAt string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity, &createdAt)
		if err != nil {
			log.Fatal(err)
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity
		product.CreatedAt = shared.FormateDate(createdAt)

		productList = append(productList, product)
	}
	defer dbConnection.Close()
	return productList
}

func CreateProduct(name, description string, price float64, quantity int) {
	dbConnection := db.ConnectToPostgres()

	createProductData, err := dbConnection.Prepare("INSERT INTO go_store.public.products(name, description, price, quantity) values ($1, $2, $3, $4);")
	if err != nil {
		log.Fatal(err)
	}

	createProductData.Exec(name, description, price, quantity)
	defer dbConnection.Close()
}
