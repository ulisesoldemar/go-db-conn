package main

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Estos métodos sirven para obtner información en una carpeta
// y se busque la información de los templates
var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	// Solicitud de acceso a la función index
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)

	log.Println("Server running...")
	// Inicia el servidor
	http.ListenAndServe(":8080", nil)
}
