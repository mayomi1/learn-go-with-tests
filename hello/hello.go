package main

import "fmt"

const englishHelloPrefix = "Hello "
// Hello function
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Mayor"))
}
