package main

import "fmt"

const cEnglishHelloPrefix = "Hello, "
const cFrenchHelloPrefix = "Bonjour, "
const cFrenchLanguageCode = "FR"

// Hello say hello
func Hello(name string, languageCode string) string {
	if name == "" {
		name = "World"
	}
	return getGreetingPrefix(languageCode) + name
}

func getGreetingPrefix(languageCode string) (greetingPrefix string) {
	greetingPrefix = cEnglishHelloPrefix
	switch languageCode {
	case cFrenchLanguageCode:
		greetingPrefix = cFrenchHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", "EN"))
}
