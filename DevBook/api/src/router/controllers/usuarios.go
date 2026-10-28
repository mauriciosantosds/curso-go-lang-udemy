package controllers

import "net/http"

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário!"))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuários!"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um Usuário!"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar Usuário!"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar Usuário!"))
}
