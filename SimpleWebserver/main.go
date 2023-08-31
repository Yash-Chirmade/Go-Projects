package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Pages int `json:"pages"`
}

func hello(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","text/html")
	w.Write([]byte("<h1 style='color:steelblue'>Hello</h1>"))
}

func getBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	book:=book{
		Title: "The Go bible",
		Author: "Alan Donovan",
		Pages: 400,
	}
	json.NewEncoder(w).Encode(book)
}

func main(){
	http.HandleFunc("/hello",hello)
	http.HandleFunc("/book",getBook)
	log.Fatal(http.ListenAndServe(":5050",nil))
}