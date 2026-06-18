package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	urlConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", urlConexao)
	if erro != nil {
		fmt.Println("Erro dentro do sql.Open")
		log.Fatal(erro)
	}

	defer db.Close()
	if erro = db.Ping(); erro != nil {
		fmt.Println("Erro dentro do db.ping")
		log.Fatal(erro)
	}
	fmt.Println("Conexão aberta!")
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()
	fmt.Println(linhas)
}
