package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()
	// registraLog("site-falso", false)

	for {
		// _, age := devolveNomeIdade()
		// fmt.Println(age)

		exibeMenu()

		command := leComando()

		// if command == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if command == 2 {
		// 	fmt.Println("Exibindo logs...")
		// } else if command == 0 {
		// 	fmt.Println("Saindo do programa.")
		// } else {
		// 	fmt.Println("Não conheço este comando.")
		// }

		switch command {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa.")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando.")
			os.Exit(-1)
		}
	}
}

func devolveNomeIdade() (string, int) {
	name := "Jackson"
	age := 27

	return name, age
}

func exibeIntroducao() {
	// var name string = "Jackson"
	// var idade int = 27
	// var version float32 = 1.1

	name := "Jackson"
	idade := 27
	version := 1.1

	fmt.Println("Hello World!", name)
	fmt.Println("Olá", name, "sua idade é", idade)
	fmt.Println("Este programa está na versão", version)
	fmt.Println("O tipo da variável name é:", reflect.TypeOf(name))
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var command int

	// fmt.Scanf("%d", &command)
	fmt.Scan(&command)
	fmt.Println("O endereço da minha variavel command é", &command)
	fmt.Println("O command escolhido foi", command)

	return command
}

func iniciarMonitoramento() {
	// site := "http://www.alura.com.br"
	// sites := []string{
	// 	"http://www.alura.com.br", "http://www.alura.com.br"}
	sites := lerSitesDoArquivo()

	fmt.Println("Monitorando...")

	// FOR antigo.
	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println((sites[i]))
	// }

	for i := 0; i < monitoramentos; i++ {
		for indice, site := range sites {
			// fmt.Println("Estou passando na posição", indice, "do meu slice e essa posicao tem o site: ", site)
			fmt.Println("Testando site", indice, ": ", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
	}
}

func exibeArray() {
	var sites [4]string
	sites[0] = "http://www.alura.com.br"
	sites[1] = "http://www.alura.com.br"
	sites[2] = "http://www.alura.com.br"
	sites[3] = "http://www.alura.com.br"

	fmt.Println("sitesArray: ", sites, reflect.TypeOf(sites))
	fmt.Println("O tamanho do array é: ", len(sites))
}

func exibeSlice() {
	nomes := []string{"Jackson", "Eu"}
	nomes = append(nomes, "Vós")

	fmt.Println("nomesSlice:", nomes, reflect.TypeOf(nomes))
	fmt.Println("O tamanho do slice é: ", len(nomes))
	fmt.Println("O capacidade do slice é: ", cap(nomes))
}

func testaSite(site string) {
	// resp, _ := http.Get(site)
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorrou um erro: ", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func lerSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println("Arquivo: ", arquivo)

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		fmt.Println("Linha:", linha)

		sites = append(sites, linha)

		if err == io.EOF {
			// fmt.Println("Ocorreu um erro: ", err)
			break
		}
	}

	fmt.Println("Sites:", sites)

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	// arquivo, err := os.Open("log.txt")
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(arquivo)
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	fmt.Println("Exibindo logs...")

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println("Arquivo:", string(arquivo))
}
