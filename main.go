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

// rota index: localhost:3000/
func routeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olar, Index Page!")
}

// rota para a segunda página: localhost:3000/two
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

	serverAndRoutes()
}
