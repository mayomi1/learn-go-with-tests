package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const yoruba = "Yoruba"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "
const yorubaHelloPrefix = "Bawoni "

// Hello function
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	
	return greetingPrefix(language) + name
}

// note the in go private functions start with lower case
func greetingPrefix(language string) (prefix string) { // named return value
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix	
	case yoruba:
		prefix = yorubaHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Mayor", "Spanish"))
}
