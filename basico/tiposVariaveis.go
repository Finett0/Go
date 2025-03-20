package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Giovanni" 
	idade := 20
	var altura float32 = 1.8 // se o tipo do float não for especificado ele sempre vai jogar para float64

	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade é", reflect.TypeOf(idade))
	fmt.Println("O tipo da variavel altura é", reflect.TypeOf(altura))

}
