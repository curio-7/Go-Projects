package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"fn"`
	LastName  string `json:"ln"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json type
	w.Header().Set("Content-Type", "appplication/json")
	//params
	params := mux.Vars(r)
	//loop over movies

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(100000000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {

	r := mux.NewRouter()

	//we are not using database just having some mock data

	movies = append(movies, Movie{ID: "1", Isbn: "1451453", Title: "Sarthak Ki Kahani", Director: &Director{FirstName: "GOD", LastName: "THE GREAT"}})
	movies = append(movies, Movie{ID: "2", Isbn: "1452461453", Title: "Sarthak Ki Kahani 2.O", Director: &Director{FirstName: "GOD", LastName: "THE GREAT 2"}})
	movies = append(movies, Movie{ID: "3", Isbn: "1451452443", Title: "Sarthak Ki Kahani 3.O", Director: &Director{FirstName: "GOD", LastName: "THE GREAT 3"}})
	movies = append(movies, Movie{ID: "4", Isbn: "1451478153", Title: "Sarthak Ki Kahani 4.O", Director: &Director{FirstName: "GOD", LastName: "THE GREATn 4"}})

	//handle function will contain routes as present in diagram
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
