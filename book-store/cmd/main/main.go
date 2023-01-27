package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abulyaev/goprojs/book-store/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))

	fmt.Println("Hello")

}
