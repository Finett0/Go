package main

import (
	"fmt"
)

func main() {
	nome := "Giovanni"
	versao := 1.0
	fmt.Println("Olá sr ", nome)
	fmt.Println("Esté programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi: ", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo...")
	case 3:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Opção não encontrada.")

	}

}
