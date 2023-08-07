package controllers

import (
	"Store/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	productList := models.GetProductList()
	temp.ExecuteTemplate(w, "Index", productList)
}
