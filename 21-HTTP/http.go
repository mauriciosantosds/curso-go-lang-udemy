package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("../22-HTML/*.html"))
	//HTTP é um protocolo de comunicação - base da comunicação da web

	//Cliente (faz requisição) - Servidor (processão requisição e envia resposta)

	// Request - Response

	// Rotas
	//URI - Identificador do Recurso
	//Método - GET, POST, PUT, DELETE
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{
			Nome:  "João",
			Email: "joaopedro@gmail.com",
		}
		templates.ExecuteTemplate(w, "home.html", u)
	})
	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pagina usuários\n\n"))
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}
