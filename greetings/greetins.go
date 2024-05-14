package greetings

import (
	"fmt"
)

// Hello irá retornar uma saudação com o nome da pessoa

func Hello(name string) string {
	//Retorna a saudação que incorpora o nome em uma mensagem
	message := fmt.Sprintf("Olá, %v. Bem vindo(a)!", name)

	return message
}
