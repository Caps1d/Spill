package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html"
	"log"
	"net"
	"net/http"
)

func main() {
	//Note: checkout JSON marshall Go

	// db connection setup
	connectionStr := "user=yegorsmertenko password=postgres dbname=blogdb port=5432 sslmode=disable"

	conn, err := sql.Open("postgres", connectionStr)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database")
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		// will output a simple html text line saying:
		// "Welcome to /home"
		fmt.Fprint(w, "Welcome to ", html.EscapeString(r.URL.Path))
		// fmt.Fprint(w, "Welcome, to my server!")
	})
	log.Println("Server starting...")

	listener, err := net.Listen("tcp", ":8080")

	if err == nil {
		fmt.Println("Server is online, listening on port:8080")
	} else {
		log.Fatal(err)
	}
	// we don't need to specify a handler since we registered a handler in HandleFunc 2nd arg
	// hence we pass nil
	http.Serve(listener, nil)

	// get rows from db post table and handle get post route
	rows, err := conn.Query("SELECT * FROM post;")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var title, content, date string
		var authorId int

		if err := rows.Scan(&id, &title, &content, &authorId, &date); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf(" ID: %d,\n Title: %s,\n Content: %s,\n AuthorId: %d,\n Date: %s\n", id, title, content, authorId, date)
	}

	// handle errors from rows.Next()
	if err := rows.Err(); err != nil {
		fmt.Println("Error during iterating over rows:", err)
	}

	// close rows explicitly
	if err := rows.Close(); err != nil {
		fmt.Println("Error closing rows:", err)
	}

	// don't forget to close the db connection
	conn.Close()
}
