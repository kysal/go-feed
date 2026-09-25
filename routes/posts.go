package routes

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/kysal/go-feed/feed"
	"github.com/kysal/go-feed/models"
)

const pageTemplate = "html/template.html"

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := models.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
		return
	}
	if post.Content == "" {
		http.Error(w, "Post cannot be empty", http.StatusBadRequest)
		return
	}

	post.FeedId = feed.FeedId

	err = post.Publish()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Could not parse post id", http.StatusBadRequest)
		return
	}

	post := models.Post{Id: id}
	err = post.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetPostsPage(w http.ResponseWriter, r *http.Request) {
	posts, err := models.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	page, err := template.ParseFiles(pageTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = page.Execute(w, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
