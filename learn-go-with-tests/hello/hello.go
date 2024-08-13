package main

import "fmt"

const helloPrefix = "Hello, "

func Hello(receiver string) string  {
	if receiver == "" {
		receiver = "World"
	}
	return helloPrefix + receiver + "!"
}

func main()  {
	fmt.Println(Hello("Dickson"))
}
