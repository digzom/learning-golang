package main

import "fmt"

func Hello(receiver string) string  {
	return "Hello, " + "Dickson" + "!"
	
}

func main()  {
	fmt.Println(Hello("Dickson"))
}
