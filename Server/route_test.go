package main

import (
	"bytes"
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

	req, _ := http.NewRequest("GET", "/api/articles", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// fmt.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
}

func TestPostArticle(t *testing.T) {
	strs := `{"Title": "test1", "Author": "test2", "Link": "test3" "Description": "test4", "Published": "test5", "Doi": "testXX"}`

	router := SetupRouter()

	req, err := http.NewRequest("POST", "/api/article", bytes.NewBuffer([]byte(strs)))
	if err != nil {
		t.Fatalf("Error Occured: %v", err.Error())
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("Error Occured: key is %v %v", strs, w.Body)
	}
}
