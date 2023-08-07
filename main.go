package main

import (
	"Store/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	productList := models.GetProductList()
	temp.ExecuteTemplate(w, "Index", productList)
}
