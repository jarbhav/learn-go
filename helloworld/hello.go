package helloworld

// We can group constants in a block instead of declaring them on their own line
const (
	spanish = "Spanish"
	french  = "French"
	hindi   = "Hindi"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	hindiHelloPrefix   = "Namaste, "
)

// Public functions start with a capital letter
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// Private functions start with a lowercase letter
func greetingPrefix(language string) (prefix string) {
	switch language {
	case hindi:
		prefix = hindiHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	// equivalent to `return prefix` since in the function signature we have made a named return value (prefix string)
	return
}
