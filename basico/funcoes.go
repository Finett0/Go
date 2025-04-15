package main

import (
	"fmt"
	"os"
)

func main() {

	exibirIntroducao()
	exibeMenu()

	comando := lerComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo..")
	case 3:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Opção não encontrada.")
		os.Exit(-1)

	}

}

func exibirIntroducao() {
	nome := "Giovanni"
	versao := 1.1
	fmt.Println("Olá sr ", nome)
	fmt.Println("Esté programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- sair do programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi: ", comandoLido)

	return comandoLido

}
