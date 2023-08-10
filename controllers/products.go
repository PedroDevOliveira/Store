package controllers

import (
	"Store/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	productList := models.GetProductList()
	temp.ExecuteTemplate(w, "Index", productList)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Fatal("Erro na conversão do preço")
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Fatal("Erro na conversão da quantidade")
		}

		models.CreateProduct(name, description, convertedPrice, convertedQuantity)

	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	convertedProductId, err := strconv.Atoi(productId)
	if err != nil {
		log.Fatal("Erro na conversão do Id")
	}
	models.DeleteProduct(convertedProductId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	convertedProductId, err := strconv.Atoi(productId)
	if err != nil {
		log.Fatal("Erro na conversão do Id")
	}
	product := models.EditProduct(convertedProductId)

	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		convertedId, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal("Erro na conversão do Id")
		}

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Fatal("Erro na conversão do preço")
		}

		convertedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Fatal("Erro na conversão da quantidade")
		}

		models.UpdateProduct(convertedId, name, description, convertedPrice, convertedQuantity)

	}

	http.Redirect(w, r, "/", 301)
}
