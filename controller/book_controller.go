package controller

import (
	"encoding/json"
	"net/http"

	"github.com/MatsumaruTsuyoshi/books_api/controller/dto"
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

func NewBookController(br repository.BookRepository) BookController {
	return &bookController{br}
}

func (bc *bookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	todos, err := bc.br.GetBooks()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var bookResponses []dto.BookResponse
	for _, v := range todos {
		bookResponses = append(bookResponses, dto.BookResponse{Id: v.Id, Title: v.Title})
	}

	var booksResponses dto.BooksResponse
	booksResponses.Books = bookResponses

	output, _ := json.MarshalIndent(booksResponses.Books, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

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
