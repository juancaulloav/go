package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
func randomFormat() string {
	formats := []string{
		"Hola, %v. bienvenido vieje!",
		"buena pelao, como va la vida %v?",
		"que onda mi bro %v, que haces?",
	}
	return formats[rand.Intn(len(formats))]
}
