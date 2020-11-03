/*
	CRUD Básico - GoLang
	author: -Arix-
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var PORT = ":1337"

// os campos precisam ser iniciados com letras maiusculas
type movie struct {
	Id       int
	Title    string
	Category string
	Year     int
}

var movies = []movie{
	movie{
		Id:       0,
		Title:    "Die Hard",
		Category: "Action",
		Year:     1989,
	},
	movie{
		Id:       1,
		Title:    "The Matrix",
		Category: "Action",
		Year:     1999,
	},
}

// função para rota index: localhost:3000/
func routeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Olar, Index Page!</h1>")
}

// função para rota da segunda página: localhost:3000/two
func routeTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olar, Página dois!!")
}

// função exibe tds os resultados
func routeAllMovies(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(movies)
}

func Routes() {
	http.HandleFunc("/", routeIndex)
	http.HandleFunc("/two", routeTwo)
	http.HandleFunc("/all-movies", routeAllMovies)
}

func serverAndRoutes() {
	Routes()
	fmt.Println(movies)
	fmt.Printf("Go Server rodando na porta%v...", PORT)
	http.ListenAndServe(PORT, nil)

}

func main() {
	/*http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
	  fmt.Fprintf(w, "Testando terceira rota.")
	})*/
	serverAndRoutes()
}
