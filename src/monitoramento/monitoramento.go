package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 2
const delay = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido.")
			os.Exit(-1)
		}

		fmt.Println("")
	}
}

func exibeIntroducao() {
	nome := "Carol"
	versao := 1.1

	fmt.Println("Olá,", nome, "!")
	fmt.Println("Bem vindo ao programa de monitoramento")
	fmt.Println("Este programa está na versão:", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Println("Digite o comando desejado:")
	fmt.Scan(&comandoLido)

	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitoramento...")

	// sites := []string{
	// 	"https://random-status-code.herokuapp.com/",
	// 	"https://alura.com.br/",
	// 	"https://google.com.br/",
	// 	"http://facebook.com.br/"}
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(arquivo)

	return sites
}
