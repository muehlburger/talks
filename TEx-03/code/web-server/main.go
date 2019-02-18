package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Movie stores data for a movie.
func main() {
	mux := http.DefaultServeMux
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/movies", movieHandler)

	fmt.Printf("The web-server is live at http://%s.", "127.0.0.1:3000")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))
}

type Movie struct {
	Title  string
	Year   int `json:"released"`
	Actors []string
}

var db = []Movie{
	{Title: "The Matrix", Year: 1996, Actors: []string{"Keanu Reeves", "Carrie Moss"}},
	{Title: "Matrix Reloaded", Year: 2003, Actors: []string{"Monica Bellucci", "Keanu Reeves"}},
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func movieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db)
}
