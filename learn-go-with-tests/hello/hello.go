package main

import "fmt"

const (
	enHelloPrefix   = "Hello, "
	ptBrHelloPrefix = "Olá, "
	frHelloPrefix   = "Bonjour, "
	esHelloPrefix   = "Ola, "
)

func Hello(receiver string, lang string) string {
	if receiver == "" {
		receiver = "World"
	}

	helloPrefix := enHelloPrefix
	switch lang {
	case "Portuguese":
		helloPrefix = ptBrHelloPrefix

	case "French":
		helloPrefix = frHelloPrefix

	case "Spanish":
		helloPrefix = esHelloPrefix
	}

	return helloPrefix + receiver + "!"
}

func main() {
	fmt.Println(Hello("Dickson", "English"))
}
