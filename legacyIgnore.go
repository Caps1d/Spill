package main

// home route handler - standard lib setup
// http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
// 	// will output a simple html text line saying:
// 	// "Welcome to /home"
// 	fmt.Fprint(w, "Welcome to ", html.EscapeString(r.URL.Path))
// 	// fmt.Fprint(w, "Welcome, to my server!")
// })
// log.Println("Server starting...")
//
// listener, err := net.Listen("tcp", ":8080")
//
// if err == nil {
// 	fmt.Println("Server is online, listening on port:8080")
// } else {
// 	log.Fatal(err)
// }
//
// // posts route handler: only GET and POST requests
// // get rows from db post table and handle get requests sent to post route
// // this is a standard lib setup
// http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/posts" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 	}
//
// 	switch r.Method {
// 	case "GET":
// 		rows, err := conn.Query("SELECT * FROM post;")
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		for rows.Next() {
// 			var id int
// 			var title, content, date string
// 			var authorId int
//
// 			if err := rows.Scan(&id, &title, &content, &authorId, &date); err != nil {
// 				fmt.Println(err)
// 				continue
// 			}
// 			fmt.Printf(" ID: %d,\n Title: %s,\n Content: %s,\n AuthorId: %d,\n Date: %s\n", id, title, content, authorId, date)
//
// 			// display the returned rows
// 			fmt.Fprintf(w, " ID: %d,\n Title: %s,\n Content: %s,\n AuthorId: %d,\n Date: %s\n", id, title, content, authorId, date)
// 		}
//
// 		// handle errors from rows.Next()
// 		if err := rows.Err(); err != nil {
// 			fmt.Println("Error during iterating over rows:", err)
// 		}
//
// 		// close rows explicitly
// 		if err := rows.Close(); err != nil {
// 			fmt.Println("Error closing rows:", err)
// 		}
//
// 	case "POST":
// 		if err := r.ParseForm(); err != nil {
// 			fmt.Fprintf(w, "ParseForm() err: %v", err)
// 			return
// 		}
//
// 		fmt.Fprintf(w, "Post from website, r.PostForm = %v\n", r.PostForm)
//
// 		title := r.FormValue("title")
// 		content := r.FormValue("content")
// 		useridString := r.FormValue("userid")
//
// 		// Convert the string to int64 because post table expects an int
// 		userid, err := strconv.ParseInt(useridString, 10, 64)
// 		if err != nil {
// 			// Handle the error (e.g., invalid input)
// 			log.Printf("Error converting userid to int64: %s", err)
// 		}
// 		// printing the FormValue's
// 		fmt.Fprintf(w, "Post title: %s is of type %T\n", title, title)
// 		fmt.Fprintf(w, "Content: %s is of type %T\n", content, content)
// 		fmt.Fprintf(w, "Created by: %d is of type %T\n", userid, userid)
//
// 		query := "INSERT INTO post (title, content, userId) VALUES ($1, $2, $3) Returning id;"
// 		// QueryRowContext runs a query that returns 1 row at most
// 		row := conn.QueryRowContext(context.Background(), query, title, content, userid)
//
// 		// since post table's pk is serial, its useful to keep track of the last
// 		// inserted key
// 		var insertedID int64
// 		err = row.Scan(&insertedID)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				log.Println("No rows were inserted.")
// 			} else {
// 				log.Fatalf("Could not retrieve last inserted id: %s", err)
// 			}
// 		} else {
// 			fmt.Printf("Inserted id: %d", insertedID)
// 		}
//
// 	default:
// 		fmt.Fprintf(w, "Sorry, this router only handles GET and POST requests ='(")
// 	}
//
// })
// // we don't need to specify a handler since we registered a handler in HandleFunc 2nd arg
// // hence we pass nil
// http.Serve(listener, nil)

// don't forget to close the db connection
