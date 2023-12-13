package main

import (
	"database/sql"
	"fmt"

	"github.com/Caps1d/Spill/internal/data"
	"github.com/Caps1d/Spill/internal/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"

	"net/http"
	"time"
)

// route handlers

func main() {
	//Note: checkout JSON marshall Go

	// db connection setup
	connectionStr := "user=yegorsmertenko password=postgres dbname=blogdb port=5432 sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)

	pm := &data.PostModel{DB: conn}

	ph := routes.PostHandler{Model: pm}

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

	r.Get("/posts", ph.GetAll)
	r.Post("/posts", ph.CreatePost)

	http.ListenAndServe(":8080", r)

	conn.Close()
}
