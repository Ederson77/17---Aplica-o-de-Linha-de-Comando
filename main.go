package main

import (
	"fmt"
	"Linha-de-comando/app"
)

func main() {
	fmt.Println("Ponto de Partida")

	aplicacao := app.Gerar()	
	fmt.Println(aplicacao)
}