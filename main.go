package main


import (
	"fmt"
	"log"
	"net/http"
	"enocoding/json"
	"math/rand"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"
	Lastname string `json:"lastname"`
}

var movies []Movie

// Returns a list of the avalible movies
func getMovies(w http.ResponseWriter, r "http.Request")  {
	w.Header().Set("Content-Type", "application/json")
	// Return the list of all the movies
	json.NewEncoder(w).Encoder(movies)
}

// Delete a movie from the movies list
func deleteMovie(w http.ResponseWriter, r "http.Request") {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// Return the list of the remaining movies after deletion
	json.NewEncoder(w).Encode(movies)
}

// 
func getMovie(w http.ResponseWriter, r "http.Request"){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range.movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}


func main() {
	r := mux.NewRouter()

	// Make these movies available on startup
	movies = append(movies, Movie(ID: "1", Isbn: "438227", Title: "Movie one", Director: &Director(Firstname: "Ilyas", Lastname: "Rufai")))
	movies = append(movies, Movie(ID: "2", Isbn: "438228", Title: "Movie two", Director: &Director(Firstname: "Maryam", Lastname: "Rufai")))
	
	// Routes functions
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/(id)", getMovie).Methods("GET")
	r.HandleFunc("/movies/", createMovie).Methods("POST")
	r.HandleFunc("/movies/(id)", updateMovie).Methods("PUT")
	r.HanldeFunc("/movies/(id)", deleteMovie).Methods("DELETE")

	// Starting the server
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}