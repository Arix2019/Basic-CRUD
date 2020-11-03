/*
	CRUD Básico - GoLang
	author: -Arix-
*/
package main

import (
	"fmt"
	"net/http"
)

var PORT = ":3000"

// função para rota index: localhost:3000/
func routeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Olar, Index Page!</h1>")
}

// função para rota da segunda página: localhost:3000/two
func routeTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olar, Página dois!!")
}

func Routes() {
	http.HandleFunc("/", routeIndex)
	http.HandleFunc("/two", routeTwo)
}

func serverAndRoutes() {
	Routes()

	fmt.Printf("Go Server rodando na porta%v...", PORT)
	http.ListenAndServe(PORT, nil)

}

func main() {
  /*http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Testando terceira rota.")
  })*/ 
	serverAndRoutes()
}
