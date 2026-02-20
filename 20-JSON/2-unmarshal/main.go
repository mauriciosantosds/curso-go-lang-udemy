package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"-"`//tracinho caso queira ignorar a tag nome na conversão
	Raca  string `json:raca`
	Idade uint   `json:idade`
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","Raca":"Dálmata","Idade":3}`
	var c cachorro
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorro2EmJson := `{"nome": "Toby", "raca":"Poodle"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJson), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}