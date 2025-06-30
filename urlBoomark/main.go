package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var bookmark []Bookmark
var nextId = 1

type Bookmark struct {
	Id int
	Name string
	Url string
}

func main() {

	http.HandleFunc("/bookmarks", bookmarks)
	http.HandleFunc("/bookmarks/",bookmarkDelete)

	fmt.Println("Listing in port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func bookmarks(w http.ResponseWriter, r *http.Request){

	switch  r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(bookmark)
	case http.MethodPost:
		var bm Bookmark
		err := json.NewDecoder(r.Body).Decode(&bm)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		
		bm.Id = nextId
		nextId++
		bookmark = append(bookmark,bm)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(bm)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func bookmarkDelete(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodDelete {
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/bookmarks/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w,"not correct ID",http.StatusBadRequest)
		return
	}

	for i, bm := range bookmark {
		if bm.Id == id {
			bookmark = append(bookmark[:i],bookmark[i+1:]...)
			return
		}
	}


}