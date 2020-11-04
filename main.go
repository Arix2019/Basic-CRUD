/*
	CRUD Básico - GoLang
	author: -Arix-
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var PORT = ":3000"

// os campos precisam ser iniciados com letras maiusculas
type movie struct {
	Id       int    `json:"id"`       // Em Golang tds as iniciais dos campos struct
	Title    string `json:"title"`    //  precisam ser maiusculas
	Category string `json:"category"` // E devem ser exibidas em minusculas
	Year     int    `json:"year"`     //
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
	movie{
		Id:       2,
		Title:    "Mad Max",
		Category: "Action",
		Year:     1979,
	},
	movie{
		Id:       3,
		Title:    "IT",
		Category: "Terror",
		Year:     1990,
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

// exibe tds os resultados
func routeAllMovies(w http.ResponseWriter, r *http.Request) {
	// seta o header da página web para exibir o .json com a sintaxe correta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// cadastra novos filmes
func cadMovie(w http.ResponseWriter, r *http.Request) {
	// seta o header da página web para exibir o .json com a sintaxe correta
	w.Header().Set("Content-Type", "application/json")

	readBody, erro := ioutil.ReadAll(r.Body) // lê o corpo da requisição -POST-
	if erro != nil {
		fmt.Println(erro)
	}
	var newMovie movie
	json.Unmarshal(readBody, &newMovie) // converte o resultado para .json
	newMovie.Id = len(movies) + 1       // atribui o valor do ultimo item ao campo 'id'
	movies = append(movies, newMovie)   // insere na lista de filmes o novo conteudo adicionado
	json.NewEncoder(w).Encode(newMovie) // envia a resposta final para o 'cliente'
}

// testa os tipos de metodos (GET,POST...)
func testMethod(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		routeAllMovies(w, r)
	} else if r.Method == "POST" {
		cadMovie(w, r)
	}
}

func Routes() {

	http.HandleFunc("/", routeIndex)
	http.HandleFunc("/two", routeTwo)
	http.HandleFunc("/all-movies", testMethod)
}

func serverAndRoutes() {
	Routes()
	fmt.Printf("Go Server rodando na porta%v...\n", PORT)
	/*if http.ListenAndServe(PORT, nil) != nil {
		fmt.Printf("Ocorreu um erro ao tentar se conectar a porta %v!\nEssa porta já pode estar em uso.\n", PORT)
	}*/
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func main() {
	/*http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
	  fmt.Fprintf(w, "Testando terceira rota.")
	})*/
	serverAndRoutes()
}
