package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var bookmarks []Bookmark
var nextID = 1

type Bookmark struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func Bookmarks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(bookmarks)

	case http.MethodPost:
		var bm Bookmark
		err := json.NewDecoder(r.Body).Decode(&bm)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		bm.Id = nextID
		nextID++
		bookmarks = append(bookmarks, bm)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bm)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func BookmarkDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/bookmarks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, bm := range bookmarks {
		if bm.Id == id {
			bookmarks = append(bookmarks[:i], bookmarks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Bookmark not found", http.StatusNotFound)
}
