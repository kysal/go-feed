package routes

import (
	"encoding/json"
	"net/http"

	"github.com/kysal/go-feed/models"
)

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
	// to be changed
	post.FeedId = 1

	err = post.Publish()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
