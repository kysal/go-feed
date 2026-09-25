package routes

import (
	"fmt"
	"net/http"
)

func InitRoutes() {
	http.HandleFunc("/api/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		switch r.Method {
		case http.MethodDelete:
			DeletePost(w, r)
		}
	})

	http.HandleFunc("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetPosts(w, r)
		case http.MethodPost:
			PublishPost(w, r)
		}
	})

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetPostsPage(w, r)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/posts", http.StatusMovedPermanently)
	})

	http.HandleFunc("/stylesheet.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/stylesheet.css")
	})
}
