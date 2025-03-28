package main

import (
	"fmt"
	"modulo/1-Pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

// Rode o comando go get <pacote externo> no terminal para importar o pacote

func main() {
	fmt.Println("Hello World!")
	auxiliar.Escrever()

	err := checkmail.ValidateFormat("123")
	if err == nil {
		fmt.Println("O email é Valido")
	} else {
		fmt.Println("O email é invalido")
	}

}
