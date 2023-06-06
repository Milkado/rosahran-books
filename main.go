package main

import (
	"github/milkado/rosharan-books/routes"
	"net/http"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8000", nil)
}
