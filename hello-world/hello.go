package main

import (
	"fmt"
	"strings"
)

const (
	spanish = "Spanish"
	french = "French"
	swedish = "Swedish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	swedishHelloPrefix = "Tjena, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if strings.HasPrefix(name, englishHelloPrefix) {
		return name
	}

	return greetingPrefix(language) + name
}

func greetingPrefix (language string) (prefix string){
switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix	
	case swedish:
		prefix = swedishHelloPrefix	
	default: 
	prefix = englishHelloPrefix	
	}
	return
}


func main() {
	fmt.Println(Hello("World", ""))
}
