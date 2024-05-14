package greetings

import (
	"errors"
	"fmt"

	"math/rand"
)

// Hello irá retornar uma saudação com o nome da pessoa

func Hello(name string) (string, error) {
	//Se nenhum nome for fornecido, retorna uma mensagem de erro
	//Retorna a saudação que incorpora o nome em uma mensagem
	if name == "" {
		return "", errors.New("empty name")
	}

	//Cria uma mensagem usando formato randômico

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat retorna uma de um conjunto de mensagens de saudação. A // mensagem retornada é selecionada aleatoriamente.
func randomFormat() string {
	//Uma fatia de formato de mensagem
	formats := []string{
		"Olá, %v! Bem vindo(a)",
		"Que bom ver você, %v",
		"Salve, %v! Bem vindo(a)",
	}
	// Retorna um formato de mensagem selecionado aleatoriamente especificando
	// um índice aleatório para a fatia de formatos.

	return formats[rand.Intn(len(formats))]

}
