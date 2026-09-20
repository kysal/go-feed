package routes

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/kysal/go-feed/models"
)

const pageTemplate = "html/template.html"

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := models.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func PublishPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if post.Content == "" {
		http.Error(w, "Post cannot be empty", http.StatusBadRequest)
	}

	// to be changed
	post.FeedId = 1

	err = post.Publish()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func GetPostsPage(w http.ResponseWriter, r *http.Request) {
	posts, err := models.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	page, err := template.ParseFiles(pageTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = page.Execute(w, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
