package main

import (
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
)

func main() {
	//JSON marshall Go
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

}
