package main

import (
	"fmt"
	"github.com/anmol/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	m := mux.NewRouter()
	routes.RegisterBookStoreRoutes(m)
	fmt.Println("Starting server at port 8010")
	if err := http.ListenAndServe(":8010", m); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
