package helloworld

// name of package same as the directoey it is in
const (
	spanish = "Spanish"
	french  = "French"
	hindi   = "Hindi"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	hindiHelloPrefix   = "Namaste, "
)

func Hello(lang, name string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

// named returned type for better readability
func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case hindi:
		prefix = hindiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
