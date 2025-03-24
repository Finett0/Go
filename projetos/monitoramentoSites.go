package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibirIntroducao()
	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo...")
		case 3:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Opção não encontrada.")
			os.Exit(-1)

		}

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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://www.binance.com/pt-BR",
		"https://www.rico.com.vc/", "https://www.coinbase.com/pt-br"}
	fmt.Println(sites)

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println(i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Minute)
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site: ", site, "está com problemas. Status Code: ", resp.StatusCode)
	}

}
