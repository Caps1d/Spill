package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Caps1d/Spill/internal/construct"
	"github.com/Caps1d/Spill/internal/data"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"

	"net/http"
	"time"
)

// route handlers
func getAll(w http.ResponseWriter, r *http.Request, pm *data.PostModel) {
	// w.Write([]byte("Calling the db and getting the posts:\n"))
	log.Println("Calling the db and getting the posts")

	posts, err := pm.All()

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

func createPost(w http.ResponseWriter, r *http.Request, pm *data.PostModel) {

	var p construct.Post

	// decode the requests body into our post struct declared as p

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error decoding JSON: %s", err)
		return
	}

	log.Printf("Received post: %v", p)

	id, err := pm.Create(&p)

	if err != nil {
		log.Printf("Could not retrieve last inserted id: %s", err)
	} else {
		log.Printf("Inserted id: %d", id)
	}
}

func main() {
	//Note: checkout JSON marshall Go

	// db connection setup
	connectionStr := "user=yegorsmertenko password=postgres dbname=blogdb port=5432 sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)

	pm := &data.PostModel{DB: conn}

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}

	// optional but good to have
	conn.SetConnMaxLifetime(time.Minute * 5)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	// router setup with chi
	r := chi.NewRouter()

	// middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})

	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		getAll(w, r, pm)
	})

	r.Post("/posts", func(w http.ResponseWriter, r *http.Request) {
		createPost(w, r, pm)
	})

	http.ListenAndServe(":8080", r)

	conn.Close()
}
