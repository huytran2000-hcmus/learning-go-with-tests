package helloworld

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	engHelloPrefix     = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = engHelloPrefix
	}

	return
}

func main() {
	x := "Huy"
	fmt.Println(Hello(x, ""))
}
