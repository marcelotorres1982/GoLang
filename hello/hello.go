package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Define propriedades do Logger predefinido, incluindo
	// o prefixo de entrada de log e um sinalizador para desabilitar a impressão
	// da hora, do arquivo de origem e do número da linha.

	// Solicita uma mensagem de saudação.
	message, err := greetings.Hello("Giovanna")

	// Se um erro foi retornado, imprima-o no console e
	// saia do programa.

	if err != nil {
		log.Fatal(err)
	}

	// Se nenhum erro foi retornado, imprima a mensagem retornada
	// no console.

	fmt.Println(message)
}
