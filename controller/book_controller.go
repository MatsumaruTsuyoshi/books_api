package controller

import (
	"net/http"

	"github.com/MatsumaruTsuyoshi/books_api/model/repository"
)

type BookController interface {
	GetBooks(w http.ResponseWriter, r *http.Request)
	InsertBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
	SearchBook(w http.ResponseWriter, r *http.Request)
}

type bookController struct {
	br repository.BookRepository
}

func NewBookRepository(br repository.BookRepository) BookController {
	return &bookController{br}
}

func (bc *bookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	//todo write process

}

func (bc *bookController) InsertBook(w http.ResponseWriter, r *http.Request) {
	// todo write process
}

func (bc *bookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// todo write process
}

func (bc *bookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// todo write process
}

func (bc *bookController) SearchBook(w http.ResponseWriter, r *http.Request) {
	// todo write process
}
