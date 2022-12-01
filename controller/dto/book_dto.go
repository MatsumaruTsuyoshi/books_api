package dto

type BookResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type BookRequest struct {
	Title string `json:"title"`
}

type BooksResponse struct {
	Books []BookResponse `json:"todos"`
}
