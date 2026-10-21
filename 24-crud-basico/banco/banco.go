package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

func Conectar() (*sql.DB, error) {
	urlConexao := "app_user:change_me_too@tcp(localhost:3306)/devbook?charset=utf8mb4&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", urlConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
