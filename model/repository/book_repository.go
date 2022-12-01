package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/MatsumaruTsuyoshi/books_api/model/entity"
)

type BookRepository interface {
	GetBooks() (books []entity.BookEntity, err error)
	InsertBook(book entity.BookEntity) (id int, err error)
	UpdateBook(book entity.BookEntity) (err error)
	DeleteBook(id int) (err error)
	SearchBook(word string) (books []entity.BookEntity, err error)
}

type bookRepositry struct {
}

func NewBookRepository() BookRepository {
	return &bookRepositry{}
}

func (tr *bookRepositry) GetBooks() (books []entity.BookEntity, err error) {
	// todo update procces
	books = []entity.BookEntity{}
	rows, err := Db.Query("SELECT id, title FROM todo ORDER BY id DESC")
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		todo := entity.BookEntity{}
		err = rows.Scan(&todo.Id, &todo.Title)
		if err != nil {
			log.Print(err)
			return
		}
		books = append(books, todo)
	}

	return
}

func (tr *bookRepositry) InsertBook(book entity.BookEntity) (id int, err error) {
	// todo update
	return
}

func (tr *bookRepositry) UpdateBook(book entity.BookEntity) (err error) {
	// todo update
	return
}

func (tr *bookRepositry) DeleteBook(id int) (err error) {
	// todo update
	return
}

func (tr *bookRepositry) SearchBook(word string) (books []entity.BookEntity, err error) {
	// todo update
	return
}
