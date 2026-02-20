package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:raca`
	Idade uint   `json:idade`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)
	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJson)
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachrro2EmJson, erro := json.Marshal(c2)

	fmt.Println(cachrro2EmJson)
	fmt.Println(bytes.NewBuffer(cachrro2EmJson))
}