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

	selectAllProducts, err := dbConnection.Query("SELECT id, name, description, price, quantity, created_at AS createdAt FROM go_store.public.products ORDER BY id ASC;")
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

func DeleteProduct(id int) {
	dbConnection := db.ConnectToPostgres()
	deleteProduct, err := dbConnection.Prepare("DELETE FROM go_store.public.products WHERE id=$1;")
	if err != nil {
		log.Fatal(err)
	}

	deleteProduct.Exec(id)

	defer dbConnection.Close()

}

func EditProduct(id int) Product {
	dbConnection := db.ConnectToPostgres()
	productFromDB, err := dbConnection.Query("SELECT * FROM go_store.public.products WHERE id=$1;", id)
	if err != nil {
		log.Fatal(err)
	}

	updatedProduct := Product{}

	for productFromDB.Next() {
		var id, quantity int
		var name, description, createdAt string
		var price float64

		err := productFromDB.Scan(&id, &name, &description, &price, &quantity, &createdAt)
		if err != nil {
			log.Fatal(err)
		}

		updatedProduct.Id = id
		updatedProduct.Name = name
		updatedProduct.Description = description
		updatedProduct.Price = price
		updatedProduct.Quantity = quantity

	}

	defer dbConnection.Close()
	return updatedProduct
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	dbConnection := db.ConnectToPostgres()

	updateProductData, err := dbConnection.Prepare("UPDATE go_store.public.products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5;")
	if err != nil {
		log.Fatal(err)
	}

	updateProductData.Exec(name, description, price, quantity, id)
	defer dbConnection.Close()
}
