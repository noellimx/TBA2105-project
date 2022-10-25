package wrangling

// identify by the hour
// get all strings and split to words (lemmatization)

import (
	"github.com/clipperhouse/jargon"
)

func Lemma() {
	text := `Let’s talk about Ruby on Rails and ASPNET MVC.`

	jargon.TokenizeString(text)

}

// table should look like

// yyyymmdd	| word | occurrence
// 20221227	| traffic	| 4
