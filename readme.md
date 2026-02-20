# rodar container go

pela primeira vez rodar:

 ```docker run --rm -it --name curso-go-lang-udemy golang```

depois:

```docker run --rm -it --name go-restful -v .:/go/src/curso-go-lang-udemy golang```

para rodar os scripts:

```go run main.go```

comandos para teste go

testa todo o projeto
 ```go test ./...```

mostra cobertura dos testes
```go test --cover```
gera arquivo com cobertura dos testes
```go test --coverprofile cobertura.txt```
```tool cover --func=cobertura.txt```
gera html detalhado com cobertura dos testes
```go tool cover --html=cobertura.txt```
 mostra detalhes da cobertura do teste
 ```go test -v```