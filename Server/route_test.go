package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Article struct {
	Id          int64
	Title       string
	Author      string
	Link        string
	Description string
	Published   string
	Doi         string
}

func TestGetArticlesAll(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/articles", nil)
	router.ServeHTTP(w, req)

	// fmt.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
}
