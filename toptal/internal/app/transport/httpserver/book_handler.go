package httpserver

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/n340r/backend-notes/toptal/internal/app/common/server"
	"github.com/n340r/backend-notes/toptal/internal/app/domain"
)

func (h HttpServer) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, err := strconv.Atoi(vars["book_id"])
	if err != nil {
		server.BadRequest("invalid-book-id", err, w, r)
		return
	}
	book, err := h.bookService.GetBook(r.Context(), bookID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("book-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}
	response := toResponseBook(book)

	server.RespondOK(response, w, r)
}

func (h HttpServer) CreateBook(w http.ResponseWriter, r *http.Request) {
	var bookRequest BookRequest
	if err := json.NewDecoder(r.Body).Decode(&bookRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}

	if err := bookRequest.Validate(); err != nil {
		server.BadRequest("invalid-request", err, w, r)
		return
	}

	book, err := toDomainBook(bookRequest)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	insertedBook, err := h.bookService.CreateBook(r.Context(), book)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}
	response := toResponseBook(insertedBook)

	server.RespondOK(response, w, r)
}

func (h HttpServer) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, err := strconv.Atoi(vars["book_id"])
	if err != nil {
		server.BadRequest("invalid-book-id", err, w, r)
		return
	}

	var bookRequest BookRequest
	if err := json.NewDecoder(r.Body).Decode(&bookRequest); err != nil {
		server.BadRequest("invalid-json", err, w, r)
		return
	}
	if err := bookRequest.Validate(); err != nil {
		server.BadRequest("invalid-request", err, w, r)
		return
	}

	_, err = h.bookService.GetBook(r.Context(), bookID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("book-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	book, err := domain.NewBook(domain.NewBookData{
		ID:         bookID,
		Title:      bookRequest.Title,
		Year:       bookRequest.Year,
		Author:     bookRequest.Author,
		Price:      bookRequest.Price,
		CategoryID: bookRequest.CategoryID,
	})
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	updateBook, err := h.bookService.UpdateBook(r.Context(), book)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := toResponseBook(updateBook)

	server.RespondOK(response, w, r)
}

func (h HttpServer) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, err := strconv.Atoi(vars["book_id"])
	if err != nil {
		server.BadRequest("invalid-book-id", err, w, r)
		return
	}

	_, err = h.bookService.GetBook(r.Context(), bookID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("book-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	err = h.bookService.DeleteBook(r.Context(), bookID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			server.NotFound("book-not-found", err, w, r)
			return
		}
		server.RespondWithError(err, w, r)
		return
	}

	server.RespondOK(map[string]bool{"deleted": true}, w, r)
}

func (h HttpServer) GetBooks(w http.ResponseWriter, r *http.Request) {
	queryCategoryIDs := r.URL.Query()["category_id"]
	var categoryIDs []int
	for _, id := range queryCategoryIDs {
		categoryID, err := strconv.Atoi(id)
		if err != nil {
			server.BadRequest("invalid-category-id", err, w, r)
			return
		}
		categoryIDs = append(categoryIDs, categoryID)
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}
	var limit, offset int
	if page > 0 {
		limit = 10
		offset = (page - 1) * limit
	}
	books, err := h.bookService.GetBooks(r.Context(), categoryIDs, limit, offset)
	if err != nil {
		server.RespondWithError(err, w, r)
		return
	}

	response := make([]BookResponse, 0, len(books))
	for _, book := range books {
		response = append(response, toResponseBook(book))
	}

	server.RespondOK(response, w, r)
}
