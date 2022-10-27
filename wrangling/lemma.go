package wrangling

// identify by the hour

// get all strings and split to words (lemmatization)

import (
	"strings"

	"github.com/clipperhouse/jargon"
	"github.com/noellimx/TBA2105-project/utils"
)

type typeLemmas struct {
	Text   string
	Lemmas []string
}

var blacklist []string = []string{"the", "of", "in", "to", " ", "(", ")", ".", "t.co", ":", ",", "http", "https", "/", "\n", "[", "]", "#"}

func shouldAppend(s string) bool {

	for _, b := range blacklist {
		if b == s {
			return false
		}
	}

	return true

}

func LemmatizeText(text string) *typeLemmas {

	jargon.TokenizeString(text)
	r := strings.NewReader(text)

	tokens := jargon.Tokenize(r)
	var lemmas []string
	for {
		token, err := tokens.Next()
		if err != nil {
			// Because the source is I/O, errors are possible
			utils.VFatal(err.Error())
		}
		if token == nil {
			break
		}
		lemma := LemmaGolemWord(token.String())

		if shouldAppend(lemma) {
			lemmas = append(lemmas, lemma)
		}
	}

	return &typeLemmas{
		Text:   text,
		Lemmas: lemmas,
	}
}

func LemmaGolemWord(word string) string {
	if err_lemma != nil {
		panic(err_lemma)
	}
	lemma := lemmatizer.Lemma(word)

	return lemma
}
