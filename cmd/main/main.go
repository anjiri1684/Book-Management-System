package main

import (
	"fmt"
	"net/http"

	"github.com/anjiri1684/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

const s string = "constant"

func main() {

	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	err := http.ListenAndServe(":80", router)
	if err != nil {
		fmt.Println(err)
	}
}