package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Ler um arquivo txt
	data, err := ioutil.ReadFile("assets/mateus-1-25.txt")

	if err != nil {
		fmt.Println("Erro ao ler o arquivo")
		return
	}

	// Converte o arquivo em String
	text := string(data)

	// Dividir o conteúdo em linhas
	lines := strings.Split(text, "\n")

	// Inicializar o gerador de números randomicos
	rand.Seed(time.Now().UnixNano())

	// Gerar um valor randomico de 0 até a quantidade de linhas - 1
	line := rand.Intn(len(lines))

	// Imprimir a linha randomica
	fmt.Println("Mateus 1\n ", lines[line])
}
