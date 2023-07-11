package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"Isbn"` 
	Title string `json:"title"`
	Director *Director `json: "director"`

}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies= append(movies,Movie{ID: "1", Isbn:"438227", Title: "Movie One", Director: &Director{Firstname: "Ram", Lastname: "Lakhan"}})
	movies= append(movies, Movie{ID: "2", Isbn:"438227", Title: "aaja Raat paryo Voli bihanai hai", Director: &Director{Firstname: "Laal", Lastname: "chadda"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))
}
