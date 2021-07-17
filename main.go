package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	rand2 "math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func removeMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item:= range movies{

		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item:= range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func addNewMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand2.Intn(1000000))

	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for index, item:= range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]

			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)

			return
		}
	}
}



func main()  {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "1216565", Title: "KGF 2", Director: &Director{Firstname: "Jhon", Lastname: "Raz"}})
	movies = append(movies, Movie{ID: "2", Isbn: "26525", Title: "Movie 2", Director: &Director{Firstname: "Test name", Lastname: "test name"}})


	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/addNew", addNewMovie).Methods("POST")
	r.HandleFunc("/update/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/remove/{id}", removeMovie).Methods("DELETE")


	fmt.Println("Starting server at Port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
	
}
