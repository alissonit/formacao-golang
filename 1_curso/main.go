package main

import (
	"io"
	"time"
	"fmt"
	"os"
	"net/http"
	"bufio"
	"strings"
	"strconv"
)

const monitors = 3
const delay = 0

func readSitesFromFile() []string {

	var sites[]string

	file, err := os.Open("sites.txt")

	if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

	read := bufio.NewReader(file)

	for {
		line, err := read.ReadString('\n')

		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}

		fmt.Println(line)
	}
	
	file.Close()

	return sites
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("ocorreu um erro: ", err)
	}
	if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
		logRegister(site, true)
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		logRegister(site, false)
    }
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	// sites := []string{
	// 	"https://random-status-code.herokuapp.com/", 
	// 	"https://www.alura.com.br",
	// 	"https://www.caelum.com.br",
	// }

	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(sites)
	// }

	sites := readSitesFromFile()
	
	for i := 0; i < monitors; i++ {
		for i, site := range sites {
			fmt.Println("Testing", i, ":", site)
			testSite(site)
		}
		time.Sleep(delay * time.Minute)
	}
}

func useSwitch(command int) {

	switch command {
    case 1:
        startMonitoring()
    case 2:
        fmt.Println("Exibindo Logs...")
		getLogs()
    case 0:
        fmt.Println("Saindo do programa...")
		os.Exit(0)
    default:
        fmt.Println("Não conheço este command")
		os.Exit(1)
    }
}

func displayIntro() {
	name := "Douglas"
    version := 1.1
    fmt.Println("Olá, sr.", name)
    fmt.Println("Este programa está na versão", version)
}

func readCommand() int {

	var commandScan int

    fmt.Scan(&commandScan)
    fmt.Println("O command escolhido foi", commandScan)

	return commandScan
}

func displayMenu() {
	fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}

func logRegister(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("ocorreu um erro: ", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func getLogs() {

	file, err := os.ReadFile("log.txt")

	if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(string(file))
}
func main() {

	displayIntro()

	for {
		displayMenu()
		command := readCommand()
		useSwitch(command)
	}
	
}