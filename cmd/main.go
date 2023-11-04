package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/billy-le/go-server/pkg/routes"
	"github.com/gorilla/mux"
)

var PORT = "localhost:8080"

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	fmt.Printf("Server running on PORT%s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}
