package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Inicio)
	http.ListenAndServe(":8080", nil)
}

// (mostrar informacion, recibir la peticion)
func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola perro")

}
