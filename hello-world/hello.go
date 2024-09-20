package main

import (
	"fmt"
	"strings"
)

const (
	spanish = "spanish"
	french  = "french"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func main() {
	fmt.Println(Hello("Dzaky", ""))
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(lang)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	language = strings.ToLower(language)

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
