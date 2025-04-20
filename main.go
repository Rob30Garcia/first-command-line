package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define uma flag chamada "uppercase"
	uppercase := flag.Bool("uppercase", false, "Exibe a mensagem em letras maiusculas")

	// Parsei as flags asntes de ler os args
	flag.Parse()

	// Pega os argumentos restantes depois das flags
	args := flag.Args()

	var name string

	if len(args) > 0 {
		name = args[0]
	} else {
		// Se o nome não for passado como argumento, pergunta interativamente
		fmt.Print("Digite seu nome: ")
		_, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Println("Erro ao ler o nome")
			os.Exit(1)
		}
	}

	message := fmt.Sprintf("Hello, %s. Welcome!", name)

	if *uppercase {
		message = strings.ToUpper(message)
	}

	fmt.Println(message)
}
