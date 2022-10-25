package wrangling

// identify by the hour

// get all strings and split to words (lemmatization)

import (
	"fmt"
	"log"
	"strings"

	"github.com/clipperhouse/jargon"
)

type typeLemmas struct {
	text   string
	lemmas []string
}

func Lemmatize(text string) *typeLemmas {

	jargon.TokenizeString(text)
	r := strings.NewReader(text)

	tokens := jargon.Tokenize(r)
	var lemmas []string
	for {
		token, err := tokens.Next()
		if err != nil {
			// Because the source is I/O, errors are possible
			log.Fatal(err)
		}
		if token == nil {
			break
		}
		lemma := LemmaGolemWord(token.String())
		lemmas = append(lemmas, lemma)
	}

	return &typeLemmas{
		text:   text,
		lemmas: lemmas,
	}
}

func LemmaGolemWord(word string) string {
	if err_lemma != nil {
		panic(err_lemma)
	}
	lemma := lemmatizer.Lemma(word)

	fmt.Printf("[LemmaGolemWord] %s -> %s\n", word, lemma)
	return lemma
}
