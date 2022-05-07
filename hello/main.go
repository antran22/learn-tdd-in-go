package hello

import (
	"fmt"
)

type Language string

const (
	English Language = "english"
	Spanish Language = "spanish"
	French  Language = "french"
)

const englishDefaultName = "World"
const spanishDefaultName = "Mundo"
const frenchDefaultName = "le Monde"
const englishPrefix = "Hello"
const spanishPrefix = "Hola"
const frenchPrefix = "Bonjour"

func Hello(name string, language Language) string {
	var prefix string
	var defaultName string

	switch language {
	case Spanish:
		defaultName = spanishDefaultName
		prefix = spanishPrefix
	case English:
		defaultName = englishDefaultName
		prefix = englishPrefix
	case French:
		defaultName = frenchDefaultName
		prefix = frenchPrefix
	}

	if name == "" {
		name = defaultName
	}

	return fmt.Sprintf("%s, %s", prefix, name)
}
