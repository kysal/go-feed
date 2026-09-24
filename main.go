package main

import (
	"net/http"

	"github.com/kysal/go-feed/db"
	"github.com/kysal/go-feed/routes"
)

func main() {
	db.InitDB()

	routes.InitRoutes()

	http.ListenAndServe(":8080", nil)
}
