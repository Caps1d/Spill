package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Caps1d/Spill/internal/construct"
)

type PostHandler struct {
	Model construct.PostService
}

func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Calling the db and getting the posts:\n"))
	log.Println("Calling the db and getting the posts")

	posts, err := h.Model.All()

	if err != nil {
		log.Printf("Error while getting rows from the db, err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}

	// for testing
	log.Printf("Posts array: %v", posts)

	// compose the json with the posts array
	js, err := json.Marshal(posts)

	if err != nil {
		log.Printf("Error while composing json, err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

	var p construct.Post

	// decode the requests body into our post struct declared as p

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding JSON: %s", err)
		return
	}

	log.Printf("Received post: %v", p)

	id, err := h.Model.Create(&p)

	if err != nil {
		log.Printf("Could not retrieve last inserted id: %s", err)
	} else {
		log.Printf("Inserted id: %d", id)
	}
}
