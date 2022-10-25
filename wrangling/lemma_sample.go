package wrangling

// identify by the hour

// get all strings and split to words (lemmatization)

import (
	"fmt"
	"log"
	"strings"

	"github.com/clipperhouse/jargon"

	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
)

var text string = `Let’s talk talked talks about Ruby on Rails and ASPNET MVC.`
var lemmatizer, err_lemma = golem.New(en.New())

func LemmaJargonSample() {

	jargon.TokenizeString(text)
	r := strings.NewReader(text)

	tokens := jargon.Tokenize(r)

	for i := 0; i < 2; i++ {

		fmt.Printf("%d", i)
		for {
			token, err := tokens.Next()
			if err != nil {
				// Because the source is I/O, errors are possible
				log.Fatal(err)
			}
			if token == nil {
				break
			}
			LemmaGolemWord(token.String())
		}
	}
}
