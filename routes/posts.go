package routes

import "net/http"

func GetPosts(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Yet to be implemented", http.StatusInternalServerError)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Yet to be implemented", http.StatusNotFound)

}

func PublishPost(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Yet to be implemented", http.StatusNotFound)

}
