package main

import "fmt"

func main() {

	//Declaração explicita
	var cidade string = "São Paulo"
	fmt.Println("Eu moro na cidade:", cidade)

	//Declarando de forma implicita, tambem chamado de inferencia
	meuEstado := "Sp"
	fmt.Println("Meu estado é:", meuEstado)

	//Declarando mais de uma variavel
	var (
		paisAtual string = "Brasil"
		paisSonho string = "Inglaterra"
	)
	fmt.Println("Meu pais atual é: ", paisAtual)
	fmt.Println("Mas meu sonho é ir para: ", paisSonho)

	//Trocando o valor da variaveis
	variavel1 := 1
	variavel2 := 2
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	variavel1, variavel2 = variavel2, variavel1
	println("O Valor da Variavel1 é: ", variavel1, "E o valor da Variavel2 é:", variavel2)

	

	const nome string = "Giovanni"
	fmt.Println("E por fim meu nome é: ", nome)

}
