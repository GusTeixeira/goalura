package main

import (
	"fmt"
)

func main() {
	//Aula 1 - Introdução
	fmt.Println("Hello World")
	//Aula 2 - Váriaveis
	var nome string = "Douglas"
	var versao float32 = 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão: ", versao)

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")

	var comando int
	fmt.Scanf("%d", &comando)

	fmt.Println("O comando escolhido foi", comando)
}
