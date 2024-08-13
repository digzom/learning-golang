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

	return greetingPrefix(lang) + receiver + "!"
}

func greetingPrefix(lang string) string {
	helloPrefix := enHelloPrefix
	switch lang {
	case "Portuguese":
		helloPrefix = ptBrHelloPrefix

	case "French":
		helloPrefix = frHelloPrefix

	case "Spanish":
		helloPrefix = esHelloPrefix
	}

	return helloPrefix
}

func main() {
	fmt.Println(Hello("Dickson", "English"))
}
