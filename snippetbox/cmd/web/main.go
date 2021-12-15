package main

import (
	"log"
	"net/http"
)

var serverPort string = ":4000"

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/snippet",showSnippet)
	mux.HandleFunc("/snippet/create",createSnippet)

	log.Printf("Starting server on",serverPort)
	err := http.ListenAndServe(serverPort,mux)
	log.Fatal(err)
}



