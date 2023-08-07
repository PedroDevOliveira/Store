package main

import (
	"Store/routes"
	"fmt"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	fmt.Print("Server is running on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
