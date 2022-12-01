package main

import (
	"net/http"

	"github.com/MatsumaruTsuyoshi/books_api/controller"
	"github.com/MatsumaruTsuyoshi/books_api/model/repository"
)

var tr = repository.NewBookRepository()
var tc = controller.NewBookController(tr)
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todos/", ro.HandleBooksRequest)
	server.ListenAndServe()
}
