package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from snippetbox!"))
}

func testfunc(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Inside the testfunc mux"))
}

func showSnippet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Displaying a specific snippet..."))
}

func createSnippet(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost {
		w.Header().Set("Allow",http.MethodPost)
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Creating a new snippet..."))

}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/test",testfunc)
	mux.HandleFunc("/snippet",showSnippet)
	mux.HandleFunc("/snippet/create",createSnippet)
	log.Println("Starting server on port tcp/4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}