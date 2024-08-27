package main

import (
	"go-sotatek/cmd/lession4/internal/data"
	"go-sotatek/cmd/lession4/internal/middlewares"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var ListBooks = map[int64]*data.Book{} // Declare a package-level map to act as our in-memory data store.

func (app *application) routes() http.Handler {
	ListBooks = make(map[int64]*data.Book) // Initialize an empty map to store the books.
	// Initialize a new httprouter router instance.
	router := httprouter.New()
	// Register the relevant methods, URL patterns and handler functions for our
	// endpoints using the HandlerFunc() method. Note that http.MethodGet and
	// http.MethodPost are constants which equate to the strings "GET" and "POST"
	// respectively.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/books", app.listBooksHandler)
	router.HandlerFunc(http.MethodPost, "/v1/books", app.createBookHandler)
	router.HandlerFunc(http.MethodGet, "/v1/books/:id", app.showBookHandler)
	router.HandlerFunc(http.MethodPut, "/v1/books/:id", app.updateBookHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/books/:id", app.deleteBookHandler)
	// Return the httprouter instance.

	// cors
	// logging
	return middlewares.RecoverPanicMiddleware(middlewares.RateLimiterMiddleware(middlewares.LoggingMiddleware(router)))
}
