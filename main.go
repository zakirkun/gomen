package main

import (

	// "github.com/gorilla/handlers"
	"gomen/core"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	args := os.Args[1:]
	core.Command(args)

	/*
		CORS ORIGIN HEADER.
		if you need this uncomment the code below, and uncomment "github.com/gorilla/handlers"
	*/

	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	// originsOk := handlers.AllowedOrigins([]string{"*"})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// set routes for http
	// http.Handle("/", web.SetRoutes())

	// set routes for socket
	// http.Handle("/socket.io/", socket.SetRoutes())

	// fmt.Println("Backend server started on port : 9000")

	// log.Fatal(http.ListenAndServe(":9000", nil))

	// log.Fatal(http.ListenAndServe(":9000", handlers.CORS(originsOk, headersOk, methodsOk)(routes.SetRoutes())))
}
