package main

import (
	"encoding/json"
	"net/http"

	"github.com/kysal/go-feed/db"
	"github.com/kysal/go-feed/routes"
)

type H map[string]any

func foo(w http.ResponseWriter, r *http.Request) {
	test := H{"message": "Hello world"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(test)
}

func main() {
	db.InitDB()

	http.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			routes.GetPosts(w, r)
		case http.MethodPost:
			routes.PublishPost(w, r)
		}
	})

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			routes.GetPostsPage(w, r)
		}
	})

	http.HandleFunc("/stylesheet.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/stylesheet.css")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/posts", http.StatusMovedPermanently)
	})

	http.ListenAndServe(":8080", nil)
}
