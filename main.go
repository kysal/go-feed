package main

import (
	"fmt"
	"net/http"

	"github.com/kysal/go-feed/db"
	"github.com/kysal/go-feed/feed"
	"github.com/kysal/go-feed/routes"
)

func main() {

	err := feed.ParseFeed()
	if err != nil {
		fmt.Println(err)
		return
	}

	db.InitDB()

	routes.InitRoutes()

	http.ListenAndServe(":8080", nil)
}
