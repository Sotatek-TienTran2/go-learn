package main

import (
	"encoding/json"
	"go-sotatek/cmd/lession4/internal/data"
	"go-sotatek/cmd/lession4/internal/validator"
	"net/http"
	"time"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	var book *data.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()

	v.Check(book.Title != "", "title", "must be provided")

	if !v.Valid() {
		http.Error(w, v.ErrToString(), http.StatusBadRequest)
		return
	}
	book.ID = time.Now().Unix()
	ListBooks[book.ID] = book
	app.writeJSON(w, http.StatusOK, book, nil)
}

func (app *application) showBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	v := validator.New()

	v.Check(id != 0, "Id", "must be provided")

	if !v.Valid() {
		http.Error(w, v.ErrToString(), http.StatusBadRequest)
		return
	}

	err = app.writeJSON(w, http.StatusOK, ListBooks[id], nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (app *application) listBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := []*data.Book{}
	for _, book := range ListBooks {
		books = append(books, book)
	}
	err := app.writeJSON(w, http.StatusOK, books, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

func (app *application) updateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book *data.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()

	v.Check(book.Title != "", "title", "must be provided")
	v.Check(book.ID != 0, "id", "must be provided")

	if !v.Valid() {
		http.Error(w, v.ErrToString(), http.StatusBadRequest)
		return
	}
	ListBooks[book.ID] = book
	app.writeJSON(w, http.StatusOK, book, nil)
}

func (app *application) deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	v := validator.New()
	v.Check(id != 0, "id", "must be provided")

	delete(ListBooks, id)

	if !v.Valid() {
		http.Error(w, v.ErrToString(), http.StatusBadRequest)
		return
	}

	err = app.writeJSON(w, http.StatusOK, "Delete successfully!", nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
