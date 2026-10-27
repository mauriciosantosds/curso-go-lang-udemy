package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		println("Erro ao converter o usuário para struct:", erro)
		w.Write([]byte("Erro ao converter o usuáiro para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao tentar conectar no banco de dados!"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		println("Erro ao criar o statement:", erro.Error())
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	//STATUS CODES

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Erro ao conectar com o banco de dados!", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		http.Error(w, "Erro ao buscar os usuários", http.StatusInternalServerError)
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}
	if erro := linhas.Err(); erro != nil {
		http.Error(w, "Erro ao buscar os usuários", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuário"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON!"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
