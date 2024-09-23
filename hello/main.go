package main

import (
	"flag"
	"fmt"
)

type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",    // Greek
	"en": "Hello world",      // English
	"fr": "Bonjour le monde", // French
}

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language")
	flag.Parse()
	greeting, err := greet(language(lang))
	if err != nil {
		panic(err)
	}
	fmt.Println(greeting)
}

func greet(l language) (string, error) {
	greeting, ok := phrasebook[l]
	if !ok {
		return "", fmt.Errorf("language not supported: %q", l)
	}

	return greeting, nil
}
