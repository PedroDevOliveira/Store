package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	productList := []Product{
		{Name: "Camiseta", Description: "Azul bem bonita", Price: 39, Quantity: 5},
		{"Tenis", "Confortável", 199, 10},
		{"Fone", "Muito bom", 866, 12},
	}

	temp.ExecuteTemplate(w, "Index", productList)
}
