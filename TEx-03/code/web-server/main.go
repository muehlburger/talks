package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Movie struct {
	Title  string
	Year   int `json:"released"`
	Actors []string
}

var movies = []Movie{
	{Title: "The Matrix", Year: 1996, Actors: []string{"Keanu Reeves", "Carrie Moss"}},
	{Title: "Matrix Reloaded", Year: 2003, Actors: []string{"Monica Bellucci", "Keanu Reeves"}},
	{Title: "The Rock", Year: 1999, Actors: []string{"Sean Connery", "Niclas Cage"}},
}

func main() {
	http.HandleFunc("/", path)
	http.HandleFunc("/movies", getMovies)
	http.ListenAndServe(":3000", nil)
}

// handler echoes the Path component of the requested URL.
func path(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
