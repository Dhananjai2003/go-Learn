package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {

	http.HandleFunc("/bookmarks", Bookmarks)
	http.HandleFunc("/bookmarks/",BookmarkDelete)

	fmt.Println("Listing in port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}


