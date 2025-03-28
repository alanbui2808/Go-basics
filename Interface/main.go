package main

import "fmt"

// interface = collection of methods

// Our program has a new type called bot
// If you are a type that has a method getGreeting() string, you are an honory member of type bot
// Thus you can use func printGreeting(b bot) below

// this gives the idea of polymorphism.
type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// printGreeting(&eb) //Go can automatically deferrences this. However this is not the convention.
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "hello!"
}

func (spanishBot) getGreeting() string {
	return "hola!"
}
