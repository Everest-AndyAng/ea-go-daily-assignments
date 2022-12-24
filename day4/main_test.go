package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/hello", nil)

	helloHandler(rec, r)

	body := rec.Result().Body
	data, _ := io.ReadAll(body)

	assert.Equal(t, string(data), "hello")
}

func TestBookHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/book", nil)

	bookHandler(rec, r)

	body := rec.Result().Body
	data, _ := io.ReadAll(body)

	var b Book
	json.Unmarshal(data, &b)
	expected := Book{1, "Learn Go", 19.90}
	assert.Equal(t, b, expected)
}

func TestGetBooksRouterHandler(t *testing.T) {
	router := setupGin()
	setupGetBooksRouter(router)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	router.ServeHTTP(w, req)

	body := w.Result().Body
	data, _ := io.ReadAll(body)
	var actual []Book
	json.Unmarshal(data, &actual)

	var expected []Book
	expected = append(expected, Book{1, "Learn Go", 19.90})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, actual, expected)
}

func TestDeleteBooksRouterHandler(t *testing.T) {
	router := setupGin()
	setupDeleteBooksRouter(router)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// ToDo: delete books (better to separate as another test case but will need to use verify)
}
