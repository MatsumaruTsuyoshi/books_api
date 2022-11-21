package controller

import "net/http"

type Router interface {
	HandleBooksRequest(w http.ResponseWriter, r *http.Request)
}

type router struct {
	bc BookController
}

func NewRouter(bc BookController) Router {
	return &router{bc}
}

func (ro *router) HandleBooksRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.bc.GetBooks(w, r)
	case "POST":
		ro.bc.InsertBook(w, r)
	case "PUT":
		ro.bc.UpdateBook(w, r)
	case "DELETE":
		ro.bc.DeleteBook(w, r)
	default:
		w.WriteHeader(405)
	}
}
